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
