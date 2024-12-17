package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt32 % int(1e9) / 1e6)
	fmt.Println(numberToWords(math.MaxInt32))
	fmt.Println("------------------------")
	fmt.Println(buildNumInHundreds(133))
	fmt.Println(buildNumInHundreds(213))
	fmt.Println(buildNumInHundreds(210))
	fmt.Println(buildNumInHundreds(200))
	fmt.Println(buildNumInHundreds(201))
	fmt.Println(buildNumInHundreds(1))
	fmt.Println(buildNumInHundreds(10))
	fmt.Println(buildNumInHundreds(12))
	fmt.Println(buildNumInHundreds(22))
	fmt.Println(buildNumInHundreds(0), "-")
}

var ones = map[int]string{
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}

var teens = map[int]string{
	1: "Eleven",
	2: "Twelve",
	3: "Thirteen",
	4: "Fourteen",
	5: "Fifteen",
	6: "Sixteen",
	7: "Seventeen",
	8: "Eighteen",
	9: "Nineteen",
}

var tens = map[int]string{
	1: "Ten",
	2: "Twenty",
	3: "Thirty",
	4: "Forty",
	5: "Fifty",
	6: "Sixty",
	7: "Seventy",
	8: "Eighty",
	9: "Ninety",
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	res := ""

	n := num / 1e9
	if n > 0 {
		res += ones[n] + " Billion"
		num %= 1e9
	}

	n = num / 1e6
	if n > 0 {
		res += buildNumInHundreds(n) + " Million"
		num %= 1e6
	}

	n = num / 1e3
	if n > 0 {
		res += buildNumInHundreds(n) + " Thousand"
		num %= 1e3
	}

	if len(res) > 0 {
		return strings.TrimSpace(res + buildNumInHundreds(num))
	}

	return strings.TrimSpace(buildNumInHundreds(num))
}

func buildNumInHundreds(num int) string {
	res := ""

	n := num / 100
	if n > 0 {
		res += " " + ones[n] + " Hundred"
		num %= 100
	}

	n = num / 10
	if n == 1 && num == 10 {
		res += " " + tens[n]
		num = 0
	} else if n == 1 && num > 10 {
		res += " " + teens[num%10]
		num = 0
	} else if n > 1 {
		res += " " + tens[n]
		num %= 10
	}

	if num > 0 {
		res += " " + ones[num]
	}

	return res
}
