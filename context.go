package rest

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
)

// abortIndex 一个极大值
// 一定比 handlers 数量大，导致 next 函数执行中断
const abortIndex int8 = math.MaxInt8 / 2

type contextImp struct {
	context.Context
	request   *http.Request
	resWriter http.ResponseWriter

	// Keys 用于在控制流中传递内容
	keys map[string]interface{}
	// 用于标志唯一请求，上下文传递
	requestID string

	// mu 保护 Keys map
	mu sync.RWMutex

	handlers []HandlerFunc
	index    int8
}

func (c *contextImp) init(w http.ResponseWriter, req *http.Request) {
	c.Context = context.Background()
	c.request = req
	c.resWriter = w
	c.keys = map[string]interface{}{}
	c.requestID = strings.Replace(uuid.Must(uuid.NewRandom()).String(), "-", "", -1)
	c.index = -1
}
func (c *contextImp) OkJSON(obj interface{}) {
	c.renderJSON(200, obj)
}
func (c *contextImp) ErrorJSON(err ErrorInterface) {
	c.renderJSON(err.ErrorCode(), Error{
		Code:    err.ErrorCode(),
		Message: err.ErrorMessage(),
	})
}
func (c *contextImp) Headers() map[string][]string {
	return c.request.Header
}
func (c *contextImp) Method() string {
	return c.request.Method
}
func (c *contextImp) Path() string {
	return c.request.URL.Path
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
		if err := c.handlers[c.index](c); err != nil {
			c.ErrorJSON(err)
		}
		c.index++
	}
}
func (c *contextImp) JSONBodyBind(ptr interface{}) error {
	return json.NewDecoder(c.request.Body).Decode(ptr)
}
func (c *contextImp) JSONQueryBind(ptr interface{}) error {
	form := map[string]interface{}{}
	for k, v := range c.request.URL.Query() {
		if len(v) > 0 {
			form[k] = v[0]
		}
	}
	bytes, _ := json.Marshal(form)
	return json.Unmarshal(bytes, ptr)
}

// private func

// abort 用于中断流
func (c *contextImp) abort() {
	c.index = abortIndex
}
func (c *contextImp) renderJSON(code int, obj interface{}) error {
	defer c.abort()
	c.resWriter.WriteHeader(code)
	header := c.resWriter.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/javascript; charset=utf-8"}
	}
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = c.resWriter.Write(jsonBytes)
	return err
}
