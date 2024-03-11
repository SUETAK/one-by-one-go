package main

func countMaxControl(n int, numbers []int) int {

	count := 0
	isFinished := false
	for isFinished == false {
		for i, number := range numbers {
			if number%2 == 0 {
				numbers[i] = number / 2
			} else {
				isFinished = true
				break
			}
			if i+1 == n {
				count++
			}
		}
	}
	return count
}
