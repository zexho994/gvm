package instructions

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/instructions/references"
)

var (
	getStatic    = &references.GET_STATIC{}
	new          = &references.NEW{}
	invokeStatic = &references.INVOKE_STATIC{}
)

func NewInstruction(opcode byte) base.Base_Instruction {
	switch opcode {
	case 0xb2:
		return getStatic
	case 0xb8:
		return invokeStatic
	case 0xbb:
		return new
	default:
		panic(fmt.Errorf("Unsupported opcode : 0x%x!", opcode))
	}
}
