package webframe

import (
	"context"
	"encoding/json"
	"net/http"
)

type M = map[string]interface{}

type Context struct {
	context.Context

	Writer  http.ResponseWriter
	Request *http.Request

	Path   string
	Method string

	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Context: context.TODO(),

		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) (string, bool) {
	q := c.Request.URL.Query()
	if q.Has(key) {
		return q.Get(key), true
	} else {
		return "", false
	}
}

func (c *Context) SetStatus(code int) {
	c.Writer.WriteHeader(code)
	c.StatusCode = code
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) StringResponse(code int, str string) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.Writer.Write([]byte(str))
}

func (c *Context) JSONResponse(code int, jsonObj M) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(jsonObj); err != nil {
		c.ErrorResponse(http.StatusInternalServerError, err)
	}
}

func (c *Context) ErrorResponse(code int, err error) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(code)
	c.Writer.Write([]byte(err.Error()))
}
