package validSudoku_test

import (
	validSudoku "github.com/rdforte/competitive-programming/LeetCode/Medium/36-valid-sudoku/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidSudoku(t *testing.T) {
	tableTest := []struct {
		expectOutput bool
		board        [][]byte
	}{
		{
			expectOutput: true,
			board: [][]byte{
				[]byte("53..7...."),
				[]byte("6..195..."),
				[]byte(".98....6."),
				[]byte("8...6...3"),
				[]byte("4..8.3..1"),
				[]byte("7...2...6"),
				[]byte(".6....28."),
				[]byte("...419..5"),
				[]byte("....8..79"),
			},
		},
		{
			expectOutput: false,
			board: [][]byte{
				[]byte("83..7....,"),
				[]byte("6..195...,"),
				[]byte(".98....6.,"),
				[]byte("8...6...3,"),
				[]byte("4..8.3..1,"),
				[]byte("7...2...6,"),
				[]byte(".6....28.,"),
				[]byte("...419..5,"),
				[]byte("....8..79,"),
			},
		},
	}

	for _, tst := range tableTest {
		got := validSudoku.IsValidSudoku(tst.board)
		assert.Equal(t, tst.expectOutput, got)
	}
}
