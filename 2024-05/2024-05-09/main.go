package main

func main() {}

func maxArea(height []int) int {
	mArea := 0

	left := 0
	right := len(height) - 1
	for left != right {
		lh := height[left]
		rh := height[right]

		var area int
		if lh < rh {
			area = (right - left) * lh
			if mArea < area {
				mArea = area
			}
			left++
		} else {
			area = (right - left) * rh
			if mArea < area {
				mArea = area
			}
			right--
		}
	}

	return mArea
}
