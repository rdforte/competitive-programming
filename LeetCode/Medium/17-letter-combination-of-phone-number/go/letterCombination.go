package letterCombination

func LetterCombinations(digits string) []string {
	letters := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}

	combinations := []string{}
	combo := []byte{}

	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(digits) {
			if len(combo) > 0 {
				combinations = append(combinations, string(combo))
			}
			return
		}

		digs := letters[digits[i]-'2']

		for _, d := range digs {
			combo = append(combo, d)
			dfs(i + 1)
			combo = combo[:len(combo)-1]
		}
	}

	dfs(0)

	return combinations
}
