/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

var nodes [][]*Node

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	nodes = make([][]*Node, 0)
	nodes = append(nodes, []*Node{})
	appendNodes(0, root)
	for index, oneNode := range nodes {
		if index == 0 {
			continue
		}
		for i := 1; i <= len(oneNode)-1; i++ {
			oneNode[i-1].Next = oneNode[i]
		}
	}
	return root
}

func appendNodes(level int, node *Node) {
	if level >= len(nodes) {
		nodes = append(nodes, []*Node{})
	}
	nodes[level] = append(nodes[level], node)
	if node.Left != nil {
		appendNodes(level+1, node.Left)
	}
	if node.Right != nil {
		appendNodes(level+1, node.Right)
	}

}