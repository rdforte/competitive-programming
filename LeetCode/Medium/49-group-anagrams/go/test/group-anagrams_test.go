package groupAnagrams_test

import (
	"sort"
	"testing"

	anagrams "github.com/rdforte/competitive-programming/LeetCode/Medium/49-group-anagrams/go"
	"github.com/stretchr/testify/assert"
)

func TestShouldGroupAllAnagramsTogether(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedOutput [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{"xxx"}, [][]string{{"xxx"}}},
		{[]string{""}, [][]string{{""}}},
	}

	for _, tc := range testCases {
		output := anagrams.GroupAnagrams(tc.input)

		// we need to sort the outputs anagram groups in order to do a comparison.
		for i := 0; i < len(output); i++ {
			sort.Slice(output[i], func(l, r int) bool {
				return output[i][l] < output[i][r]
			})
		}

		assert.ElementsMatch(t, tc.expectedOutput, output)
	}
}
