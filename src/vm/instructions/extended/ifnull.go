package extended

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/runtime"

type IFNULL struct{ base.BranchInstruction } // Branch if reference is null

type IFNONNULL struct{ base.BranchInstruction } // Branch if reference not null

func (self *IFNULL) Execute(frame *runtime.Frame) {
	ref := frame.OperandStack().PopRef()
	// 如果是null就跳转
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNULL) Execute(frame *runtime.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
