package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	r := csv.NewReader(os.Stdin)
	r.Comma = '	'
	table, _ := r.ReadAll()

	iterations, distance := findLoop(toNumbers(table[0]))

	fmt.Println(iterations)
	fmt.Println(distance)
}

func toNumbers(row []string) []int {
	numbers := make([]int, len(row))
	for i, s := range row {
		numbers[i], _ = strconv.Atoi(s)
	}
	return numbers
}

func findLoop(numbers []int) (int, int) {
	var seenBefore [][]int

	for {
		oldNumbers := make([]int, len(numbers))
		copy(oldNumbers, numbers)
		seenBefore = append(seenBefore, oldNumbers)

		reorderBanks(numbers)

		for i, r := range seenBefore {
			if reflect.DeepEqual(r, numbers) {
				return len(seenBefore), (len(seenBefore) - i)
			}

		}
	}
}

func reorderBanks(numbers []int) {
	var imax, max int = 0, numbers[0]
	for i, n := range numbers {
		if n > max {
			imax, max = i, n
		}
	}
	numbers[imax] = 0

	for ; max > 0; max-- {
		imax = (imax + 1) % len(numbers)
		numbers[imax]++
	}
}
