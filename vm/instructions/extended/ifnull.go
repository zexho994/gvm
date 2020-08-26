package extended

import "../../instructions/base"
import "../../rtda"

type IFNULL struct{ base.BranchInstruction }    // Branch if reference is null
type IFNONNULL struct{ base.BranchInstruction } // Branch if reference not null

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}
