package groupAnagrams

import "sort"

func GroupAnagrams(strs []string) [][]string {
	anagrams := map[string][]string{}

	for _, str := range strs {
		strSorted := orderStringRunesAscending(str)
		anagrams[strSorted] = append(anagrams[strSorted], str)
	}

	groupedAnagrams := [][]string{}

	for _, group := range anagrams {
		groupedAnagrams = append(groupedAnagrams, group)
	}

	return groupedAnagrams
}

func orderStringRunesAscending(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
