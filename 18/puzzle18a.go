package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Registry map[string]int64

func main() {
	reader := csv.NewReader(os.Stdin)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1 // variable number of fields
	instructions, _ := reader.ReadAll()

	reg := make(Registry)
	pc, rval, recovered := 0, int64(0), false
	for pc >= 0 && pc < len(instructions) && !recovered {
		rval, recovered = reg.program(instructions, &pc)
	}

	fmt.Println(rval)
}

func (reg Registry) program(instructions [][]string, pc *int) (rval int64, recovered bool) {
	cmd := instructions[*pc]
	switch cmd[0] {
	case "snd":
		reg["soundBlaster"] = reg.toInt(cmd[1])
	case "set":
		reg[cmd[1]] = reg.toInt(cmd[2])
	case "add":
		reg[cmd[1]] += reg.toInt(cmd[2])
	case "mul":
		reg[cmd[1]] *= reg.toInt(cmd[2])
	case "mod":
		reg[cmd[1]] %= reg.toInt(cmd[2])
	case "rcv":
		if reg.toInt(cmd[1]) != 0 {
			rval = reg["soundBlaster"]
			recovered = true
		}
	case "jgz":
		if reg.toInt(cmd[1]) > 0 {
			*pc += int(reg.toInt(cmd[2])) - 1 // compensate for pc++ below
		}
	}
	*pc++
	return
}

func (reg Registry) toInt(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}
	return reg[s]
}
