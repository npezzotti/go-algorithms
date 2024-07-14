package contains_duplicates

/*
Given an integer array nums, return true if any value
appears at least twice in the array, and return false
if every element is distinct.
*/

func containsDuplicate(nums []int) bool {
	countInts := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := countInts[v]; ok {
			return true
		}
		countInts[v] = struct{}{}
	}

	return false
}
