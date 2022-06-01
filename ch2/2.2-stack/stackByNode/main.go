package stackByNode

type ElementType = interface {
}

type SNode struct {
	Data ElementType
	Next *SNode
}

func (_ *SNode) MakeEmpty() *SNode {
	return &SNode{}
}

func (node *SNode) Push(x ElementType) {
	newNode := &SNode{
		Data: x,
		Next: node.Next,
	}
	node.Next = newNode
}

func (node *SNode) Pop() (ok bool, pop ElementType) {
	ok = true
	if node.Next == nil {
		ok = false
		return
	}
	next := node.Next
	pop = next.Data
	node.Next = next.Next
	return
}
