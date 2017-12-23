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
	pc, mul := 0, 0
	for pc >= 0 && pc < len(instructions) {
		mul += reg.program(instructions, &pc)
	}

	fmt.Println(mul)
}

func (reg Registry) program(instructions [][]string, pc *int) (mul int) {
	cmd := instructions[*pc]
	switch cmd[0] {
	case "set":
		reg[cmd[1]] = reg.toInt(cmd[2])
	case "sub":
		reg[cmd[1]] -= reg.toInt(cmd[2])
	case "mul":
		reg[cmd[1]] *= reg.toInt(cmd[2])
		mul = 1
	case "jnz":
		if reg.toInt(cmd[1]) != 0 {
			*pc += int(reg.toInt(cmd[2])) - 1 // compensate for pc++ below
		}
	}
	*pc++
	return mul
}

func (reg Registry) toInt(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}
	return reg[s]
}
