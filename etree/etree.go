package etree

import (
	"errors"
)

type Node struct {
	cache   map[error]struct{}
	parent  []*Node
	message string
}

func New(message string, parents ...*Node) *Node {
	n := &Node{
		parent:  nil,
		cache:   make(map[error]struct{}),
		message: message,
	}
	for _, parent := range parents {
		n.parent = append(n.parent, parent)

	}
	return n
}

func (n *Node) Error() string {
	return n.message
}

func Cover(base, target error) bool {
	baseNode, ok := base.(*Node)
	if ok {
		if _, ok := baseNode.cache[target]; ok {
			return true
		}
	} else {
		return errors.Is(base, target)
	}
	targetNode, ok := target.(*Node)
	if !ok {
		return false
	}
	visited := make(map[*Node]bool)
	next := make([]*Node, 0, len(targetNode.parent)*2)
	next = append(next, baseNode)
	for len(next) > 0 {
		err := next[0]
		next = next[1:]
		if visited[err] {
			continue
		}
		visited[err] = true
		if err.message == targetNode.message {
			baseNode.cache[target] = struct{}{}
			return true
		}
		next = append(next, err.parent...)
	}
	return false
}
