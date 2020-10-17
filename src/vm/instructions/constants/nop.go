package constants

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
