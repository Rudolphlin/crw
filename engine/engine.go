package engine

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}

}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) registRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) Run(port string) (err error) {
	http.ListenAndServe("port", engine)
}

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.registRoute("GET", pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandlerFunc) {
	e.registRoute("POST", pattern, handler)
}

func (e *Engine) Put(pattern string, handler HandlerFunc) {
	e.registRoute("PUT", pattern, handler)
}

func (e *Engine) Delete(pattern string, handler HandlerFunc) {
	e.registRoute("DEL", pattern, handler)
}
