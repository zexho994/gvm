package base

import (
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/runtime"
)

// InitClass 初始化<clinit>方法
func InitClass(k *klass.Klass, thread *runtime.Thread) {
	clinit, exist := k.Methods.Clinit()
	k.IsInit = true
	if exist {
		attrCode, err := clinit.Attributes().AttrCode()
		if err != nil {
			panic(err.Error())
		}

		frame := runtime.NewFrame(attrCode.MaxLocals, attrCode.MaxStack, clinit, thread)
		thread.Push(frame)
	}

	// 如果父类也还未初始化，则先初始化父类
	// 因为栈的原因，所以父类的初始化frame要后push
	if k.SuperClass != nil && !k.SuperClass.IsInit {
		InitClass(k.SuperClass, thread)
	}
}