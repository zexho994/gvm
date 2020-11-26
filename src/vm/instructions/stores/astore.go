package stores

import "github.com/zouzhihao-994/gvm/src/vm/instructions/base"
import "github.com/zouzhihao-994/gvm/src/vm/runtime"

/*

 */
type ASTORE struct {
	base.Index8Instruction
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func _astore(frame *runtime.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	//fmt.Printf("[gvm][PushInt] %v 存储到局部变量表[%v]中\n", val, index)
	frame.LocalVars().SetRef(index, val)
}

func (self *ASTORE) Execute(frame *runtime.Frame) {
	_astore(frame, self.Index)
}

func (self *ASTORE_0) Execute(frame *runtime.Frame) {
	_astore(frame, 0)
}

func (self *ASTORE_1) Execute(frame *runtime.Frame) {
	//fmt.Println("[gvm][astore_1] 操作数栈存储数到局部变量表[1]中")
	_astore(frame, 1)
}

func (self *ASTORE_2) Execute(frame *runtime.Frame) {
	_astore(frame, 2)
}

func (self *ASTORE_3) Execute(frame *runtime.Frame) {
	_astore(frame, 3)
}
