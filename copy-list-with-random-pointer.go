/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

var nodesMap map[*Node]*Node
var nodesMapNew map[*Node]*Node

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodesMap = make(map[*Node]*Node)
	nodesMapNew = make(map[*Node]*Node)
	cloneRandomNodes(head)
	var cpl Node
	cpl.Val = head.Val
	nodesMapNew[head] = &cpl
	cpl.Next = copyNodeN(head.Next)
	for k, v := range nodesMapNew {
		v.Random = nodesMapNew[nodesMap[k]]
	}
	return &cpl
}

func copyNodeN(node *Node) *Node {
	if node == nil {
		return nil
	}
	res := new(Node)
	res.Val = node.Val
	res.Next = copyNodeN(node.Next)
	nodesMapNew[node] = res
	return res
}

func cloneRandomNodes(head *Node) {
	p := head
	for p != nil {
		nodesMap[p] = p.Random
		p = p.Next
	}
}
 