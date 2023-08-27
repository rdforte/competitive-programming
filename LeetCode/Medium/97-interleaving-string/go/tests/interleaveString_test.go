package interleaveString_test

import (
	interleave "github.com/rdforte/competitive-programming/LeetCode/Medium/97-interleaving-string/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInterleaveTrue(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	isInterleave := interleave.IsInterleave(s1, s2, s3)
	expected := true

	assert.Equal(t, expected, isInterleave)
}

func TestIsInterleaveFalse(t *testing.T) {
	s1 := "bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa"
	s2 := "babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab"
	s3 := "babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab"
	isInterleave := interleave.IsInterleave(s1, s2, s3)
	expected := false

	assert.Equal(t, expected, isInterleave)
}

func BenchmarkIsInterleave(b *testing.B) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	for i := 0; i < b.N; i++ {
		interleave.IsInterleave(s1, s2, s3)
	}
}
