/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/list"

func middleNode(head *ListNode) *ListNode {

	var length int
	// initialize a queue
	queue := list.New()

	for head != nil {

		queue.PushBack(head)
		length++

		// pop for every two elements
		if (length % 2) == 0 {
			queue.Remove(queue.Front())
		}

		head = head.Next
	}

	return (queue.Front()).Value.(*ListNode)

}
