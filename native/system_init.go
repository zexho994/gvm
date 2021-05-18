package native

import (
	"github.com/zouzhihao-994/gvm/native/java/lang"
)

// InitSystemClass execute the method of system calss
func InitSystemClass() {
	lang.Init()
}
