package engine

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *Router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) Run(port string) (err error) {
	return http.ListenAndServe(port, e)
}

func (e *Engine) registRouter(method, pattern string, handler HandlerFunc) {
	e.router.registRouter(method, pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.registRouter("GET", pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandlerFunc) {
	e.registRouter("POST", pattern, handler)
}

func (e *Engine) Put(pattern string, handler HandlerFunc) {
	e.registRouter("PUT", pattern, handler)
}

func (e *Engine) Delete(pattern string, handler HandlerFunc) {
	e.registRouter("DEL", pattern, handler)
}
