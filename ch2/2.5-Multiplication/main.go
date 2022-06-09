package __5_Multiplication

import (
	"DataStructureByGolang/ch2/2.4-polynomial"
	"fmt"
)

type PolyNode = __4_polynomial.PolyNode

//type PolyNode struct {
//	Coef  int // 系数
//	Expon int // 指数
//	Link  *PolyNode
//}

func ReadPoly() (newNode *PolyNode, len int) {
	len, _ = fmt.Scanf("%s")
	rare := newNode
	for count := 1; count <= len; count++ {
		rare.Coef, _ = fmt.Scanf("%d")
		rare.Expon, _ = fmt.Scanf("%d")
		rare.Link = &PolyNode{}
		rare = rare.Link
		fmt.Printf("第%d个多项式", count)
	}
	return
}

func MultiplySimple(l1, l2 *PolyNode) *PolyNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	m1 := &PolyNode{}
	m1.Expon = l1.Expon + l2.Expon
	m1.Coef = l1.Coef * l2.Coef
	return m1
}

func Insert(pre, cur, node1, node2 *PolyNode) {
	// 安全性检查，保证pre.Link为cur
	if pre == nil || cur == nil || pre.Link != cur {
		return
	}
	newNode := MultiplySimple(node1, node2)
	// 系数为0时返回
	if newNode == nil || newNode.Coef == 0 {
		return
	}
	for cur.Expon > newNode.Expon && cur != nil {
		cur = cur.Link
		pre = pre.Link
	}
	switch {
	case cur == nil:
		{
			pre.Link = newNode
		}
	case cur.Expon == newNode.Expon:
		{
			cur.Coef = cur.Coef + newNode.Coef
		}
	default:
		{
			pre.Link = newNode
			newNode.Link = cur
		}
	}
}

func Multiply(l1, l2 *PolyNode) *PolyNode {
	if l1 == nil || l2 == nil {
		return &PolyNode{Coef: 0, Expon: 0}
	}
	p1 := l1
	p2 := l2
	sentry := &PolyNode{}
	head := sentry
	// 构建初始多项式
	for p2 != nil {
		head.Link = MultiplySimple(p1, p2)
		head = head.Link
		p2 = p2.Link
	}
	t1 := p1.Link
	for t1 != nil {
		t2 := p2
		pre := sentry
		cur := sentry.Link
		for t2 != nil {
			Insert(pre, cur, t1, t2)
			t2 = t2.Link
		}
		t1 = t1.Link
	}
	return sentry.Link
}

func main() {
	l1, _ := ReadPoly()
	l2, _ := ReadPoly()
	addNode := __4_polynomial.Add(l1, l2)
	fmt.Println(addNode)
}
