package random_pick_with_weight

import (
	"math"
	"testing"
)

func TestPickIndex(t *testing.T) {
	t.Run("single element", func(t *testing.T) {
		nums := []int{1}
		s := Constructor(nums)

		counts := make(map[int]int)
		rounds := 1000

		for range rounds {
			counts[s.PickIndex()]++
		}

		if len(counts) != len(nums) {
			t.Fatalf("expected counts to be %d", len(nums))
		}

		count, ok := counts[0]
		if !ok {
			t.Fatalf("index 0 not in counts")
		}

		if count != rounds {
			t.Fatalf("got count of %d, expected %d", count, rounds)
		}
	})

	t.Run("weighted distribution", func(t *testing.T) {
		nums := []int{1, 2, 3}

		sum := sumSlice(nums)
		p0, p1, p2 := float64(nums[0])/float64(sum), float64(nums[1])/float64(sum), float64(nums[2])/float64(sum)

		s := Constructor(nums)

		counts := make(map[int]int)
		rounds := 1000
		for range rounds {
			counts[s.PickIndex()]++
		}

		if len(counts) != len(nums) {
			t.Fatalf("expected counts to be %d", len(nums))
		}

		count0, ok := counts[0]
		if !ok {
			t.Errorf("index 0 not in counts")
		}

		count1, ok := counts[1]
		if !ok {
			t.Errorf("index 1 not in counts")
		}

		count2, ok := counts[2]
		if !ok {
			t.Errorf("index 2 not in counts")
		}

		t.Log(p0, p1, p2)
		t.Log(count0, count1, count2)

		res0, res1, res2 := float64(count0)/float64(rounds), float64(count1)/float64(rounds), float64(count2)/float64(rounds)
		t.Log(res0, res1, res2)

		margin := .03
		if math.Abs(p0-res0) > margin {
			t.Errorf("got %f for index 0, expected %f", p0, res0)
		}

		if math.Abs(p1-res1) > margin {
			t.Errorf("got %f for index 1, expected %f", p1, res1)
		}

		if math.Abs(p2-res2) > margin {
			t.Errorf("got %f for index 2, expected %f", p2, res2)
		}
	})
}

func sumSlice(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum
}
