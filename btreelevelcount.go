package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 1
	leftCount := BTreeLevelCount(root.Left)
	rightCount := BTreeLevelCount(root.Right)

	if leftCount > rightCount {
		count += leftCount
	} else {
		count += rightCount
	}
	return count
}
