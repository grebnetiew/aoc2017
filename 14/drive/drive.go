package drive

import "fmt"

type Drive [][]byte

func (d Drive) GetBit(addr uint) bool {
	return d[addr/128][(addr%128)/8]&(128>>(addr%8)) != 0
}
func (d Drive) ZeroBit(addr uint) {
	d[addr/128][(addr%128)/8] &^= 128 >> (addr % 8)
}
func (d Drive) OneBit(addr uint) {
	d[addr/128][(addr%128)/8] |= 128 >> (addr % 8)
}
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
