package main

import (
	"fmt"
	"strings"
)

// Node represents a node in tree
type Node struct {
	Name       string
	Parent     *Node
	depth      int
	depthKnown bool
	onPath     bool
}

// GetDepth returns the depth of this node in the tree
func (n *Node) GetDepth() int {
	if n.Parent == nil {
		return 0
	}
	if !n.depthKnown {
		n.depth = n.Parent.GetDepth() + 1 // memoize
		n.depthKnown = true
	}
	return n.depth
}

// SetRoot sets a Node as the tree root
func (n *Node) SetRoot() {
	n.depth = 0
	n.depthKnown = true
}

// SetParent sets a node's parent
func (n *Node) SetParent(p *Node) {
	n.Parent = p
}

// Detach will orphan a node
func (n *Node) Detach() {
	n.Parent = nil
}

// MarkPath flags nodes upward to a root or another mark
func (n *Node) MarkPath() {
	if n.onPath {
		n.Detach()
		return
	}
	n.onPath = true
	if n.Parent != nil {
		n.Parent.MarkPath()
	}
}

// ROOT is the label of the root node
const ROOT string = "COM"

func main() {
	var (
		input string
		tree  map[string]*Node
		total int
		me    *Node
		santa *Node
	)
	tree = make(map[string]*Node)
	for {
		n, _ := fmt.Scanln(&input)
		if n == 0 {
			break
		}
		orbit := strings.Split(input, ")")
		locus, lok := tree[orbit[0]]
		if !lok {
			locus = &Node{Name: orbit[0]}
			if orbit[0] == ROOT {
				locus.SetRoot()
			}
			tree[orbit[0]] = locus
		}
		if orbit[1] == "YOU" {
			me = locus
		} else if orbit[1] == "SAN" {
			santa = locus
		} else {
			orbital, ook := tree[orbit[1]]
			if !ook {
				orbital = &Node{Name: orbit[1]}
				tree[orbit[1]] = orbital
			}
			orbital.SetParent(locus)
		}
	}
	fmt.Println("Found", len(tree), "nodes")
	me.MarkPath()
	santa.MarkPath()
	total = me.GetDepth() + santa.GetDepth()
	fmt.Println("Total:", total)
}
