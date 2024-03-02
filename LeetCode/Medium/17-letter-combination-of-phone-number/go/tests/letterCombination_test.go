package letterCombination_test

import (
	letterCombination "github.com/rdforte/competitive-programming/LeetCode/Medium/17-letter-combination-of-phone-number/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinationsOfPhoneNumber(t *testing.T) {
	t.Run("should return all possible letter combinations that the number could represent 23", func(t *testing.T) {
		digits := "23"
		got := letterCombination.LetterCombinations(digits)
		want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
		assert.ElementsMatch(t, got, want)
	})

	t.Run("should return all possible letter combinations that the number could represent ''", func(t *testing.T) {
		digits := ""
		got := letterCombination.LetterCombinations(digits)
		want := []string{}
		assert.ElementsMatch(t, got, want)
	})

	t.Run("should return all possible letter combinations that the number could represent 2", func(t *testing.T) {
		digits := "2"
		got := letterCombination.LetterCombinations(digits)
		want := []string{"a", "b", "c"}
		assert.ElementsMatch(t, got, want)
	})
}
