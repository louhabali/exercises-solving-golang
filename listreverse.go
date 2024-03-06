package piscine

func ListReverse(l *List) {
	var previous *NodeL
	c := l.Head
	l.Tail = l.Head
	for c != nil {
		nx := c.Next
		c.Next = previous
		previous = c
		c = nx
	}
	l.Head = previous
}
