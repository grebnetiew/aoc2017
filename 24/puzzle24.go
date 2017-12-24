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

	strongest, _, _ := findBest(parts, 0, make([]bool, len(parts)), stronger)
	fmt.Println(strongest)

	strongest, length, _ := findBest(parts, 0, make([]bool, len(parts)), longer)
	fmt.Printf("%d (%d)\n", strongest, length)
}

type fBetter func(length, weight, bestlength, bestweight int) bool

func findBest(parts []Part, start int, visited []bool, compare fBetter) (bestweight int, bestlength int, bestchain []Part) {
	for i, p := range parts {
		// Skip this part if it's used or doesn't fit
		if visited[i] || (p[0] != start && p[1] != start) {
			continue
		}
		// Determine the 'open' end of this candidate part
		newstart := p[1]
		if p[1] == start {
			newstart = p[0]
		}
		// Add the candidate part to the list of visited parts
		visited[i] = true
		// Find the best results if starting from this candidate
		w, l, chain := findBest(parts, newstart, visited, compare)
		// If better, keep it as the new best options
		if compare(l+1, w+p[0]+p[1], bestlength, bestweight) {
			bestweight = w + p[0] + p[1]
			bestlength = l + 1
			bestchain = append([]Part{p}, chain...)
		}
		// "unreserve" the candidate node
		visited[i] = false
	}
	return bestweight, bestlength, bestchain
}

func stronger(length, weight, bestlength, bestweight int) bool {
	return weight > bestweight
}
func longer(length, weight, bestlength, bestweight int) bool {
	return length > bestlength || (length == bestlength && weight > bestweight)
}
