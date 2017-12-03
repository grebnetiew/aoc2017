package main

import (
	"fmt"
)

type Table struct {
	grid [100][100]int
	x, y int
}
func NewTable() *Table {
	t := new(Table)
	t.x = 1
	t.y = 0
	t.grid[50][50] = 1
	return t
}
func (t *Table) Step() (int) {
	// fill the current spot
	result := t.grid[t.x + 51][t.y + 49] + t.grid[t.x + 51][t.y + 50] + t.grid[t.x + 51][t.y + 51] +
	          t.grid[t.x + 50][t.y + 49]                              + t.grid[t.x + 50][t.y + 51] +
	          t.grid[t.x + 49][t.y + 49] + t.grid[t.x + 49][t.y + 50] + t.grid[t.x + 49][t.y + 51]
	t.grid[t.x + 50][t.y + 50] = result

	// move the cursor
	switch {
	case t.x > 0 && t.x > abs(t.y):
		t.y -= 1 // move up
	case t.y > 0 && t.y >= abs(t.x):
		t.x += 1 // move right
	case t.x < 0 && t.x <= t.y:
		t.y += 1 // move down
	default:
		t.x -= 1 // move left
	}

	return result

}

func main() {
	var square int
	fmt.Scanf("%d", &square)

	
	t := NewTable()
	r := 1
	for r <= square {
		r = t.Step()
	}
	fmt.Println(r)
}


func abs(n int) int { // why does this only exist on floats
	if n > 0 {
		return n
	}
	return -n
}