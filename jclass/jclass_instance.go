package jclass

import (
	"github.com/zouzhihao-994/gvm/classfile"
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/jclass/attribute"
	"github.com/zouzhihao-994/gvm/jclass/constant_pool"
	"github.com/zouzhihao-994/gvm/utils"
)

// 存储在方法区中的对象，也是 JClass 经过链接步骤后得到的对象
// 同一个的类或接口的所有子类/实现类对该部分的依赖都会是同一个对象，即不会存在两个一样的 JClassInstance 对象
type JClassInstance struct {
	// 常量池
	ConstantPool constant_pool.ConstantPool
	// 类访问标志,表示是类还是接口,public还是private等
	AccessFlags uint16
	// 本类
	ThisClass string
	// 父类
	SuperClass *JClassInstance
	// 接口
	Interfaces []*JClassInstance
	// 字段表,用于表示接口或者类中声明的变量
	Fields Fields
	// 方法表
	Methods Methods
	// 属性表
	Attributes attribute.AttributesInfo
	// 初始化标识
	IsInit     bool
	StaticVars *StaticFieldVars
}

// TODO 如果后面什么时候引入多线程了，这个地方要注意线程安全问题，可能存在多个线程同时执行一个 JClassInstance 的解析
func ParseInstance(jclass *JClass) *JClassInstance {
	jci := &JClassInstance{}
	// todo 解析运行时常量池,将 CONSTANT_Class_info,CONSTANT_Fieldref_info,CONSTANT_Methodref_info等类型符号引用的常量转换为直接引用
	jci.ConstantPool = jclass.ConstantPool
	// 类访问符不变
	jci.AccessFlags = jclass.AccessFlags
	// 获取到全限定名
	jci.ThisClass = jci.ConstantPool.GetClassName(jclass.ThisClassIdx)
	// 加载父类
	jci.SuperClass = parseSuper(jclass)
	// 加载接口
	jci.Interfaces = parseInterfaces(jclass)
	// TODO parse field
	jci.Fields = jclass.Fields
	// TODO parse methods
	jci.Methods = jclass.Methods
	jci.Attributes = jclass.Attributes
	// 默认未初始化，只有在进行实际调用该类的时候才进行初始化
	jci.IsInit = false
	// 保存到方法区
	GetPerm().Space[jci.ThisClass] = jci
	// 执行链接步骤
	jci.Linked()
	return jci
}

func ParseInstanceByClassName(className string) *JClassInstance {
	if perm := GetPerm().Space[className]; perm != nil {
		return perm
	}
	bytecode := classfile.ClaLoader.Loading(className)
	jclass := ParseToJClass(bytecode)
	jci := ParseInstance(jclass)
	perm.Space[className] = jci
	return jci
}

// 递归解析父类
// todo: parseSuper 和 parseInterfaces 都需要对访问权限进行判断
func parseSuper(jclass *JClass) *JClassInstance {
	thisName := jclass.ConstantPool.GetClassName(jclass.ThisClassIdx)
	if thisName == "java/lang/Object" {
		return nil
	}
	// 判断是否存在父类
	superName := jclass.ConstantPool.GetClassName(jclass.SuperClassIdx)
	// 方法区存在该类结构
	perm := GetPerm()
	if supre := perm.Space[superName]; supre != nil {
		return supre
	}
	return ParseInstanceByClassName(superName)
}

// 递归解析接口
func parseInterfaces(jclass *JClass) []*JClassInstance {
	if jclass.InterfacesCount < 1 {
		return nil
	}

	interfaces := make([]*JClassInstance, jclass.InterfacesCount)
	for i := range jclass.Interfaces {
		iIdx := jclass.Interfaces[i]
		iName := jclass.ConstantPool.GetClassName(iIdx)
		iInstance := &JClassInstance{}
		// 如果方法区中已经有直接引用
		if iInstance = perm.Space[iName]; iInstance != nil {
			interfaces[i] = iInstance
			continue
		}
		// 没有的情况，进行接口类的加载
		instance := ParseInstanceByClassName(iName)
		// 接口类型验证
		if !IsInterface(instance.AccessFlags) {
			panic("[gvm] 接口解析错误 :" + iName + "的父接口对象不为 interface 类型")
		}
		interfaces[i] = instance
	}

	return interfaces
}

