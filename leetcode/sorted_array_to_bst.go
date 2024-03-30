package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}

	mid := len(nums) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}
	createNodes(root, nums[:mid], nums[mid+1:])
	return root
}

func createNodes(parent *TreeNode, smaller []int, bigger []int) {
	if len(smaller) > 0 {
		mid := len(smaller) / 2
		parent.Left = &TreeNode{
			Val:   smaller[mid],
			Left:  nil,
			Right: nil,
		}
		createNodes(parent.Left, smaller[:mid], smaller[mid+1:])
	}
	if len(bigger) > 0 {
		mid := len(bigger) / 2
		parent.Right = &TreeNode{
			Val:   bigger[mid],
			Left:  nil,
			Right: nil,
		}
		createNodes(parent.Right, bigger[:mid], bigger[mid+1:])
	}
}
