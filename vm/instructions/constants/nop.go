package constants

import "../base"
import "../../rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) execute(frame *rtda.Frame) {

}
