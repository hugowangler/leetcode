package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lca, _ := lca(root, p, q)
	return lca
}

func lca(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	if root.Val == p.Val {
		return p, true
	}
	if root.Val == q.Val {
		return q, true
	}
	left, foundL := lca(root.Left, p, q)
	right, foundR := lca(root.Right, p, q)
	if foundL && foundR {
		return root, true
	} else if foundL {
		return left, true
	} else if foundR {
		return right, true
	}
	return nil, false
}

func buildTree(nodes []*int, i int) *TreeNode {
	if i >= len(nodes) {
		return nil
	}
	if nodes[i] == nil {
		return nil
	}
	return &TreeNode{*nodes[i], buildTree(nodes, 2*i+1), buildTree(nodes, 2*i+2)}
}

func findNode(node int, root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == node {
		return root
	}
	left := findNode(node, root.Left)
	if left != nil {
		return left
	}
	right := findNode(node, root.Right)
	return right
}
