package main

import (
	"fmt"
	"encoding/csv"
	"strconv"
	"reflect"
	"os"
)

func main() {
	r := csv.NewReader(os.Stdin)
	r.Comma = '	';
	table, _ := r.ReadAll();

	iterations, distance := findLoop(toNumbers(table[0]))

	fmt.Println(iterations)
	fmt.Println(distance)
}

func toNumbers(row []string) []int {
	numbers := make([]int, len(row))
	for i, s := range row {
		numbers[i], _ = strconv.Atoi(s);
	}
	return numbers
}

func findLoop(numbers []int) (int, int) {
	var seenBefore [][]int
	it := 0

	for {
		oldNumbers := make([]int, len(numbers))
		copy(oldNumbers, numbers)
		seenBefore = append(seenBefore, oldNumbers)

		reorderBanks(numbers)
		it++
		
		for i, r := range seenBefore {
			if reflect.DeepEqual(r, numbers) {
				return it, (len(seenBefore) - i)
			}

		}
	}
}

func reorderBanks(numbers []int) {
	imax := 0;
	max := numbers[0]
	for i, n := range numbers {
		if n > max {
			max = n
			imax = i
		}
	}
	numbers[imax] = 0;

	for ; max > 0; max-- {
		imax = (imax + 1) % len(numbers)
		numbers[imax]++
	}
}
