package counting_bits

func countBitsDp(n int) []int {
	res := make([]int, n+1)

	for i := 1; i < len(res); i++ {
		res[i] = res[i>>1] + (i & 1)
	}

	return res
}

func countBitsNaive(n int) []int {
	res := make([]int, n+1)
	for i := 0; i < len(res); i++ {
		var count int
		n := i
		for n > 0 {
			count += n & 1
			n >>= 1
		}
		res[i] = count
	}
	return res
}
