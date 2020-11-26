package oops

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/vm/classfile"
	"strings"
)

type Class struct {
	// 类访问标志，public,protect,private
	classFile *classfile.ClassFile
	// 类完全限定名
	name string
	// 父类完全限定名
	superClassName string
	// 接口完全限定名
	interfaceNames []string
	// 常量池指针
	constantPool *ConstantPool
	// 字段表
	fields []*Field
	// 方法表
	methods []*Method
	// 类加载器
	loader *ClassLoader
	// 父类
	superClass *Class
	// 接口
	interfaces []*Class
	// 实例字段数量
	instanceSlotCount uint
	// 静态字段数量
	staticSlotCount uint
	// 静态变量数组
	staticVars Slots
	// 判断类是否被加载过
	initStarted bool
	// 在Class中可以获取到类对象，为了完成反射
	jClass *Object // java.lang.Class
}

func (class *Class) JClass() *Object {
	return class.jClass
}

func (class *Class) JavaName() string {
	return strings.Replace(class.name, "/", ".", -1)
}

/*
判断class有没有初始化过
*/
func (class *Class) InitStarted() bool {
	return class.initStarted
}

/*
初始化类
*/
func (class *Class) StartInit() {
	class.initStarted = true
}

func (class *Class) Loader() *ClassLoader {
	return class.loader
}

func (class *Class) Fields() []*Field {
	return class.fields
}

func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}

func (class Class) StaticVars() Slots {
	return class.staticVars
}

func (class Class) IsAbstract() bool {
	return 0 != class.classFile.AccessFlags()&ACC_ABSTRACT
}

/*
将字节码信息转化成Class结构
*/
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.classFile = cf
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields().BaseInfo())
	class.methods = newMethods(class, cf.Methods().BaseInfo())
	return class
}

func (class *Class) IsPublic() bool {
	return 0 != class.classFile.AccessFlags()&ACC_PUBLIC
}

func (class *Class) IsPrivate() bool {
	return 0 != class.classFile.AccessFlags()&ACC_PRIVATE
}

func (class *Class) IsProtected() bool {
	return 0 != class.classFile.AccessFlags()&ACC_PROTECTED
}

func (class *Class) IsInterface() bool {
	return 0 != class.classFile.AccessFlags()&ACC_INTERFACE
}

/**
一个类想要访问另一个类，必须含有
1. 两个类在同一个包下
2. 要访问的类public
*/
func (class *Class) isAccessibleTo(other *Class) bool {
	return class.IsPublic() || class.getPackageName() == other.getPackageName()

}

/**
获取包名
lastIndex : "/" 在 name中出现的最后所以下标
如果name = "java/lang/Object","/"出现的最后位置就是9
*/
func (class *Class) getPackageName() string {
	if i := strings.LastIndex(class.name, "/"); i >= 0 {
		fmt.Printf("[gvm][getPackgeName] 包名 %v \n", class.name[:i])
		return class.name[:i]
	}
	return ""
}

func (class *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := class; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

func (class *Class) GetPackageName() string {
	if i := strings.LastIndex(class.name, "/"); i >= 0 {
		return class.name[:i]
	}
	return ""
}

func (class *Class) NewObject() *Object {
	return newObject(class)
}

func (class *Class) GetMainMethod() *Method {
	return class.getStaticMethod("main", "([Ljava/lang/String;)V")
}

/*
获取静态方法
*/
func (class *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range class.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (class *Class) SuperClass() *Class {
	return class.superClass
}

/*
判断方法的ACC_SUPER是否有被标记
*/
func (class *Class) IsSuper() bool {
	return 0 != class.classFile.AccessFlags()&ACC_SUPER
}

func (class *Class) Name() string {
	return class.name
}

/*

 */
func (class *Class) GetClinitMethod() *Method { return class.getStaticMethod("<clinit>", "()V") }

/*
返回与类对应的数组类
*/
func (class *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(class.name)
	return class.loader.LoadClass(arrayClassName)
}

func (class *Class) isJlObject() bool {
	return class.name == "java/lang/Object"
}

func (class *Class) isJlCloneable() bool {
	return class.name == "java/lang/Cloneable"
}

func (class *Class) isJioSerializable() bool {
	return class.name == "java/io/Serializable"
}

func (class *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[class.name]
	return ok
}
