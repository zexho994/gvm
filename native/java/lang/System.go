package lang

import (
	"github.com/zouzhihao-994/gvm/native"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

func init() {
	_system(setOut0, "setOut0", "(Ljava/io/printStream;)V")
}

func _system(method native.Method, name, desc string) {
	native.Register("java/lang/System", name, desc, method)
}

func setOut0(frame *runtime.Frame) {
	out := frame.LocalVars().GetRef(0)
	sysClass := frame.Method().JClass()
	slots := make([]utils.Slot, 2)
	slots[0].Ref = out
	sysClass.StaticVars.SetField("in", slots)
}