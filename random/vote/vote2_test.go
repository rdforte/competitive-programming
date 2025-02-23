package vote_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rdforte/competitive-programming/random/vote"
)

func TestVote2_RankTeams(t *testing.T) {
	tableTest := []struct {
		name     string
		teams    []string
		wantRank string
	}{
		{
			name:     "should produce rank ACB",
			teams:    []string{"ABC", "ACB", "ABC", "ACB", "ACB"},
			wantRank: "ACB",
		},
		{
			name:     "should produce rank XWYZ",
			teams:    []string{"WXYZ", "XYZW"},
			wantRank: "XWYZ",
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			gotRank := vote.RankTeams(tt.teams)
			assert.Equal(t, tt.wantRank, gotRank)
		})
	}
}
