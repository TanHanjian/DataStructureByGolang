package queueByNode

type ElementType = interface {
}

type Node struct {
	Data ElementType
	Next *Node
}

type QNode struct {
	Front *Node
	Rare  *Node
}

func (node *QNode) Delete() (val ElementType, ok bool) {
	if node.Front == nil {
		return
	}
	front := node.Front
	if front == node.Rare {
		node.Rare = nil
		node.Front = nil
	} else {
		node.Front = node.Front.Next
	}
	val = front.Data
	*front = Node{}
	return val, true
}

func (node *QNode) Add(x ElementType) bool {
	newNode := &Node{Data: x}
	if node.Rare == nil {
		node.Rare = newNode
		node.Front = newNode
		return true
	}
	node.Rare.Next = newNode
	node.Rare = node.Rare.Next
	return true
}
