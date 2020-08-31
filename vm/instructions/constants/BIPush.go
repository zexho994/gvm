package constants

//
//import (
//	"../../instructions/base"
//	"../../rtda"
//)
//
//// Push byte
//type BIPush struct {
//	Val int32 // TODO
//}
//
//func (instr *BIPush) FetchOperands(reader *base.CodeReader) {
//	instr.Val = int32(reader.ReadInt8())
//}
//func (instr *BIPush) Execute(frame *rtda.Frame) {
//	frame.PushInt(instr.Val)
//}
//
//// Push short
//type SIPush struct {
//	Val int32
//}
//
//func (instr *SIPush) FetchOperands(reader *base.CodeReader) {
//	instr.Val = int32(reader.ReadInt16())
//}
//func (instr *SIPush) Execute(frame *rtda.Frame) {
//	frame.PushInt(instr.Val)
//}
//
