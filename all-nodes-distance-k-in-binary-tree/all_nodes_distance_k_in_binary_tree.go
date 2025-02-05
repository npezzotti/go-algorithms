package all_nodes_distance_k_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parentMap := make(map[*TreeNode]*TreeNode)
	var populateParentMap func(node *TreeNode)
	populateParentMap = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			parentMap[node.Left] = node
		}

		if node.Right != nil {
			parentMap[node.Right] = node
		}

		populateParentMap(node.Left)
		populateParentMap(node.Right)
	}

	populateParentMap(root)

	var res []int
	visited := make(map[*TreeNode]struct{})
	var dfs func(node *TreeNode, count int)
	
	dfs = func(node *TreeNode, count int) {
		if node == nil {
			return
		}

		if _, ok := visited[node]; ok {
			return
		}

		visited[node] = struct{}{}

		if count == k {
			res = append(res, node.Val)
			return
		}

		dfs(node.Left, count+1)
		dfs(node.Right, count+1)
		dfs(parentMap[node], count+1)
	}

	dfs(target, 0)

	return res
}
