package main

import (
	"fmt"
)

type CircList []int

func (l *CircList) get(idx int) int {
	return (*l)[idx%len(*l)]
}
func (l *CircList) set(idx, val int) {
	(*l)[idx%len(*l)] = val
}
func (l *CircList) insertAfter(idx, val int) {
	idx = idx % len(*l)
	if idx == len(*l)-1 {
		*l = append(*l, val)
		return
	}
	*l = append(*l, 0)
	copy((*l)[idx+2:], (*l)[idx+1:])
	(*l)[idx+1] = val
}

func main() {
	var delta int
	fmt.Scanf("%d", &delta)

	buffer := CircList{0}
	pos := 0
	for i := 1; i < 2018; i++ {
		pos = (pos + delta) % len(buffer)
		buffer.insertAfter(pos, i)
		pos++
	}
	fmt.Println(buffer.get(pos + 1))

	lastInsertedAfterZero := 0
	for length := 2018; length <= 50000000; length++ {
		pos = (pos + delta) % length
		if pos == 0 {
			// i will insert (i) after zero, so print it
			lastInsertedAfterZero = length
		}
		pos++
	}
	fmt.Println(lastInsertedAfterZero)
}
