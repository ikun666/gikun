package gikun

type RouterGroup struct {
	prefix      string        //分组前缀
	middlewares []HandlerFunc // 中间件
	engine      *Engine       // 引擎指针
	router      *router       //全部路由
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix:      g.prefix + prefix,
		middlewares: g.middlewares,
		engine:      g.engine,
		router:      g.router,
	}
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
