package heap

import (
	"github.com/zouzhihao-994/gvm/src/share/oops"
	"sync"
)

// 等同于方法区的概念
// 专门存储 JClass_Instance 对象
type Heap struct {
	size  uint64
	Space map[string]*oops.Oop_Instance
}

var heap *Heap
var once sync.Once

func GetHeap() *Heap {
	once.Do(func() {
		heap = &Heap{Space: make(map[string]*oops.Oop_Instance)}
	})
	return heap
}

func (h *Heap) Add(instance *oops.Oop_Instance) {

}
