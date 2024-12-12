package maximum_product_subarray

import "fmt"

func maxProductNaive(nums []int) int {
	var res int = nums[0]

	for i := range nums {
		mult := 1
		for j := i; j < len(nums); j++ {
			mult *= nums[j]

			res = max(res, mult)
		}
	}

	return res
}

func maxProductDP(nums []int) int {
	curMin, curMax, maxProd := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		curMaxTemp := curMax
		curMax = max(nums[i], curMax*nums[i], curMin*nums[i])
		fmt.Println("curMax", curMax)
		curMin = min(nums[i], curMaxTemp*nums[i], curMin*nums[i])
		fmt.Println("curMin", curMin)

		if curMax > maxProd {
			maxProd = curMax
		}

		fmt.Println("maxProd", maxProd)
	}

	return maxProd
}
