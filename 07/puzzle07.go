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
	reader := csv.NewReader(os.Stdin)
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1 // variable # of fields

	programList := make(ProgramList)
	var p string
	row, err := reader.Read()
	for err == nil {
		var w int
		if len(row) == 1 {
			fmt.Sscanf(row[0], "%s (%d)", &p, &w)
		} else {
			fmt.Sscanf(row[0], "%s (%d) -> %s", &p, &w, &row[0])
		}
		programList[p] = Program{Weight: w, Tower: row}

		row, err = reader.Read()
	}

	// Find the root of the tree. This is extremely inefficient :(
	root := p
search:
	for {
		for key, val := range programList {
			for _, name := range val.Tower {
				if root == name {
					root = key
					continue search
				}
			}
		}
		break search
	}

	fmt.Println(root)

	// Find the unbalanced disk
	balance(programList, root)
}

// returns total weight
func balance(pl ProgramList, name string) int {
	p := pl[name]
	if p.Tower == nil {
		return p.Weight
	}
	w := make([]int, len(p.Tower))
	different := 0
	for i, sp := range p.Tower {
		w[i] = balance(pl, sp)
		if w[i] != w[0] {
			different = i
		}
	}
	if different != 0 {
		var goodWeight = w[0]
		// they might've been all different - check for this case
		if different == len(w)-1 && w[0] != w[1] {
			// it was the first one
			different = 0
			goodWeight = w[1]
		}

		difference := w[different] - goodWeight
		fmt.Println(-difference + pl[p.Tower[different]].Weight)
	}

	return w[0]*len(w) + p.Weight
}
