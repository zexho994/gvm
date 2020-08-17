package rtda

type Stack struct {
	// 当前栈的容量,最多可以容纳多少桢
	maxSize uint
	// 当前桢的大小
	size uint
	// 栈顶指针
	_top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("[gvm][Stack.push] StackOverflowError!")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self.top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("[gvm][Stack.pop] stack is empty")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("[gvm][Stack.top] stack is empty")
	}
	return self.top()
}
