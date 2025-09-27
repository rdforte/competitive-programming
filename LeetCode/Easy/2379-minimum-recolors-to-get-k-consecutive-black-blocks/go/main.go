package main

func main() {
	blocks := "WBWBBBW"
	println(minimumRecolors(blocks, 2))
	blocks = "WBBWWBBWBW"
	println(minimumRecolors(blocks, 7))
	blocks = "BWWWBB"
	println(minimumRecolors(blocks, 6))
}

func minimumRecolors(blocks string, k int) int {
	minFlips := len(blocks)
	curFlips := 0

	for left, right := 0, 0; right < len(blocks); right++ {
		if right-left == k {
			if blocks[left] == 'W' {
				curFlips--
			}
			left++
		}

		if blocks[right] == 'W' {
			curFlips++
		}

		if right-left == k-1 {
			minFlips = min(minFlips, curFlips)
		}
	}

	return minFlips
}
