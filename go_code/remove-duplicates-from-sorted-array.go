package go_code

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/

func removeDuplicates(nums []int) int {
	k := 1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[i+1] = nums[j]
				k++
				break
			}
			if j == len(nums)-1 {
				return k
			}
		}
	}
	return k
}

func removeDuplicatesV2(nums []int) int {
	slow, fast := 1, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
