package gikun

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type router struct {
	roots    map[string]*Node       //roots 来存储每种请求方式的Trie 树根节点 GET/POST
	handlers map[string]HandlerFunc //handlers 存储每种请求方式的 HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*Node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed 路径按/划分
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &Node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
	log.Printf("[%s] %s\n", method, pattern)
}

func (r *router) getRoute(method string, pattern string) (*Node, map[string]string) {
	searchParts := parsePattern(pattern)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

// 处理handle
func (r *router) handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.pattern
		t1 := time.Now()
		r.handlers[key](c)
		t2 := time.Now()
		log.Printf("[%s] %s %dms\n", c.Method, c.Path, (t2.Nanosecond()-t1.Nanosecond())/1e6)
	} else {
		c.SendString(http.StatusNotFound, fmt.Sprintf("404 NOT FOUND: %s\n", c.Path))
	}
}
