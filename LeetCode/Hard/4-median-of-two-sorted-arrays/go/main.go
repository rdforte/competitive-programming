package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(findMedianSortedArrays([]int{1, 2, 4, 8, 12, 21}, []int{1, 2, 2, 3, 8, 13}))
	// fmt.Println(findMedianSortedArrays1([]int{1, 3}, []int{2}))
	// fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
}

// With these bigger questions you have to break it down into steps and focus on solving those smaller steps.
// 1 step at a time

func findMedianSortedArrays(a, b []int) float64 {
	totalCombinedLen := len(a) + len(b)
	partitionSize := (totalCombinedLen + 1) / 2

	// we want to do binary search on the smallest
	if len(b) < len(a) {
		a, b = b, a
	}

	l, r := 0, len(a)-1

	for {
		aMid := l + (r-l)/2
		bMid := partitionSize - aMid - 2 // 1 index from nums1 and 1 index from nums2 thats why -2

		// set defaults to infinity so its easy to calculate when index out of bounds
		aLeft, aRight, bLeft, bRight := math.Inf(-1), math.Inf(1), math.Inf(-1), math.Inf(1)
		if len(a) > 0 && aMid >= 0 {
			aLeft = float64(a[aMid])
		}
		if (aMid + 1) < len(a) {
			aRight = float64(a[aMid+1])
		}
		if bMid >= 0 {
			bLeft = float64(b[bMid])
		}
		if (bMid + 1) < len(b) {
			bRight = float64(b[bMid+1])
		}

		// partition is correct
		if aLeft <= bRight && bLeft <= aRight {
			// Even len
			if totalCombinedLen%2 == 0 {
				return (math.Max(aLeft, bLeft) + math.Min(bRight, aRight)) / 2
			}
			// Odd len
			return math.Max(aLeft, bLeft)
		}

		// partition is not correct
		if aLeft > bRight {
			r = aMid - 1
		} else {
			// n2Left > n1Right
			l = aMid + 1
		}
	}
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	total := len(nums1) + len(nums2)
	half := (total + 1) / 2

	var Aleft, Aright float64
	var Bleft, Bright float64

	if len(B) < len(A) {
		A, B = B, A
	}

	l, r := 0, len(A)-1
	for {
		i := (l + r) >> 1 // A
		j := half - i - 2 // B

		if i >= 0 {
			Aleft = float64(A[i])
		} else {
			Aleft = math.Inf(-1)
		}

		if (i + 1) < len(A) {
			Aright = float64(A[i+1])
		} else {
			Aright = math.Inf(1)
		}

		if j >= 0 {
			Bleft = float64(B[j])
		} else {
			Bleft = math.Inf(-1)
		}

		if (j + 1) < len(B) {
			Bright = float64(B[j+1])
		} else {
			Bright = math.Inf(1)
		}

		// partition is correct
		if Aleft <= Bright && Bleft <= Aright {
			// odd
			if total%2 == 1 {
				return max(Aleft, Bleft)
			}
			// even
			return (max(Aleft, Bleft) + min(Aright, Bright)) / 2
		} else if Aleft > Bright {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
