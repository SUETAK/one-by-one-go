package main

func main() {}

// 最小公倍数
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

// 最大公約数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
