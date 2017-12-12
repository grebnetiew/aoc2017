package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Vertices []Vertex
type Vertex struct {
	Neighbours []int
}
type Set map[int]bool

func main() {
	// Technically, "abc <-> def, ghi, jkl" are comma separated values. Abuse this
	reader := csv.NewReader(os.Stdin)
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1 // variable # of fields

	// Put everything in our list
	var programList Vertices

	row, err := reader.Read()
	for err == nil {
		var p int
		// We replace the "abc <-> def" part by just "def"
		// so row is exactly a slice of the children
		fmt.Sscanf(row[0], "%d <-> %s", &p, &row[0])
		children := make([]int, len(row))
		for i, v := range row {
			children[i], _ = strconv.Atoi(v)
		}

		programList = append(programList, Vertex{Neighbours: children})

		row, err = reader.Read()
	}

	// Find the size of the first component
	visited := make(Set)
	fmt.Println(connected(visited, programList, 0))

	// Find the other components
	components := 1
	for i := range programList {
		if visited[i] {
			continue
		}
		connected(visited, programList, i)
		components++
	}
	fmt.Println(components)
}

// connected counts the new nodes connected to the 'name' one.
// it uses the Set to track which ones were counted already.
func connected(visited Set, pl Vertices, name int) int {
	if visited[name] {
		return 0
	}
	visited[name] = true

	total := 1
	for _, prog := range pl[name].Neighbours {
		total += connected(visited, pl, prog)
	}

	return total
}
