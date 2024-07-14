package merge_sorted_array

/*
You are given two integer arrays nums1 and nums2, sorted
in non-decreasing order, and two integers m and n, representing
the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing
order.

The final sorted array should not be returned by the function, but
instead be stored inside the array nums1. To accommodate this, nums1
has a length of m + n, where the first m elements denote the elements
that should be merged, and the last n elements are set to 0 and should
be ignored. nums2 has a length of n.
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p3 := m-1, n-1, n+m-1

	for ; p1 >= 0 && p2 >= 0; p3-- {
		if nums1[p1] >= nums2[p2] {
			nums1[p3] = nums1[p1]
			p1--
		} else {
			nums1[p3] = nums2[p2]
			p2--
		}
	}

	if p2 >= 0 {
		copy(nums1[:p3+1], nums2[:p2+1])
	}
}
