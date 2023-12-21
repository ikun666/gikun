package gikun

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

// 添加路由
func (r *router) addRoute(method, path string, handler HandlerFunc) {
	log.Printf("Route [%s] - %s", method, path)
	key := method + "-" + path
	r.handlers[key] = handler
}

// 处理handle
func (r *router) handle(c *Context) {
	//获取方法-url路径
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		t1 := time.Now()
		handler(c)
		t2 := time.Now()
		log.Printf("[%s] %s %dms\n", c.Method, c.Path, (t2.Nanosecond()-t1.Nanosecond())/1e6)
	} else {
		c.SendString(http.StatusNotFound, fmt.Sprintf("404 NOT FOUND: %s\n", c.Path))
	}
}
