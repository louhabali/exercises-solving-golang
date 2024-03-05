package piscine

import "fmt"

func PrintList(l *List) {
	if l.Head == nil {
		fmt.Println("nil")
		return
	}
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}

func ListClear(l *List) {
	l.Head = nil
	l.Tail = nil
}
