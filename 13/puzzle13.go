package main

import "fmt"

func main() {
	var depths, lengths []int
	for {
		var depth, length int
		if _, err := fmt.Scanf("%d: %d\n", &depth, &length); err != nil {
			break
		}
		depths = append(depths, depth)
		lengths = append(lengths, length)
	}

	fmt.Println(getSeverity(depths, lengths))

	wait := 1
	for {
		if !getCaught(depths, lengths, wait) {
			break
		}
		wait++
	}
	fmt.Println(wait)
}

func getSeverity(depths, lengths []int) int {
	score := 0
	for i := range depths {
		cycleLength := lengths[i]*2 - 2
		if cycleLength == 0 {
			cycleLength = 1
		}
		if depths[i]%cycleLength == 0 {
			// get caught!
			score += depths[i] * lengths[i]
		}
	}
	return score
}

func getCaught(depths, lengths []int, wait int) bool {
	for i := range depths {
		if (depths[i]+wait)%(lengths[i]*2-2) == 0 {
			// get caught!
			return true
		}
	}
	return false
}
