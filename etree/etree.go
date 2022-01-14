package etree

import (
	"errors"
)

type Node struct {
	cache  map[error]struct{}
	parent []*Node
	err    error
}

func New(err error, parents ...error) error {
	n := &Node{
		parent: nil,
		cache:  make(map[error]struct{}),
		err:    err,
	}
	for _, parent := range parents {
		if p, ok := parent.(*Node); ok {
			n.parent = append(n.parent, p)
		}
	}
	return n
}

func (n *Node) Error() string {
	return n.err.Error()
}

func Cover(base, target error) bool {
	if _, ok := base.(*Node); ok {
		if _, ok := base.(*Node).cache[target]; ok {
			return true
		}
	} else {
		return errors.Is(base, target)
	}
	if _, ok := target.(*Node); !ok {
		return false
	}
	visited := make(map[error]bool)
	next := make([]*Node, 0, 10)
	next = append(next, base.(*Node))
	for len(next) > 0 {
		err := next[0]
		next = next[1:]
		if visited[err] {
			continue
		}
		visited[err] = true
		if err.err == target.(*Node).err {
			base.(*Node).cache[target] = struct{}{}
			return true
		}
		next = append(next, err.parent...)
	}
	return false
}