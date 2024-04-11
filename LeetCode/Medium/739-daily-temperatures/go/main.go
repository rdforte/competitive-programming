package main

func main() {

}

func dailyTemperatures(temperatures []int) []int {
	var stack [][]int
	stack = append(stack, []int{temperatures[0], 0})

	res := make([]int, len(temperatures))

	for i := 1; i < len(temperatures); i++ {
		for len(stack) > 0 && stack[len(stack)-1][0] < temperatures[i] {
			lastTempIndex := stack[len(stack)-1][1]
			res[lastTempIndex] = i - lastTempIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, []int{temperatures[i], i})
	}

	return res
}
