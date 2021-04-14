/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return getNextNode(l1, l2, 0)

}

func calCol(val1 int, val2 int, carryOut int) (int, int) {
	num := val1 + val2 + carryOut
	if num < 10 {
		return num, 0
	}
	if num == 10 {
		return 0, 1
	}
	return num % 10, num / 10

}

func getNextNode(l1 *ListNode, l2 *ListNode, carryout int) *ListNode {

	if l1 == nil && l2 == nil && carryout == 0 {
		return nil
	}
	if l1 == nil && l2 == nil && carryout != 0 {
		return &ListNode{
			carryout,
			nil,
		}
	}
	var NextVal1 *ListNode
	var NextVal2 *ListNode
	val1 := 0
	val2 := 0
	if l1 == nil {
		val1 = 0
		NextVal1 = nil
	} else {
		val1 = l1.Val
		NextVal1 = l1.Next
	}
	if l2 == nil {
		val2 = 0
		NextVal2 = nil
	} else {
		val2 = l2.Val
		NextVal2 = l2.Next
	}
	val, carryOut := calCol(val1, val2, carryout)

	return &ListNode{
		val,
		getNextNode(NextVal1, NextVal2, carryOut),
	}
}