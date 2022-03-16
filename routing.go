package webframe

// 目前暂不用通配符等模糊匹配的路由规则，只做路由的精准匹配。
// 后续可以考虑增加实用的模糊通配路由逻辑来增强场景支持。

type routingNode struct {
	children map[string]*routingNode
	pattern  string
	part     string
	handler  HandlerFunc
}

func (n *routingNode) insert(pattern string, parts []string, handler HandlerFunc, height int) {
	if len(parts) == height {
		n.pattern = pattern
		n.handler = handler
		return
	}
	part := parts[height]
	child := n.matchNode(part)
	if child == nil {
		child = &routingNode{
			children: make(map[string]*routingNode),
			part:     part,
		}
		n.children[part] = child
	}
	child.insert(pattern, parts, handler, height+1)
}

func (n *routingNode) search(parts []string, height int) *routingNode {
	if len(parts) == height {
		return n
	}
	part := parts[height]
	child := n.matchNode(part)
	if child == nil {
		return nil
	}
	return child.search(parts, height+1)
}

func (n *routingNode) matchNode(part string) *routingNode {
	child := n.children[part]
	return child
}
