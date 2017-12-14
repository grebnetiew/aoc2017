package main

import (
	"./drive"
	"./knothash"
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	var d drive.Drive
	for i := 0; i < 128; i++ {
		h := knothash.KnotHash([]byte(fmt.Sprintf("%s-%d", input, i)))
		bytes := make([]byte, len(h))
		for i, v := range h {
			bytes[i] = byte(v)
		}
		d = append(d, bytes)
	}

	regions := 0
	for i := uint(0); i < 128*128; i++ {
		if d.GetBit(i) {
			regions++
			d.RemoveRegion(i)
		}
	}
	fmt.Println(regions)
}
