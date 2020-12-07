package runtime

import (
	"github.com/zouzhihao-994/gvm/src/share/jclass"
)

type Slot struct {
	num int32
	ref *jclass.JClass_Instance
}
