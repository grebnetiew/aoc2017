package main

import (
	"fmt"
	"math"
)

func main() {
	var square int
	fmt.Scanf("%d", &square)

	fmt.Println(coordsum(square))
}

func coordsum(n int) (sum int) {                               // say 21
	if n == 1 {
		return 0
	}
	maxNumInRingSR := ceilOdd(math.Sqrt(float64(n)))           // 5
	distanceAlongRing := maxNumInRingSR * maxNumInRingSR - n   // 4
	positionInLine := distanceAlongRing % (maxNumInRingSR - 1) // 0
	otherCoord := maxNumInRingSR / 2                           // 2 (the y in this case)
	lineCoord := abs(positionInLine - otherCoord)              // 2 (the x in this case)
	return lineCoord + otherCoord                              // 4
}

func ceilOdd(n float64) int {
	// Round n up to nearest odd number
	if math.Mod(math.Ceil(n), 2) == 0 {
		return int(math.Ceil(n)) + 1
	}
	return int(math.Ceil(n))
}

func abs(n int) int { // why does this only exist on floats
	if n > 0 {
		return n
	}
	return -n
}