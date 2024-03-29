package klass

import (
	"github.com/zouzhihao-994/gvm/klass/attribute"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/loader"
	"github.com/zouzhihao-994/gvm/utils"
)

type Methods []*MethodKlass

type MethodKlass struct {
	accessFlag    uint16
	nameIdx       uint16
	descriptorIdx uint16
	attrCount     uint16
	*attribute.AttributesInfo
	*constant_pool.ConstantPool
	argSlotCount uint
	descriptor   *MethodDescriptor
	*Klass
}

// InjectCodeAttrIfNative injected a code attribute for method
func (m *MethodKlass) InjectCodeAttrIfNative() {
	if !utils.IsNative(m.accessFlag) {
		return
	}

	tmpMaxStack := uint16(4)
	tmpMaxLocal := uint16(4)
	attributes := make(attribute.AttributesInfo, 1)
	methodDescriptor := ParseMethodDescriptor(m.MethodDescriptor())
	var codeAttr *attribute.AttrCode

	switch methodDescriptor.returnTypt {
	case "V":
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xb1}, m.ConstantPool) // return
	case "D":
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xaf}, m.ConstantPool) // dreturn
	case "F":
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xae}, m.ConstantPool) // freturn
	case "J":
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xad}, m.ConstantPool) // lreturn
	case "L", "[":
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xb0}, m.ConstantPool) // areturn
	default:
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xac}, m.ConstantPool) // ireturn
	}

	attributes[0] = codeAttr
	m.AttributesInfo = &attributes
}

func (m MethodKlass) MethodDescriptor() string {
	return m.GetUtf8(m.descriptorIdx)
}

func (m MethodKlass) Descriptor() *MethodDescriptor {
	return m.descriptor
}

// MethodDescriptorArray 将方法描述符由字符串转换成数组
// (Ljava/lang/String;IDJLPrintLong;) -> [1,1,2,2,1,1] : 对应参数的类型需要几个slot槽
//func (m MethodKlass) MethodDescriptorArray() []string {
//}

func (m MethodKlass) DescriptorIdx() uint16 {
	return m.descriptorIdx
}

func (m MethodKlass) MethodName() string {
	return m.GetUtf8(m.nameIdx)
}

func (m MethodKlass) NameIdx() uint16 {
	return m.nameIdx
}

func (m MethodKlass) AccessFlag() uint16 {
	return m.accessFlag
}

func (m MethodKlass) ArgSlotCount() uint {
	return m.argSlotCount
}

func (ms Methods) GetClinitMethod() (*MethodKlass, bool) {
	for idx := range ms {
		i := ms[idx].nameIdx
		nameStr := ms[idx].GetUtf8(i)
		if nameStr == "<clinit>" {
			return ms[idx], true
		}
	}
	return nil, false
}

func (ms Methods) FindMethod(name, desc string) (*MethodKlass, bool) {
	for idx := range ms {
		nameStr := ms[idx].GetUtf8(ms[idx].nameIdx)
		descStr := ms[idx].GetUtf8(ms[idx].descriptorIdx)
		if nameStr == name && descStr == desc {
			return ms[idx], true
		}
	}
	return nil, false
}

// 解析方法表
func parseMethod(count uint16, reader *loader.ClassReader, pool *constant_pool.ConstantPool, k *Klass) Methods {
	methods := make([]*MethodKlass, count)
	for i := range methods {
		method := &MethodKlass{}
		//设置常量池
		method.ConstantPool = pool
		//设置访问标记
		method.accessFlag = reader.ReadUint16()
		//name常量池索引
		method.nameIdx = reader.ReadUint16()
		//描述符常量池索引
		method.descriptorIdx = reader.ReadUint16()
		//属性表数量
		method.attrCount = reader.ReadUint16()
		// 解析方法表中的属性表字段
		method.AttributesInfo = attribute.ParseAttributes(method.attrCount, reader, pool)
		methods[i] = method
		method.descriptor = ParseMethodDescriptor(method.MethodDescriptor())
		method.argSlotCount = method.descriptor.ParamsCount()
		// 绑定klass&method
		method.Klass = k
		// 本地方法注入字节码
		method.InjectCodeAttrIfNative()
	}
	return methods
}

func (m *MethodKlass) IsRegisterNatives() bool {
	return utils.IsStatic(m.accessFlag) && m.MethodName() == "registerNatives" && m.MethodDescriptor() == "()V"
}

func (m *MethodKlass) IsInitIDs() bool {
	return utils.IsStatic(m.accessFlag) && m.MethodName() == "initIDs" && m.MethodDescriptor() == "()V"
}
