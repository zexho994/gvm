package klass

import (
	"github.com/zouzhihao-994/gvm/exception"
	"github.com/zouzhihao-994/gvm/klass/attribute"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/loader"
	"github.com/zouzhihao-994/gvm/utils"
)

// Klass 存储在方法区中的对象，也是 Klass 经过链接步骤后得到的对象
// 同一个的类或接口的所有子类/实现类对该部分的依赖都会是同一个对象，即不会存在两个一样的 Klass 对象
type Klass struct {
	// 魔术
	Magic uint32
	// 次版本
	MinorVersion uint16
	// 主版本
	MajorVersion uint16
	// 常量池
	ConstantPoolCount uint16
	*constant_pool.ConstantPool
	// 类访问标志,表示是类还是接口,public还是private等
	AccessFlags uint16
	// 本类
	ThisClassIdx uint16
	ThisClass    string
	// 父类
	SuperClassIdx uint16
	SuperClass    *Klass
	// 接口
	InterfacesCount uint16
	InterfaceUints  []uint16
	Interfaces      []*Klass
	// 字段表,用于表示接口或者类中声明的变量
	FieldsCount uint16
	Fields
	// 方法表
	MethodsCount uint16
	Methods      Methods
	// 属性表
	AttributesCount uint16
	Attributes      attribute.AttributesInfo

	// 初始化标识
	IsInit     bool
	StaticVars *StaticFieldVars
}

// ParseToKlass
// TODO 如果后面什么时候引入多线程了，这个地方要注意线程安全问题，可能存在多个线程同时执行一个 Klass 的解析
func ParseToKlass(reader *loader.ClassReader) *Klass {
	kl := &Klass{}
	// CAFEBABY
	kl.Magic = parseMagic(reader)
	// jdk version
	kl.MinorVersion = parseMinorVersion(reader)
	kl.MajorVersion = paresMajorVersion(reader)
	// 常量池
	kl.ConstantPoolCount = reader.ReadUint16()
	cp := constant_pool.ReadConstantPool(kl.ConstantPoolCount, reader)
	kl.ConstantPool = &cp
	// 类访问符
	kl.AccessFlags = reader.ReadUint16()
	// 本类
	kl.ThisClassIdx = reader.ReadUint16()
	kl.ThisClass = kl.ConstantPool.GetClassName(kl.ThisClassIdx)
	// 父类
	kl.SuperClassIdx = reader.ReadUint16()
	kl.SuperClass = kl.parseSuper()
	// 接口数量 & 列表
	kl.InterfacesCount = reader.ReadUint16()
	kl.InterfaceUints = reader.ReadUint16Array(kl.InterfacesCount)
	kl.Interfaces = kl.parseInterfaces()
	// 字段数量 & 列表
	kl.FieldsCount = reader.ReadUint16()
	kl.Fields = parseFields(kl.FieldsCount, reader, kl.ConstantPool)
	// 方法数量 & 列表
	kl.MethodsCount = reader.ReadUint16()
	kl.Methods = parseMethod(kl.MethodsCount, reader, kl.ConstantPool, kl)
	// 属性数量 & 列表
	kl.AttributesCount = reader.ReadUint16()
	kl.Attributes = attribute.ParseAttributes(kl.AttributesCount, reader, kl.ConstantPool)
	// 默认未初始化，只有在进行实际调用该类的时候才进行初始化
	kl.IsInit = false
	// 保存到方法区
	Perm().Space[kl.ThisClass] = kl
	// 执行链接步骤
	kl.Linked()
	kl.init()
	return kl
}

func parseMagic(reader *loader.ClassReader) uint32 {
	magic := reader.ReadUint32()
	if magic != 0xCAFEBABE {
		panic("[gvm] this file is not support")
	}
	return magic
}

func parseMinorVersion(reader *loader.ClassReader) uint16 {
	return reader.ReadUint16()
}

func paresMajorVersion(reader *loader.ClassReader) uint16 {
	return reader.ReadUint16()

}

func ParseByClassName(className string) *Klass {
	if k := Perm().Space[className]; k != nil {
		return k
	}

	bytecode := loader.Loading(className)
	reader := &loader.ClassReader{Bytecode: bytecode}
	klass := ParseToKlass(reader)

	Perm().Space[className] = klass
	return klass
}

