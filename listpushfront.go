package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	nouvelle := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = nouvelle
		l.Tail = nouvelle
	} else {
		l.Tail.Next = nouvelle
		l.Tail = nouvelle
	}
}
