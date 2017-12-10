package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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
	reader := csv.NewReader(os.Stdin)
	reader.TrimLeadingSpace = true
	row, _ := reader.ReadAll()

	var circList CircList = make([]int, 256)
	for i := range circList {
		circList[i] = i
	}

	var cursor, skipSize int
	for _, v := range row[0] {
		length, _ := strconv.Atoi(v)
		circList.reverse(cursor, length)
		cursor += skipSize + length
		skipSize++
	}

	fmt.Println(circList[0] * circList[1])
}
