package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
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

func ListReverse(l *List) {

}

func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, 2)
	ListPushBack(link, 3)
	ListPushBack(link, 4)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}
// $ go run .
// 4
// 3
// 2
// 1
// Tail &{1 <nil>}
// Head &{4 0xc42000a140}
// $