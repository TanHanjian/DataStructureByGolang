// Package queueByArray 顺环队列
package queueByArray

const maxQueueSize = 100

type ElementType = interface {
}

type QNode struct {
	Data  []ElementType
	Rear  int
	Front int
}

func (_ *QNode) MakeEmpty() *QNode {
	return &QNode{
		Data:  make([]ElementType, maxQueueSize, maxQueueSize),
		Rear:  -1,
		Front: -1,
	}
}

func (queue *QNode) Add(x ElementType) bool {
	if (queue.Front+1)%maxQueueSize == queue.Rear {
		return false
	}
	queue.Rear = (queue.Rear + 1) % maxQueueSize
	queue.Data[queue.Rear] = x
	return true
}

func (queue *QNode) Delete() (ok bool, delete ElementType) {
	ok = true
	if queue.Front == queue.Rear {
		ok = false
		return
	}
	delete = queue.Data[queue.Front]
	queue.Front = (queue.Front + 1) % maxQueueSize
	return
}
