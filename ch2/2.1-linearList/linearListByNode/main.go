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

func (head *LNode) FindKth(index int) (p *LNode, ok bool) {
	ok = true
	p = head
	for i := 1; i < index && p != nil; i++ {
		p = p.Next
	}
	if p == nil {
		ok = false
	}
	return
}

func (head *LNode) Find(x ElementType) (p *LNode, ok bool) {
	ok = true
	p = head
	for p != nil && p.Val != x {
		p = p.Next
	}
	if p == nil {
		ok = false
	}
	return
}

// InsertBySentry 有可能会在头部新增节点，此时需要把head节点变更，Insert方法需要返回一个新节点
func (head *LNode) InsertBySentry(x ElementType, i int) (res *LNode, ok bool) {
	p := head
	// 前驱节点，保证开始时pre一定不为nil
	sentry := &LNode{Next: p}
	res = sentry
	for n := 1; n < i && p != nil; n++ {
		sentry = p
		p = p.Next
	}
	// pre不为nil时，说明链表长度大于等于i，此时可插入数据
	if p != nil {
		newNode := LNode{Val: x, Next: p}
		sentry.Next = &newNode
		ok = true
	}
	return
}

func (head *LNode) Insert(x ElementType, i int) (res *LNode, ok bool) {
	ok = true
	res = head
	if i == 1 {
		newNode := &LNode{Val: x, Next: head}
		res = newNode
		return
	}
	pre, isFound := head.FindKth(i - 1)
	if isFound {
		newNode := &LNode{Val: x, Next: pre.Next}
		pre.Next = newNode
	}
	ok = isFound
	return
}

func (head *LNode) DeleteKth(i int) (res *LNode, ok bool) {
	ok = true
	res = head
	if head == nil {
		ok = false
		return
	}
	if i == 1 {
		res = res.Next
		return
	}
	pre, isFound := head.FindKth(i - 1)
	ok = isFound
	if !isFound {
		return
	}
	deleteNode := pre.Next
	if deleteNode == nil {
		ok = false
		return
	}
	pre.Next = deleteNode.Next
	deleteNode = nil
	return
}
