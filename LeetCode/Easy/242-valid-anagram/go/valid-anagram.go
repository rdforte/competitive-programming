package validAnagram

import (
	"sort"
)

func IsAnagram(s string, t string) bool {
	return sortStringAscending(s) == sortStringAscending(t)
}

func sortStringAscending(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
