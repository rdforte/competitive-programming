package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.2", "1.10") == -1)
	fmt.Println(compareVersion("1.01", "1.001") == 0)
	fmt.Println(compareVersion("1.0", "1.0.0.0") == 0)
}

// Time complexity:
// len v1 = n and len v2 = m
// Splitting strings is m + n
// we then range the strings with the the loop ranging the longest so max(m, n)
// therefore time = m + n + max(m, n)
// Space = O(m + n)
func compareVersion(version1, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")

	i := 0
	for i < len(v1) && i < len(v2) {
		v1Num, _ := strconv.Atoi(v1[i])
		v2Num, _ := strconv.Atoi(v2[i])

		if v1Num > v2Num {
			return 1
		}

		if v2Num > v1Num {
			return -1
		}
		i++
	}

	// process remaining
	for i < len(v1) {
		v1Num, _ := strconv.Atoi(v1[i])
		if v1Num > 0 {
			return 1
		}
		i++
	}

	for i < len(v2) {
		v2Num, _ := strconv.Atoi(v2[i])
		if v2Num > 0 {
			return -1
		}
		i++
	}

	return 0
}
