package removeDuplicates

func RemoveDuplicates(s string) string {
	var stack []rune

	for _, r := range s {
		if len(stack) > 0 && r == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}

	return string(stack)
}
