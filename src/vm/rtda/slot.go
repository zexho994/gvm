package rtda

import "github.com/zouzhihao-994/gvm/src/vm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
