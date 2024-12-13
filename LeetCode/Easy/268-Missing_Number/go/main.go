package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}) == 2)
	fmt.Println(missingNumber([]int{0, 1}) == 2)
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) == 8)

	fmt.Println("<< XOR >>")
	fmt.Println(missingNumberXOR([]int{3, 0, 1}) == 2)
	fmt.Println(missingNumberXOR([]int{0, 1}) == 2)
	fmt.Println(missingNumberXOR([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) == 8)

	fmt.Println("<< GAUSS' FORMULA EXAMPLE>>")
	fmt.Println("Find sum of first 100 numbers")

	n := 100

	sumIterative := 0
	for i := 1; i <= n; i++ {
		sumIterative += i
	}

	sumGauss := n * (n + 1) / 2

	fmt.Println(sumIterative == sumGauss)

	fmt.Println(^29 == -30)
}

// Space O(1)
// Time = let n be the length of nums and s be all the nums that make up the length
// O(n+s)
func missingNumber(nums []int) int {
	sum := 0
	for i := 1; i <= len(nums); i++ {
		sum += i
	}
	for _, n := range nums {
		sum -= n
	}

	return sum
}

// We can improve the above by instead of looping through all 1..len(n) to get the sum we can do this in constant time O(1)
// Gaus Formula = n * (n + 1) / 2
// Now time complexity is len(nums)
func missingNumberGaussFormula(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, n := range nums {
		sum -= n
	}

	return sum
}

// XOR produces bits where bit in same position is different
func missingNumberXOR(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= i ^ nums[i]
	}
	return res
}
