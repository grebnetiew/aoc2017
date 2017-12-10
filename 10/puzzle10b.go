package main

import (
	"bufio"
	"fmt"
	"os"
)

type CircList []int

func (l CircList) reverse(begin, length int) {
	for i := 0; i < length/2; i++ {
		temp := l.get(begin + i)
		l.set(begin+i, l.get(begin+length-i-1))
		l.set(begin+length-i-1, temp)
	}
}

func (l CircList) get(idx int) int {
	return l[idx%len(l)]
}
func (l CircList) set(idx, val int) {
	l[idx%len(l)] = val
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		bytes, err := r.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		// The last one is now '\n', so..
		arbitrarySuffix := []byte{17, 31, 73, 47, 23}
		bytes = append(bytes[0:len(bytes)-1], arbitrarySuffix...)

		hash := knotHash(bytes)

		for i := 0; i < 16; i++ {
			fmt.Printf("%02x", hash[i])
		}
		fmt.Println()
	}
}

func knotHash(bytes []byte) []int {
	var circList CircList = make([]int, 256)
	for i := range circList {
		circList[i] = i
	}

	// The 64 hashes
	var cursor, skipSize int
	for i := 0; i < 64; i++ {
		for _, v := range bytes {
			length := int(v)
			circList.reverse(cursor, length)
			cursor += skipSize + length
			skipSize++
		}
	}

	// Now reduce it
	denseHash := make([]int, 16)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			denseHash[i] ^= circList[16*i+j]
		}
	}
	return denseHash
}
