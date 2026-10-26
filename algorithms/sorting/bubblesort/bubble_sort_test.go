package bubblesort

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("easy case", func(t *testing.T) {
		sample1 := []int{4, 2, 7, 1, 9, 3}
		expected := []int{1, 2, 3, 4, 7, 9}

		BubbleSort(sample1)

		if !slices.Equal(sample1, expected) {
			t.Error("easy case not ordered\n")
		}
	})
	t.Run("normal case", func(t *testing.T) {
		sample2 := []int{2, 0, 67, 89, 2, 1, 999, 1001, 34}
		expected := []int{0, 1, 2, 2, 34, 67, 89, 999, 1001}

		BubbleSort(sample2)

		if !slices.Equal(sample2, expected) {
			t.Error("normal case not ordered\n")
		}
	})
	t.Run("best case", func(t *testing.T) {
		sample3 := []int{1, 1, 2, 3, 5, 8, 13, 21}
		expected := []int{1, 1, 2, 3, 5, 8, 13, 21}

		BubbleSort(sample3)

		if !slices.Equal(sample3, expected) {
			t.Error("best case not ordered\n")
		}
	})
	t.Run("worst case", func(t *testing.T) {
		sample4 := []int{999, 998, 997, 996, 995, 994, 993, 992, 991, 990}
		expected := []int{990, 991, 992, 993, 994, 995, 996, 997, 998, 999}

		BubbleSort(sample4)

		if !slices.Equal(sample4, expected) {
			t.Error("worst case not ordered\n")
		}
	})
}
