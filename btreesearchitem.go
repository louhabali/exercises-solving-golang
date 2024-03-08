package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}

	if elem < root.Data {
		found := BTreeSearchItem(root.Left, elem)
		if found != nil {
			return found
		}
	}
	return BTreeSearchItem(root.Right, elem)
}
