package main

import "fmt"

type Coor [3]int

func (c *Coor) Add(d Coor) {
	for i, v := range d {
		(*c)[i] += v
	}
}
func (c Coor) Equal(d Coor) bool {
	return c[0] == d[0] && c[1] == d[1] && c[2] == d[2]
}

func (c *Coor) Norm() int {
	return Abs((*c)[0]) + Abs((*c)[1]) + Abs((*c)[2])
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func Delete(slice *[]Coor, i int) {
	(*slice)[i] = (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
}

func main() {
	var positions, velocities, accelerations []Coor
	for {
		var x, y, z, u, v, w, a, b, c int
		n, err := fmt.Scanf("p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>\n", &x, &y, &z, &u, &v, &w, &a, &b, &c)
		if n < 9 || err != nil {
			break
		}
		positions = append(positions, Coor{x, y, z})
		velocities = append(velocities, Coor{u, v, w})
		accelerations = append(accelerations, Coor{a, b, c})
	}

	for i := 0; i < 100000; i++ {
		simulate(&positions, &velocities, &accelerations)
	}

	fmt.Println(len(positions))
}

func simulate(positions, velocities, accelerations *[]Coor) {
	for i, v := range *accelerations {
		(*velocities)[i].Add(v)
	}
	for i, v := range *velocities {
		(*positions)[i].Add(v)
	}
	delete := make([]bool, len(*positions))
	for i, v := range *positions {
		for j := 0; j < i; j++ {
			if v.Equal((*positions)[j]) {
				delete[i] = true
				delete[j] = true
			}
		}
	}
	for i := len(delete) - 1; i >= 0; i-- {
		if delete[i] {
			Delete(positions, i)
			Delete(velocities, i)
			Delete(accelerations, i)
		}
	}
}
