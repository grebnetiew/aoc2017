package main

import (
	"fmt"
)

const (
	MinInt32 = -1 << 31
)

type Registry map[string]int

func main() {
	// Get the input

	reg := make(Registry)
	record := MinInt32
	for {
		r, ins, amt, cr, cond, co := "", "", 0, "", "", 0
		_, err := fmt.Scanf("%s %s %d if %s %s %d\n",
			&r, &ins, &amt, &cr, &cond, &co)
		if err != nil {
			break
		}
		program(reg, r, ins, amt, cr, cond, co)
		if record < largest(reg) {
			record = largest(reg)
		}
	}

	fmt.Println(largest(reg))
	fmt.Println(record)
}

func program(reg Registry, r string, ins string, amt int, cr string, cond string, co int) {
	// See if we should do anything
	var result bool
	switch cond {
	case ">":
		result = reg[cr] > co
	case "<":
		result = reg[cr] < co
	case ">=":
		result = reg[cr] >= co
	case "<=":
		result = reg[cr] <= co
	case "==":
		result = reg[cr] == co
	case "!=":
		result = reg[cr] != co
	}
	if !result {
		return
	}
	// then do it
	switch ins {
	case "inc":
		reg[r] += amt
	case "dec":
		reg[r] -= amt
	}
}

func largest(r Registry) int {
	max := MinInt32
	for _, val := range r {
		if val > max {
			max = val
		}
	}
	return max
}
