package attribute

import (
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/loader"
)

// AttrExceptions Exceptions属性是变长属性，位于 MethodKlass 结构中
// 表示一个方法可能抛出的受检查异常(checked exception)
// 一个 MethodKlass 结果的属性表中最多一个 AttrExceptions 属性
// 如果一个方法要抛出异常，需要满足3个条件之一
// 		1. 抛出的是RuntimeException或者子类的实例
// 		2. 要抛出的是Error或子类的实例
// 		3. 要抛出的是exTable中声明的异常类或者子类
type AttrExceptions struct {
	// 名称索引
	nameIdx uint16
	// 不包括在bytecode中，nameIdx解析后的数据
	name string
	// 当前属性长度，不包括初始的6字节
	attrlen uint32
	// table 长度
	exCount uint16
	// 每个成员都是常量池中一个有效的索引
	// 而且类型是 Constant_Class_info，表示要抛出的类的类型
	exTable []uint16
	*constant_pool.ConstantPool
}

func (attr *AttrExceptions) parse(reader *loader.ClassReader) {
	exCount := reader.ReadUint16()
	attr.exTable = make([]uint16, exCount)
	for i := range attr.exTable {
		attr.exTable[i] = reader.ReadUint16()
	}
}

func (attr AttrExceptions) Name() string {
	return attr.name
}
