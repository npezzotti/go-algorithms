package find_duplicate

import "fmt"

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		fmt.Println(slow, fast)

		if slow == fast {
			break
		}
	}

	slow = nums[0]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
