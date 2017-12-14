package knothash

type circList []int

func (l circList) reverse(begin, length int) {
	for i := 0; i < length/2; i++ {
		temp := l.get(begin + i)
		l.set(begin+i, l.get(begin+length-i-1))
		l.set(begin+length-i-1, temp)
	}
}

func (l circList) get(idx int) int {
	return l[idx%len(l)]
}
func (l circList) set(idx int, val int) {
	l[idx%len(l)] = val
}

func KnotHash(bytes []byte) []int {
	arbitrarySuffix := []byte{17, 31, 73, 47, 23}
	bytes = append(bytes, arbitrarySuffix...)

	var cl circList = make([]int, 256)
	for i := range cl {
		cl[i] = i
	}

	// The 64 hashes
	var cursor, skipSize int
	for i := 0; i < 64; i++ {
		for _, v := range bytes {
			length := int(v)
			cl.reverse(cursor, length)
			cursor += skipSize + length
			skipSize++
		}
	}

	// Now reduce it
	denseHash := make([]int, 16)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			denseHash[i] ^= cl[16*i+j]
		}
	}
	return denseHash
}
