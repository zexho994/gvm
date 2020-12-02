package attribute

type Attr_StackMapTable struct {
	nameIdx         uint16
	name            string
	attrLen         uint32
	numberOfEntries uint16
	stackMapFrame   []Entries
}
