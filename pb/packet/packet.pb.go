// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.24.4
// source: packet.proto

package packet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PacketBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packets []*Packet `protobuf:"bytes,1,rep,name=packets,proto3" json:"packets,omitempty"`
}

func (x *PacketBatch) Reset() {
	*x = PacketBatch{}
	mi := &file_packet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PacketBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketBatch) ProtoMessage() {}

func (x *PacketBatch) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketBatch.ProtoReflect.Descriptor instead.
func (*PacketBatch) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{0}
}

func (x *PacketBatch) GetPackets() []*Packet {
	if x != nil {
		return x.Packets
	}
	return nil
}

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Meta *Meta  `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	mi := &file_packet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{1}
}

func (x *Packet) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Packet) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size        uint64       `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Addr        string       `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Port        uint32       `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Flags       *PacketFlags `protobuf:"bytes,4,opt,name=flags,proto3" json:"flags,omitempty"`
	SenderStake uint64       `protobuf:"varint,5,opt,name=sender_stake,json=senderStake,proto3" json:"sender_stake,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	mi := &file_packet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{2}
}

func (x *Meta) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Meta) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Meta) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Meta) GetFlags() *PacketFlags {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *Meta) GetSenderStake() uint64 {
	if x != nil {
		return x.SenderStake
	}
	return 0
}

type PacketFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Discard        bool `protobuf:"varint,1,opt,name=discard,proto3" json:"discard,omitempty"`
	Forwarded      bool `protobuf:"varint,2,opt,name=forwarded,proto3" json:"forwarded,omitempty"`
	Repair         bool `protobuf:"varint,3,opt,name=repair,proto3" json:"repair,omitempty"`
	SimpleVoteTx   bool `protobuf:"varint,4,opt,name=simple_vote_tx,json=simpleVoteTx,proto3" json:"simple_vote_tx,omitempty"`
	TracerPacket   bool `protobuf:"varint,5,opt,name=tracer_packet,json=tracerPacket,proto3" json:"tracer_packet,omitempty"`
	FromStakedNode bool `protobuf:"varint,6,opt,name=from_staked_node,json=fromStakedNode,proto3" json:"from_staked_node,omitempty"`
}

func (x *PacketFlags) Reset() {
	*x = PacketFlags{}
	mi := &file_packet_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PacketFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketFlags) ProtoMessage() {}

func (x *PacketFlags) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketFlags.ProtoReflect.Descriptor instead.
func (*PacketFlags) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{3}
}

func (x *PacketFlags) GetDiscard() bool {
	if x != nil {
		return x.Discard
	}
	return false
}

func (x *PacketFlags) GetForwarded() bool {
	if x != nil {
		return x.Forwarded
	}
	return false
}

func (x *PacketFlags) GetRepair() bool {
	if x != nil {
		return x.Repair
	}
	return false
}

func (x *PacketFlags) GetSimpleVoteTx() bool {
	if x != nil {
		return x.SimpleVoteTx
	}
	return false
}

func (x *PacketFlags) GetTracerPacket() bool {
	if x != nil {
		return x.TracerPacket
	}
	return false
}

func (x *PacketFlags) GetFromStakedNode() bool {
	if x != nil {
		return x.FromStakedNode
	}
	return false
}

var File_packet_proto protoreflect.FileDescriptor

var file_packet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x37, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x22,
	0x3e, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22,
	0x90, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x6b, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x70, 0x61,
	0x69, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x76, 0x6f, 0x74,
	0x65, 0x5f, 0x74, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x54, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x74, 0x61,
	0x6b, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x41, 0x49, 0x4f, 0x2d,
	0x44, 0x65, 0x76, 0x43, 0x6f, 0x72, 0x65, 0x2f, 0x6a, 0x69, 0x74, 0x6f, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_packet_proto_rawDescOnce sync.Once
	file_packet_proto_rawDescData = file_packet_proto_rawDesc
)

func file_packet_proto_rawDescGZIP() []byte {
	file_packet_proto_rawDescOnce.Do(func() {
		file_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_packet_proto_rawDescData)
	})
	return file_packet_proto_rawDescData
}

var file_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_packet_proto_goTypes = []any{
	(*PacketBatch)(nil), // 0: packet.PacketBatch
	(*Packet)(nil),      // 1: packet.Packet
	(*Meta)(nil),        // 2: packet.Meta
	(*PacketFlags)(nil), // 3: packet.PacketFlags
}
var file_packet_proto_depIdxs = []int32{
	1, // 0: packet.PacketBatch.packets:type_name -> packet.Packet
	2, // 1: packet.Packet.meta:type_name -> packet.Meta
	3, // 2: packet.Meta.flags:type_name -> packet.PacketFlags
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_packet_proto_init() }
func file_packet_proto_init() {
	if File_packet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_packet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packet_proto_goTypes,
		DependencyIndexes: file_packet_proto_depIdxs,
		MessageInfos:      file_packet_proto_msgTypes,
	}.Build()
	File_packet_proto = out.File
	file_packet_proto_rawDesc = nil
	file_packet_proto_goTypes = nil
	file_packet_proto_depIdxs = nil
}
