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

	reg0 := make(Registry)
	reg0["p"] = 0
	reg1 := make(Registry)
	reg1["p"] = 1

	c01 := make([]int64, 0)
	c10 := make([]int64, 0)
	pc0, pc1, valuesSent := 0, 0, 0
	for {
		keepgoing := true
		for pc0 >= 0 && pc0 < len(instructions) && keepgoing {
			keepgoing, _ = reg0.program(instructions, &pc0, &c10, &c01)
		}
		keepgoing = true
		for pc1 >= 0 && pc1 < len(instructions) && keepgoing {
			var v int
			keepgoing, v = reg1.program(instructions, &pc1, &c01, &c10)
			valuesSent += v
		}

		// 1 is waiting, 0 too. see if there is new info for 0
		if len(c10) == 0 {
			break
		}
	}
	fmt.Println(valuesSent)
}

func (reg Registry) program(instructions [][]string, pc *int, cin, cout *[]int64) (bool, int) {
	cmd := instructions[*pc]
	valuesSent := 0
	switch cmd[0] {
	case "snd":
		*cout = append(*cout, reg.toInt(cmd[1]))
		valuesSent++
	case "set":
		reg[cmd[1]] = reg.toInt(cmd[2])
	case "add":
		reg[cmd[1]] += reg.toInt(cmd[2])
	case "mul":
		reg[cmd[1]] *= reg.toInt(cmd[2])
	case "mod":
		reg[cmd[1]] %= reg.toInt(cmd[2])
	case "rcv":
		if len(*cin) == 0 {
			return false, valuesSent
		}
		reg[cmd[1]] = (*cin)[0]
		(*cin) = (*cin)[1:]
	case "jgz":
		if reg.toInt(cmd[1]) > 0 {
			*pc += int(reg.toInt(cmd[2])) - 1 // compensate for pc++ below
		}
	}
	*pc++
	return true, valuesSent
}

func (reg Registry) toInt(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}
	return reg[s]
}
