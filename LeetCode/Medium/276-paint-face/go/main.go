package main

import "fmt"

func main() {
	fmt.Println(numWays(2, 3))
	fmt.Println(numWays(3, 3))
	fmt.Println(numWays(4, 3))
	fmt.Println(numWays(3, 2))
}

/*
  https://www.youtube.com/watch?v=_E5dTlV8-Ss

  0 pickets = 0
  1 picket = colours
  2 pickets = Same + Different. Same = colours and Different = colours * (colours-1)
  3+ pickets = Same + Different
            Same = prev different
            Different = (prevSame + prevDiff) * (colours-1)
*/

func numWays(pickets int, colours int) int {
	if pickets == 0 {
		return 0
	}
	if pickets == 1 {
		return colours
	}

	same := colours
	diff := (colours * (colours - 1))

	if pickets == 2 {
		return same + diff
	}

	prevSame, prevDiff := same, diff

	for post := 3; post <= pickets; post++ {
		curSame := prevDiff
		curDiff := (prevSame + prevDiff) * (colours - 1)
		prevSame, prevDiff = curSame, curDiff
	}

	return prevSame + prevDiff
}
