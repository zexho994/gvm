package constant_pool

import "github.com/zouzhihao-994/gvm/classloader"

type ConstantType interface {
	ReadInfo(reader *classloader.ClassReader)
}
