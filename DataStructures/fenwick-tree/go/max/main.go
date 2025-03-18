package main

import "fmt"

func main() {
	// 0

	f := New(21)
	f.Add(0, 0)
	f.Add(3, 3)
	f.Add(7, 4)
	f.Add(15, 8)
	f.Add(19, 4)
	f.Add(20, 1)

	fmt.Println(f.Max(20))
}

type FenwickTree []int

func New(size int) FenwickTree {
	return make(FenwickTree, size)
}

func (f FenwickTree) Add(pos, val int) {
	for ; pos < len(f); pos |= (pos + 1) {
		f[pos] = max(f[pos], val)
	}
}

// return sum [begin, end]
// func (f FenwickTree) Max(begin, end int) (maxNum int) {
// maxNum = f.prefixMax(end) - f.prefixMax(begin-1)
// return
// }

func (f FenwickTree) Max(pos int) (mx int) {
	for ; pos >= 0; pos = (pos & (pos + 1)) - 1 {
		mx = max(f[pos], mx)
	}
	return
}
