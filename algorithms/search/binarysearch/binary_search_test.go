package binarysearch

import "testing"

type Input struct {
	nums   []int
	target int
}

type TestCase struct {
	inputs   Input
	expected int
}

func TestBinarySearch(t *testing.T) {
	var table []TestCase = []TestCase{
		{
			inputs: Input{
				nums:   []int{2, 2, 4, 7, 11, 21, 23, 30, 41},
				target: 30,
			},
			expected: 7,
		},
		{
			inputs: Input{
				nums:   []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89},
				target: 2,
			},
			expected: 2,
		},
		{
			inputs: Input{
				nums:   []int{13, 14, 17, 22, 67, 69},
				target: 68,
			},
			expected: -1,
		},
	}

	t.Run("easy case", func(t *testing.T) {
		for _, testCase := range table {
			result := BinarySearch(testCase.inputs.nums, testCase.inputs.target)

			if result != testCase.expected {
				t.Errorf("Target not found as expected")
			}
		}
	})
}
