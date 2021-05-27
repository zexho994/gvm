package klass

import "github.com/zouzhihao-994/gvm/config"

type PrimitiveKlass struct {
	Descriptor       string
	ArrayClassName   string
	Name             string
	WrapperClassName string
}

var PrimitiveKlasses = []PrimitiveKlass{
	{"V", "[V", "void", config.JPrimitiveVoid},
	{"Z", "[Z", "boolean", config.JPrimitiveBoolean},
	{"B", "[B", "byte", config.JPrimitiveByte},
	{"C", "[C", "char", config.JPrimitiveChar},
	{"S", "[S", "short", config.JPrimitiveShort},
	{"I", "[I", "int", config.JPrimitiveInteger},
	{"J", "[J", "long", config.JPrimitiveLong},
	{"F", "[F", "float", config.JPrimitiveFloat},
	{"D", "[D", "double", config.JPrimitiveDouble},
}
