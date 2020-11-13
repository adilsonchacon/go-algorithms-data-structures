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
	Tail *ListNode
}

// ListNode

func (node *ListNode) Add(head *ListNode, value int) {
	if node.Next == head {
		aux := &ListNode{Value: value, Next: head}
		node.Next = aux
	} else if node.Next.Value > value {
		aux := &ListNode{Value: value, Next: node.Next}
		node.Next = aux
	} else {
		node.Next.Add(head, value)
	}
}

func (node *ListNode) Print(head *ListNode, times int) {
	if head == node {
		times--
		if times >= 0 {
			fmt.Printf("countdown %d\n", times+1)
		} else {
			return
		}
	}
	fmt.Println(node.Value)
	node.Next.Print(head, times)
}

// List

func (list *List) GetTail() *ListNode {
	node := list.Tail
	if node == nil {
		node = list.Head
	}

	if node == nil {
		return nil
	}

	for {
		if node.Next == nil || node.Next == list.Head {
			break
		}
		node = node.Next
	}

	return node
}

func (list *List) Add(value int) {
	if list.Head == nil {
		list.Head = &ListNode{Value: value}
		list.Head.Next = list.Head
		list.Tail = list.Head
	} else if list.Head.Value > value {
		aux := &ListNode{Value: value, Next: list.Head}
		list.Tail = list.GetTail()
		list.Tail.Next = aux
		list.Head = aux
	} else {
		list.Head.Add(list.Head, value)
		list.Tail = list.GetTail()
		list.Tail.Next = list.Head
	}
}

func (list *List) Print() {
	if list.Head == nil {
		fmt.Println("List is empty")
	} else {
		list.Head.Print(list.Head, 5)
	}
}

func main() {
	list := &List{}
	list.Add(6)
	list.Add(2)
	list.Add(7)
	list.Add(4)
	list.Add(5)

	list.Print()
}
