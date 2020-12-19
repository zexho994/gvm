package references

import (
	"github.com/zouzhihao-994/gvm/src/vm/instructions/base"
)
import "github.com/zouzhihao-994/gvm/src/vm/runtime"
import "github.com/zouzhihao-994/gvm/src/vm/oops"

/*
用于调用构造函数
*/
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *runtime.Frame) {
	// 拿到当前类的类，常量池，方法符号引用
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*oops.MethodRef)

	// 解析类和方法
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	// 判断条件1： 如果方法是构造方法，那么resolvedClass必须是resolvedMethod的类
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	// 判断条件2：如果是静态方法，异常
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 获取this引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	// protected方法只能被本类或者子类调用
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := resolvedMethod
	// 当前类的ACC_SUPER被标记，而且调用方法的类是本类的父类，不是构造函数
	// 就需要新建一个过程去从父类的方法表中寻找最终的调用方法
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = oops.LookupMethodInClass(
			currentClass.SuperClass(),
			methodRef.Name(),
			methodRef.Descriptor(),
		)
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
