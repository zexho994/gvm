package native

import (
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

func Init() {
	_system(setOut0, "setOut0", "(Ljava/io/printStream;)V")
}

func _system(method Method, name, desc string) {
	Register("java/lang/System", name, desc, method)
}

func setOut0(frame *runtime.Frame) {
	out := frame.LocalVars().GetRef(0)
	sysClass := frame.Method().Klass()
	slots := make([]utils.Slot, 2)
	slots[0].Ref = out
	sysClass.StaticVars.SetField("in", slots)
}
