package vote

import (
	"sort"
)

func RankTeamsV2(ranks []string) string {
	if len(ranks) == 0 {
		return ""
	}

	ranking := [26][26]int{} // space O(1)

	// v = voters
	// t = teams
	// O(rv)
	for _, rank := range ranks {
		for i := range rank {
			name := rank[i] - 'A'
			ranking[name][i]++
		}
	}

	teams := []byte(ranks[0])

	// t*log*t
	// t^2*log*t
	sort.Slice(teams, func(i, j int) bool {
		teamI := teams[i] - 'A'
		teamJ := teams[j] - 'A'

		for i := range len(teams) {
			if ranking[teamI][i] > ranking[teamJ][i] {
				return true
			}

			if ranking[teamI][i] < ranking[teamJ][i] {
				return false
			}
		}

		return teamI < teamJ
	})

	return string(teams)
}

func RankTeams(ranks []string) string {
	if len(ranks) == 0 {
		return ""
	}

	ranking := make(map[byte][]int, len(ranks))
	for _, rank := range ranks {
		for i := range rank {
			name := rank[i]

			if _, ok := ranking[name]; !ok {
				ranking[name] = make([]int, len(rank))
			}
			ranking[name][i]++
		}
	}

	teams := []byte(ranks[0])

	sort.Slice(teams, func(i, j int) bool {
		teamI := teams[i]
		teamJ := teams[j]

		for i := range len(teams) {
			if ranking[teamI][i] > ranking[teamJ][i] {
				return true
			}

			if ranking[teamI][i] < ranking[teamJ][i] {
				return false
			}
		}

		return teamI < teamJ
	})

	return string(teams)
}
