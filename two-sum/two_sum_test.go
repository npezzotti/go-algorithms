package two_sum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	got := TwoSum([]int{5, 7, 2, 3, 9}, 5)
	want := []int{2, 3}

	errMsg := fmt.Sprintf("want %#v got %#v", want, got)

	if len(got) != len(want) {
		t.Errorf(errMsg)
	}

	for i, num := range got {
		if num != want[i] {
			t.Errorf(errMsg)
		}
	}

}
