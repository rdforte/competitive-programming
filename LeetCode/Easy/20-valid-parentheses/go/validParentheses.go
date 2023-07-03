package validParentheses

func IsValid(s string) bool {
	var stack []string
	brkB := map[string]string{
		"]": "[",
		")": "(",
		"}": "{",
	}

	brkF := map[string]struct{}{
		"[": {},
		"{": {},
		"(": {},
	}

	for i := 0; i < len(s); i++ {
		if _, ok := brkF[string(s[i])]; ok {
			stack = append(stack, string(s[i]))
			continue
		}

		if v, ok := brkB[string(s[i])]; ok && len(stack) > 0 {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if v == b {
				continue
			}
		}
		return false
	}

	return len(stack) == 0
}
