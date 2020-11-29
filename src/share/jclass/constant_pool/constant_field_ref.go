package constant_pool

import "github.com/zouzhihao-994/gvm/src/share/jclass"

type ConstantFieldRefInfo struct {
	tag              uint8
	cp               jclass.ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}
