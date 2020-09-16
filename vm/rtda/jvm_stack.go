package rtda

/*
虚拟机栈是JVM运行时数据区的一部分，线程私有
主要存储方法的栈桢 Frame
*/
type Stack struct {
	// 当前栈的容量,最多可以容纳多少桢
	maxSize uint
	// 当前桢的大小
	size uint
	// 栈顶指针
	_top *Frame
}

/*
新的虚拟机栈
构造方法中只会设置最大栈字段
*/
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
	self._top = frame
	self.size++
}

/**
弹出栈桢

*/
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

func (self *Stack) isEmpty() bool {
	return self._top == nil
}
