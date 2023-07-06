package removeDuplicates_test

import (
	removeDuplicates "github.com/rdforte/competitive-programming/LeetCode/Easy/1047-remove-all-adjacent-duplicates-in-string/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tableTest := []struct {
		input        string
		expectOutput string
	}{
		{"abbaca", "ca"},
		{"azxxzy", "ay"},
		{"aaaaaaaa", ""},
		{"aaaaaaaaa", "a"},
	}

	for _, test := range tableTest {
		t.Run(test.input, func(t *testing.T) {
			gotOutput := removeDuplicates.RemoveDuplicates(test.input)
			assert.Equal(t, test.expectOutput, gotOutput)
		})
	}
}
