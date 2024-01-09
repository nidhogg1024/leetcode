package go_code

import "testing"

func Test_merge(t *testing.T) {
	merge([]int{-1, 0, 0, 3, 3, 3, 0, 0, 0}, 6, []int{1, 2, 2}, 3)
}
