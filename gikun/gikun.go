package gikun

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)
type Engine struct {
	*RouterGroup                //引擎为根分组
	groups       []*RouterGroup //保存所有路由分组

}

// 创建Engine实例
func New() *Engine {
	engine := &Engine{
		groups: make([]*RouterGroup, 0),
	}
	engine.RouterGroup = &RouterGroup{
		prefix:      "",
		middlewares: make([]HandlerFunc, 0),
		engine:      engine,
		router:      newRouter(),
	}
	//engine为根分组
	engine.groups = append(engine.groups, engine.RouterGroup)
	return engine
}
func Default() *Engine {
	engine := &Engine{
		groups: make([]*RouterGroup, 0),
	}
	engine.RouterGroup = &RouterGroup{
		prefix:      "",
		middlewares: make([]HandlerFunc, 0),
		engine:      engine,
		router:      newRouter(),
	}
	//engine为根分组
	engine.groups = append(engine.groups, engine.RouterGroup)
	engine.Use(Recovery())
	return engine
}

// // 添加路由
// func (e *Engine) addRoute(method, path string, hander HandlerFunc) {
// 	e.router.addRoute(method, path, hander)
// }

// // 封装GET
// func (e *Engine) GET(path string, hander HandlerFunc) {
// 	e.addRoute("GET", path, hander)
// }

// // 封装POST
// func (e *Engine) POST(path string, hander HandlerFunc) {
// 	e.addRoute("POST", path, hander)
// }

// Engine Run
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// 实现ServeHTTP接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//获取请求所需中间件
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(w, req)
	c.handler = middlewares
	e.router.handle(c)
}
