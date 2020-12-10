package launcher

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	jclass "github.com/zouzhihao-994/gvm/src/share/jclass"
)

// 通过debug模式启动gvm
func StartGvmByDebug() {
	loader := classfile.InitClassLoader(JrePath, UserClassPath)
	bytecode := loader.Loading("classFile")
	jc := jclass.ParseToJClass(bytecode)
	instance := jclass.ParseInstance(jc)
	// 执行main方法
	method, err := instance.FindStaticMethod("main")
	if err != nil || method == nil {
		panic(err)
	}
	code, err := method.Attributes().Code()
	if err != nil {
		panic(err)
	}
	if code == nil {

	}

}
