package main

import "fmt"

func reverseBits(num uint32) uint32 {

	var res uint32 = 0
	for i := 32; i > 0; i-- {
		res = res << 1
		res = res | (num & 1)
		num >>= 1
	}
	return res
}

func main() {
	n := 43261596
	fmt.Println(reverseBits(uint32(n)))
}
