package candidate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	candidate "github.com/rdforte/competitive-programming/at/rnd2"
)

func TestCandidateBallot(t *testing.T) {
	ballots := [][]string{
		{"c1", "c2", "c3"},
		{"c1", "c2", "c3"},
	}

	got := candidate.GetResult(ballots)
	want := []string{"c1", "c2", "c3"}
	assert.Equal(t, want, got)
}

func TestCandidateBallot_DifferentOrderCandidates(t *testing.T) {
	ballots := [][]string{
		{"c1", "c3", "c2"},
		{"c1", "c2", "c3"},
		{"c1", "c3", "c2"},
	}

	got := candidate.GetResult(ballots)
	want := []string{"c1", "c3", "c2"}
	assert.Equal(t, want, got)
}

func TestCandidateBallot_DifferentOrderCandidatesWithFallback(t *testing.T) {
	ballots := [][]string{
		{"c1", "c3"},
		{"c2", "c1"},
		{"c2", "c1"},
		{"c2", "c1"},
	}

	got := candidate.GetResult(ballots)
	want := []string{"c2", "c1", "c3"}
	assert.Equal(t, want, got)
}

func TestCandidateBallot_DifferentOrderCandidatesWithSamePositions(t *testing.T) {
	ballots := [][]string{
		{"c1"}, // 0
		{"c2"}, // 1
	}

	got := candidate.GetResult(ballots)
	want := []string{"c1", "c2"}
	assert.Equal(t, want, got)
}

func TestCandidateBallot_DifferentOrderCandidatesWithSamePositions2(t *testing.T) {
	ballots := [][]string{
		{"c1", "c2"},
		{"c2", "c1"}, // c2 reaches the weight of 5 before c1 does
	}

	got := candidate.GetResult(ballots)
	want := []string{"c2", "c1"}
	assert.Equal(t, want, got)
}
