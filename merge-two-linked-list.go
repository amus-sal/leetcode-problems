/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		return l2
	}
	if l2 != nil && l1.Val >= l2.Val {
		check(l2, l1)
		return l2
	} else {
		check(l1, l2)
		return l1
	}
}

func check(node1 *ListNode, node2 *ListNode) {
	if node2 == nil {
		return
	}
	if node1 == nil {
		return
	}
	rem := node2

	if node2.Val >= node1.Val && (node1.Next == nil || node2.Val <= node1.Next.Val) {
		rem = node2.Next
		node2.Next = node1.Next
		node1.Next = node2
		check(node1.Next, rem)
	} else {
		check(node1.Next, node2)
	}

}