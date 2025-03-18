package main

import "fmt"

func main() {
	// Prefix Sum =  3 5 4 10 15 19 16 19 26 28 31

	f := New(11)
	f.Add(0, 3)
	f.Add(1, 2)
	f.Add(2, -1)
	f.Add(3, 6)
	f.Add(4, 5)
	f.Add(5, 4)
	f.Add(6, -3)
	f.Add(7, 3)
	f.Add(8, 7)
	f.Add(9, 2)
	f.Add(10, 3)

	fmt.Println("--------------------")
	fmt.Println(f.Sum(0, 0))
	fmt.Println(f.Sum(0, 1))
	fmt.Println(f.Sum(0, 2))
	fmt.Println(f.Sum(0, 3))
	fmt.Println(f.Sum(0, 4))
	fmt.Println(f.Sum(0, 5))
	fmt.Println(f.Sum(0, 6))
	fmt.Println(f.Sum(0, 7))
	fmt.Println(f.Sum(0, 8))
	fmt.Println(f.Sum(0, 9))
	fmt.Println(f.Sum(0, 10))
	fmt.Println("--------------------")
	fmt.Println(f.Sum(5, 10)) // 31 - 15 = 16
	f.Add(5, 2)
	fmt.Println(f.Sum(5, 10)) // 33 - 15 = 18
}

type FenwickTree []int

func New(size int) FenwickTree {
	return make(FenwickTree, size)
}

func (f FenwickTree) Add(pos, val int) {
	for ; pos < len(f); pos |= (pos + 1) {
		f[pos] += val
	}
}

// return sum [begin, end]
func (f FenwickTree) Sum(begin, end int) (sum int) {
	sum = f.prefixSum(end) - f.prefixSum(begin-1)
	return
}

func (f FenwickTree) prefixSum(pos int) (sum int) {
	for ; pos >= 0; pos = (pos & (pos + 1)) - 1 {
		sum += f[pos]
	}
	return
}
