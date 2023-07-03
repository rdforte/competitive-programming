package validParentheses_test

import (
	validParentheses "github.com/rdforte/competitive-programming/LeetCode/Easy/20-valid-parentheses/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tableTests := []struct {
		input      string
		wantOutput bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"{[]}", true},
		{"[]", true},
		{"[[]]", true},
		{"[[[]]]", true},
		{"[[{}]]", true},
		{"[([{}])]", true},
		{"[[[]]", false},
		{"][[]]]", false},
		{"[[]]]", false},
	}

	for _, test := range tableTests {
		got := validParentheses.IsValid(test.input)
		assert.Equal(t, test.wantOutput, got)
	}
}
