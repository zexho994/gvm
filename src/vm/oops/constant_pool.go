package oops

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/vm/classfile"
)

type Constant interface{}

// 运行时常量池
type ConstantPool struct {
	class  *Class
	consts []Constant
}

/*
将class文件中的常量池转换成运行时常量池
*/
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	// 获取常量池的长度
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	// 常量池结构体
	rtCp := &ConstantPool{class, consts}
	// copy常量池
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		// 对于单数据位的直接提取常量池存到constas
		// 对于两个数据位的索引需要特殊处理
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()

		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClssRef(rtCp, classInfo)

		case *classfile.ConstantFieldRefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)

		case *classfile.ConstantMethodRefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)

		case *classfile.ConstantInterfaceMethodRefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		}
	}
	return rtCp
}

/*
根据索引返回常量
*/
func (self *ConstantPool) GetConstant(index uint) Constant {
	// 在Constant[index]中获取常量
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
