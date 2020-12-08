package instructions

import "github.com/zouzhihao-994/gvm/src/share/runtime"

// 指令接口
type Base_Instruction interface {
	// 获取操作数
	fetch(reader *CodeReader)
	// 执行指令
	execute(frame *runtime.Frame)
}
