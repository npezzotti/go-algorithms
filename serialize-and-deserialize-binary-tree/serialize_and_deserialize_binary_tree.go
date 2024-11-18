package serialize_and_deserialize_binary_tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var s []string

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			s = append(s, "nil")
			return
		}

		s = append(s, strconv.Itoa(node.Val))

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return strings.Join(s, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, ",")

	var idx int
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if s[idx] == "nil" {
			idx++
			return nil
		}

		v, _ := strconv.Atoi(s[idx])
		node := &TreeNode{Val: v}

		idx++
		node.Left = dfs()
		node.Right = dfs()

		return node
	}

	return dfs()
}
