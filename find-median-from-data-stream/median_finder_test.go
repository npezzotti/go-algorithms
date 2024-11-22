package median_finder

import "testing"

func TestMedianFinder(t *testing.T) {
	t.Run("even length", func(t *testing.T) {
		obj := Constructor()
		obj.AddNum(1)
		obj.AddNum(2)
		obj.AddNum(3)
		obj.AddNum(4)

		if median := obj.FindMedian(); median != 2.5 {
			t.Errorf("got %f, expected %f", median, 2.5)
		}
	})

	t.Run("odd length", func(t *testing.T) {
		obj := Constructor()
		obj.AddNum(1)
		obj.AddNum(2)
		obj.AddNum(3)
		obj.AddNum(4)
		obj.AddNum(5)

		if median := obj.FindMedian(); median != 3 {
			t.Errorf("got %f, expected %d", median, 3)
		}
	})
}
