package constant_pool

import "github.com/zouzhihao-994/gvm/classfile"

type ConstantType interface {
	ReadInfo(reader *classfile.ClassReader)
}
