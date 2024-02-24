package happyNumber

import "math"

func Ishappy(n int) bool {
	seen := make(map[int]struct{})
	for n != 1 {
		num := 0
		for cur := n; cur > 0; cur /= 10 {
			remainder := cur % 10
			num += int(math.Pow(float64(remainder), 2))
		}
		if _, ok := seen[num]; ok {
			return false
		}
		seen[num] = struct{}{}
		n = num
		// do check to see if same as n, if so break
	}
	return true
}
