package base

import (
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 指令接口
type Base_Instruction interface {
	// 获取操作数
	FetchOperands(reader *MethodCodeReader)
	// 执行指令
	Execute(frame *runtime.Frame)
}
