package main

import (
	"container/list"
)

type FirstUnique struct {
	// number frequency counter
	num_map map[int]int
	// queue for preserving order
	queue list.List
}

func Constructor(nums []int) FirstUnique {

	m := new(FirstUnique)
	m.num_map = make(map[int]int)

	for _, v := range nums {
		m.num_map[v]++
		if m.num_map[v] == 1 {
			m.queue.PushBack(v)
		}
	}

	return *m
}

func (this *FirstUnique) ShowFirstUnique() int {

	// for strange reasons queue.Remove doesn't work
	// had to use this method instead
	var y = this.queue.Front()
	for i := 0; i < this.queue.Len(); i++ {
		if this.num_map[y.Value.(int)] == 1 {
			return y.Value.(int)
		}
		y = y.Next()
	}

	return -1
}

func (this *FirstUnique) Add(value int) {

	// if new element add to queue
	if _, found := this.num_map[value]; !found {
		this.queue.PushBack(value)

	}
	this.num_map[value]++
}