// 递归解析父类
// todo: parseSuper 和 parseInterfaces 都需要对访问权限进行判断
func (k *Klass) parseSuper() *Klass {
	thisName := k.ConstantPool.GetClassName(k.ThisClassIdx)
	if thisName == "java/lang/Object" {
		return nil
	}
	// 判断是否存在父类
	superName := k.ConstantPool.GetClassName(k.SuperClassIdx)
	// 方法区存在该类结构
	perm := Perm()
	if supre := perm.Space[superName]; supre != nil {
		return supre
	}
	return ParseByClassName(superName)
}

// 递归解析接口
func (k *Klass) parseInterfaces() []*Klass {
	if k.InterfacesCount < 1 {
		return nil
	}

	interfaces := make([]*Klass, k.InterfacesCount)
	for i := range k.Interfaces {
		iIdx := k.InterfaceUints[i]
		iName := k.ConstantPool.GetClassName(iIdx)
		iInstance := &Klass{}
		// 如果方法区中已经有直接引用
		if iInstance = Perm().Space[iName]; iInstance != nil {
			interfaces[i] = iInstance
			continue
		}
		// 没有的情况，进行接口类的加载
		instance := ParseByClassName(iName)
		// 接口类型验证
		if !utils.IsInterface(instance.AccessFlags) {
			panic("[gvm] 接口解析错误 :" + iName + "的父接口对象不为 interface 类型")
		}
		interfaces[i] = instance
	}

	return interfaces
}

func (k *Klass) FindStaticMethod(name, descriptor string) (*MethodInfo, error) {
	for i := range k.Methods {
		methodInfo := k.Methods[i]
		if !utils.IsStatic(methodInfo.AccessFlag()) {
			continue
		}
		if name != k.ConstantPool.GetUtf8(methodInfo.NameIdx()) ||
			descriptor != k.ConstantPool.GetUtf8(methodInfo.DescriptorIdx()) {
			continue
		}
		methodInfo.SetKlass(k)
		return methodInfo, nil
	}
	return nil, exception.GvmError{Msg: "not find static method it name " + name}
}

// FindMethod TODO:可以从父类中加载出方法，并检查权限
// name: method method
// @return the MethodInfo belong to the Klass
func (k *Klass) FindMethod(name, descriptor string) (*MethodInfo, error, *Klass) {
	if k == nil {
		return nil, nil, nil
	}
	for _, methodInfo := range k.Methods {
		if utils.IsStatic(methodInfo.AccessFlag()) {
			continue
		}
		mName := k.ConstantPool.GetUtf8(methodInfo.NameIdx())
		mDesc := k.ConstantPool.GetUtf8(methodInfo.DescriptorIdx())
		if mName == name && mDesc == descriptor {
			return methodInfo, nil, k
		}
	}

	// 在父类中遍历查找
	m, err, jc := k.SuperClass.FindMethod(name, descriptor)
	if err == nil {
		return m, nil, jc
	}

	// 在接口中遍历查找
	for i := range k.Interfaces {
		m, err, jc := k.Interfaces[i].FindMethod(name, descriptor)
		if err == nil {
			return m, nil, jc
		}
	}

	return nil, exception.GvmError{Msg: "not find method it name " + name}, nil
}

// Linked 链接阶段，分为3部分，验证，准备，解析
func (k *Klass) Linked() {
	k.verify()
	k.prepare()
	k.parse()
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
func (k *Klass) verify() {

}

// 准备阶段主要是为准备静态变量
// 主要负责两件事：
// 	1. 分配内存
// 	2. 设置零值
//
// 注意点：如果静态变量类型为引用类型，则零值为null，且不需要判断引用类是否进行类类加载过程。
//        这部分的判断逻辑在putstatic,getstatic指令时再执行.
func (k *Klass) prepare() {
	jFields := k.Fields
	vars := NewStaticFieldVars()
	for idx := range jFields {
		// 不处理实例变量
		if !utils.IsStatic(k.Fields[idx].AccessFlags) {
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
	k.StaticVars = vars
}

func (k *Klass) parse() {

}

func (k *Klass) init() {

}
