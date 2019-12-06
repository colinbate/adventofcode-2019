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

// ROOT is the label of the root node
const ROOT string = "COM"

func main() {
	var (
		input string
		tree  map[string]*Node
		total int
	)
	tree = make(map[string]*Node)
	for {
		n, _ := fmt.Scanln(&input)
		if n == 0 {
			break
		}
		// fmt.Println("Scanned:", input)
		orbit := strings.Split(input, ")")
		locus, lok := tree[orbit[0]]
		if !lok {
			locus = &Node{Name: orbit[0]}
			if orbit[0] == ROOT {
				locus.SetRoot()
			}
			tree[orbit[0]] = locus
		}
		orbital, ook := tree[orbit[1]]
		if !ook {
			orbital = &Node{Name: orbit[1]}
			tree[orbit[1]] = orbital
		}
		orbital.SetParent(locus)
	}
	fmt.Println("Found", len(tree), "nodes")
	for _, info := range tree {
		nd := info.GetDepth()
		// fmt.Println("Depth of", name, "is", nd)
		total += nd
	}
	fmt.Println("Total:", total)
}
