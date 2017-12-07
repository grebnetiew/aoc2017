package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type ProgramList map[string]Program
type Program struct {
	Weight int
	Tower []string
}

func main() {
	var err error = nil
	programList := make(ProgramList)
	reader := bufio.NewReader(os.Stdin)
	var p string
	for err == nil {
		var w int
		var sp []string
		var text string
		text, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Sscanf(text, "%s (%d)", &p, &w)
		if n := strings.Index(text, "->"); n != -1 {
			var sps string
			sps = text[(n + 3):(len(text) - 2)]
			sp = strings.Split(sps, ", ")
		}
		programList[p] = Program{Weight: w, Tower: sp}
	}
	
	root := p
	// Find the root of the tree. This is extremely inefficient :(
	search: for {
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
		if different == len(w) - 1 && w[0] != w[1] {
			// it was the first one
			different = 0
			goodWeight = w[1]
		}
		
		difference := w[different] - goodWeight
		fmt.Println(-difference + pl[p.Tower[different]].Weight)
	}

	return w[0] * len(w) + p.Weight
}
