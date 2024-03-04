package piscine

func ListPushFront(l *List, data interface{}) {
	nouvelle := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = nouvelle
		l.Tail = nouvelle
	} else {
		nouvelle.Next = l.Head
		l.Head = nouvelle
	}
}
