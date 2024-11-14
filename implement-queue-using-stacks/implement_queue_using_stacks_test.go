package implement_queue_using_stacks

import "testing"

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	if num1 := obj.Peek(); num1 != 1 {
		t.Errorf("expected %d, got %d", 1, num1)
	}

	if num1 := obj.Pop(); num1 != 1 {
		t.Errorf("expected %d, got %d", 1, num1)
	}

	if obj.Empty() {
		t.Errorf("queue should not be empty")
	}
}
