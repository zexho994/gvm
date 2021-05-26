package runtime

// Thread 映射到java中的一个thread
type Thread struct {
	pc uint
	*Stack
}

func (t Thread) ThreadPC() uint {
	return t.pc
}

func (t *Thread) SetThreadPC(pc uint) {
	t.pc = pc
}

func (t *Thread) RevertFramePC() {
	t.PeekFrame().RevertPC()
}

//func (thread *Thread) InvokeMethodWithShim(method *klass.MethodKlass, args []utils.Slot) {
//	shimFrame := newShimFrame(thread, args)
//	thread.PushFrame(shimFrame)
//	thread.InvokeMethod(method)
//}
//func (thread *Thread) InvokeMethod(method *heap.Method) {
//	//thread._logInvoke(thread.stack.size, method)
//	currentFrame := thread.CurrentFrame()
//	newFrame := thread.NewFrame(method)
//	thread.PushFrame(newFrame)
//	if n := method.ParamSlotCount; n > 0 {
//		_passArgs(currentFrame, newFrame, n)
//	}
//
//	if method.IsSynchronized() {
//		var monitor *heap.Monitor
//		if method.IsStatic() {
//			classObj := method.Class.JClass
//			monitor = classObj.Monitor
//		} else {
//			thisObj := newFrame.GetThis()
//			monitor = thisObj.Monitor
//		}
//
//		monitor.Enter(thread)
//		newFrame.AppendOnPopAction(func(*Frame) {
//			monitor.Exit(thread)
//		})
//	}
//}
