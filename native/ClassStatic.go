package native

import (
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
)

func InitClassStatic() {
	_class(getPrimitiveClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;")
}

func _class(method Method, name, desc string) {
	Register("java/lang/Class", name, desc, method)
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *runtime.Frame) {
	objRef := frame.GetRef(0)
	k := klass.Perm.GetPrimitive(objRef.JString())
	instance := oops.NewOopInstance(k)
	frame.PushRef(instance)
}
