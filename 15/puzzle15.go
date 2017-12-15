package main

import "fmt"

const (
	factorA = 16807
	factorB = 48271
	modulus = 2147483647
)

func main() {
	var startA, startB uint64
	fmt.Scanf("%d %d\n", &startA, &startB)

	currentA, currentB := startA, startB
	score := 0
	for i := 0; i < 40000000; i++ {
		currentA = (currentA * factorA) % modulus
		currentB = (currentB * factorB) % modulus
		if currentA&0xffff == currentB&0xffff {
			score++
		}
	}
	fmt.Println(score)

	currentA, currentB, score = startA, startB, 0
	for i := 0; i < 5000000; i++ {
		for {
			if currentA = (currentA * factorA) % modulus; currentA%4 == 0 {
				break
			}
		}
		for {
			if currentB = (currentB * factorB) % modulus; currentB%8 == 0 {
				break
			}
		}
		if currentA&0xffff == currentB&0xffff {
			score++
		}
	}
	fmt.Println(score)
}
