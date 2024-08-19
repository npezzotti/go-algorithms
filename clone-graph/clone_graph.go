package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	nodeMap := make(map[*Node]*Node)
	return clone(node, nodeMap)
}

func clone(node *Node, nodeMap map[*Node]*Node) *Node {
	if _, ok := nodeMap[node]; ok {
		return nodeMap[node]
	}

	copy := &Node{Val: node.Val}
	nodeMap[node] = copy

	for _, neighbor := range node.Neighbors {
		copy.Neighbors = append(copy.Neighbors, clone(neighbor, nodeMap))
	}

	return copy
}
