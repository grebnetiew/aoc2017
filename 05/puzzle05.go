package main

import (
	"fmt"
)

func main() {
	var numbers1 []int
	var numbers2 []int
	var v int
	n, _ := fmt.Scanf("%d\n", &v)
	for n != 0 {
		numbers1 = append(numbers1, v)
		numbers2 = append(numbers2, v)
		n, _ = fmt.Scanf("%d\n", &v)
	}
	fmt.Println(mazeOne(numbers1))
	fmt.Println(mazeTwo(numbers2))
}

func mazeOne(numbers []int) int {
	it := 0
	cur := 0
	for cur >= 0 && cur < len(numbers) {
		numbers[cur]++
		cur += numbers[cur] - 1
		it++
	}
	return it
}

func mazeTwo(numbers []int) int {
	it := 0
	cur := 0
	for cur >= 0 && cur < len(numbers) {
		old := numbers[cur]
		if numbers[cur] < 3 {
			numbers[cur]++
		} else {
			numbers[cur]--
		}
		cur += old
		it++
	}
	return it
}


