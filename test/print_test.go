package test

import (
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/launcher"
	"testing"
)

func TestAllPrintOut(t *testing.T) {
	config.ClassName = "PrintOut"
	config.JrePath = config.JrePathDefault
	config.ClassPath = config.UserClassPathDefault + "/print"
	launcher.StartVM()
}

func TestDoublePrint(t *testing.T) {
	config.ClassName = "PrintLong"
	config.JrePath = config.JrePathDefault
	config.ClassPath = config.UserClassPathDefault + "/print"
	launcher.StartVM()
}
