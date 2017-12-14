package main

import (
	"./knothash"
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	totalBits := 0
	for i := 0; i < 128; i++ {
		h := knothash.KnotHash([]byte(fmt.Sprintf("%s-%d", input, i)))
		totalBits += countBits(h)
	}
	fmt.Println(totalBits)
}

func countBits(hash []int) int {
	total := 0
	for _, h := range hash {
		v := byte(h)
		total += int(v&1 + (v>>1)&1 + (v>>2)&1 + (v>>3)&1 +
			(v>>4)&1 + (v>>5)&1 + (v>>6)&1 + (v>>7)&1)
	}
	return total
}
