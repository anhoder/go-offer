package main

import "fmt"

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
//示例 1：
//
//输入：head = [1,3,2]
//输出：[2,3,1]

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {

	head := &ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				2,
				nil,
			},
		},
	}

	fmt.Println(reversePrint(head))

}


func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	// 方法一
	//var stack, result []int
	//
	//
	//for head != nil {
	//	stack = append(stack, head.Val)
	//	head = head.Next
	//}
	//
	//for i := len(stack) - 1; i >= 0; i-- {
	//	result = append(result, stack[i])
	//}
	//
	//return result

	// 方法二 递归
	//var result []int
	//reverse(head, &result)
	//
	//return result
	
	var length int
	node := head
	for node != nil {
		length++
		node = node.Next
	}

	result := make([]int, length)
	for i := length - 1; head != nil; i-- {
		result[i] = head.Val
		head = head.Next
	}

	return result
}

func reverse(node *ListNode, result *[]int)  {
	if node == nil {
		return
	}
	reverse(node.Next, result)
	*result = append(*result, node.Val)
}