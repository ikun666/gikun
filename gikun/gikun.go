package gikun

import (
	"fmt"
	"net/http"
	"time"
)

type HandleFunc func(http.ResponseWriter, *http.Request)
type Engine struct {
	router map[string]HandleFunc
}

// 创建Engine实例
func New() *Engine {
	return &Engine{
		router: make(map[string]HandleFunc),
	}
}

// 添加路由
func (e *Engine) addRoute(method, path string, hander HandleFunc) {
	key := method + "-" + path
	e.router[key] = hander
}

// 封装GET
func (e *Engine) GET(path string, hander HandleFunc) {
	e.addRoute("GET", path, hander)
}

// 封装POST
func (e *Engine) POST(path string, hander HandleFunc) {
	e.addRoute("POST", path, hander)
}

// Engine Run
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// 实现ServeHTTP接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//获取方法-url路径
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		t1 := time.Now()
		handler(w, req)
		t2 := time.Now()
		fmt.Printf("[%s] %s %vms\n", req.Method, req.URL.Path, (t2.Nanosecond()-t1.Nanosecond())/1e6)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND:%s", key)
	}
}
