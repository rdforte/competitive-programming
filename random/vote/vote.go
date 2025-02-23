package vote

import "sort"

func CalculateVoteOrder(votes []string) []string {
	// v
	candidateVoteCount := make(map[string]int)
	for _, v := range votes {
		candidateVoteCount[v]++
	}

	// c
	candidates := make([]string, 0, len(candidateVoteCount))
	for name := range candidateVoteCount {
		candidates = append(candidates, name)
	}

	// c = unique candidates
	// v = total num of votes
	// v * c log c
	sort.Slice(candidates, func(i, j int) bool {
		nameI := candidates[i]
		nameJ := candidates[j]
		if candidateVoteCount[nameI] == candidateVoteCount[nameJ] {
			nameIVotes := candidateVoteCount[nameI]
			nameJVotes := candidateVoteCount[nameJ]

			for _, name := range votes {
				if name == nameI {
					nameIVotes--
				}

				if name == nameJ {
					nameJVotes--
				}

				if nameIVotes == 0 {
					return true
				}

				if nameJVotes == 0 {
					return false
				}
			}
		}

		return candidateVoteCount[nameI] > candidateVoteCount[nameJ]
	})

	return candidates
}

// Total time = O(vc*log*c + v + c)
