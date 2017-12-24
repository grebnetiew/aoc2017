package main

import "fmt"

type Part [2]int
type ComparisonFunction func(length, weight, bestlength, bestweight int) bool

func main() {
	var parts []Part
	for {
		var p, q int
		_, err := fmt.Scanf("%d/%d\n", &p, &q)
		if err != nil {
			break
		}
		parts = append(parts, [2]int{p, q})
	}

	strength, _ := findBest(&parts, 0, nil, stronger)
	fmt.Println(strength)

	strength, length := findBest(&parts, 0, nil, longer)
	fmt.Printf("%d (%d)\n", strength, length)
}

func findBest(parts *[]Part, start int, visited *[]bool, cmp ComparisonFunction) (bestweight, bestlength int) {
	// Make a list of visited nodes, if this is a top level invocation
	if visited == nil {
		v := make([]bool, len(*parts))
		visited = &v
	}

	// Try to connect each part and recurse
	for i, p := range *parts {
		// Skip this part if it's used or doesn't fit
		if (*visited)[i] || (p[0] != start && p[1] != start) {
			continue
		}
		// Determine the 'open' end of this candidate part
		newstart := p[1]
		if p[1] == start {
			newstart = p[0]
		}
		// Add the candidate part to the list of visited parts
		(*visited)[i] = true
		// Find the best results if starting from this candidate
		w, l := findBest(parts, newstart, visited, cmp)
		// If better, keep it as the new best options
		if cmp(l+1, w+p[0]+p[1], bestlength, bestweight) {
			bestweight = w + p[0] + p[1]
			bestlength = l + 1
		}
		// "unreserve" the candidate node
		(*visited)[i] = false
	}
	return bestweight, bestlength
}

func stronger(length, weight, bestlength, bestweight int) bool {
	return weight > bestweight
}
func longer(length, weight, bestlength, bestweight int) bool {
	return length > bestlength || (length == bestlength && weight > bestweight)
}
