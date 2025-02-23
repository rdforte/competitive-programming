package vote_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rdforte/competitive-programming/random/vote"
)

func Test_Vote_SuccessCalculatedVotes(t *testing.T) {
	taleTest := []struct {
		name      string
		votes     []string
		wantOrder []string
	}{
		{
			name:      "should return order A B C",
			votes:     []string{"A", "B", "A", "B", "C"},
			wantOrder: []string{"A", "B", "C"},
		},
		{
			name:      "should return order A",
			votes:     []string{"A"},
			wantOrder: []string{"A"},
		},
		{
			name:      "should return order A",
			votes:     []string{"A", "B", "A", "A", "B", "B", "B", "A"},
			wantOrder: []string{"B", "A"},
		},
		{
			name:      "should return order A",
			votes:     []string{"C", "B", "A"},
			wantOrder: []string{"C", "B", "A"},
		},
	}

	for _, tt := range taleTest {
		t.Run(tt.name, func(t *testing.T) {
			gotWinnindOrder := vote.CalculateVoteOrder(tt.votes)
			assert.Equal(t, tt.wantOrder, gotWinnindOrder)
		})
	}
}
