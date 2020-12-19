package runtime

import (
	"github.com/zouzhihao-994/gvm/src/share/oops"
)

type Slot struct {
	num int32
	ref *oops.Oop_Instance
}
