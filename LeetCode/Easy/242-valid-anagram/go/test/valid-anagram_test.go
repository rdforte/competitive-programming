package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	anagrams "github.com/rdforte/competitive-programming/LeetCode/Easy/242-valid-anagram/go"
)

func TestShouldReturnTrueIfValidAnagram(t *testing.T) {
	testCases := []struct {
		caseOne string
		caseTwo string
	}{
		{"a", "a"},
		{"ab", "ba"},
	}

	for _, tc := range testCases {
		isAnagram := anagrams.IsAnagram(tc.caseOne, tc.caseTwo)
		assert.True(t, isAnagram)
	}
}

func TestShouldReturnFalseWhenNotValidAnagram(t *testing.T) {
	testCases := []struct {
		caseOne string
		caseTwo string
	}{
		{"a", "b"},
		{"b", "a"},
		{"aa", "a"},
	}

	for _, tc := range testCases {
		isAnagram := anagrams.IsAnagram(tc.caseOne, tc.caseTwo)
		assert.False(t, isAnagram)
	}
}