func (j *JClassInstance) FindStaticMethod(name, descriptor string) (*MethodInfo, error) {
	for i := range j.Methods {
		methodInfo := j.Methods[i]
		if !IsStatic(methodInfo.accessFlag) {
			continue
		}
		if name != j.ConstantPool.GetUtf8(methodInfo.nameIdx) ||
			descriptor != j.ConstantPool.GetUtf8(methodInfo.descriptorIdx) {
			continue
		}
		methodInfo.jclass = j
		return methodInfo, nil
	}
	return nil, exception.GvmError{Msg: "not find static method it name " + name}
}

// TODO:可以从父类中加载出方法，并检查权限
// name: method method
// @return the MethodInfo belong to the JClassInstance
func (j *JClassInstance) FindMethod(name, descriptor string) (*MethodInfo, error, *JClassInstance) {
	for i := range j.Methods {
		methodInfo := j.Methods[i]
		if IsStatic(methodInfo.accessFlag) {
			continue
		}
		mName := j.ConstantPool.GetUtf8(methodInfo.nameIdx)
		mDesc := j.ConstantPool.GetUtf8(methodInfo.descriptorIdx)
		if mName == name && mDesc == descriptor {
			return j.Methods[i], nil, j
		}
	}
	// 在父类中遍历查找
	m, err, jc := j.SuperClass.FindMethod(name, descriptor)
	if err == nil {
		return m, nil, jc
	}
	// 在接口中遍历查找
	for i := range j.Interfaces {
		m, err, jc := j.Interfaces[i].FindMethod(name, descriptor)
		if err == nil {
			return m, nil, jc
		}
	}
	return nil, exception.GvmError{Msg: "not find method it name " + name}, nil
}

// 链接阶段，分为3部分，验证，准备，解析
func (j *JClassInstance) Linked() {
	j.jci_verify()
	j.jci_prepare()
	j.jci_parse()
}

// 验证，目的是确保类或者接口的二进制表示在结构上是正确的
// 验证过程可能导致某些额外的类或者接口被加载进来，但不一定会导致它们也需要验证或准备
// 具体的验证范围包括一下几种：
// 1. 方法的访问控制
// 2. 参数和静态类型检查
// 3. 堆栈是否滥用
// 4. 变量是否初始化
// 5. 变量是否赋予正确类型
// 6. 异常表必须引用类合法的指令
// 7. 验证局部变量表
// 8. 逐一验证每个字节码的合法性
func (j *JClassInstance) jci_verify() {

}

// 准备阶段主要是为准备静态变量
// 主要负责两件事：
// 	1. 分配内存
// 	2. 设置零值
//
// 注意点：如果静态变量类型为引用类型，则零值为null，且不需要判断引用类是否进行类类加载过程。
//        这部分的判断逻辑在putstatic,getstatic指令时再执行.
func (j *JClassInstance) jci_prepare() {
	jFields := j.Fields
	vars := NewStaticFieldVars()
	for idx := range jFields {
		// 不处理实例变量
		if !IsStatic(j.Fields[idx].AccessFlags) {
			continue
		}
		var slot utils.Slot
		desc := jFields[idx].Descriptor()
		switch desc {
		case "I":
			slot = utils.Slot{Num: 0, Type: utils.SlotInt, Ref: nil}
		case "B":
			slot = utils.Slot{Num: 0, Type: utils.SlotByte, Ref: nil}
		case "D":
			slot = utils.Slot{Num: 0, Type: utils.SlotDouble, Ref: nil}
		case "F":
			slot = utils.Slot{Num: 0, Type: utils.SlotFloat, Ref: nil}
		case "J":
			slot = utils.Slot{Num: 0, Type: utils.SlotLong, Ref: nil}
		case "S":
			slot = utils.Slot{Num: 0, Type: utils.SlotShort, Ref: nil}
		case "Z":
			slot = utils.Slot{Num: 0, Type: utils.SlotBoolean, Ref: nil}
		case "L":
			exception.GvmError{Msg: "jclass prepare Error"}.Throw()
		default:
			slot = utils.Slot{Num: 0, Type: utils.SlotRef, Ref: nil}
		}

		vars.AddField(jFields[idx].Name(), slot)
	}
	j.StaticVars = vars
}

func (j *JClassInstance) jci_parse() {

}

func (j JClassInstance) Name() string {
	return j.ThisClass
}
