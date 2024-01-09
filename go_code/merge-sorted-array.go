package go_code

// https://leetcode.cn/problems/merge-sorted-array/description/

// 逆向双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; (p1 >= 0 || p2 >= 0) && tail >= 0; tail-- {
		if p1 == -1 {
			nums1[tail] = nums2[p2]
			p2--
		} else if p2 == -1 {
			nums1[tail] = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			nums1[tail] = nums1[p1]
			p1--
		} else {
			nums1[tail] = nums2[p2]
			p2--
		}
	}
}

// 正向双指针
func mergeV2(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, len(nums1))
	for p1, p2, cur := 0, 0, 0; p1 < m || p2 < n; cur++ {
		if p1 == m {
			res[cur] = nums2[p2]
			p2++
		} else if p2 == n {
			res[cur] = nums1[p1]
			p1++
		} else if nums1[p1] < nums2[p2] {
			res[cur] = nums1[p1]
			p1++
		} else {
			res[cur] = nums2[p2]
			p2++
		}
	}
	copy(nums1, res)
}
