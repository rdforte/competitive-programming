package main

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABCABCABC", "ABCABC") == gcdOfStringsOpt("ABCABCABC", "ABCABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB") == gcdOfStringsOpt("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("AA", "A") == gcdOfStringsOpt("AA", "A"))
	fmt.Println(gcdOfStrings("LEET", "CODE") == gcdOfStringsOpt("LEET", "CODE"))
	fmt.Println(gcdOfStrings("ABC", "ABCDEF") == gcdOfStringsOpt("ABC", "ABCDEF"))
}

// Brute Force
// Time: min(str1, str2).(str1 + str2) = We loop throught the smalles string and for each iteration if its a match we build a substring for str1 and str2.
// Space: for each iteration we hold the substring of the smallest string min(str1, str2)
func gcdOfStrings(str1 string, str2 string) string {
	smallest := min(len(str1), len(str2))
	largest := max(len(str1), len(str2))

	for i := smallest; i > 0; i-- {
		lModI := largest%i == 0
		sModI := smallest%i == 0
		same := str1[:i] == str2[:i]
		if same && lModI && sModI {
			// check to see if substring adds up to largest string
			subStr := str1[:i]
			if subStrSameConcatenated(subStr, str1) && subStrSameConcatenated(subStr, str2) {
				return subStr
			}
		}
	}

	return ""
}

func subStrSameConcatenated(subStr, word string) bool {
	str := ""
	for i := 0; i < (len(word) / len(subStr)); i++ {
		str += subStr
	}
	return str == word
}

// OPTIMAL SOLUTION using GCD - Greatest Common Divisor
func gcdOfStringsOpt(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLen := gcd(len(str1), len(str2))
	return str1[:gcdLen]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
