package min_stack

import "testing"

func TestMinStack(t *testing.T) {
	obj := Constructor()

	vals := []int{-2, 0, -3}
	for _, val := range vals {
		obj.Push(val)
	}

	if min := obj.GetMin(); min != -3 {
		t.Fatalf("expected min to be %d, got %d", -3, min)
	}

	obj.Pop()

	if top := obj.Top(); top != 0 {
		t.Fatalf("expected last item to be %d, got %d", 0, top)
	}

	if newMin := obj.GetMin(); newMin != -2 {
		t.Fatalf("expected new min to be %d, got %d", -2, newMin)
	}
}
