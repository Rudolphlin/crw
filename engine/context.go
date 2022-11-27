package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Req *http.Request
	Rsp http.ResponseWriter

	// request info
	Path   string
	Method string

	StatusCode int
}

func newContext(rsp http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Rsp:    rsp,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Rsp.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Rsp.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Rsp.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Rsp)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Rsp, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Rsp.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Rsp.Write([]byte(html))
}
