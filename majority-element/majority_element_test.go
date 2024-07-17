package majority_element

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "majority should be 2",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "majority should be 3",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "majority should be 5",
			nums: []int{6,5,5},
			want: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := majorityElement(tc.nums)

			if got != tc.want {
				t.Fatalf("got %d, want %d\n", got, tc.want)
			}
		})
	}
}
