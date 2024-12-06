package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Brute Force ----------------------------------------")
	fmt.Println(GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(GroupAnagrams([]string{""}))
	fmt.Println(GroupAnagrams([]string{"a"}))
	fmt.Println("Optimized ----------------------------------------")
	fmt.Println(GroupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(GroupAnagrams2([]string{""}))
	fmt.Println(GroupAnagrams2([]string{"a"}))
}

// Space = hold each str in memory = O(n)
// Time = range each string = O(n) *
// sort each string k = k.log.k
// loop each group at the end n
// therefore  n * k.log.k + n
// = n * k.log.k
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string, len(strs))

	for _, str := range strs {
		b := SortedByte(str)
		sort.Sort(b)
		groups[string(b)] = append(groups[string(b)], str)
	}

	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		res = append(res, g)
	}

	return res
}

type SortedByte []byte

func (a SortedByte) Len() int { return len(a) }

func (a SortedByte) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a SortedByte) Less(i, j int) bool { return a[i] < a[j] }

// OPTIMIZED -----------------------------------------------------------------------------------
// Space = hold each str in memory = O(n)
// Time = loop through each string n times
// for each string k we loop through its characters.
// therefore Time = nk
func GroupAnagrams2(strs []string) [][]string {
	groups := make(map[string][]string, len(strs))

	for _, str := range strs {
		count := make([]int, 26)
		for i := 0; i < len(str); i++ {
			count[str[i]-'a']++
		}
		key := buildKey(count)
		groups[key] = append(groups[key], str)
	}

	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		res = append(res, g)
	}

	return res
}

func buildKey(count []int) string {
	builder := strings.Builder{}

	for _, c := range count {
		builder.WriteString("#")
		builder.WriteString(string(c))
	}

	return builder.String()
}
