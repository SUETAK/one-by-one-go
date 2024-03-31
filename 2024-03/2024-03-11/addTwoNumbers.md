# やること
leetCode の問題の復習と解説の作成


# [Add Two Numbers の解説](https://leetcode.com/problems/add-two-numbers/)
問題文（要約）
2つのint型 linkedList が渡される。それぞれのlinkedList の値を合算して、linkedList で返却する


# 解法1
linkedList は下記のような構造体になっている。1つのオブジェクトにValue と次の値のオブジェクトを持っている
```go
package main
type ListNode struct {
	Val  int
	Next *ListNode
}
```
例えばl1 = [2,4,3] である場合は下記のように値が入れ子になっているような構造になる
```go
&ListNode{
    Val: 2,
    Next: &ListNode{
        Val: 4,
        Next: &ListNode{
            Val:  3,
            Next: nil,
        },
    },
}
```
最後の要素はNext がnil になる点も特徴の一つであると言える。

これを踏まえて、2つのlinkedList のループを回し、筆算のように計算する。
１桁ずつ計算し、繰り上がりがあれば保存。次の桁に移動する時に１を足す。

繰り上がりの場合には１桁目は答えの桁として保存し、NextオブジェクトのVal に1を足しておくことで表現する

```go
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
			// 1桁目を答えのVal に保存
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

```
