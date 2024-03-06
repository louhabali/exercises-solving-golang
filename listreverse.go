package piscine

func ListReverse(l *List) {
	reserve := l.Head
	var previous *NodeL
	c := l.Head
	for c != nil {
		previous, c, c.Next = c, c.Next, previous
	}
	l.Head = previous
	l.Tail = reserve
}
