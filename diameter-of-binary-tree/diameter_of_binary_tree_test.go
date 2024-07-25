package diameter_of_binary_tree

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tcases := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "test",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			want: 3,
		},
		{
			name: "diameter of 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			want: 1,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := diameterOfBinaryTree(tc.root)

			if got != tc.want {
				t.Fatalf("want %d, got %d\n", tc.want, got)
			}
		})
	}
}
