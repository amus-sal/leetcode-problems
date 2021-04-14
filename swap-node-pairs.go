/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var parents map[*ListNode]*ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	parents = make(map[*ListNode]*ListNode)
	simPar := &ListNode{
		0,
		head,
	}
	parents[head] = simPar
	fill(head)
	swap(head)

	return simPar.Next
}
func swap(node *ListNode) {
	if node == nil {
		return
	}
	if node.Next == nil {
		return
	}
	forw := node.Next
	node.Next = node.Next.Next
	mainPar := parents[node]
	parents[node].Next = forw
	parents[node].Next.Next = node

	parents[node] = forw
	parents[forw] = mainPar
	parents[node.Next] = node
	swap(node.Next)

}

func fill(head *ListNode) {

	if head.Next == nil {
		return
	}
	parents[head.Next] = head
}