package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const (
	clean    = iota
	weakened = iota
	infected = iota
	flagged  = iota
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
func (c *Carrier) uturn() {
	c.dir *= -1
}
func (c *Carrier) fw() {
	c.pos += c.dir
}

func main() {
	c := Carrier{pos: 0, dir: 1i}
	m := make(map[complex128]int)

	rows, _ := csv.NewReader(os.Stdin).ReadAll()
	for j, v := range rows {
		// v[0] is a line of input
		for k, c := range v[0] {
			if c == '#' {
				m[complex128(complex(float64(k-len(rows)/2), float64(len(rows)/2-j)))] = infected
			}
		}
	}

	infections := 0
	for k := 0; k < 10000000; k++ {
		switch m[c.pos] {
		case clean:
			c.left()
		case weakened:
			infections++
		case infected:
			c.right()
		case flagged:
			c.uturn()
		}
		m[c.pos] = (m[c.pos] + 1) % 4
		c.fw()
	}

	fmt.Println(infections)
}
