package test

import (
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/launcher"
	"testing"
)

func TestIntPrintOut(t *testing.T) {
	config.ClassName = "PrintOut"
	config.JrePath = config.JrePathDefault
	config.ClassPath = config.UserClassPathDefault + "/print"
	launcher.StartVM()
}
