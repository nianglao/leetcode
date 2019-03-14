package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// PrintTree print the tree in preorder
func PrintTree(t *TreeNode) {
	printTree(t)
	fmt.Printf("\n")
}

func printTree(t *TreeNode) {
	if t != nil {
		fmt.Printf("%d->", t.Val)
		printTree(t.Left)
		printTree(t.Right)
	}
	return
}

func main() {
	fmt.Println("vim-go")
}
