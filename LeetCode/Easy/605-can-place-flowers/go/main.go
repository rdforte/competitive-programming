package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		noLeft := i == 0 || flowerbed[i-1] == 0
		noRight := i == len(flowerbed)-1 || flowerbed[i+1] == 0
		if flowerbed[i] == 0 && noLeft && noRight {
			flowerbed[i] = 1
			n--
		}
	}

	return n == 0
}
