/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {

	current := head
	var element *ListNode

	for current.Next != nil {

		if current.Val <= current.Next.Val { //already in order
			current = current.Next //go to the next value
			continue
		} else {
			element = current.Next           //copy the next element
			current.Next = current.Next.Next //remove the element from the list

			if element.Val <= head.Val { //check against the head value
				element.Next = head
				head = element
				continue
			}

			insert := head

			for insert.Next != nil { //check against all other values
				if element.Val < insert.Next.Val {
					element.Next = insert.Next
					insert.Next = element
					break
				} else {
					insert = insert.Next
				}
			}
		}
	}

	return head
}