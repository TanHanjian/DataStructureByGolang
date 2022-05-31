package stackByArray

import "fmt"

const MaxSize = 100

type ElementType = interface {
}

type SNode struct {
	Data []ElementType
	Top  int
}

func (_ *SNode) MakeEmpty() *SNode {
	return &SNode{
		Data: make([]ElementType, 0, MaxSize),
		// Top为-1时堆栈为空
		Top: -1,
	}
}

func (list *SNode) Push(x ElementType) (ok bool) {
	ok = true
	if list.Top >= MaxSize-1 {
		fmt.Println("栈越界！无法插入！")
		ok = false
		return
	}
	list.Data = append(list.Data, x)
	list.Top++
	return
}

func (list *SNode) Pop() (ok bool) {
	ok = true
	if list.Top == -1 {
		ok = false
		return
	}
	list.Data = list.Data[:len(list.Data)-1]
	list.Top--
	return
}
