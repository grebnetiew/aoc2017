package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var maze []string
	for line, err := r.ReadString('\n'); err == nil; line, err = r.ReadString('\n') {
		maze = append(maze, line)
	}

	var x, y, dir int
	for maze[0][x] != '|' {
		x++
	}
	dir = 3 // down

	keepgoing, letters, steps := true, "", 0
	for keepgoing {
		var letter byte
		keepgoing, letter = step(maze, &x, &y, &dir)
		if letter != 0 {
			letters += string(letter)
		}
		steps++
	}

	fmt.Println(letters)
	fmt.Println(steps)
}

func step(maze []string, px, py, pdir *int) (keepgoing bool, letter byte) {
	x, y, dir := *px, *py, *pdir

	switch dir {
	case 0:
		x, y = x+1, y
	case 1:
		x, y = x, y-1
	case 2:
		x, y = x-1, y
	case 3:
		x, y = x, y+1
	}

	if maze[y][x] == '+' {
		switch {
		case x != 0 && maze[y][x-1] != ' ' && dir != 0:
			dir = 2
		case x != len(maze[0])-1 && maze[y][x+1] != ' ' && dir != 2:
			dir = 0
		case y != 0 && maze[y-1][x] != ' ' && dir != 3:
			dir = 1
		case y != len(maze)-1 && maze[y+1][x] != ' ' && dir != 1:
			dir = 3
		}
	}

	if maze[y][x] != ' ' {
		keepgoing = true
	}

	if maze[y][x] >= 'A' && maze[y][x] <= 'Z' {
		letter = maze[y][x]
	}
	*px, *py, *pdir = x, y, dir
	return keepgoing, letter
}
