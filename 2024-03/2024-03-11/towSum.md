# やること
leetCode の問題の復習と解説の作成


# [Two sum の解説](https://leetcode.com/problems/two-sum/)
問題文（要約）
1つのint型配列と1つのint が与えられる
配列の中で、target になる数字の組み合わせを探し、そのindex を配列にして返却する

条件
- 1つの解しか存在しない
- 同じ数字は配列内に登場しない

## 解法１
解は一つしか存在しないので、前から順にtarget かどうかを確認する
オーダーは最大O(n**2) かな…？

nums = [2,7,11,15], target = 9 であれば
- nums をfor 文で回す
- for 文内部でnums のループの中で、次の値を始まりとしたfor文を作成する

```go
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	var answer []int
	label:
	for i := 0; i < len(nums); i++ {
		currentNum := nums[i]
		for j := i + 1; i < len(nums); j++ {
			if currentNum+nums[j] == target {
				answer = []int{i, j}
				// 解は1つしかないので見つかった時点でbreak。２重ループはbreak では抜けられないのでlabel を使う
				break label
			}
		}
	}
	fmt.Println(answer)
}

```

## 解法２
解になる2つの数字はnums の中に確実に存在することを利用して map に値とindex を保存する
map に解になる数字が存在しなければmap に保存し、あればそれが回答になる

```go
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	// int のマップを作成
	numMap := make(map[int]int)

	for i, v := range nums {
		// 9 - 2 = 7 のように答えになる数字を確認する
		comp := target - v

		// numMap[comp] を使って、もし値が存在するなら配列を返す numMap[7]
		// 解は1つしかない→num の中に絶対に2つ選べる数字がある
		// [2, 7, 11, 15] であれば、 9 - 2 = 7 なので numMap[2] = 1
		// 9 - 7 = 2 なので numMap[2] を探せば確実に値を取得できる
		// オーダーはO(n)
		if _, exist := numMap[comp]; exist {
			result := []int{numMap[comp], i}
			fmt.Println(result)
		}
		// 配列が存在しないなら、map に格納する
		numMap[v] = i
	}
}
```


