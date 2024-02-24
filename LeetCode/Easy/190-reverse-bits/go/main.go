package main

import "fmt"

func main() {
	fmt.Println(reverseBits(43261596) == 964176192)
	fmt.Println(reverseBits(4294967293) == 3221225471)
}

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 31; i >= 0; i-- {
		if num&(1<<i) > 0 {
			res += (1 << (31 - i))
		}
	}
	return res
}
