package runtime

import "github.com/zouzhihao-994/gvm/src/vm/oops"

type Slot struct {
	num int32
	ref *oops.Object
}
