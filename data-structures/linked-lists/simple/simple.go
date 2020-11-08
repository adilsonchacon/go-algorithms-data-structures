package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

type List struct {
	Head *ListNode
}

// ListNode

func (node *ListNode) Add(value int) {
	if node.Next == nil {
		node.Next = &ListNode{Value: value}
	} else {
		node.Next.Add(value)
	}
}

func (node *ListNode) Print() {
	if node != nil {
		fmt.Println(node.Value)
		node.Next.Print()
	}
}

// List

func (list *List) Add(value int) {
	if list.Head == nil {
		list.Head = &ListNode{Value: value}
	} else {
		list.Head.Add(value)
	}
}

func (list *List) Print() {
	if list.Head == nil {
		fmt.Println("List is empty")
	} else {
		list.Head.Print()
	}
}

func main() {
	list := &List{}
	list.Add(2)
	list.Add(6)
	list.Add(7)
	list.Add(4)
	list.Add(3)

	list.Print()
}
