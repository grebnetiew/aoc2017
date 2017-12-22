package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var p2, q3, p3, q4 []string
	for {
		var p, q string
		_, err := fmt.Scanf("%s => %s\n", &p, &q)
		if err != nil {
			break
		}
		if len(p) == 5 {
			p2 = append(p2, p)
			q3 = append(q3, q)
		} else {
			p3 = append(p3, p)
			q4 = append(q4, q)
		}
	}

	bitmap := []byte(".#...####")
	for i := 0; i < 18; i++ {
		if sz := sqrt(len(bitmap)); sz%2 == 0 {
			nSz := sz / 2 * 3
			newBitmap := make([]byte, nSz*nSz)
			// Split it up in blocks of two
			for j := 0; j < sz/2; j++ {
				for k := 0; k < sz/2; k++ {
					// take this 2x2 bit and make a corresponding 3x3 bit
					s := q3[find(p2, string([]byte{bitmap[sz*(2*j)+2*k], bitmap[sz*(2*j)+2*k+1], '/', bitmap[sz*(2*j+1)+2*k], bitmap[sz*(2*j+1)+2*k+1]}))]
					for u := 0; u < 3; u++ {
						for v := 0; v < 3; v++ {
							if newBitmap[(3*j+u)*nSz+3*k+v] != 0 {
								fmt.Println("!!! Assigning to nonempty")
							}
							newBitmap[(3*j+u)*nSz+3*k+v] = s[u*(3+1)+v] // the slashes..
						}
					}
				}
			}
			bitmap = newBitmap
		} else {
			// it's a multiple of three
			nSz := sz / 3 * 4
			newBitmap := make([]byte, nSz*nSz)
			// Split it up in blocks of three
			for j := 0; j < sz/3; j++ {
				for k := 0; k < sz/3; k++ {
					// take this 3x3 bit and make a corresponding 4x4 bit
					s := q4[find(p3, string(bitmap[sz*(3*j)+3*k:sz*(3*j)+3*k+3])+"/"+string(bitmap[sz*(3*j+1)+3*k:sz*(3*j+1)+3*k+3])+"/"+string(bitmap[sz*(3*j+2)+3*k:sz*(3*j+2)+3*k+3]))]
					for u := 0; u < 4; u++ {
						for v := 0; v < 4; v++ {
							if newBitmap[(j*4+u)*nSz+k*4+v] != 0 {
								fmt.Println("!!! Assigning to nonempty")
							}
							newBitmap[(j*4+u)*nSz+k*4+v] = s[u*(4+1)+v] // the slashes..
						}
					}
				}
			}
			bitmap = newBitmap
		}
		// fmt.Println("----------------------")
		// printBitmap(bitmap)
		// fmt.Println("----------------------")
	}

	fmt.Println(strings.Count(string(bitmap), "#"))
}

func hflip(s string) string {
	if len(s) == 5 {
		return string([]byte{s[1], s[0], '/', s[4], s[3]})
	}
	return string([]byte{s[2], s[1], s[0], '/', s[6], s[5], s[4], '/', s[10], s[9], s[8]})
}
func vflip(s string) string {
	if len(s) == 5 {
		return string(s[3:5] + "/" + s[0:2])
	}
	return string(s[8:11] + "/" + s[4:7] + "/" + s[0:3])
}
func rot(s string) string {
	if len(s) == 5 {
		return string([]byte{s[3], s[0], '/', s[4], s[1]})
	}
	return string([]byte{s[8], s[4], s[0], '/', s[9], s[5], s[1], '/', s[10], s[6], s[2]})
}
func find(hay []string, s string) int {
	// just generate the group
	s1 := hflip(s)
	s2 := vflip(s)
	s3 := rot(s)
	s4 := rot(s3)
	s5 := rot(s4)
	s6 := hflip(s3)
	s7 := vflip(s3)
	for i, v := range hay {
		if s == v || s1 == v || s2 == v || s3 == v || s4 == v || s5 == v || s6 == v || s7 == v {
			return i
		}
	}
	fmt.Println("Can not find ", s)
	return -1
}
func split(s string) (string, string, string, string) {
	return string(s[0:2] + "/" + s[5:7]), string(s[2:4] + "/" + s[7:9]), string(s[10:12] + "/" + s[15:17]), string(s[12:14] + "/" + s[17:19])
}
func sqrt(i int) int {
	return int(math.Sqrt(float64(i)) + 0.5)
}
func printBitmap(b []byte) {
	sz := sqrt(len(b))
	for i := 0; i < sz; i++ {
		fmt.Println(string(b[sz*i : sz*(i+1)]))
	}
}
