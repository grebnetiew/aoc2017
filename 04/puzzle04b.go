package main

import (
	"fmt"
	"encoding/csv"
	"sort"
	"os"
	"strings"
)

func main() {
	r := csv.NewReader(os.Stdin)
	r.Comma = ' ';
	r.FieldsPerRecord = -1; // variable
	table, _ := r.ReadAll();

	sum := 0
	for _, row := range table {
		if (!hasAnagrams(row)) {
			sum++;
		}
	}

	fmt.Println(sum)
}

func hasAnagrams(r []string) bool {
	for i, s := range r {
		for j := 0; j < i; j++ {
			if anagram(s, r[j]) {
				return true
			}
		}
	}
	return false
}

func anagram(a, b string) bool {
	sA := strings.Split(a, "");
	sort.Strings(sA);
	sB := strings.Split(b, "");
	sort.Strings(sB);
	a = strings.Join(sA, "");
	b = strings.Join(sB, "");
	return a == b
}