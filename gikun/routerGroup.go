package gikun

type RouterGroup struct {
	prefix      string        //分组前缀
	middlewares []HandlerFunc // 中间件
	engine      *Engine       // 引擎指针
	router      *router       //全部路由
}

// 使用中间件
func (g *RouterGroup) Use(handlers ...HandlerFunc) *RouterGroup {
	g.middlewares = append(g.middlewares, handlers...)
	return g
}

// 分组
func (g *RouterGroup) Group(prefix string) *RouterGroup {
	group := &RouterGroup{
		prefix:      g.prefix + prefix,
		middlewares: make([]HandlerFunc, 0),
		engine:      g.engine,
		router:      g.router,
	}
	g.engine.groups = append(g.engine.groups, group)
	return group
}

func (g *RouterGroup) addRoute(method, pattern string, handler HandlerFunc) {
	p := g.prefix + pattern
	g.router.addRoute(method, p, handler)
}
func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}
func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}
