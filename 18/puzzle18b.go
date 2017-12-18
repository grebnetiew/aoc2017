package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Registry map[string]int64

// type WaitGroup struct {
// 	c chan bool
// 	w int
// 	m int
// }

// func (w *WaitGroup) Wait() {
// 	w.w++
// 	if w.w == w.m {
// 		w.c <- true
// 	}
// }
// func (w *WaitGroup) Done() {
// 	w.w--
// }

func main() {
	reader := csv.NewReader(os.Stdin)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1 // variable number of fields
	instructions, _ := reader.ReadAll()

	reg0 := make(Registry)
	reg0["p"] = 0
	reg1 := make(Registry)
	reg0["p"] = 1

	c01 := make(chan int64, 100000) // 0 to 1
	c1m := make(chan int64, 100000) // 1 to me -- I need to count them
	cm0 := make(chan int64, 100000) // me to 0

	//cFail := make(chan bool)
	//w := WaitGroup{m: 2, w: 0, c: cFail}

	go mainLoop(reg0, instructions, cm0, c01)
	go mainLoop(reg1, instructions, c01, c1m)

	valuesSent := 0
	defer fmt.Println(valuesSent) // in case deadlock
	for v := range c1m {
		// select {
		// case v := <-c1m:
		// 	cm0 <- v
		// 	valuesSent++
		// case <-cFail:
		// 	break
		// }
		cm0 <- v
		valuesSent++
		fmt.Println(valuesSent)
	}
}

func mainLoop(reg Registry, instructions [][]string, cin, cout chan int64) {
	pc := 0
	for pc >= 0 && pc < len(instructions) {
		reg.program(instructions, &pc, cin, cout)
	}
	close(cout)
}

func (reg Registry) program(instructions [][]string, pc *int, cin, cout chan int64) {
	cmd := instructions[*pc]
	switch cmd[0] {
	case "snd":
		cout <- reg.toInt(cmd[1])
	case "set":
		reg[cmd[1]] = reg.toInt(cmd[2])
	case "add":
		reg[cmd[1]] += reg.toInt(cmd[2])
	case "mul":
		reg[cmd[1]] *= reg.toInt(cmd[2])
	case "mod":
		reg[cmd[1]] %= reg.toInt(cmd[2])
	case "rcv":
		reg[cmd[1]] = <-cin
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
