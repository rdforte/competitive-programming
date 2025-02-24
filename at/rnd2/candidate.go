/*
Atlassian rnd2
* Given a list of ballots where each ballot can contain at most 3 candidates
* Rank the candidates based on votes. A candidate at pos 1 has a weight 3, candidate at 2 has
weight 2 and candidate at pos 3 has weight 1
*/
package candidate

import "sort"

func GetResult(ballots [][]string) []string {
	ballotResults := make(map[string][]int)
	candidateWeights := make(map[string]int)
	candidateLastBallotPos := make(map[string]int)

	for bPos, ballot := range ballots {
		for i, candidate := range ballot {
			if _, ok := ballotResults[candidate]; !ok {
				ballotResults[candidate] = make([]int, 3)
			}
			ballotResults[candidate][i]++
			candidateWeights[candidate] += (3 - i)
			candidateLastBallotPos[candidate] = (bPos * i) + i
		}
	}

	candidates := make([]string, 0, len(ballotResults))
	for candidate := range ballotResults {
		candidates = append(candidates, candidate)
	}

	sort.Slice(candidates, func(i, j int) bool {
		candidateI := candidates[i]
		candidateJ := candidates[j]

		if candidateWeights[candidateI] == candidateWeights[candidateJ] {
			// check the positions

			for i := range 3 {
				if ballotResults[candidateI][i] == ballotResults[candidateJ][i] {
					continue
				}

				if ballotResults[candidateI][i] > ballotResults[candidateJ][i] {
					return true
				}

				return false
			}

			return candidateLastBallotPos[candidateI] < candidateLastBallotPos[candidateJ]
		}

		return candidateWeights[candidateI] > candidateWeights[candidateJ]
	})

	return candidates
}
