package main

import (
	"fmt"
	"encoding/csv"
	"os"
)

func main() {
	r := csv.NewReader(os.Stdin)
	r.Comma = ' ';
	r.FieldsPerRecord = -1; // variable
	table, _ := r.ReadAll();

	sum := 0
	for _, row := range table {
		if (!hasDuplicates(row)) {
			sum++;
		}
	}

	fmt.Println(sum)
}

func hasDuplicates(r []string) bool {
	for i, s := range r {
		for j := 0; j < i; j++ {
			if s == r[j] {
				return true
			}
		}
	}
	return false
}