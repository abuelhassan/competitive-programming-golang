// https://leetcode.com/problems/linked-list-cycle-ii
// resource: https://cp-algorithms.com/others/tortoise_and_hare.html

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func move(ptr *ListNode, moves int) *ListNode {
	for moves > 0 {
		if ptr == nil {
			return nil
		}
		moves--
		ptr = ptr.Next
	}
	return ptr
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow, fast := head.Next, head.Next.Next
	for fast != nil && slow != fast {
		slow, fast = move(slow, 1), move(fast, 2)
	}
	if fast == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow, fast = move(slow, 1), move(fast, 1)
	}
	return slow
}
