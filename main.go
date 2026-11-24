package main

import "fmt"

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Ceil returns the smallest value in the BST that is ≥ key; ok is false
// when every value is below the key.
func Ceil(root *TreeNode, key int) (int, bool) {
	ceil, ok := 0, false
	for n := root; n != nil; {
		switch {
		case n.Val == key:
			return n.Val, true
		case n.Val > key:
			ceil, ok = n.Val, true // candidate; a tighter one may sit left
			n = n.Left
		default:
			n = n.Right
		}
	}
	return ceil, ok
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(8, tn(4, tn(2, nil, nil), tn(6, nil, nil)), tn(12, tn(10, nil, nil), tn(14, nil, nil)))
	fmt.Println(Ceil(root, 11))
}
