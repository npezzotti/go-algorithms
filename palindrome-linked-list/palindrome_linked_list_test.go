package palindrome_linked_list

import "testing"

func Test_isPalindrome(t *testing.T) {
	tcases := []struct {
		name   string
		head   *ListNode
		output bool
	}{
		{
			name: "[1,2,2]",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 1},
					},
				},
			},
			output: true,
		},
		{
			name: "[1,2]",
			head: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			},
			output: false,
		},
		{
			name: "[1,1,2,1]",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 1},
					},
				},
			},
			output: false,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			if isPalindrome(tc.head) != tc.output {
				t.Errorf("got %t, expected %t", !tc.output, tc.output)
			}
		})

		t.Run(tc.name, func(t *testing.T) {
			if isPalindromeOptimized(tc.head) != tc.output {
				t.Errorf("got %t, expected %t", !tc.output, tc.output)
			}
		})
	}
}
