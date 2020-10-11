package rtda

import "./heap"

type Slot struct {
	num int32
	ref *heap.Object
}
