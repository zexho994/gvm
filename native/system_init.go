package native

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/jclass"
	"github.com/zouzhihao-994/gvm/runtime"
)

// InitSystemClass execute the method of system calss
func InitSystemClass(frame *runtime.Frame) {
	initClass := jclass.ParseInstanceByClassName("java/lang/System")
	staticMethod, _ := initClass.FindStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, staticMethod, true)
}
