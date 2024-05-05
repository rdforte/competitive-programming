package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumOperationsToMakeKPeriodic("leetcodeleet", 4))
	fmt.Println(minimumOperationsToMakeKPeriodic("leetleetleet", 4))
	fmt.Println(minimumOperationsToMakeKPeriodic("leetcoleet", 2))
}

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	maxRepWord := make(map[string]float64)

	var maxRep float64
	for i, j := 0, k-1; j < len(word); i, j = i+k, j+k {
		maxRepWord[word[i:j+1]]++
		maxRep = math.Max(maxRepWord[word[i:j+1]], maxRep)
	}

	return (len(word) - int(maxRep)*k) / k
}
