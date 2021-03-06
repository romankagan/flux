// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UnsignedIntegerLiteral struct {
	_tab flatbuffers.Table
}

func GetRootAsUnsignedIntegerLiteral(buf []byte, offset flatbuffers.UOffsetT) *UnsignedIntegerLiteral {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UnsignedIntegerLiteral{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *UnsignedIntegerLiteral) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UnsignedIntegerLiteral) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UnsignedIntegerLiteral) Loc(obj *SourceLocation) *SourceLocation {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SourceLocation)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *UnsignedIntegerLiteral) Value() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *UnsignedIntegerLiteral) MutateValue(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *UnsignedIntegerLiteral) TypType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *UnsignedIntegerLiteral) MutateTypType(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *UnsignedIntegerLiteral) Typ(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func UnsignedIntegerLiteralStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func UnsignedIntegerLiteralAddLoc(builder *flatbuffers.Builder, loc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(loc), 0)
}
func UnsignedIntegerLiteralAddValue(builder *flatbuffers.Builder, value uint64) {
	builder.PrependUint64Slot(1, value, 0)
}
func UnsignedIntegerLiteralAddTypType(builder *flatbuffers.Builder, typType byte) {
	builder.PrependByteSlot(2, typType, 0)
}
func UnsignedIntegerLiteralAddTyp(builder *flatbuffers.Builder, typ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(typ), 0)
}
func UnsignedIntegerLiteralEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
