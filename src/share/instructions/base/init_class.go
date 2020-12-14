package base

import (
	"github.com/zouzhihao-994/gvm/src/share/jclass"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

// 初始化<clinit>方法
func InitClass(j *jclass.JClass_Instance, thread *runtime.Thread) {
	// 获取<clinit>方法
	clinit := j.Methods.Clinit()
	attrCode, err := clinit.Attributes().AttrCode()
	if err != nil {
		panic(err.Error())
	}

	frame := runtime.NewFrame(attrCode.MaxLocals, attrCode.MaxStack, &clinit, thread)
	thread.Push(frame)

	// 如果父类也还未初始化，则先初始化父类
	// 因为栈的原因，所以父类的初始化frame要后push
	if j.SuperClass != nil && !j.SuperClass.IsInit {
		super := j.SuperClass
		InitClass(super, thread)
	}
	j.IsInit = true

}
