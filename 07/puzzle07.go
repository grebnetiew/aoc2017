package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type ProgramList map[string]Program
type Program struct {
	Weight int
	Tower  []string
}

func main() {
	// Technically, "abc (12) -> def, ghi, jkl" are comma separated values. Abuse this
	reader := csv.NewReader(os.Stdin)
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1 // variable # of fields

	// Put everything in our hashmap
	programList := make(ProgramList)
	var p string // remember it so we can find a root later

	row, err := reader.Read()
	for err == nil {
		var w int
		if len(row) == 1 { // a leaf node
			fmt.Sscanf(row[0], "%s (%d)", &p, &w)
		} else {
			// We replace the "abc (12) -> def" part by just "def"
			// so row is exactly a slice of the children
			fmt.Sscanf(row[0], "%s (%d) -> %s", &p, &w, &row[0])
		}
		programList[p] = Program{Weight: w, Tower: row}

		row, err = reader.Read()
	}

	root := findRoot(programList)
	fmt.Println(root)

	// Find the unbalanced disk
	balance(programList, root)
}

func findRoot(pl ProgramList) string {
	// Now that the map is loaded, make a "reverse" tree
	parents := make(map[string]string)
	var key string
	var val Program
	for key, val = range pl {
		for _, name := range val.Tower {
			parents[name] = key
		}
	}

	// Find the root of the tree
	for ; parents[key] != ""; key = parents[key] {
	}
	return key
}

// This function returns total weight of this node and any subtree
// and prints the weight difference if a set of children is not balanced
func balance(pl ProgramList, name string) int {
	p := pl[name]
	if p.Tower == nil { // no subtree
		return p.Weight
	}

	// Keep a list of the subtree weights, and try to keep track of the odd one out
	w := make([]int, len(p.Tower))
	different := 0
	subWeight := 0
	for i, sp := range p.Tower {
		w[i] = balance(pl, sp)
		if w[i] != w[0] {
			different = i
		}
		subWeight += w[i]
	}

	difference := 0
	// If the odd one out is still 0, they are all the same
	if different != 0 {
		// If not, make sure we found the odd one out
		// (if it's the 0th one, all the others will be "different"..)
		// check for this. Also keep track of the others' weight
		var goodWeight = w[0]
		if different == len(w)-1 && w[0] != w[1] {
			// it was the first one
			different = 0
			goodWeight = w[1]
		}

		// Find the difference, and subtract it from the offending node's
		// personal weight to see what it should weigh
		difference = w[different] - goodWeight
		fmt.Println(-difference + pl[p.Tower[different]].Weight)
	}

	return subWeight + p.Weight - difference
}
