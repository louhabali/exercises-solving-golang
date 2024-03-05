package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	thelast := l.Head
	for thelast.Next != nil {
		thelast = thelast.Next
	}
	return thelast.Data
}
