/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if (head == nil) || (k < 1) {
		return head
	}

	size := 1
	node := head
	for node.Next != nil {
		size++
		node = node.Next
	}
	a
	k = k % size
	if k == 0 {
		return head
	}

	node.Next = head
	limit := size - k
	for idx := 0; idx < limit; idx++ {
		node = node.Next
	}

	res := node.Next
	node.Next = nil
	return res
}