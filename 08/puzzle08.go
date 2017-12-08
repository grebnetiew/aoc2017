package main

import "fmt"

const MinInt32 = -1 << 31

type Registry map[string]int

func main() {
	reg := make(Registry)
	record := MinInt32
	for {
		r, ins, amt, cr, cond, co := "", "", 0, "", "", 0
		if _, err := fmt.Scanf("%s %s %d if %s %s %d\n",
			&r, &ins, &amt, &cr, &cond, &co); err != nil {
			break
		}
		reg.program(r, ins, amt, cr, cond, co)
		if max := reg.largest(); record < max {
			record = max
		}
	}

	fmt.Println(reg.largest())
	fmt.Println(record)
}

func (reg Registry) program(r string, ins string, amt int, cr string, cond string, co int) {
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

func (reg Registry) largest() int {
	max := MinInt32
	for _, val := range reg {
		if val > max {
			max = val
		}
	}
	return max
}
