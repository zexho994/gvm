package runtime

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

func (stack *Stack) push(frame *Frame) {
	//fmt.Printf("[gvm][jvm_stack.push] start push. stack.size = %v , stack.maxSize = %v \n", stack.size, stack.maxSize)
	if stack.size >= stack.maxSize {
		panic("[gvm][Stack.push] StackOverflowError!")
	}
	if stack._top != nil {
		frame.lower = stack._top
	}
	stack._top = frame
	stack.size++
	//fmt.Printf("[gvm][jvm_stack.push] push done. stack.size = %v , stack.maxSize = %v \n", stack.size, stack.maxSize)
}

/**
弹出栈桢

*/
func (stack *Stack) pop() *Frame {
	if stack._top == nil {
		panic("[gvm][Stack.pop] stack is empty")
	}
	top := stack._top
	stack._top = top.lower
	top.lower = nil
	stack.size--
	return top
}

func (stack *Stack) top() *Frame {
	if stack._top == nil {
		panic("[gvm][Stack.top] stack is empty")
	}
	return stack._top
}

func (stack *Stack) isEmpty() bool {
	return stack._top == nil
}
