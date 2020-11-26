package oops

const (
	// public
	ACC_PUBLIC = 0x0001
	// private
	ACC_PRIVATE = 0x0002
	// protected
	ACC_PROTECTED = 0x0004
	// static
	ACC_STATIC = 0x0008
	// final
	ACC_FINAL = 0x0010
	// synchronized
	ACC_SYNCHRONIZED = 0x0020
	// volatile
	ACC_VOLATILE = 0x0020
	// super
	ACC_SUPER = 0x0020
	// bridge
	ACC_BRIDGE = 0x0040
	// transient
	ACC_TRANSIENT = 0x0040
	// varargs
	ACC_VARARGS = 0x0080
	// native
	ACC_NATIVE = 0x0100
	// interface
	ACC_INTERFACE = 0x0200
	// abstract
	ACC_ABSTRACT = 0x0400
	// strict
	ACC_STRICT = 0x0800
	// synthetic
	ACC_SYNTHETIC = 0x1000
	// annotation
	ACC_ANNOTATION = 0x2000
	// enum
	ACC_ENUM = 0x4000
)

func checkClassAccess(class *Class) {
	if 0 != class.classFile.AccessFlags()&ACC_FINAL {
		//fmt.Printf("[gvm][checkClassAccess] 类%v 是 final 的\n", class.name)
	}
	if 0 != class.classFile.AccessFlags()&ACC_ABSTRACT {
		//fmt.Printf("[gvm][checkClassAccess] 类%v 是 abstract 的\n", class.name)
	}
	if 0 != class.classFile.AccessFlags()&ACC_INTERFACE {
		//fmt.Printf("[gvm][checkClassAccess] 类%v 是 Interface 的\n", class.name)
	}
	if 0 != class.classFile.AccessFlags()&ACC_ENUM {
		//fmt.Printf("[gvm][checkClassAccess] 类%v 是 enum 的\n", class.name)
	}
	if 0 != class.classFile.AccessFlags()&ACC_PUBLIC {
		//fmt.Printf("[gvm][checkClassAccess] 类%v 是 public 的\n", class.name)
	}
	if 0 != class.classFile.AccessFlags()&ACC_PRIVATE {
		//fmt.Printf("[gvm][checkClassAccess] 类%v 是 private 的\n", class.name)
	}
	if 0 != class.classFile.AccessFlags()&ACC_PROTECTED {
		//fmt.Printf("[gvm][checkClassAccess] 类%v 是 protected 的\n", class.name)
	}
	if 0 != class.classFile.AccessFlags()&ACC_ANNOTATION {
		//fmt.Printf("[gvm][checkClassAccess] 类%v 是 annotation 的\n", class.name)
	}
}
