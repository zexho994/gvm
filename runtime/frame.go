package runtime

import (
	"github.com/zouzhihao-994/gvm/klass"
)

// Frame 一个Frame对应着一个已调用而且未结束的方法
// TODO：栈的大小支持自动 扩/缩 , 如果扩到极限仍然发送内容不足的情况抛出 OutOfMemoryError 异常
type Frame struct {
	framePC   uint
	nextFrame *Frame
	*LocalVars
	*OperandStack
	*klass.MethodKlass
	*Thread
}

func (f *Frame) SetFramePC(pc uint) {
	f.framePC = pc
}

func (f *Frame) FramePC() uint {
	return f.framePC
}

// RevertPC 重置帧指针
// 在执行 inst.Execute() 方法之前会将 frame 的 framePC 指针后移
// 而在某些 Execute() 方法中，发送类例如 pushFrame() 操作，为了保证新加入frame会在下次执行
// 就将frame的指针重置为thread的pc，
// 选择重置为thread.pc的而不是简单的进行pc--，因为除了获取操作码会进行pc++,在读取操作数的时候也会进行不同长度的pc++
func (f *Frame) RevertPC() {
	f.framePC = f.Thread.pc
}

func NewFrame(method *klass.MethodKlass, thread *Thread) *Frame {
	attrCode, _ := method.AttrCode()
	maxlocals := attrCode.MaxLocals
	maxStack := attrCode.MaxStack
	return &Frame{
		LocalVars:    NewLocalVars(maxlocals),
		OperandStack: NewOperandStack(maxStack),
		MethodKlass:  method,
		Thread:       thread,
	}
}
