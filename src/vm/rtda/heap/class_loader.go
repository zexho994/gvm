package heap

import (
	"github.com/zouzhihao-994/gvm/src/vm/classfile"
	"github.com/zouzhihao-994/gvm/src/vm/loader"
)

/*
类加载器
类加载器依赖ClassPath来搜寻和读取class文件
*/
type ClassLoader struct {
	// 保存cp指针
	loader *loader.Loader

	// 已经加载的类，key是类的全限定名
	classMap map[string]*Class

	// 是否控制台打印
	verboseFlag bool
}

/*
创建一个加载器实例
*/
func NewClassLoader(loader *loader.Loader, verboseFlag bool) *ClassLoader {
	classLoader := &ClassLoader{
		loader:      loader,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}

	//  加载基础类
	classLoader.loadBasicClasses()
	classLoader.loadPrimitiveClasses()
	return classLoader
}

// 加载基础类
func (classLoader *ClassLoader) loadBasicClasses() {
	jlClassClass := classLoader.LoadClass("java/lang/Class")
	// 类加载到map中
	for _, class := range classLoader.classMap {
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
func (classLoader *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		classLoader.loadPrimitiveClass(primitiveType)
	}
}

/*
生成void和基本类型类
*/
func (classLoader *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		name:        className,
		loader:      classLoader,
		initStarted: true,
	}
	class.jClass = classLoader.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	classLoader.classMap[className] = class
}

/*
在 classMap 中根据 classpath 查询类
然后将类加载到方法区中
*/
func (classLoader *ClassLoader) LoadClass(classPath string) *Class {
	if class, ok := classLoader.classMap[classPath]; ok {
		// already loaded
		return class
	}

	var class *Class
	// '['的为数组类型，否则为非数组（普通）类型
	if classPath[0] == '[' {
		class = classLoader.loadArrayClass(classPath)
	} else {
		class = classLoader.loadNonArrayClass(classPath)
	}

	if jlClassClass, ok := classLoader.classMap["java/lang/Class"]; ok {
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
func (classLoader *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		name:        name,
		loader:      classLoader,
		initStarted: true,
		superClass:  classLoader.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			classLoader.LoadClass("java/lang/Cloneable"),
			classLoader.LoadClass("java/io/Serializable"),
		},
	}
	classLoader.classMap[name] = class
	return class
}

// 非数组类的加载
// 调用classpath的readClass方法，
// 该方法会按顺序从bootClasspath,extClassapath，userClasspath中根据name查找class文件
// data是class的二进制数据
func (classLoader *ClassLoader) loadNonArrayClass(classPath string) *Class {
	data, entry := classLoader.readClass(classPath)
	if entry == nil {
		panic("entry is nil")
	}
	// 将二进制数据解析成Class结构体
	class := classLoader.defineClass(data)
	// 类的链接
	link(class)
	return class
}

/*
在classpath中搜索名称为name的类
*/
func (classLoader *ClassLoader) readClass(classPath string) ([]byte, loader.Entry) {
	data, entry, err := classLoader.loader.ReadClass(classPath)
	if err != nil {
		panic("java.lang.ClassNotFoundException:" + classPath)
	}
	return data, entry
}

/*
将二进制数据解析成Class结构体
*/
func (classLoader *ClassLoader) defineClass(data []byte) *Class {
	// 将类的数据转换成类结构体
	class := parseClass(data)
	// 设置类的加载器
	// 所以判断一个类是否相等还需要判断类加载器是否相等
	class.loader = classLoader
	// 解析父类以及接口
	resolveSuperClass(class)
	resolveInterfaces(class)
	// classMap相当于方法区
	// key为class的全限制定名，value为class结构体
	classLoader.classMap[class.name] = class
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
}

/*
准备阶段做两件事
设置初始值和分配内存
设置初始值是给静态变量设置初始值，非final修饰的
*/
func prepare(class *Class) {
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
