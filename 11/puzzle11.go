package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	reader := csv.NewReader(os.Stdin)
	rows, _ := reader.ReadAll()

	x, y, record := 0, 0, 0
	for _, v := range rows[0] {
		switch v {
		case "n":
			x -= 1
			y -= 1
		case "ne":
			y -= 1
		case "nw":
			x -= 1
		case "s":
			x += 1
			y += 1
		case "se":
			x += 1
		case "sw":
			y += 1
		}
		record = max(record, distance(x, y))
	}

	fmt.Println(distance(x, y))
	fmt.Println(record)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func distance(x, y int) int {
	if x*y <= 0 {
		return abs(x) + abs(y)
	} else {
		return max(abs(x), abs(y))
	}
}
