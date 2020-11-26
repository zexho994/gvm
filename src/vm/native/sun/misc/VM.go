package misc

//import "../../../instructions/base"
//import "../../../native"
//import "../../../runtime"
//import "../../../runtime/oops"
//
//func init() {
//	native.Register("sun/misc/VM", "initialize", "()V", initialize)
//}
//
//// private static native void initialize();
//// ()V
//func initialize(frame *runtime.Frame) { // hack: just make VM.savedProps nonempty
//	vmClass := frame.Method().Class()
//	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
//	key := oops.JString(vmClass.Loader(), "foo")
//	val := oops.JString(vmClass.Loader(), "bar")
//
//	frame.OperandStack().PushRef(savedProps)
//	frame.OperandStack().PushRef(key)
//	frame.OperandStack().PushRef(val)
//
//	propsClass := vmClass.Loader().LoadClass("java/util/Properties")
//	setPropMethod := propsClass.GetInstanceMethod("setProperty",
//		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
//	base.InvokeMethod(frame, setPropMethod)
//}
