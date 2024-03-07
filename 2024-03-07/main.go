package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// {1, X}, {2, Y} -> {3, Z}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result

	// linked list の頭からお尻までやってく。お尻に辿り着くとNext はnil になる
	// 両方のlinked list がnil になるまでfor を回す
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}

		// 繰り上がり処理
		if tmp.Val >= 10 {
			tmp.Val %= 10
			// 繰り上がっていることを保存
			tmp.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
			tmp = tmp.Next
			continue

		}

		if l1 != nil || l2 != nil {
			tmp.Next = &ListNode{}
		}

		tmp = tmp.Next
	}
	return result
}
