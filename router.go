package webframe

import (
	"errors"
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*routingNode
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*routingNode),
	}
}

func parseRoutingPattern(pattern string) []string {
	rawParts := strings.Split(pattern, "/")
	parts := make([]string, 0, len(rawParts))
	for _, part := range rawParts {
		if part == "" {
			continue
		}
		parts = append(parts, part)
	}
	return parts
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parseRoutingPattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &routingNode{
			children: make(map[string]*routingNode),
		}
	}
	r.roots[method].insert(pattern, parts, handler, 0)
}

func (r *router) getRoute(method, path string) *routingNode {
	parts := parseRoutingPattern(path)
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	return root.search(parts, 0)
}

func (r *router) handle(c *Context) {
	n := r.getRoute(c.Method, c.Path)
	if n != nil {
		n.handler(c)
	} else {
		c.ErrorResponse(http.StatusNotFound, errors.New("route not found"))
	}
}
