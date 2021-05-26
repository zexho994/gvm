package native

import (
	"sync"
)

var once sync.Once

// InitNativeMethod when invoke StartVM()
func InitNativeMethod() {
	once.Do(func() {
		InitVM()
		InitSystem()
		InitClassStatic()
		InitFloat()
	})
}
