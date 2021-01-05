package test

import (
	"github.com/zouzhihao-994/gvm/src/share/launcher"
	"testing"
)

// print test
func TestPrint(t *testing.T) {
	launcher.StartGvmByDebug("PrintTest", "/Library/Java/JavaVirtualMachines/jdk1.8.0_261.jdk/Contents/Home/jre", "/Users/zexho/GolandProjects/gvm/java/src")
}
