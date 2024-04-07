package main

import (
	"math/rand"
)

func main() {}

type RandomizedSet struct {
	valueMap  map[int]int
	valueKeys []int
}

func Constructor() RandomizedSet {

	return RandomizedSet{valueMap: make(map[int]int), valueKeys: make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valueMap[val]; ok {
		return false
	} else {
		this.valueMap[val] = len(this.valueKeys)
		newValueKeys := make([]int, 0, len(this.valueMap))
		for i, _ := range this.valueMap {
			newValueKeys = append(newValueKeys, i)
		}
		this.valueKeys = newValueKeys
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.valueMap[val]; ok {
		delete(this.valueMap, val)
		newValueKeys := make([]int, 0, len(this.valueMap))
		for i, _ := range this.valueMap {
			newValueKeys = append(newValueKeys, i)
		}
		this.valueKeys = newValueKeys
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.valueKeys))
	return this.valueKeys[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
