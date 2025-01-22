package largest_number

import (
	"slices"
	"strconv"
)

func largestNumber(nums []int) string {
	slices.SortFunc(nums, func(a, b int) int {
		aStr := strconv.Itoa(a)
		bStr := strconv.Itoa(b)

		num1, _ := strconv.Atoi(aStr + bStr)
		num2, _ := strconv.Atoi(bStr + aStr)

		if num1 > num2 {
			return -1
		} else if num1 < num2 {
			return 1
		}

		return 0
	})

	var res string
	for _, num := range nums {
		res += strconv.Itoa(num)
	}

	// if largest num is 0, all input nums are 0
	if res[0] == '0' {
		return "0"
	}

	return res
}
