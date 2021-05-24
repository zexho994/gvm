package native

import (
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
	"sort"
)

func InitSystem() {
	_system(setOut0, "setOut0", "(Ljava/io/PrintStream;)V")
	_system(initProperties, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;")
}

func _system(method Method, name, desc string) {
	Register("java/lang/System", name, desc, method)
}

func setOut0(frame *runtime.Frame) {
	out := frame.GetRef(0)
	sysClass := frame.Method().Klass
	slots := make([]utils.Slot, 2)
	slots[0].Ref = out
	sysClass.StaticFields.SetField("in", slots)
}

func initProperties(frame *runtime.Frame) {
	//props := frame.LocalVars().GetRef(0)
	//frame.OperandStack().PushRef(props)
	//setPropMethod, _, _ := props.Klass().FindMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")

	//sysPropMap := _getSysProps()
	//sysPropKeys := _getSysPropKeys(sysPropMap)

	//for _, key := range sysPropKeys {
	//val := sysPropMap[key]
	//jKey := frame.

	//}
}

func _getSysProps() map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jvm.go",
		"java.vendor.url":      "https://github.com/zxh0/jvm.go",
		"java.home":            "home",
		"java.class.version":   "52.0",
		"java.class.path":      config.ClassPath, // TODO
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              "",   // todo
		"os.arch":              "",   // todo
		"os.version":           "",   // todo
		"file.separator":       "/",  // todo os.PathSeparator
		"path.separator":       ":",  // todo os.PathListSeparator
		"line.separator":       "\n", // todo
		"user.name":            "",   // todo
		"user.home":            "",   // todo
		"user.dir":             ".",  // todo
		"user.country":         "CN", // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
	}
}

func _getSysPropKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
