package klass

type PrimitiveKlass struct {
	Descriptor       string
	ArrayClassName   string
	Name             string
	WrapperClassName string
}

var PrimitiveKlasses = []PrimitiveKlass{
	{"V", "[V", "void", "java/lang/Void"},
	{"Z", "[Z", "boolean", "java/lang/Boolean"},
	{"B", "[B", "byte", "java/lang/Byte"},
	{"C", "[C", "char", "java/lang/Character"},
	{"S", "[S", "short", "java/lang/Short"},
	{"I", "[I", "int", "java/lang/Integer"},
	{"J", "[J", "long", "java/lang/Long"},
	{"F", "[F", "float", "java/lang/Float"},
	{"D", "[D", "double", "java/lang/Double"},
}
