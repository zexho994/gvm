package klass

type Klass struct {
	// The "layout helper" is a combined descriptor of object layout.
	// For klasses which are neither instance nor array, the value is zero.
	//
	// For instances, layout helper is a positive number, the instance size.
	// This size is already passed through align_object_size and scaled to bytes.
	// The low order bit is set if instances of this class cannot be
	// allocated using the fastpath.
	//
	// For arrays, layout helper is a negative number, containing four
	// distinct bytes, as follows:
	//    MSB:[tag, hsz, ebt, log2(esz)]:LSB
	// where:
	//    tag is 0x80 if the elements are oops, 0xC0 if non-oops
	//    hsz is array header size in bytes (i.e., offset of first element)
	//    ebt is the BasicType of the elements
	//    esz is the element size in bytes
	// This packed word is arranged so as to be quickly unpacked by the
	// various fast paths that use the various subfields.
	//
	// The esz bits can be used directly by a SLL instruction, without masking.
	//
	// Note that the array-kind tag looks like 0x00 for instance klasses,
	// since their length in bytes is always less than 24Mb.
	//
	// Final note:  This comes first, immediately after C++ vtable,
	// because it is frequently queried.
	layoutHelper uint32
}
