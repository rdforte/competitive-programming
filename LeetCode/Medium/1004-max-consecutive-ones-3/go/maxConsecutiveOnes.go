package maxOnes

import "math"

// time complexity = O(2n) = O(n)
// space = O(1)
func LongestOnes(nums []int, k int) int {
	longest := 0
	// loop through every num atleast once therefore O(n)
	for l, r := 0, 0; r < len(nums); r++ {
		if nums[r] == 0 {
			k--
		}

		// we visit each element here at most 1 for the entirety of the algorithm therefore O(n)
		for k < 0 {
			if nums[l] == 0 {
				k++
			}
			l++
		}

		longest = int(math.Max(float64(longest), float64((r-l)+1)))
	}

	return longest
}
