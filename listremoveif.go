package piscine 

type List struct {
	Head *NodeL
	Tail *NodeL
}
type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListPushBack(l *List, data interface{}) {
	nouvelle := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = nouvelle
		l.Tail = nouvelle
	} else {
		l.Tail.Next = nouvelle
		l.Tail = nouvelle
	}
}

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	current := l.Head
	for current != nil && current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	if l.Head == nil {
		l.Tail = nil
	}
}
