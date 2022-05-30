package linearListByNode

type ElementType = interface {
}

type LNode struct {
	Val  ElementType
	Next *LNode
}

func (head *LNode) Length() (len int) {
	p := head
	if p == nil {
		return
	}
	for p != nil {
		len++
		p = p.Next
	}
	return
}

func (head *LNode) FindKth(index int) *LNode {
	p := head
	for i := 1; i < index && p != nil; i++ {
		p = p.Next
	}
	return p
}

func (head *LNode) Find(x ElementType) *LNode {
	p := head
	for p != nil && p.Val != x {
		p = p.Next
	}
	return p
}

func (head *LNode) Insert(x ElementType, i int) bool {
	p := head
	// 前驱节点，保证开始时pre一定不为nil
	pre := &LNode{Next: p}
	for n := 1; n <= i && pre != nil; n++ {
		pre = p
		p = p.Next
	}
	// pre不为nil时，说明链表长度大于等于i，此时可插入数据
	if pre != nil {
		newNode := LNode{Val: x, Next: p}
		pre.Next = &newNode
		return true
	}
	return false
}
