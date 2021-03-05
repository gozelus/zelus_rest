package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gozelus/zelus_rest/core/bindding"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type jwtUtils interface {
	NewToken(uid int64) (string, error)
}

var _ context.Context = &contextImp{}
var _ Context = &contextImp{}

type contextImp struct {
	gc *gin.Context

	validate *validator.Validate
	jwtUtils jwtUtils
}

func (c *contextImp) Deadline() (deadline time.Time, ok bool) {
	return c.gc.Deadline()
}

func (c *contextImp) Done() <-chan struct{} {
	return c.gc.Done()
}

func (c *contextImp) Err() error {
	return c.gc.Errors
}

func (c *contextImp) Value(key interface{}) interface{} {
	return c.gc.Value(key)
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
	// copy request body
	if c.gc != nil {
		if c.gc.Request != nil {
			if _, ok := c.gc.Get("rest-request-body-json-str"); !ok {
				bodyBytes, _ := ioutil.ReadAll(c.gc.Request.Body)
				c.gc.Set("rest-request-body-json-str", string(bodyBytes))
				_ = c.gc.Request.Body.Close() //  must close
				c.gc.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}
			if _, ok := c.gc.Get("rest-query-map"); !ok {
				// let query to map
				queryMap := map[string]string{}
				for k := range c.gc.Request.URL.Query() {
					if v, ok := c.gc.GetQuery(k); ok {
						queryMap[k] = v
					}
				}
				c.gc.Set("rest-query-map", queryMap)
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
	return c.gc.GetString("rest-response-body-json-str")
}
func (c *contextImp) QueryMap() map[string]string {
	if val, ok := c.gc.Get("rest-query-map"); ok {
		return val.(map[string]string)
	}
	return nil
}
func (c *contextImp) RequestBodyJsonStr() string {
	return c.gc.GetString("rest-request-body-json-str")
}
func (c *contextImp) HttpCode() int {
	return c.gc.Writer.Status()
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
func (c *contextImp) setError(err error) {
	c.gc.Set("rest-error", err)
}
func (c *contextImp) setTimeout(duration time.Duration) {
	ctx := c.gc.Request.Context()
	ctx, _ = context.WithTimeout(ctx, duration)
	c.gc.Request = c.gc.Request.WithContext(ctx)
}
func (c *contextImp) setRequestID(id string) {
	c.gc.Set("rest-request-id", id)
}
func (c *contextImp) GetRequestID() string {
	return c.gc.GetString("rest-request-id")
}
func (c *contextImp) GetError() error {
	if err, ok := c.gc.Get("rest-error"); ok {
		return err.(error)
	}
	return nil
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
	c.setError(err)
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
	var err error
	if c.gc.Request.ContentLength > 0 && strings.Contains(c.gc.Request.Header.Get("Content-Type"), "application/json") {
		err = binding.JSON.Bind(c.gc.Request, ptr)
	} else {
		err = binding.Form.Bind(c.gc.Request, ptr)
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
	if err := binding.Query.Bind(c.gc.Request, ptr); err != nil {
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
	c.gc.AbortWithStatusJSON(code, obj)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	c.gc.Set("rest-response-body-json-str", string(jsonBytes))
	return err
}
