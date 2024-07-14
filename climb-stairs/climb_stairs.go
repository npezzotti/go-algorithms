package climb_stairs

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct
ways can you climb to the top?
*/

// memoization

var mem map[int]int = map[int]int{}

func climbStairsRecursive(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if val, ok := mem[n]; ok {
		return val
	}

	res := climbStairs(n-1) + climbStairs(n-2)
	mem[n] = res
	return res
}

func climbStairs(n int) int {
	res := 0

	last := 0
	secondLast := 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			res = 1
		} else if i == 2 {
			res = 2
		} else {
			res = last + secondLast
		}

		secondLast = last
		last = res
	}

	return res
}
