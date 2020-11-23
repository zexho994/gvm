package references

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
	"github.com/zouzhihao-994/gvm/src/vm/rtda"
)

// 进入监视器，用来实现synchronized关键字
type MONITOR_ENTRY struct {
	base.Index16Instruction
}

// 退出监视器
type MONITOR_EXIT struct {
	base.Index16Instruction
}

func (monitorEntry MONITOR_ENTRY) Execute(frame *rtda.Frame) {
	fmt.Println("monitor entry")
}

func (monitorExit MONITOR_EXIT) Execute(frame *rtda.Frame) {
	fmt.Println("monitor exit")
}
