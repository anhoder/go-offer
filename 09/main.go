package main

import "fmt"

//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )
//
//示例 1：
//
//输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]
//示例 2：
//
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]


func main() {
	queue := Constructor()
	fmt.Println(queue.DeleteHead())
	queue.AppendTail(5)
	queue.AppendTail(2)
	fmt.Println(queue.DeleteHead())
	fmt.Println(queue.DeleteHead())
}

type Stack struct {
	vendor []int
}

func (s *Stack) pop() int {

	stackLen := len(s.vendor)

	if stackLen == 0 {
		return -1
	}

	last := s.vendor[stackLen-1]
	s.vendor = s.vendor[:stackLen-1]

	return last
}

func (s *Stack) push(item int) {
	s.vendor = append(s.vendor, item)
}

func (s *Stack) len() int {
	return len(s.vendor)
}


type CQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func Constructor() CQueue {
	return CQueue{
		&Stack{},
		&Stack{},
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.stack1.push(value)
}


func (this *CQueue) DeleteHead() int {
	if this.stack2.len() == 0 {
		var item int
		for {
			item = this.stack1.pop()
			if item < 0 {
				break
			}
			this.stack2.push(item)
		}
	}

	return this.stack2.pop()
}