package rest

import (
	"context"
	"math"
	"net/http"
	"sync"
)

// abortIndex 一个极大值
// 一定比 handlers 数量大，导致 next 函数执行中断
const abortIndex int8 = math.MaxInt8 / 2

type Context struct {
	context.Context
	// Request 
	Request   *http.Request
	ResWriter http.ResponseWriter

	// Keys 用于在控制流中传递内容
	Keys map[string]interface{}
	// 用于标志唯一请求，上下文传递
	RequestID string

	// mu 保护 Keys map
	mu sync.RWMutex

	handlers []func(c *Context)
	index    int8
}

func (c *Context) Reset() {
	c.index = -1
}
func (c *Context) ErrorJSON() {
}
func (c *Context) OkJSON() {
}
func (c *Context) renderJSON(code int, jsonValue interface{}) {
}
func (c *Context) Abort() {
	c.index = abortIndex
}
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}