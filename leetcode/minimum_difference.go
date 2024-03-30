package leetcode

func GetMinimumDifference_v2(root *TreeNode) int {
	var min_diff = 100001
	var last_value = -100001
	traverse_v2(root, &last_value, &min_diff)
	return min_diff
}

func traverse_v2(node *TreeNode, last_value *int, min_diff *int) {
	if node == nil {
		return
	}
	traverse_v2(node.Left, last_value, min_diff)
	if node.Val-*last_value < *min_diff {
		*min_diff = node.Val - *last_value
	}
	*last_value = node.Val
	traverse_v2(node.Right, last_value, min_diff)
}

func GetMinimumDifference(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	var array []int
	traverse(root, &array)
	var min_diff = array[1] - array[0]
	for i, v := range array[2:] {
		if v-array[1+i] < min_diff {
			min_diff = v - array[1+i]
		}
	}
	return min_diff
}

func traverse(node *TreeNode, array *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, array)
	*array = append(*array, node.Val)
	traverse(node.Right, array)
}
