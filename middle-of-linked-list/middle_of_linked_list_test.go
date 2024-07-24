package middle_of_linked_list

import "testing"

func TestMiddleNode(t *testing.T) {
	tcases := []struct {
		name string
		head *ListNode
		want int
	}{
		{
			name: "single list item",
			head: &ListNode{
				Val: 1,
				Next: nil,
			},
			want: 1,
		},
		{
			name: "odd length list",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			want: 3,
		},
		{
			name: "even length list",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: 3,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := middleNode(tc.head)

			if got.Val != tc.want {
				t.Fatalf("expected 3, got %d\n", got.Val)
			}
		})
	}
}
