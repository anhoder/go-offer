package main

import "fmt"

//输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
//例如，给出
//
//前序遍历 preorder =[3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//
//3
/// \
//9  20
///  \
//15   7

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main() {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	head := buildTree(preorder, inorder)

	printTree(head)
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	preLen := len(preorder)
	if preLen == 0 {
		return nil
	}

	if preLen == 1 {
		return &TreeNode{
			preorder[0],
			nil,
			nil,
		}
	}

	var inorderIndex int
	var head, leftTree, rightTree *TreeNode

	for i, v := range inorder {
		if v == preorder[0] {
			inorderIndex = i
			break
		}
	}

	leftTree = buildTree(preorder[1:inorderIndex+1], inorder[0:inorderIndex])
	rightTree = buildTree(preorder[inorderIndex+1:], inorder[inorderIndex+1:])

	head = &TreeNode{
		preorder[0],
		leftTree,
		rightTree,
	}

	return head
}

func printTree(head *TreeNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	printTree(head.Left)
	printTree(head.Right)
}