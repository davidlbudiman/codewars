package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
			return 0
	}
	leftLevel, rightLevel := MaxDepth(root.Left) + 1, MaxDepth(root.Right) + 1
	if leftLevel < rightLevel {
			return rightLevel
	} else {
			return leftLevel
	}
}

func MaxDepth_v1(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}
	leftLevel, rightLevel := 0, 0
	if root.Left != nil {
		leftLevel = MaxDepth_v1(root.Left) + 1
	}
	if root.Right != nil {
		rightLevel = MaxDepth_v1(root.Right) + 1
	}
	if leftLevel < rightLevel {
		return rightLevel
	} else {
		return leftLevel
	}
}
