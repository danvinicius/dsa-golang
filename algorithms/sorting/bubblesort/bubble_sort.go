package bubblesort

func BubbleSort(nums []int) {
	for range len(nums) - 1 {
		for j := range len(nums) - 1 { // o(n^2)
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
