package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	score, depth, garbage, removed := 0, 0, false, 0
	for {
		var b byte
		var err error
		if b, err = r.ReadByte(); err != nil {
			break
		}
		if !garbage {
			switch b {
			case '{':
				depth++
			case '}':
				score += depth
				depth--
			case '<':
				garbage = true
			}
		} else {
			switch b {
			case '!':
				r.ReadByte()
			case '>':
				garbage = false
			default:
				removed++
			}
		}
	}
	if depth != 0 {
		fmt.Println("Unclosed { in stream")
	}
	fmt.Println(score)
	fmt.Println(removed)
}
