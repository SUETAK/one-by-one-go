package main

func twoSum(nums []int, target int) []int {
	var answer []int
	//　2つの数字から、target になる数字を探す
	for i := 0; i < len(nums); i++ {
		// 同じ数字は2つ存在しない
		// 回答は1つ
		// 1つ目のvalue を2、3、4... と回していって、target と同じになったら答え
		// 2つめ以降のindex との組み合わせを見れば良い
		currentNum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if currentNum+nums[j] == target {
				answer = []int{i, j}
			}
		}
	}
	return answer
}

// TODO 考え方は明日書く
func twoSumCorrect(nums []int, target int) []int {
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if j, ok := hash[nums[i]]; ok {
			return []int{i, j}
		}

		hash[target-nums[i]] = i
	}

	return nil
}

func twoSumByMap(nums []int, target int) []int {
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
			return []int{numMap[comp], i}
		}
		// 配列が存在しないなら、map に格納する
		numMap[v] = i

	}

	return []int{}

}
