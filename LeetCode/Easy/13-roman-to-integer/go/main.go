package main

import "fmt"

func main() {
	// fmt.Println(romanToInt("III"))
	// fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	numerals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := numerals[s[0]]

	for i := 1; i < len(s); i++ {
		// refactored. We subtract if the numeral before is less than the current.
		if numerals[s[i-1]] < numerals[s[i]] {
			res = res - numerals[s[i-1]] + (numerals[s[i]] - numerals[s[i-1]])
		} else {
			res += numerals[s[i]]
		}
		// if s[i] == 'V' && s[i-1] == 'I' {
		// res = res - numerals['I'] + (numerals['V'] - numerals['I'])
		// } else if s[i] == 'L' && s[i-1] == 'X' {
		// res = res - numerals['X'] + (numerals['L'] - numerals['X'])
		// } else if s[i] == 'D' && s[i-1] == 'C' {
		// res = res - numerals['C'] + (numerals['D'] - numerals['C'])
		// } else if s[i] == 'M' && s[i-1] == 'C' {
		// res = res - numerals['C'] + (numerals['M'] - numerals['C'])
		// } else if s[i] == 'C' && s[i-1] == 'X' {
		// res = res - numerals['X'] + (numerals['C'] - numerals['X'])
		// } else if s[i] == 'X' && s[i-1] == 'I' {
		// res = res - numerals['I'] + (numerals['X'] - numerals['I'])
		// } else {
		// res += numerals[s[i]]
		// }
	}

	return res
}
