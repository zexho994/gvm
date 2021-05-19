package base

import (
	"github.com/zouzhihao-994/gvm/runtime"
)

// Instruction 指令接口
type Instruction interface {
	// FetchOperands 获取操作数
	FetchOperands(reader *MethodCodeReader)
	// Execute 执行指令
	Execute(frame *runtime.Frame)
}
