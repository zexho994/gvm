package runtime

// Stack 栈的主要作用就是存储栈帧
type Stack struct {
	// 栈最大大小
	maxSize uint
	// 当前栈的大小
	size uint
	// 顶层帧
	top *Frame
}

func (s *Stack) IsEmtpy() bool {
	return s.size == 0
}

func (s *Stack) Push(newFrame *Frame) {
	if s.size == s.maxSize {
		panic("[gvm] system error StackOverflow")
	}

	if s.top != nil {
		newFrame.nextFrame = s.top
	}

	s.top = newFrame
	s.size++
}

func (s *Stack) Pop() *Frame {
	if s.top == nil || s.size == 0 {
		panic("[gvm] stack is empty")
	}
	p := s.top
	s.top = s.top.nextFrame
	p.nextFrame = nil
	s.size--
	return p
}

func (s *Stack) Peek() *Frame {
	return s.top
}

func NewStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}
