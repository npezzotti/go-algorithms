package pow

// Based on formulas = a^n = (a^(n/2))^2 (even
// exponent and a^n = a*(a^((n-1)/2))^2 (odd exponent)
func myPow(x float64, n int) float64 {
	pow := func(x float64, n int) float64 {
		ans := float64(1)
		for n > 0 {
			// check if exponent is odd
			if n&1 == 1 {
				ans *= x
			}

			x *= x
			n >>= 1
		}

		return ans
	}

	if n < 0 {
		return 1 / pow(x, -n)
	}

	return pow(x, n)
}
