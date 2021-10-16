package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Value   int
	balance int
	Left    *TreeNode
	Right   *TreeNode
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
			if node.Left == nil {
				node.balance = -1
			} else {
				node.balance = 0
			}
		} else {
			err := node.Right.Insert(newTreeNode)
			if err == nil {
				if math.Abs(float64(node.Right.balance)) > 1 {
					fmt.Println("(Right) node needs balancing")
					node.Balance(node.Right)
				} else {
					node.balance--
				}
			}
		}
	} else if newTreeNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newTreeNode
			if node.Right == nil {
				node.balance = 1
			} else {
				node.balance = 0
			}
		} else {
			err := node.Left.Insert(newTreeNode)
			if err == nil {
				if math.Abs(float64(node.Left.balance)) > 1 {
					fmt.Println("(Left) node needs balancing")
					node.Balance(node.Left)
				} else {
					node.balance++
				}
			}
		}
	}
	return nil
}

func (node *TreeNode) Balance(targetNode *TreeNode) {
	if targetNode.Right != nil && targetNode.Right.balance == -1 {
		fmt.Println("RotateLeft")
		node.RotateLeft(targetNode)
	} else if targetNode.Left != nil && targetNode.Left.balance == 1 {
		fmt.Println("RotateRight")
		node.RotateRight(targetNode)
	} else if targetNode.Left != nil && targetNode.Left.balance == -1 {
		fmt.Println("RotateLeftRight")
		node.RotateLeftRight(targetNode)
	} else if targetNode.Right != nil && targetNode.Right.balance == 1 {
		fmt.Println("RotateRightLeft")
		node.RotateRightLeft(targetNode)
	}
}

func (node *TreeNode) RotateLeft(targetNode *TreeNode) {
	aux := targetNode.Right
	targetNode.Right = aux.Left
	aux.Left = targetNode

	if targetNode == node.Left {
		node.Left = aux
	} else {
		node.Right = aux
	}

	targetNode.balance = 0
	aux.balance = 0
}

func (node *TreeNode) RotateRight(targetNode *TreeNode) {
	aux := targetNode.Left
	targetNode.Left = aux.Right
	aux.Right = targetNode

	if targetNode == node.Left {
		node.Left = aux
	} else {
		node.Right = aux
	}

	targetNode.balance = 0
	aux.balance = 0
}

func (node *TreeNode) RotateLeftRight(targetNode *TreeNode) {
	// targetNode.Left.Right.balance = 1
	targetNode.RotateLeft(targetNode.Left)
	targetNode.Left.balance = 1

	if targetNode.Left.Left != nil && targetNode.Left.Left.Left != nil && targetNode.Left.Left.Right == nil {
		targetNode.Left.Left.balance = 1
	}

	node.RotateRight(targetNode)
}

func (node *TreeNode) RotateRightLeft(targetNode *TreeNode) {
	// targetNode.Right.Left.balance = -1
	targetNode.RotateRight(targetNode.Right)
	targetNode.Right.balance = -1

	if targetNode.Right.Right != nil && targetNode.Right.Right.Right != nil && targetNode.Right.Right.Left == nil {
		targetNode.Right.Right.balance = -1
	}

	node.RotateLeft(targetNode)
}

func (node *TreeNode) Print(i int) {
	if node.Left != nil {
		node.Left.Print(i + 1)
	}

	if i > 0 {
		fmt.Printf(`%`+strconv.Itoa(i*2)+`s`, ` `)
	}
	fmt.Printf("Value: %d -> balance: %d\n", node.Value, node.balance)

	if node.Right != nil {
		node.Right.Print(i + 1)
	}
}

func (node *TreeNode) Find(value int) (*TreeNode, bool) {
	if node == nil {
		return nil, false
	}

	if node.Value == value {
		return node, true
	}

	if value > node.Value {
		return node.Right.Find(value)
	} else {
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
		if err == nil {
			if math.Abs(float64(tree.Root.balance)) > 1 {
				fmt.Println("(Root) node needs balancing")
				tree.Balance()
			}
		} else {
			fmt.Println(err)
		}
	}
}

func (tree *Tree) Balance() {
	aux := &TreeNode{Left: tree.Root, Value: -1}
	aux.Balance(tree.Root)
	tree.Root = aux.Left
}

func (tree *Tree) Print() {
	if tree.Root == nil {
		fmt.Println("Empty tree!")
	} else {
		tree.Root.Print(0)
	}
}

func (tree *Tree) Find(value int) (*TreeNode, bool) {
	node, ok := tree.Root.Find(value)

	return node, ok
}

func main() {
	tree := &Tree{}
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(0)
	tree.Insert(3)
	tree.Insert(4)
	tree.Print()
}
