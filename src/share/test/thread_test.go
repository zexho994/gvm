package test

import (
	"github.com/zouzhihao-994/gvm/src/share/launcher"
	"testing"
)

func TestThread(t *testing.T) {
	launcher.StartGvmByDebug("ThreadTest", "/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre", "/Users/zexho/GolandProjects/gvm/java/src/test")
}
