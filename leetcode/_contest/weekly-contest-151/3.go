/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	var res, last *ListNode

	node := head
	for node != nil {
		s := node.Val
		next := node.Next

		for next != nil {
			s += next.Val
			next = next.Next
		}
		node = node.Next
	}

	return res
}