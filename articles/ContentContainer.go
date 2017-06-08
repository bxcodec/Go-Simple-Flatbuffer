// automatically generated by the FlatBuffers compiler, do not modify

package articles

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ContentContainer struct {
	_tab flatbuffers.Table
}

func GetRootAsContentContainer(buf []byte, offset flatbuffers.UOffsetT) *ContentContainer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ContentContainer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ContentContainer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ContentContainer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ContentContainer) ContentsType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ContentContainer) MutateContentsType(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *ContentContainer) Contents(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func ContentContainerStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ContentContainerAddContentsType(builder *flatbuffers.Builder, contentsType byte) {
	builder.PrependByteSlot(0, contentsType, 0)
}
func ContentContainerAddContents(builder *flatbuffers.Builder, contents flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(contents), 0)
}
func ContentContainerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
