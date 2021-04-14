/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var nodesMap map[*ListNode]*ListNode
var nodes []*ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodesMap = make(map[*ListNode]*ListNode)
	nodes = make([]*ListNode, 0)
	addParent(head, nil)
	if n == 1 && len(nodes) == 1 {
		return nil
	}
	if n == 1 {
		trgNode := nodes[len(nodes)-1]
		nodesMap[trgNode].Next = nil
		return head

	}
	trgNode := nodes[len(nodes)-n]
	if nodesMap[trgNode] == nil {
		return trgNode.Next
	}
	nodesMap[trgNode].Next = nodes[len(nodes)-n+1]
	return head

}
func addParent(node *ListNode, parent *ListNode) {
	if node == nil {
		return
	}
	nodesMap[node] = parent
	nodes = append(nodes, node)
	addParent(node.Next, node)
}
 
 