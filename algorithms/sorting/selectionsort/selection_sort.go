package selectionsort

func SelectionSort(nums []int) {
	n := len(nums)
	for i := range n - 1 {
		indexMin := i

		for j := i + 1; j < n; j++ {
			if nums[j] < nums[indexMin] {
				indexMin = j
			}
		}

		if nums[i] > nums[indexMin] {
			nums[i], nums[indexMin] = nums[indexMin], nums[i]
		}
	}
}
