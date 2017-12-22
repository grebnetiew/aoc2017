package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Carrier struct {
	pos, dir complex128 // it is crazy, but I'm betting it works >:)
}

func (c *Carrier) left() {
	c.dir *= 1i
}
func (c *Carrier) right() {
	c.dir *= -1i
}
func (c *Carrier) fw() {
	c.pos += c.dir
}

func main() {
	c := Carrier{pos: 0, dir: 1i}
	m := make(map[complex128]bool)

	rows, _ := csv.NewReader(os.Stdin).ReadAll()
	for j, v := range rows {
		// v[0] is a line of input
		for k, c := range v[0] {
			if c == '#' {
				m[complex128(complex(float64(k-len(rows)/2), float64(len(rows)/2-j)))] = true
			}
		}
	}

	infections := 0
	for k := 0; k < 10000; k++ {
		if m[c.pos] {
			c.right()
			m[c.pos] = false
		} else {
			c.left()
			m[c.pos] = true
			infections++
		}
		c.fw()
	}

	fmt.Println(infections)
}
