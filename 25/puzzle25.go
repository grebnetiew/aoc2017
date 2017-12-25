package main

import (
	"bufio"
	"fmt"
	"os"
)

type State struct {
	Oval, Odir, Ogoto, Ival, Idir, Igoto int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var line string
	r.ReadString('\n')

	checkAfter := 0
	line, _ = r.ReadString('\n')
	fmt.Sscanf(line, "Perform a diagnostic checksum after %d steps.\n", &checkAfter)

	states := make([]State, 0)
	for {
		r.ReadString('\n')
		var s string
		var oval, odir, ogoto, ival, idir, igoto int
		line, _ = r.ReadString('\n')
		_, err := fmt.Sscanf(line, "In state %s", &s)
		if err != nil {
			break
		}
		r.ReadString('\n')
		line, _ = r.ReadString('\n')
		fmt.Sscanf(line, "    - Write the value %d", &oval)
		line, _ = r.ReadString('\n')
		fmt.Sscanf(line, "    - Move one slot to the %s", &s)
		if s[0:4] == "left" {
			odir = -1
		} else {
			odir = 1
		}
		line, _ = r.ReadString('\n')
		fmt.Sscanf(line, "    - Continue with state %s", &s)
		ogoto = int(s[0] - 'A')
		r.ReadString('\n')
		line, _ = r.ReadString('\n')
		fmt.Sscanf(line, "    - Write the value %d", &ival)
		line, _ = r.ReadString('\n')
		fmt.Sscanf(line, "    - Move one slot to the %s", &s)
		if s[0:4] == "left" {
			idir = -1
		} else {
			idir = 1
		}
		line, _ = r.ReadString('\n')
		fmt.Sscanf(line, "    - Continue with state %s", &s)
		igoto = int(s[0] - 'A')

		states = append(states, State{oval, odir, ogoto, ival, idir, igoto})
	}

	tape := make([]int, 100000)
	cursor := 50000
	state := 0

	for i := 0; i != checkAfter; i++ {
		if tape[cursor] == 0 {
			tape[cursor] = states[state].Oval
			cursor += states[state].Odir
			state = states[state].Ogoto
		} else {
			tape[cursor] = states[state].Ival
			cursor += states[state].Idir
			state = states[state].Igoto
		}
	}

	// Make checksum
	sum := 0
	for _, v := range tape {
		sum += v
	}
	fmt.Println(sum)
}
