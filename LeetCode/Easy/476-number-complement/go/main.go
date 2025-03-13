package main

func main() {
}

/*
NOTE:
num = 1025
= 0000010000000001

1.
left shift by 1 = 0000001000000000
OR with bitmask = 0000011000000001

2.
left shift by 2 = 0000000110000000
OR with bitmask = 0000011110000001

3.
left shift by 2 = 0000000001111000
OR with bitmask = 0000011111111001

4.
left shift by 2 = 0000000000000111
OR with bitmask = 0000011111111111
*/

func findComplement(num int) int {
	bitmask := num
	bitmask |= (bitmask >> 1)  // 1.
	bitmask |= (bitmask >> 2)  // 2.
	bitmask |= (bitmask >> 4)  // 3.
	bitmask |= (bitmask >> 8)  // 4.
	bitmask |= (bitmask >> 16) // 5.

	return bitmask ^ num
}
