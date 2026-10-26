package selectionsort

import (
	"slices"
	"testing"
)

type TestCase struct {
	input    []int
	expected []int
}

func TestSelectionSort(t *testing.T) {

	var table []TestCase = []TestCase{
		{
			input:    []int{4, 2, 7, 1, 9, 3},
			expected: []int{1, 2, 3, 4, 7, 9},
		},
		{
			input:    []int{2, 0, 67, 89, 2, 1, 999, 1001, 34},
			expected: []int{0, 1, 2, 2, 34, 67, 89, 999, 1001},
		},
		{
			input:    []int{1, 1, 2, 3, 5, 8, 13, 21},
			expected: []int{1, 1, 2, 3, 5, 8, 13, 21},
		},
		{
			input:    []int{999, 998, 997, 996, 995, 994, 993, 992, 991, 990},
			expected: []int{990, 991, 992, 993, 994, 995, 996, 997, 998, 999},
		},
	}

	for _, testCase := range table {

		t.Run("easy case", func(t *testing.T) {
			SelectionSort(testCase.input)

			if !slices.Equal(testCase.input, testCase.expected) {
				t.Error("list not ordered as expected\n")
			}
		})
	}
}
