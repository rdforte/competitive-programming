package main

import "fmt"

func main() {
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}) == 4)
	fmt.Println(totalFruit([]int{1, 2, 1}) == 3)
	fmt.Println(totalFruit([]int{0, 1, 2, 2}) == 3)
	fmt.Println(totalFruit([]int{0, 0, 0, 0}) == 4)

	fmt.Println("Improved Space")
	fmt.Println(totalFruitImprovedSpace([]int{1, 2, 3, 2, 2}) == 4)
	fmt.Println(totalFruitImprovedSpace([]int{1, 2, 1}) == 3)
	fmt.Println(totalFruitImprovedSpace([]int{0, 1, 2, 2}) == 3)
	fmt.Println(totalFruitImprovedSpace([]int{0, 0, 0, 0}) == 4)
}

// Sliding Window
// Time: O(n)
// Space: O(n) keep track of picked fruits
func totalFruit(fruits []int) int {
	res, baskets := 0, 0
	picked := make(map[int]int)

	for l, r := 0, 0; r < len(fruits); r++ {
		if v, ok := picked[fruits[r]]; !ok || v == 0 {
			baskets++
		}
		picked[fruits[r]]++

		for baskets > 2 {
			picked[fruits[l]]--
			if v, _ := picked[fruits[l]]; v == 0 {
				baskets--
			}
			l++
		}
		res = max(res, (r-l)+1)
	}

	return res
}

// Sliding Window - Optimised
// Time: O(n)
// Space: O(1) keep track of at most 3 picked fruits.
// In this case here because we delete picked items from our map we can use its
// length to work out the number of baskets.
func totalFruitImprovedSpace(fruits []int) int {
	res := 0
	picked := make(map[int]int, 3) // at most types of picked fruit = 3

	for l, r := 0, 0; r < len(fruits); r++ {
		picked[fruits[r]]++

		for len(picked) > 2 {
			picked[fruits[l]]--
			if v, _ := picked[fruits[l]]; v == 0 {
				delete(picked, fruits[l])
			}
			l++
		}
		res = max(res, (r-l)+1)
	}

	return res
}
