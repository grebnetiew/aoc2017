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

	strength, _ := FindBest(&parts, stronger)
	fmt.Println(strength)

	strength, length := FindBest(&parts, longer)
	fmt.Printf("%d (%d)\n", strength, length)
}

func FindBest(parts *[]Part, cmp ComparisonFunction) (bestweight, bestlength int) {
	return findBest(parts, cmp, 0, make([]bool, len(*parts)))
}

func findBest(parts *[]Part, cmp ComparisonFunction, start int, visited []bool) (bestweight, bestlength int) {
	for i, p := range *parts {
		var newstart int
		switch {
		case visited[i]:
			continue // Skip this part if already used
		case p[0] == start:
			newstart = p[1]
		case p[1] == start:
			newstart = p[0]
		default:
			continue // Skip this part if it doesn't fit
		}
		// Add the candidate part to the list of visited parts
		visited[i] = true
		// Find the best results if starting from this candidate
		w, l := findBest(parts, cmp, newstart, visited)
		// If better, keep it as the new best options
		if cmp(l+1, w+p[0]+p[1], bestlength, bestweight) {
			bestweight = w + p[0] + p[1]
			bestlength = l + 1
		}
		// "unreserve" the candidate node
		visited[i] = false
	}
	return bestweight, bestlength
}

func stronger(length, weight, bestlength, bestweight int) bool {
	return weight > bestweight
}
func longer(length, weight, bestlength, bestweight int) bool {
	return length > bestlength || (length == bestlength && weight > bestweight)
}
