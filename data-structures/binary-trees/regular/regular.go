package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

// TreeNode

func (node *TreeNode) Insert(newTreeNode *TreeNode) error {
	if newTreeNode.Value == node.Value {
		return errors.New(fmt.Sprintf("value %d is already in tree"))
	} else if newTreeNode.Value > node.Value {
		if node.Right == nil {
			node.Right = newTreeNode
		} else {
			node.Right.Insert(newTreeNode)
		}

	} else if newTreeNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newTreeNode
		} else {
			node.Left.Insert(newTreeNode)
		}
	}

	return nil
}

func (node *TreeNode) Print() {
	if node.Left != nil {
		node.Left.Print()
	}

	fmt.Printf("Value: %d\n", node.Value)

	if node.Right != nil {
		node.Right.Print()
	}
}

func (node *TreeNode) Find(value int) (int, bool) {
	if node == nil {
		return 0, false
	}

	if node.Value == value {
		return value, true
	}

	if value > node.Value {
		return node.Right.Find(value)
	} else { // value < node.Value
		return node.Left.Find(value)
	}

}

// Tree

func (tree *Tree) Insert(value int) {
	node := &TreeNode{Value: value}

	if tree.Root == nil {
		tree.Root = node
	} else {
		err := tree.Root.Insert(node)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (tree *Tree) Print() {
	if tree.Root == nil {
		fmt.Println("Empty tree!")
	} else {
		tree.Root.Print()
	}
}

func (tree *Tree) Find(value int) {
	_, ok := tree.Root.Find(value)

	if ok {
		fmt.Printf("Success! Value %d was found!\n", value)
	} else {
		fmt.Printf("Oops! Value %d was NOT found!\n", value)
	}
}

func main() {
	root := &Tree{}
	root.Insert(10)
	root.Insert(7)
	root.Insert(14)
	root.Insert(2)
	root.Insert(0)
	root.Insert(19)
	root.Insert(5)
	root.Insert(1)

	root.Print()

	root.Find(2)
	root.Find(211)
}
