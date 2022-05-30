// Package linearListByArray 线性表 By list
package linearListByArray

import "fmt"

const MaxListSize = 1000

type ElementType = interface {
}

type LNode struct {
	DATA []ElementType
	LAST int
}

func (_ *LNode) MakeEmpty() *LNode {
	// LAST为-1时代表没有元素
	return &LNode{
		LAST: -1,
		DATA: make([]ElementType, 0, MaxListSize),
	}
}

func (node *LNode) Find(x ElementType) (i int) {
	for i <= node.LAST && node.DATA[i] != x {
		i++
	}
	if i > node.LAST {
		i = -1
	}
	return
}

// Insert 在能插入元素的空位的第i个空位中插入元素
func (node *LNode) Insert(x ElementType, i int) bool {
	if node.LAST >= MaxListSize-1 {
		fmt.Println("线性表已满!")
		return false
	}
	if i < 1 || i > node.LAST+2 {
		fmt.Println("插入的位置不合法!")
		return false
	}
	// 需要替换的位置是i-1，即node.DATA[i-1] = x
	// 原有的从i-1开始的元素全往后挪，所以当n>=i-1时，继续循环
	for n := node.LAST; n >= i-1; n-- {
		node.DATA[n+1] = node.DATA[n]
	}
	node.DATA[i-1] = x
	node.LAST++
	return true
}

func (node *LNode) Delete(i int) bool {
	if i > node.LAST+1 || i < 1 {
		fmt.Println("删除的位置不合法!")
		return false
	}
	for n := i; n <= node.LAST; n++ {
		node.DATA[n-1] = node.DATA[n]
	}
	node.LAST--
	return true
}
