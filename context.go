package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gozelus/zelus_rest/core/bindding"
	"io"
	"math"
	"net/http"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gozelus/zelus_rest/logger"
	"github.com/pkg/errors"

	"github.com/google/uuid"
)

// abortIndex 一个极大值
// 一定比 handlers 数量大，导致 next 函数执行中断
const abortIndex int8 = math.MaxInt8 / 2

var _ Context = &contextImp{}

type contextImp struct {
	context.Context
	request   *http.Request
	resWriter http.ResponseWriter

	// err 用于存储中间可能发生的错误
	err error
	// Keys 用于在控制流中传递内容
	keys map[string]interface{}
	// 用于标志唯一请求，上下文传递
	requestID string

	// mu 保护 Keys map
	mu sync.RWMutex

	validate *validator.Validate
	handlers []HandlerFunc
	index    int8
}

// contextImp 的构造函数
func newContext() *contextImp {
	c := contextImp{}
	c.validate = validator.New()
	return &c
}
func (c *contextImp) init(w http.ResponseWriter, req *http.Request) {
	c.Context = context.Background()
	c.request = req
	c.resWriter = w
	c.keys = map[string]interface{}{}
	c.requestID = strings.Replace(uuid.Must(uuid.NewRandom()).String(), "-", "", -1)
	c.index = -1
}
func (c *contextImp) RenderOkJSON(data interface{}) {
	_ = c.renderJSON(http.StatusOK, struct {
		Code      int         `json:"code"`
		Message   string      `json:"message"`
		RequestID string      `json:"request_id"`
		Data      interface{} `json:"data"`
	}{
		Code:      200,
		Message:   "success",
		Data:      data,
		RequestID: c.GetRequestID(),
	})
}

func (c *contextImp) GetError() error {
	return c.err
}

func (c *contextImp) RenderErrorJSON(data interface{}, err error) {
	var theError StatusError = statusInternalServerError
	if val, ok := err.(StatusError); ok {
		theError = val
	}
	c.err = theError
	_ = c.renderJSON(theError.GetCode(), struct {
		Code      int         `json:"code"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
		RequestID string      `json:"request_id"`
	}{
		Code:      theError.GetCode(),
		Message:   theError.GetMessage(),
		Data:      data,
		RequestID: c.GetRequestID(),
	})
}

func (c *contextImp) Headers() map[string][]string {
	return c.request.Header
}
func (c *contextImp) Method() string {
	return c.request.Method
}
func (c *contextImp) Path() string {
	return c.request.URL.String()
}
func (c *contextImp) GetRequestID() string { return c.requestID }
func (c *contextImp) Set(key string, v interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.keys[key] = v
}

func (c *contextImp) Get(key string) (v interface{}, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok = c.keys[key]
	return
}

func (c *contextImp) setHandlers(hs ...HandlerFunc) {
	c.handlers = hs
}
func (c *contextImp) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}
func (c *contextImp) JSONBodyBind(ptr interface{}) error {
	var reader io.Reader
	if c.request.ContentLength > 0 && strings.Contains(c.request.Header.Get("Content-Type"), "application/json") {
		reader = c.request.Body
	} else {
		reader = strings.NewReader("{}")
	}
	err := json.NewDecoder(reader).Decode(ptr)
	if err != nil {
		return err
	}
	if err := c.validate.Struct(ptr); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logger.Warnf("eer : %s ", err.Error())
			return err
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Namespace = %s \n", err.Namespace())
			fmt.Printf("Field = %s \n", err.Field())
			fmt.Printf("StructNamespace = %s \n", err.StructNamespace())
			fmt.Printf("StructField = %s \n", err.StructField())
			fmt.Printf("Tag = %s \n", err.Tag())
			fmt.Printf("ActualTag = %s \n", err.ActualTag())
			fmt.Printf("Kind = %s \n", err.Kind())
			fmt.Printf("Type = %s \n", err.Type())
			fmt.Printf("Value = %s \n", err.Value())
			fmt.Printf("Param = %s \n", err.Param())
			fmt.Println()
		}
		return errors.New("?")
	}
	return nil
}
func (c *contextImp) JSONQueryBind(ptr interface{}) error {
	if err := binding.Query.Bind(c.request, ptr); err != nil {
		return err
	}
	if err := c.validate.Struct(ptr); err != nil {
		return err
	}
	return nil
}

// private func

// abort 用于中断流
func (c *contextImp) abort() {
	c.index = abortIndex
}
func (c *contextImp) renderJSON(code int, obj interface{}) error {
	defer c.abort()
	c.resWriter.Header().Add("Content-Type", "application/json; charset=utf-8")
	c.resWriter.WriteHeader(code)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = c.resWriter.Write(jsonBytes)
	return err
}
