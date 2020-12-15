package instructions

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/instructions/control"
	"github.com/zouzhihao-994/gvm/src/share/instructions/references"
)

var (
	voidReturn   = &control.RETURN{}           //177
	getStatic    = &references.GET_STATIC{}    // 178
	invokeStatic = &references.INVOKE_STATIC{} // 184
	_new         = &references.NEW{}           //187
)

func NewInstruction(opcode byte) base.Base_Instruction {
	switch opcode {
	case 0xb1:
		return voidReturn
	case 0xb2:
		return getStatic
	case 0xb8:
		return invokeStatic
	case 0xbb:
		return _new
	default:
		panic(fmt.Errorf("Unsupported opcode : 0x%x!", opcode))
	}
}
