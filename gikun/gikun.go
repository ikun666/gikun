package gikun

import (
	"net/http"
)

type HandlerFunc func(*Context)
type Engine struct {
	router *router
}

// 创建Engine实例
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// 添加路由
func (e *Engine) addRoute(method, path string, hander HandlerFunc) {
	e.router.addRoute(method, path, hander)
}

// 封装GET
func (e *Engine) GET(path string, hander HandlerFunc) {
	e.addRoute("GET", path, hander)
}

// 封装POST
func (e *Engine) POST(path string, hander HandlerFunc) {
	e.addRoute("POST", path, hander)
}

// Engine Run
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// 实现ServeHTTP接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	c := NewContext(w, req)
	e.router.handle(c)
}
