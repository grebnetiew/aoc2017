package main

import "fmt"

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
func (l CircList) set(idx int, val int) {
	l[idx%len(l)] = val
}

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	totalBits := 0
	for i := 0; i < 128; i++ {
		h := knotHash([]byte(fmt.Sprintf("%s-%d", input, i)))
		totalBits += countBits(h)
	}
	fmt.Println(totalBits)
}

func knotHash(bytes []byte) []int {
	arbitrarySuffix := []byte{17, 31, 73, 47, 23}
	bytes = append(bytes, arbitrarySuffix...)

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

func countBits(hash []int) int {
	total := 0
	for _, h := range hash {
		v := byte(h)
		total += int(v&1 + (v>>1)&1 + (v>>2)&1 + (v>>3)&1 +
			(v>>4)&1 + (v>>5)&1 + (v>>6)&1 + (v>>7)&1)
	}
	return total
}
