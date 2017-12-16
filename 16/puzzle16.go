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

	fmt.Println(string(dance(dance([]byte("abcdefghijklmnop"), input), input)))
	fmt.Println(string(dancecheat(dancecheat([]byte("abcdefghijklmnop"), input), input)))

	perm := dancecheat([]byte("abcdefghijklmnop"), input)
	//these two should be equal
	fmt.Println(string(perm))
	fmt.Println(string(permutation([]byte("abcdefghijklmnop"), perm)))

	fmt.Println(string(dance(line, input)))
	fmt.Println(string(permutation(perm, perm)))

	for i := 0; i < 1000000000-1; i++ {
		line = permutation(line, perm)
		if i%10000000 == 0 {
			fmt.Printf("%d percent\n", i/10000000)
		}
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

func dancecheat(line []byte, input []string) []byte {
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
		}
	}
	return line
}

func permutation(line, perm []byte) []byte {
	newline := make([]byte, len(line))
	for i, v := range perm {
		j := v - byte('a')
		newline[i] = line[j]
	}
	return newline
}
