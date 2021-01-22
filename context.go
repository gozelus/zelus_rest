package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gozelus/zelus_rest/core/bindding"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"
)

// abortIndex 一个极大值
// 一定比 handlers 数量大，导致 next 函数执行中断
const abortIndex int8 = math.MaxInt8 / 2

type jwtUtils interface {
	NewToken(uid int64) (string, error)
}

var _ context.Context = &contextImp{}

type contextImp struct {
	context.Context
	request   *http.Request
	resWriter http.ResponseWriter

	queryMap           map[string]string
	requestBodyJsonStr string
	httpCode           int

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
	jwtUtils jwtUtils
}

func (c *contextImp) QueryMap() map[string]string {
	return c.queryMap
}

func (c *contextImp) RequestBodyJsonStr() string {
	return c.requestBodyJsonStr
}

func (c *contextImp) HttpCode() int {
	return c.httpCode
}

func (c *contextImp) setJwtUtils(utils jwtUtils) {
	c.jwtUtils = utils
}

func (c *contextImp) setJwtToken(tokenStr string) {
	c.Set("jwt-token", tokenStr)
}
func (c *contextImp) setUserID(uid int64) {
	c.Set("jwt-user-id", uid)
}

func (c *contextImp) SetUserID(uid int64) {
	token, err := c.jwtUtils.NewToken(uid)
	if err != nil {
		c.RenderErrorJSON(nil, statusUnauthorized)
		return
	}
	c.setJwtToken(token)
}

func (c *contextImp) UserID() int64 {
	if val, ok := c.Get("jwt-user-id"); ok {
		return val.(int64)
	}
	return 0
}

// contextImp 的构造函数
func newContext() *contextImp {
	c := contextImp{}
	c.validate = validator.New()
	return &c
}
func (c *contextImp) init(w http.ResponseWriter, req *http.Request, timeOut *time.Duration) {
	c.Context = context.Background()
	if timeOut != nil {
		c.Context, _ = context.WithTimeout(c.Context, *timeOut)
	}
	c.request = req
	c.resWriter = w
	c.keys = map[string]interface{}{}
	c.requestID = strings.Replace(uuid.Must(uuid.NewRandom()).String(), "-", "", -1)
	c.index = -1
	c.queryMap = map[string]string{}

	// copy request body
	if req != nil {
		bodyBytes, _ := ioutil.ReadAll(req.Body)
		c.requestBodyJsonStr = string(bodyBytes)
		req.Body.Close() //  must close
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// let query to map
		for k := range req.URL.Query() {
			c.queryMap[k] = req.URL.Query().Get(k)
		}
	}
}

type standResp struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"request_id"`
	Reason    struct {
		Message string
		Code    int
	} `json:"reason"`
	Token string `json:"token"`
}

func (c *contextImp) RenderOkJSON(data interface{}) {
	resp := standResp{
		Code:      200,
		Message:   "success",
		Data:      data,
		RequestID: c.GetRequestID(),
	}
	if token, ok := c.Get("jwt-token"); ok {
		resp.Token = token.(string)
	}
	_ = c.renderJSON(http.StatusOK, resp)
}

func (c *contextImp) GetError() error {
	return c.err
}

func (c *contextImp) RenderErrorJSON(data interface{}, err error) {
	var theError StatusError = statusInternalServerError
	c.err = err
	if val, ok := err.(StatusError); ok {
		theError = val
	}
	resp := standResp{
		Code:      theError.GetCode(),
		Message:   theError.GetMessage(),
		Data:      data,
		RequestID: c.GetRequestID(),
	}

	if theError.GetReason() != nil {
		resp.Reason.Code = theError.GetReason().GetReasonCode()
		resp.Reason.Message = theError.GetReason().GetReasonMessage()
	}

	if token, ok := c.Get("jwt-token"); ok {
		resp.Token = token.(string)
	}

	_ = c.renderJSON(theError.GetCode(), resp)
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
func (c *contextImp) File(name string) (io.Reader, error) {
	f, _, err := c.request.FormFile(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}
func (c *contextImp) JSONBodyBind(ptr interface{}) error {
	var err error
	if c.request.ContentLength > 0 && strings.Contains(c.request.Header.Get("Content-Type"), "application/json") {
		err = binding.JSON.Bind(c.request, ptr)
	} else {
		err = binding.Form.Bind(c.request, ptr)
	}
	if err != nil {
		return err
	}
	if err = c.validate.Struct(ptr); err != nil {
		return err
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
	c.httpCode = code

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = c.resWriter.Write(jsonBytes)
	return err
}
