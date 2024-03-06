package piscine

func ListForEachIf(l *List, f func(*NodeL), condition func(*NodeL) bool) {
	it := l.Head
	for it != nil {
		if condition(it) {
			f(it)
		}
		it = it.Next
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
