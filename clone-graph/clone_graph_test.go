package clone_graph

import (
	"testing"
)

func TestCloneGraph(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}

	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n3, n1}
	n3.Neighbors = []*Node{n4, n2}
	n4.Neighbors = []*Node{n1, n3}

	clonedNode := cloneGraph(n1)

	// verify cloned graph is structurally identical to original
	if !isGraphEqual(n1, clonedNode) {
		t.Fatalf("cloned graph does not match original graph")
	}

	// verify the nodes are different (not shallow copies)
	if n1 == clonedNode {
		t.Fatalf("cloned node references original node")
	}
}

func isGraphEqual(node1, node2 *Node) bool {
	nodeMap := make(map[*Node]*Node)
	return compareDfs(node1, node2, nodeMap)
}

func compareDfs(node1, node2 *Node, visited map[*Node]*Node) bool {
	if node1 == nil || node2 == nil {
		return node1 == node2
	}

	if _, ok := visited[node1]; ok {
		return visited[node1] == node2
	}

	if node1.Val != node2.Val || len(node1.Neighbors) != len(node2.Neighbors) {
		return false
	}

	visited[node1] = node2

	for i := range node1.Neighbors {
		if !compareDfs(node1.Neighbors[i], node2.Neighbors[i], visited) {
			return false
		}
	}

	return true
}
