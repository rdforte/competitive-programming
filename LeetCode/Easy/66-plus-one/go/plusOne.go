package plusOne

func PlusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; carry > 0 && i >= 0; i-- {
		digit := digits[i] + carry
		if digit == 10 {
			digits[i] = 0
		} else {
			digits[i] = digit
			carry = 0
		}
	}

	if carry == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}
