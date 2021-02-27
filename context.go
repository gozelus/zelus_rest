package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
	"time"
)

type jwtUtils interface {
	NewToken(uid int64) (string, error)
}

var _ context.Context = &contextImp{}
var _ Context = &contextImp{}

type contextImp struct {
	gc *gin.Context

	queryMap            map[string]string
	requestBodyJsonStr  string
	responseBodyJsonStr string
	httpCode            int

	// err 用于存储中间可能发生的错误
	err error

	validate *validator.Validate
	jwtUtils jwtUtils
}

func (c *contextImp) Deadline() (deadline time.Time, ok bool) {
	return c.gc.Request.Context().Deadline()
}

func (c *contextImp) Done() <-chan struct{} {
	return c.gc.Request.Context().Done()
}

func (c *contextImp) Err() error {
	return c.gc.Request.Context().Err()
}

func (c *contextImp) Value(key interface{}) interface{} {
	return c.gc.Request.Context().Value(key)
}

// ************************ jwt auth imp begin
func (c *contextImp) setJwtUtils(utils jwtUtils) {
	c.jwtUtils = utils
}
func (c *contextImp) setJwtToken(tokenStr string) {
	c.gc.Set("jwt-token", tokenStr)
}
func (c *contextImp) GetJwtToken() string {
	return c.gc.GetString("jwt-token")
}
func (c *contextImp) setUserID(uid int64) {
	c.gc.Set("jwt-user-id", uid)
}
func (c *contextImp) UserID() int64 {
	return c.gc.GetInt64("jwt-user-id")
}
func (c *contextImp) SetUserID(uid int64) {
	token, err := c.jwtUtils.NewToken(uid)
	if err != nil {
		c.RenderErrorJSON(nil, statusUnauthorized)
		return
	}
	c.setJwtToken(token)
}

// ************************ jwt auth imp end

// contextImp 的构造函数
func newContext(gc *gin.Context) Context {
	c := contextImp{
		gc:       gc,
		validate: validator.New(),
	}
	c.queryMap = map[string]string{}

	// copy request body
	if c.gc != nil {
		if c.gc.Request != nil {
			bodyBytes, _ := ioutil.ReadAll(c.gc.Request.Body)
			c.requestBodyJsonStr = string(bodyBytes)
			_ = c.gc.Request.Body.Close() //  must close
			c.gc.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			// let query to map
			for k := range c.gc.Request.URL.Query() {
				if v, ok := c.gc.GetQuery(k); ok {
					c.queryMap[k] = v
				}
			}
		}
	}
	return &c
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

func (c *contextImp) ResponseBodyJsonStr() string {
	return c.responseBodyJsonStr
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
func (c *contextImp) Headers() map[string][]string {
	return c.gc.Request.Header
}
func (c *contextImp) SetResponseHeader(key, value string) {
	c.gc.Writer.Header().Set(key, value)
}
func (c *contextImp) Method() string {
	return c.gc.Request.Method
}
func (c *contextImp) Path() string {
	return c.gc.Request.URL.Path
}
func (c *contextImp) setTimeout(duration time.Duration) {
	ctx := c.gc.Request.Context()
	c.gc.Request = c.gc.Request.WithContext(ctx)
}
func (c *contextImp) setRequestID(id string) {
	c.gc.Set("rest-request-id", id)
}
func (c *contextImp) GetRequestID() string {
	return c.gc.GetString("rest-request-id")
}
func (c *contextImp) GetError() error {
	return c.err
}

// ********************* control follow begin
func (c *contextImp) RenderOkJSON(data interface{}) {
	resp := standResp{
		Code:      200,
		Message:   "success",
		Data:      data,
		RequestID: c.GetRequestID(),
	}
	resp.Token = c.GetJwtToken()
	_ = c.renderJSON(http.StatusOK, resp)
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
	resp.Token = c.GetJwtToken()
	_ = c.renderJSON(theError.GetCode(), resp)
}
func (c *contextImp) Next() {
	c.gc.Next()
}
func (c *contextImp) JSONBodyBind(ptr interface{}) error {
	if err := c.gc.ShouldBind(ptr); err != nil {
		return err
	}
	if err := c.validate.Struct(ptr); err != nil {
		return err
	}
	return nil
}
func (c *contextImp) JSONQueryBind(ptr interface{}) error {
	if err := c.gc.BindQuery(ptr); err != nil {
		return err
	}
	if err := c.validate.Struct(ptr); err != nil {
		return err
	}
	return nil
}

// ******************** control follow end

// private func
func (c *contextImp) renderJSON(code int, obj interface{}) error {
	c.httpCode = code
	c.gc.AbortWithStatusJSON(code, obj)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	c.responseBodyJsonStr = string(jsonBytes)
	return err
}
