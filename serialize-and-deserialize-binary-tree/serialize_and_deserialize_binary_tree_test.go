package serialize_and_deserialize_binary_tree

import "testing"

func TestCodec(t *testing.T) {
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}

	c := Constructor()

	s := "1,2,nil,nil,3,4,nil,nil,5,nil,nil"

	res := c.serialize(root)
	if res != s {
		t.Errorf("got %s, expected %s", res, s)
	}

	if newTree := c.deserialize(res); !areIdentical(newTree, root) {
		t.Errorf("deserialized tree does not match original")
	}
}

func areIdentical(t1, t2 *TreeNode) bool {
	// Both trees are nil
	if t1 == nil && t2 == nil {
		return true
	}

	// One tree is nil, the other is not
	if t1 == nil || t2 == nil {
		return false
	}

	// Check if values match and recursively check subtrees
	return t1.Val == t2.Val &&
		areIdentical(t1.Left, t2.Left) &&
		areIdentical(t1.Right, t2.Right)
}
