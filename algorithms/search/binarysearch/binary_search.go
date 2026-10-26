package binarysearch

func BinarySearch(nums []int, target int) int {
	n := len(nums)
	if n <= 0 {
		return -1
	}
	l, r := 0, n-1

	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}
