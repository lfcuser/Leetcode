/*
404 Sum of Left Leaves

Given the root of a binary tree, return the sum of all left leaves.

A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.

Constraints:

	The number of nodes in the tree is in the range [1, 1000].
	-1000 <= Node.val <= 1000
*/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}

	return sum + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func main() {
	root := &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	fmt.Println(sumOfLeftLeaves(root), 24)
}

// go run ./golang/SumOfLeftLeaves/SumOfLeftLeaves.go
