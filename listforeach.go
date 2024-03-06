package piscine

func ListForEach(l *List, f func(*NodeL)) {
	c := l.Head
	for c != nil {
		f(c)
		c = c.Next
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}
