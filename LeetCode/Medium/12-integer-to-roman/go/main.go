package main

import "fmt"

func main() {
	fmt.Println(intToRoman(0) == "")
	fmt.Println(intToRoman(1) == "I")
	fmt.Println(intToRoman(2) == "II")
	fmt.Println(intToRoman(3) == "III")
	fmt.Println(intToRoman(4) == "IV")
	fmt.Println(intToRoman(5) == "V")
	fmt.Println(intToRoman(6) == "VI")
	fmt.Println(intToRoman(7) == "VII")
	fmt.Println(intToRoman(8) == "VIII")
	fmt.Println(intToRoman(9) == "IX")
	fmt.Println(intToRoman(10) == "X")
	fmt.Println(intToRoman(11) == "XI")
	fmt.Println(intToRoman(58) == "LVIII")
	fmt.Println(intToRoman(50) == "L")
	fmt.Println(intToRoman(3749) == "MMMDCCXLIX")
	fmt.Println(intToRoman(1994) == "MCMXCIV")
}

func intToRoman(num int) string {
	places := [][]string{
		{"I", "V"},
		{"X", "L"},
		{"C", "D"},
		{"M", "E"},
	}

	res := ""
	idx := 0
	for num > 0 {
		digit := num % 10
		num /= 10

		repeates := digit % 5
		place := digit / 5

		switch repeates {
		case 0:
			if digit > 0 {
				res = places[idx][1] + res
			}
		case 1:
			if place == 1 {
				res = places[idx][1] + places[idx][0] + res
			} else {
				res = places[idx][0] + res
			}
		case 2:
			if place == 1 {
				res = places[idx][1] + places[idx][0] + places[idx][0] + res
			} else {
				res = places[idx][0] + places[idx][0] + res
			}
		case 3:
			if place == 1 {
				res = places[idx][1] + places[idx][0] + places[idx][0] + places[idx][0] + res
			} else {
				res = places[idx][0] + places[idx][0] + places[idx][0] + res
			}
		case 4:
			if place < 1 {
				res = places[idx][0] + places[idx][1] + res
			} else {
				res = places[idx][0] + places[idx+1][0] + res
			}
		}
		idx++
	}

	return res
}

// The cleanest way to go about it in code is to have 4 separate arrays; one for each place value.
// Then, extract the digits, look up their symbols in the relevant array, and append them all together.
func intToRomanOptimal(num int) string {
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{
		"",
		"C",
		"CC",
		"CCC",
		"CD",
		"D",
		"DC",
		"DCC",
		"DCCC",
		"CM",
	}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}
