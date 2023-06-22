package twoSum

func TwoSum(nums []int, target int) []int {
	mp := map[int]int{}

	for i := 0; i < len(nums); i++ {
		neededToEqualTarget := target - nums[i]
		if index, ok := mp[neededToEqualTarget]; ok {
			return []int{index, i}
		}
		mp[nums[i]] = i
	}

	return []int{}
}
