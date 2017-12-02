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
		sum += checksum(row);
	}

	fmt.Println(sum)
}

func checksum(numbers []string) int {
	min, _ := strconv.Atoi(numbers[0]);
	max := min;
	for _, s := range numbers {
		n, _ := strconv.Atoi(s);
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return max - min
}
