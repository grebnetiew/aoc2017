package main

import (
	"fmt"
	"log"
)

func main() {
	var numbersMason string
	if _, err := fmt.Scanln(&numbersMason); err != nil {
		log.Fatalf("What do the numbers mean Mason\n%s\n", err)
	}
	log.Printf("There are %d numbers\n", len(numbersMason))
	fmt.Printf("%d\n", sumRepeatedNumbers(numbersMason))
	fmt.Printf("%d\n", sumCircularNumbers(numbersMason))
}

func sumRepeatedNumbers(s string) int {
	total := 0
	prev := s[len(s)-1]
	for i := 0; i < len(s); i++ {
		if s[i] == prev {
			var digit int
			if _, err := fmt.Sscan(string(prev), &digit); err != nil {
				log.Fatalf("Error scanning %s, %s\n", s[i], err)
			}
			total += digit
		}
		prev = s[i]
	}
	return total
}

func sumCircularNumbers(s string) int {
	total := 0
	halfLen := len(s) / 2
	for i := 0; i < halfLen; i++ {
		if s[i] == s[i+halfLen] {
			var digit int
			if _, err := fmt.Sscan(string(s[i]), &digit); err != nil {
				log.Fatalf("Error scanning %s, %s\n", s[i], err)
			}
			total += digit
		}
	}
	return total * 2
}
