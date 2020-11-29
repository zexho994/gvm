package constant_pool

import "github.com/zouzhihao-994/gvm/src/share/classfile"

type ConstantType interface {
	readInfo(reader *classfile.ClassReader)
}
