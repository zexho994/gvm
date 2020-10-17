package base

import (
	"github.com/zouzhihao-994/gvm/src/vm/rtda"
	"github.com/zouzhihao-994/gvm/src/vm/rtda/heap"
)

/*
初始化类
*/
func InitClass(thread *rtda.Thread, class *heap.Class) {
	// 标志已初始化
	class.StartInit()
	// 准备开始初始化
	// 起名为schedule的而不是invokeClient的原因是不一定是立即执行该类的初始化，而是先执行父类的初始化
	scheduleClinit(thread, class)
	// 初始化父类
	initSuperClass(thread, class)
}

/*
准备初始化
*/
func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	// 获取<client>方法
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

/*
如果存在父类，而且父类还未初始化过，先执行父类的初始化
*/
func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
