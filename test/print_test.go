package test

import (
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/launcher"
	"testing"
)

// test i++ ++i
func TestAutoInnc(t *testing.T) {
	config.ClassName = "AutoInnc"
	config.JrePath = config.JrePathDefault
	config.ClassPath = config.UserClassPathDefault + "/print"
	launcher.StartVM()
}

// test if...else..., while, for
func TestCompare(t *testing.T) {
	config.ClassName = "LogicTest"
	config.JrePath = config.JrePathDefault
	config.ClassPath = config.UserClassPathDefault + "/compare"
	launcher.StartVM()
}

// test print
func TestAllPrintOut(t *testing.T) {
	config.ClassName = "PrintOut"
	config.JrePath = config.JrePathDefault
	config.ClassPath = config.UserClassPathDefault + "/print"
	launcher.StartVM()
}

// 测试double print
func TestDoublePrint(t *testing.T) {
	config.ClassName = "PrintLong"
	config.JrePath = config.JrePathDefault
	config.ClassPath = config.UserClassPathDefault + "/print"
	launcher.StartVM()
}

// 测试泛型
func TestPersonGenericity(t *testing.T) {
	config.ClassName = "Person"
	config.JrePath = config.JrePathDefault
	config.ClassPath = config.UserClassPathDefault + "/genericity"
	launcher.StartVM()
}

// 测试泛型上界
func TestSonGenericity(t *testing.T) {
	config.ClassName = "Son"
	config.JrePath = config.JrePathDefault
	config.ClassPath = config.UserClassPathDefault + "/genericity"
	launcher.StartVM()
}

// test invokevirtual
func TestInvoke(t *testing.T) {
	config.ClassName = "InvokeVirtual"
	config.JrePath = config.JrePathDefault
	config.ClassPath = config.UserClassPathDefault + "/invoke"
	launcher.StartVM()
}
