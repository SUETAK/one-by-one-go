package main

func main() {}

func candy(ratings []int) int {

	candyCount := 0
	// 両側がある場所だけみる
	addCount := 0
	for i := 0; i < len(ratings); i++ {
		if i == 0 {
			if ratings[0] > ratings[1] {
				addCount++
			}
			candyCount += addCount
			continue
		}
		if i == len(ratings)-1 {
			if ratings[len(ratings)-1] == ratings[len(ratings)-2] {
				continue
			} else if ratings[len(ratings)-1] > ratings[len(ratings)-2] {
				addCount++
			} else if ratings[len(ratings)-1] < ratings[len(ratings)-2] {
				if addCount > 0 {
					addCount--
				}
			}
			candyCount += addCount
			continue
		}
		if ratings[i] == ratings[i-1] && ratings[i] == ratings[i+1] {
			continue
		}
		if ratings[i] > ratings[i-1] || ratings[i] > ratings[i+1] {
			addCount++
		}

		if ratings[i] < ratings[i-1] || ratings[i] < ratings[i+1] {
			if addCount > 0 {
				addCount--
			}
		}
		candyCount += addCount
	}
	return candyCount + len(ratings)
}
