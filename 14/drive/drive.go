package drive

import "fmt"

// The drive is a matrix of bytes, 128 rows by 16 columns
// Filling is is an exercise for the reader
type Drive [][]byte

// Get the value of a bit. Addressing is left-right top-bottom.
// First select the row, this is a/128
// Then the column, 128/16 = 8 so divide the low bits of a by 8
// Then the bit, which is a%8. Select the bit with &,
// 128 is the leftmost bit so we shift it right as needed.
func (d Drive) GetBit(addr uint) bool {
	return d[addr/128][(addr%128)/8]&(128>>(addr%8)) != 0
}

// Set a bit to zero. Addressing is as above, but we use AND NOT
// to clear the relevant bit.
func (d Drive) ZeroBit(addr uint) {
	d[addr/128][(addr%128)/8] &^= 128 >> (addr % 8)
}

// Set a bit to one. Using OR this time.
func (d Drive) OneBit(addr uint) {
	d[addr/128][(addr%128)/8] |= 128 >> (addr % 8)
}

// Flood fill recursively with zeros from a position addr.
func (d Drive) RemoveRegion(addr uint) {
	d.ZeroBit(addr)
	if addr%128 < 127 && d.GetBit(addr+1) {
		d.RemoveRegion(addr + 1)
	}
	if addr%128 > 0 && d.GetBit(addr-1) {
		d.RemoveRegion(addr - 1)
	}
	if addr < 128*127-1 && d.GetBit(addr+128) {
		d.RemoveRegion(addr + 128)
	}
	if addr > 128 && d.GetBit(addr-128) {
		d.RemoveRegion(addr - 128)
	}
}

// Print the entire contents using the exercise's #/. notation
func (d Drive) Print() {
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if d.GetBit(uint(128*i + j)) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
