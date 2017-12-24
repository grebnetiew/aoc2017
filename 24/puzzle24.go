package main

import (
	"fmt"
	"sort"
)

type Part [2]int

func main() {
	var parts []Part
	for {
		var p, q int
		_, err := fmt.Scanf("%d/%d\n", &p, &q)
		if err != nil {
			break
		}
		if p < q {
			parts = append(parts, [2]int{p, q})
		} else {
			parts = append(parts, [2]int{q, p})
		}
	}
	sort.Slice(parts, func(i, j int) bool { return parts[i][0] < parts[j][0] })

	fmt.Println(strongest(parts, 0, make([]bool, len(parts))))
	fmt.Println(longest(parts, 0, make([]bool, len(parts))))
}

func strongest(parts []Part, start int, visited []bool) (int, []Part) {
	best := 0
	var bestchain []Part
	for i, p := range parts {
		if !visited[i] && (p[0] == start || p[1] == start) {
			newstart := p[1]
			if p[1] == start {
				newstart = p[0]
			}
			newvisited := append([]bool(nil), visited...)
			newvisited[i] = true
			w, chain := strongest(parts, newstart, newvisited)
			if w+p[0]+p[1] > best {
				best = w + p[0] + p[1]
				bestchain = append([]Part{p}, chain...)
			}
		}
	}
	return best, bestchain
}

func longest(parts []Part, start int, visited []bool) (weight int, length int, _ []Part) {
	best := 0
	maxlength := 0
	var bestchain []Part
	for i, p := range parts {
		if !visited[i] && (p[0] == start || p[1] == start) {
			newstart := p[1]
			if p[1] == start {
				newstart = p[0]
			}
			newvisited := append([]bool(nil), visited...)
			newvisited[i] = true
			w, l, chain := longest(parts, newstart, newvisited)
			if l > maxlength || (l == maxlength && w+p[0]+p[1] > best) {
				best = w + p[0] + p[1]
				maxlength = l
				bestchain = append([]Part{p}, chain...)
			}
		}
	}
	return best, maxlength + 1, bestchain
}

func print(p []Part) {
	for _, v := range p {
		fmt.Printf("%d/%d--", v[0], v[1])
	}
	fmt.Println()
}
