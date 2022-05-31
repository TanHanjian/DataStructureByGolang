package DStack

const maxStackSize = 100

type ElementType = interface {
}

type DStack struct {
	Data []ElementType
	Top1 int
	Top2 int
}

func (_ *DStack) MakeEmpty() *DStack {
	return &DStack{
		Data: make([]ElementType, maxStackSize, maxStackSize),
		Top1: -1,
		Top2: maxStackSize,
	}
}

func (dStack *DStack) Push(x ElementType, tag int) (ok bool) {
	if dStack.Top1+1 >= dStack.Top2 {
		return
	}
	ok = true
	switch {
	case tag == 0:
		{
			dStack.Top1++
			dStack.Data[dStack.Top1] = x
		}
	case tag == 1:
		{
			dStack.Top2--
			dStack.Data[dStack.Top2] = x
		}
	}
	return
}

func (dStack *DStack) Pop(tag int) bool {
	switch {
	case dStack.Top1 > -1 && tag == 0:
		{
			// pop操作不清空原有数据可能也没关系
			// dStack.Data[dStack.Top1] = nil
			dStack.Top1--
			return true
		}
	case dStack.Top2 < maxStackSize && tag == 1:
		{
			dStack.Top2++
			return true
		}
	}
	return false
}
