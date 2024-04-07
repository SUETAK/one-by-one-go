package main

import (
	"math/rand"
)

func main() {}

type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {

	return RandomizedSet{m: make(map[int]int), a: make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.a)
	this.a = append(this.a, val)
	return true

}

func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.m[val]; ok {
		lastElement := this.a[len(this.a)-1]
		this.a[i], this.m[lastElement] = lastElement, i
		this.a = this.a[:len(this.a)-1]
		delete(this.m, val)
		return true
	}
	return false

}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.a))
	return this.a[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
