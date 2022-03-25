package webframe

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	*routerGroup
	router *router
	groups []*routerGroup
}

func NewEngine() *Engine {
	engine := &Engine{router: newRouter()}
	engine.routerGroup = &routerGroup{engine: engine}
	engine.groups = []*routerGroup{engine.routerGroup}
	return engine
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc, group *routerGroup) {
	e.router.addRoute(method, pattern, handler, group)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// middlewares := make([]HandlerFunc, 0)
	// for _, group := range e.groups {
	// 	if strings.HasPrefix(r.URL.Path, group.routePrefix) {
	// 		middlewares = append(middlewares, group.middlewares...)
	// 	}
	// }
	c := newContext(w, r)
	// c.handlers = middlewares
	e.router.handle(c)
}
