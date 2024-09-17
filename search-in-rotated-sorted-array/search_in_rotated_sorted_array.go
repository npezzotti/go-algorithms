package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2

		if target == nums[mid] {
			return mid
		}

		// left section
		if nums[l] <= nums[mid] {
			if target < nums[l] || target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			// right section
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}
