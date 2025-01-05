package main

func main() {
}

func maxLength(nums []int) int {
	maxLen := 0

	for i := 1; i < len(nums); i++ {
		prod := nums[i]
		l := nums[i]
		g := nums[i]
		for j := i - 1; j >= 0; j-- {
			prod = prod * nums[j]
			l = lcm(nums[j], l)
			g = gcd(nums[j], g)

			if prod == l*g {
				maxLen = max(maxLen, (i-j)+1)
			}
		}
	}

	return maxLen
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
