package main_test

import (
	"testing"

	validAnagram "github.com/rdforte/competitive-programming/LeetCode/Easy/242-valid-anagram/go"
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
		isAnagram := validAnagram.IsAnagram(tc.caseOne, tc.caseTwo)

		if !isAnagram {
			t.Errorf("got %v and wanted %v", isAnagram, true)
		}
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
		isAnagram := validAnagram.IsAnagram(tc.caseOne, tc.caseTwo)

		if isAnagram {
			t.Errorf("got %v and wanted %v", isAnagram, false)
		}
	}
}
