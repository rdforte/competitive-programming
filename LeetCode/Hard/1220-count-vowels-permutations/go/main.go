package main

import "fmt"

func main() {
	fmt.Println("BRUTE FORCE")
	fmt.Println(countVowelsPermutations(1))
	fmt.Println(countVowelsPermutations(2))
	fmt.Println(countVowelsPermutations(5))
	fmt.Println("TOP DOWN")
	fmt.Println(countVowelsPermutationsTopDown(1))
	fmt.Println(countVowelsPermutationsTopDown(2))
	fmt.Println(countVowelsPermutationsTopDown(5))
	fmt.Println(countVowelsPermutationsTopDown(144))
}

// Brute Force no memoisation
func countVowelsPermutations(n int) int {
	var dp func(prevChar byte, strLen int) int
	dp = func(prevChar byte, strLen int) int {
		if strLen == n-1 {
			return 1
		}

		var sum int
		switch prevChar {
		case 'a':
			sum += dp('e', strLen+1)
		case 'e':
			sum += dp('a', strLen+1)
			sum += dp('i', strLen+1)
		case 'i':
			sum += dp('a', strLen+1)
			sum += dp('e', strLen+1)
			sum += dp('o', strLen+1)
			sum += dp('u', strLen+1)
		case 'o':
			sum += dp('i', strLen+1)
			sum += dp('u', strLen+1)
		case 'u':
			sum += dp('a', strLen+1)
		default:
			panic("shouldnt reach this case")
		}

		return sum % (1e9 + 7)
	}

	total := 0

	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, v := range vowels {
		total += dp(v, 0)
	}

	return total
}

// Top Down
func countVowelsPermutationsTopDown(n int) int {
	mod := int(1e9 + 7)
	memo := make([][]int, 0, n)
	for range n {
		chars := make([]int, 0, 26)
		for range 26 {
			chars = append(chars, -1)
		}
		memo = append(memo, chars)
	}

	var dp func(prevChar byte, strLen int) int
	dp = func(prevChar byte, strLen int) int {
		if strLen == n-1 {
			return 1
		}

		if memo[strLen][prevChar-'a'] == -1 {
			var sum int
			switch prevChar {
			case 'a':
				sum += dp('e', strLen+1)
			case 'e':
				sum += dp('a', strLen+1)
				sum += dp('i', strLen+1)
			case 'i':
				sum += dp('a', strLen+1)
				sum += dp('e', strLen+1)
				sum += dp('o', strLen+1)
				sum += dp('u', strLen+1)
			case 'o':
				sum += dp('i', strLen+1)
				sum += dp('u', strLen+1)
			case 'u':
				sum += dp('a', strLen+1)
			default:
				panic("shouldnt reach this case")
			}

			memo[strLen][prevChar-'a'] = sum % mod
		}

		return memo[strLen][prevChar-'a']
	}

	total := 0

	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, v := range vowels {
		// note that total += dp(v, 0) % mod is not the same as
		// total = (total + dp(v, 0)) % mod
		total = (total + dp(v, 0)) % mod
	}

	return total
}
