package main

import (
	"fmt"
	"log"
)

const (
	factorA = 16807
	factorB = 48271
	modulus = 2147483647
)

func main() {
	var startA, startB uint64
	fmt.Scanf("%d %d\n", &startA, &startB)

	currentA, currentB, score := startA, startB, 0
	log.Println("Now starting 50M iterations, old fashioned way")
	for i := 0; i < 50000000; i++ {
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
	log.Println(score)

	log.Println("Now starting 50M iterations, coroutine way")
	cA := make(chan uint64, 10000)
	cB := make(chan uint64, 10000)
	go generator(cA, startA, factorA, 4)
	go generator(cB, startB, factorB, 8)
	score = 0
	for rA, ok := <-cA; ok; rA, ok = <-cA {
		rB := <-cB
		if rA&0xffff == rB&0xffff {
			score++
		}
	}
	log.Println(score)
}

func generator(c chan uint64, current uint64, factor uint64, mod uint64) {
	for i := 0; i < 50000000; i++ {
		for {
			if current = (current * factor) % modulus; current%mod == 0 {
				break
			}
		}
		c <- current
	}
	close(c)
}
