package __4_polynomial

type PolyNode struct {
	Coef  int // 系数
	Expon int // 指数
	Link  *PolyNode
}

func Add(p1, p2 *PolyNode) *PolyNode {
	sentry := &PolyNode{}
	p3 := sentry
	for p1 != nil && p2 != nil {
		switch {
		case p1.Expon == p2.Expon:
			{
				p3.Expon = p1.Expon
				p3.Coef = p1.Coef + p2.Coef
				p1 = p1.Link
				p2 = p2.Link
			}
		case p1.Expon > p2.Expon:
			{
				p3.Expon = p1.Expon
				p3.Coef = p1.Coef
				p1 = p1.Link
			}
		case p1.Expon < p2.Expon:
			{
				p3.Expon = p2.Expon
				p3.Coef = p2.Coef
				p1 = p2.Link
			}
			p3.Link = &PolyNode{}
			p3 = p3.Link
		}
		if p1 != nil {
			for p1 != nil {
				p3.Link = &PolyNode{
					Coef:  p1.Coef,
					Expon: p1.Expon,
				}
				p1 = p1.Link
				p3 = p3.Link
			}
		}
		if p2 != nil {
			for p2 != nil {
				p3.Link = &PolyNode{
					Coef:  p2.Coef,
					Expon: p2.Expon,
				}
				p2 = p2.Link
				p3 = p3.Link
			}
		}
	}
	return sentry
}

func main() {

}
