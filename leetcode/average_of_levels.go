package leetcode

func AverageOfLevels_v1(root *TreeNode) []float64 {
	var res = []float64{float64(root.Val)}
	var queue = []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		res, queue = averageOfLevel(res, queue)
	}
	return res
}

func averageOfLevel(res []float64, queue []*TreeNode) ([]float64, []*TreeNode) {
	var sum float64
	var length int
	var newQueue []*TreeNode
	for _, item := range queue {
		if item != nil {
			length++
			sum = float64(item.Val) + sum
			newQueue = append(newQueue, item.Left, item.Right)
		}
	}
	if length > 0 {
		res = append(res, sum/float64(length))
	}
	return res, newQueue
}
