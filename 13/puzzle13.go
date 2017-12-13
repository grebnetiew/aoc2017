package main

import "fmt"

func main() {
	firewall := make(map[int]int)
	for {
		var depth, length int
		if _, err := fmt.Scanf("%d: %d\n", &depth, &length); err != nil {
			break
		}
		firewall[depth] = length
	}

	fmt.Println(getSeverity(firewall, 0))

	wait := 1
	for {
		if _, c := getSeverity(firewall, wait); !c {
			break
		}
		wait++
	}
	fmt.Println(wait)
}

func getSeverity(fw map[int]int, wait int) (int, bool) {
	score := 0
	caught := false
	for depth, length := range fw {
		cycleLength := length*2 - 2
		if cycleLength == 0 {
			cycleLength = 1
		}
		if (depth+wait)%cycleLength == 0 {
			// get caught!
			score += depth * length
			caught = true
		}
	}
	return score, caught
}
