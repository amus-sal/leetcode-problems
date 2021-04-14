/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

var oldNodes map[*Node]*Node

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	oldNodes = make(map[*Node]*Node)
	newNode := &Node{
		node.Val,
		[]*Node{},
	}
	oldNodes[node] = newNode
	build(newNode, node)
	fmt.Println(newNode)
	return newNode
}
func build(node *Node, oldOne *Node) {
	fmt.Println(node)
	if node == nil {
		return
	}
	if oldOne == nil {
		return
	}
	for _, nd := range oldOne.Neighbors {
		if val, found := oldNodes[nd]; found {
			node.Neighbors = append(node.Neighbors, val)
		} else {
			newNex := &Node{
				nd.Val,
				[]*Node{},
			}
			oldNodes[nd] = newNex
			node.Neighbors = append(node.Neighbors, newNex)
			build(newNex, nd)
		}
	}
}