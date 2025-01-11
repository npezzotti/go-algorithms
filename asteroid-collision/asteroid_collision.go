package asteroid_collision

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, asteroid := range asteroids {
		for len(stack) > 0 && asteroid < 0 && stack[len(stack)-1] > 0 {
			if abs(asteroid) > stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else if abs(asteroid) == stack[len(stack)-1] {
				asteroid = 0
				stack = stack[:len(stack)-1]
				break
			} else {
				asteroid = 0
				break
			}
		}

		if asteroid != 0 {
			stack = append(stack, asteroid)
		}
	}

	return stack
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// func asteroidCollision(asteroids []int) []int {
// 	stack := make([]int, 0)
// 	for _, asteroid := range asteroids {
// 		if asteroid > 0 || (asteroid < 0 && (len(stack) == 0 || stack[len(stack)-1] < 0)) {
// 			stack = append(stack, asteroid)
// 		} else {
// 			for len(stack) > 0 && stack[len(stack)-1] > 0 {
// 				poppedAsteroid := stack[len(stack)-1]
// 				if abs(asteroid) >= abs(poppedAsteroid) {
// 					stack = stack[:len(stack)-1]

// 					if abs(asteroid) == abs(poppedAsteroid) {
// 						break
// 					}

// 					if len(stack) == 0 || (len(stack) > 0 && stack[len(stack)-1] < 0) {
// 						stack = append(stack, asteroid)
// 					}
// 				} else {
// 					break
// 				}
// 			}
// 		}
// 	}

// 	return stack
// }
