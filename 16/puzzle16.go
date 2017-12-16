package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	reader := csv.NewReader(os.Stdin)
	input, _ := reader.Read()

	line := dance([]byte("abcdefghijklmnop"), input)
	fmt.Println(string(line))

	var i int
	for i = 1; !isIdentity(line); i++ {
		line = dance(line, input)
	}
	fmt.Println("Cycle length is", i)

	for j := 0; j < 1000000000%i; j++ {
		line = dance(line, input)
	}
	fmt.Println(string(line))
}

func dance(line []byte, input []string) []byte {
	for _, command := range input {
		switch command[0] {
		case 's':
			var amount int
			fmt.Sscanf(command, "s%d", &amount)
			line = append(line[len(line)-amount:], line[:len(line)-amount]...)
		case 'x':
			var a, b int
			fmt.Sscanf(command, "x%d/%d", &a, &b)
			tmp := line[a]
			line[a] = line[b]
			line[b] = tmp
		case 'p':
			a, b := command[1], command[3]
			for k, v := range line {
				switch v {
				case a:
					line[k] = b
				case b:
					line[k] = a
				}
			}
		}
	}
	return line
}

func isIdentity(line []byte) bool {
	for i, v := range line {
		if int(v-'a') != i {
			return false
		}
	}
	return true
}
