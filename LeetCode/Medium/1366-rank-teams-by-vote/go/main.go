package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rankTeamsV2([]string{"ABC", "ACB", "ABC", "ACB", "ACB"}))
	fmt.Println(rankTeamsV2([]string{"WXYZ", "XYZW"}))
}

/*
  By laying out the postions like below it allows us to do an easy comparison.

    0   1    2
--------------
A - 5   0    0  - A came first 5 times.
B - 0   2    3  - B came first 0 times, 2nd 2 times, 3rd 3 times.
C - 0   3    2  - C came first 0 times, 2nd 3 times, 3rd 2 times.

Now when we sort we compare each characters positions. For example when sorting all characters,
when B is compared to C we look at the first place values and they are the same so we continue.
We then look at 2nd place values. C came second more times than B so it should come before B.
*/

func rankTeamsV2(votes []string) string {
	numContestents := len(votes[0])
	ranking := make([][]int, 0, 26)
	for range 26 {
		ranking = append(ranking, make([]int, numContestents))
	}

	for _, v := range votes {
		for i, c := range v {
			ranking[c-'A'][i]++
		}
	}

	res := []byte(votes[0]) // lets sort the characters

	sort.Slice(res, func(i, j int) bool {
		char1 := res[i]
		char2 := res[j]

		// compare each characters positions
		for i := range numContestents {
			if ranking[char1-'A'][i] == ranking[char2-'A'][i] {
				// same rank
				continue
			}
			return ranking[char1-'A'][i] > ranking[char2-'A'][i]
		}

		return char1 < char2
	})

	return string(res)
}

// Because there is a max of 26 characters and each is unique there can only be a max of 26 positions so we
// can create a 2D array of 26 * 26.
func rankTeams(votes []string) string {
	var count [26][26]int
	for _, v := range votes {
		for i, ch := range v {
			count[ch-'A'][i]++
		}
	}
	cand := []byte(votes[0])
	sort.Slice(cand, func(i, j int) bool {
		a := count[cand[i]-'A']
		b := count[cand[j]-'A']
		for rank := 0; rank < 26; rank++ {
			if a[rank] == b[rank] {
				continue
			}
			return a[rank] > b[rank]
		}
		return cand[i] < cand[j]
	})
	return string(cand)
}
