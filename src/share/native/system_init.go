package native

import (
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// execute the method of system calss
func InitSystemClass(frame *runtime.Frame) {
	// load class

	initClass := jclass.ParseInstanceByClassName("java/lang/System")
	staticMethod, _ := initClass.FindStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, staticMethod, true)

	// execute method
}
