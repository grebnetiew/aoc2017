package main

import (
	"fmt"
	"encoding/csv"
	"strconv"
	"os"
)

func main() {
	r := csv.NewReader(os.Stdin)
	r.Comma = '	';
	table, _ := r.ReadAll();

	sum := 0
	for _, row := range table {
		sum += checksum(toNumbers(row));
	}

	fmt.Println(sum)
}

func toNumbers(row []string) []int {
	numbers := make([]int, len(row))
	for i, s := range row {
		numbers[i], _ = strconv.Atoi(s);
	}
	return numbers
}

func checksum(numbers []int) int {
	for i, n := range numbers {
		for _, m := range numbers[(i+1):] {
			if n % m == 0 {
				return n/m
			}
			if m % n == 0 {
				return m/n
			}
		}
	}
	return 0
}
