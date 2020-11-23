package heap

import (
	"github.com/zouzhihao-994/gvm/src/vm/classfile"
	"github.com/zouzhihao-994/gvm/src/vm/classpath"
)

/*
类加载器
类加载器依赖ClassPath来搜寻和读取class文件
*/
type ClassLoader struct {
	// 保存cp指针
	cp *classpath.Loader

	// 已经加载的类，key是类的全限定名
	classMap map[string]*Class

	// 是否控制台打印
	verboseFlag bool
}

/*
创建一个加载器实例
*/
func NewClassLoader(cp *classpath.Loader, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}

	//  加载基础类
	loader.loadBasicClasses()
	//
	loader.loadPrimitiveClasses()
	return loader
}

// 加载基础类
func (cl *ClassLoader) loadBasicClasses() {
	jlClassClass := cl.LoadClass("java/lang/Class")
	// 类加载到map中
	for _, class := range cl.classMap {
		if class.jClass == nil {
			class.jClass = jlClassClass.NewObject()
			class.jClass.extra = class
		}
	}
}

/*
加载void和基本类型的类
在基本类型的包装类中，例如Integer，都有一个Type字段。
Type字段存放的就是基本类型的类
*/
func (cl *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		cl.loadPrimitiveClass(primitiveType)
	}
}

/*
生成void和基本类型类
*/
func (cl *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        className,
		loader:      cl,
		initStarted: true,
	}
	class.jClass = cl.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	cl.classMap[className] = class
}

/*
在classMap中根据name查询类
然后将将类数据加载到方法区中
*/
func (cl *ClassLoader) LoadClass(name string) *Class {
	if class, ok := cl.classMap[name]; ok {
		// already loaded
		return class
	}

	var class *Class
	if name[0] == '[' { // array class
		class = cl.loadArrayClass(name)
	} else {
		class = cl.loadNonArrayClass(name)
	}

	if jlClassClass, ok := cl.classMap["java/lang/Class"]; ok {
		class.jClass = jlClassClass.NewObject()
		class.jClass.extra = class
	}

	return class
}

/*
数组类型的加载方法
超类是Object类
父接口是Cloneable和Serializable
*/
func (cl *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        name,
		loader:      cl,
		initStarted: true,
		superClass:  cl.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			cl.LoadClass("java/lang/Cloneable"),
			cl.LoadClass("java/io/Serializable"),
		},
	}
	cl.classMap[name] = class
	return class
}

/*
非数组类的加载
*/
func (cl *ClassLoader) loadNonArrayClass(name string) *Class {
	// 调用classpath的readClass方法，
	// 该方法会按顺序从bootClasspath,extClassapath，userClasspath中根据name查找class文件
	// data是class的二进制数据
	data, entry := cl.readClass(name)
	if entry == nil {
		panic("entry is nil")
	}
	// 将二进制数据解析成Class结构体
	class := cl.defineClass(data)
	// 类的链接
	link(class)
	if cl.verboseFlag {
	}
	return class
}

/*
在classpath中搜索名称为name的类
*/
func (cl *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := cl.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException:" + name)
	}
	return data, entry
}

/*
将二进制数据解析成Class结构体
*/
func (cl *ClassLoader) defineClass(data []byte) *Class {
	// 将类的数据转换成类结构体
	class := parseClass(data)
	// 设置类的加载器
	// 所以判断一个类是否相等还需要判断类加载器是否相等
	class.loader = cl
	// 解析父类以及接口
	resolveSuperClass(class)
	resolveInterfaces(class)
	// classMap相当于方法区
	// key为class的全限制定名，value为class结构体
	cl.classMap[class.name] = class
	return class
}

/*
将类数据解析成类结构体
*/
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

/*
解析超类的符号引用
如果父类不是Object类，就获取父类的名称然后调用类加载器加载父类
*/
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

/*
解析接口的符号引用
获取接口名数组 -> 对每一个父接口进行加载
*/
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

/*
链接阶段可以分为3个步骤
1 验证：检查class文件
2 准备：
3 解析：
*/
func link(class *Class) {
	verify(class)
	prepare(class)
	//resolution()
}

/*
验证阶段
*/
func verify(class *Class) {
	//fmt.Printf("[gvm][verify] 类加载-验证阶段\n")
}

/*
准备阶段做两件事
设置初始值和分配内存
设置初始值是给静态变量设置初始值，非final修饰的
*/
func prepare(class *Class) {
	//fmt.Printf("[gvm][verify] 类加载-准备阶段\n")
	// 计算实例字段数量
	calcInstanceFieldSlotIds(class)
	// 计算静态字段数量
	calcStaticFieldSlotIds(class)
	// 分配空间
	allocAndInitStaticVars(class)
}

/**
计算实例字段数量
1. 父类的字段都属于子类。子类的字段表需要加上父类的字段
2.
*/
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		//fmt.Printf("[gvm][calcInstanceFieldSlotIds]%v,%v ", field.name, slotId)
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	//fmt.Printf("[gvm][calcInstanceFieldSlotIds] 实例字段数量: %v\n", slotId)
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	//fmt.Printf("[gvm][calcStaticFieldSlotIds] 静态字段数量 %v\n", slotId)

	class.staticSlotCount = slotId
}

/**
给类变量分配空间，然后赋予初始值
*/
func allocAndInitStaticVars(class *Class) {
	//fmt.Printf("[gvm][allocAndInitStaticVars] class: %v, staticCount %v 分配空间 \n", class, class.staticSlotCount)
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		// 对于常量类型，值在编译时期已经存在了字段的attribute表里面
		// 所以在初始化的时候直接给常量赋值
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

/**
类变量的值在编译时候就已知，所以可以直接从class文件常量池中获取
*/
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}
