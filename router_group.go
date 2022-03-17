package webframe

import "net/http"

type routerGroup struct {
	routePrefix string
	engine      *Engine
	middlewares []HandlerFunc
}

func (g *routerGroup) NewGroup(prefix string, middlewares ...HandlerFunc) *routerGroup {
	newGroup := &routerGroup{
		routePrefix: g.routePrefix + prefix,
		engine:      g.engine,
		middlewares: make([]HandlerFunc, 0),
	}
	newGroup.middlewares = append(newGroup.middlewares, g.middlewares...)
	return newGroup
}

func (g *routerGroup) WithMiddleware(mws ...HandlerFunc) *routerGroup {
	g.middlewares = append(g.middlewares, mws...)
	return g
}

func (g *routerGroup) addRoute(method, pattern string, handler HandlerFunc) {
	g.engine.addRoute(method, g.routePrefix+pattern, handler)
}

func (g *routerGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodGet, pattern, handler)
}

func (g *routerGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodPost, pattern, handler)
}
