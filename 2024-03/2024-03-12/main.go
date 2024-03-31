package main

import "sort"

func main() {}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := append(nums1, nums2...)
	sort.Ints(m)

	if len(m)%2 == 0 {
		i1 := m[(len(m)/2)-1]
		i2 := m[(len(m) / 2)]
		return float64(i1+i2) / 2
	}
	return float64(m[len(m)/2])
}
