package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	return searchItemHelper(root, elem, nil)
}

func searchItemHelper(root *TreeNode, elem string, parent *TreeNode) *TreeNode {
	if root == nil || root.Data == elem {
		if root != nil {
			root.Parent = parent
		}
		return root
	}

	if elem < root.Data {
		return searchItemHelper(root.Left, elem, root)
	}

	return searchItemHelper(root.Right, elem, root)
}
