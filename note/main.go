package main

import (
	"fmt"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func addNode(head *TreeNode, value int) *TreeNode {
	if head == nil {
		return &TreeNode{value, nil, nil}
	} else if value >= head.value {
		head.right = addNode(head.right, value)
	} else {
		head.left = addNode(head.left, value)
	}
	return head
}

func inOrderTraversal(root *TreeNode) {
	var cur = root
	var tmp *TreeNode

	for cur != nil {
		if cur.left == nil {
			fmt.Println(cur.value)
			cur = cur.right
		} else {
			tmp = cur.left
			for tmp.right != nil && tmp.right != cur {
				tmp = tmp.right
			}

			if tmp.right == nil {
				tmp.right = cur
				cur = cur.left
			} else {
				tmp.right = nil
				fmt.Println(cur.value)
				cur = cur.right
			}
		}
	}
}

func main() {
	root := addNode(nil, 1)
	addNode(root, 1)
	addNode(root, 7)
	addNode(root, 2)
	addNode(root, 6)
	addNode(root, 5)
	addNode(root, 3)
	inOrderTraversal(root)
}
