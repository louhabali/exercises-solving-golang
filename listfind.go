package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	init := l.Head
	for init != nil {
		if comp(ref, init.Data) { // comparaison
			return &init.Data
		}
		init = init.Next
	}
	return nil
}
