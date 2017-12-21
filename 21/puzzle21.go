package main

import (
	"fmt"
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

	bitmap := [][]string{{".#./..#/###"}}
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			// the inner sizes are 3x3
			// make new bitmap
			newBitmap := make([][]string, len(bitmap)*2)
			for i := range newBitmap {
				newBitmap[i] = make([]string, len(bitmap)*2)
			}
			// fill it
			for i, v := range bitmap {
				for j, w := range v {
					s1, s2, s3, s4 := split(q4[find(p3, w)])
					newBitmap[2*j][2*i] = s1
					newBitmap[2*j][2*i+1] = s2
					newBitmap[2*j+1][2*i] = s3
					newBitmap[2*j+1][2*i+1] = s4
				}
			}
			bitmap = newBitmap
		} else {
			// the inner sizes are 2x2
			for i, v := range bitmap {
				for j, w := range v {
					bitmap[i][j] = q3[find(p2, w)]
				}
			}
		}
		fmt.Println(bitmap)
	}

	c := 0
	for _, v := range bitmap {
		for _, w := range v {
			c += strings.Count(w, "#")
		}
	}
	fmt.Println(c)
}

func hflip(s string) string {
	if len(s) == 5 {
		return string([]byte{s[1], s[0], '/', s[4], s[3]})
	}
	return string([]byte{s[2], s[1], s[0], '/', s[6], s[5], s[4], '/', s[10], s[9], s[8]})
}
func vflip(s string) string {
	if len(s) == 5 {
		return string(s[3:4] + "/" + s[0:1])
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
