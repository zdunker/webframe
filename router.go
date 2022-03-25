package webframe

import (
	"errors"
	"net/http"
	"strings"
)

type router struct {
	roots map[string]*routingNode
}

func newRouter() *router {
	return &router{
		roots: make(map[string]*routingNode),
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

func (r *router) addRoute(method, pattern string, handler HandlerFunc, group *routerGroup) {
	parts := parseRoutingPattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &routingNode{
			children: make(map[string]*routingNode),
		}
	}
	r.roots[method].insert(group, pattern, parts, handler, 0)
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
		c.handlers = append(n.group.middlewares, n.handler)
	} else {
		c.handlers = append(c.handlers, Logger(), func(ctx *Context) {
			c.ErrorResponse(http.StatusNotFound, errors.New("route not found"))
		})
	}
	c.Next()
}
