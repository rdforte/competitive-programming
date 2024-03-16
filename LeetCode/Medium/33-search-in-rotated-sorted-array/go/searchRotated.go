package searchRotated

func Search(nums []int, target int) int {
	l, u := 0, len(nums)-1
	for l < u {
		mid := l + ((u - l) / 2)

		if nums[mid] < nums[u] {
			u = mid
		} else {
			l = mid + 1
		}
	}

	if piv := binarySearch(nums, l, len(nums)-1, target); piv != -1 {
		return piv
	}

	if piv := binarySearch(nums, 0, l-1, target); piv != -1 {
		return piv
	}

	return -1
}

func binarySearch(nums []int, l, u, target int) int {
	for l <= u {
		mid := l + (u-l)/2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			l = mid + 1
		} else {
			u = mid - 1
		}
	}

	return -1
}
