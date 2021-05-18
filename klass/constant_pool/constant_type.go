package constant_pool

import "github.com/zouzhihao-994/gvm/loader"

type ConstantType interface {
	ReadInfo(reader *loader.ClassReader)
}
