package heap

import (
	"github.com/zouzhihao-994/gvm/src/share/oops"
	"sync"
)

// 等同于方法区的概念
// 专门存储 JClass_Instance 对象
type Heap struct {
	size  uint64
	Space map[*oops.OopInstance]*oops.OopInstance
}

var heap *Heap
var once sync.Once

func GetHeap() *Heap {
	once.Do(func() {
		heap = &Heap{Space: make(map[*oops.OopInstance]*oops.OopInstance)}
	})
	return heap
}

func (h *Heap) Add(instance *oops.OopInstance) {
	h.Space[instance] = instance
}

func (h *Heap) Get(instance *oops.OopInstance) *oops.OopInstance {
	return h.Space[instance]
}
