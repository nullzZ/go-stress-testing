// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: msg/accelerate.proto

package message

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

//    使用物品 物品id，物品数量，额外参数
//
//    额外参数 ：地上物修复、升级---地上物sid
//    物品生产 ---地上物sid，物品索引
//    地组解锁--- 地组id
//    许愿树订单加速--
//    火车返回加速--
type Accelerate_C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapId  int32 `protobuf:"varint,1,opt,name=mapId,proto3" json:"mapId,omitempty"`
	Typ    int32 `protobuf:"varint,2,opt,name=typ,proto3" json:"typ,omitempty"`       //加速类型
	UseTyp int32 `protobuf:"varint,3,opt,name=useTyp,proto3" json:"useTyp,omitempty"` //使用类型
	// 各个功能的索引不同，
	Param []int32 `protobuf:"varint,4,rep,packed,name=param,proto3" json:"param,omitempty"` //如果使用物品 则是 id，物品id，物品数量
}

func (x *Accelerate_C) Reset() {
	*x = Accelerate_C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_accelerate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accelerate_C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accelerate_C) ProtoMessage() {}

func (x *Accelerate_C) ProtoReflect() protoreflect.Message {
	mi := &file_msg_accelerate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accelerate_C.ProtoReflect.Descriptor instead.
func (*Accelerate_C) Descriptor() ([]byte, []int) {
	return file_msg_accelerate_proto_rawDescGZIP(), []int{0}
}

func (x *Accelerate_C) GetMapId() int32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *Accelerate_C) GetTyp() int32 {
	if x != nil {
		return x.Typ
	}
	return 0
}

func (x *Accelerate_C) GetUseTyp() int32 {
	if x != nil {
		return x.UseTyp
	}
	return 0
}

func (x *Accelerate_C) GetParam() []int32 {
	if x != nil {
		return x.Param
	}
	return nil
}

type Accelerate_S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	MapId  int32 `protobuf:"varint,2,opt,name=mapId,proto3" json:"mapId,omitempty"`
	Typ    int32 `protobuf:"varint,3,opt,name=typ,proto3" json:"typ,omitempty"` //加速类型
	Time   int32 `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Accelerate_S) Reset() {
	*x = Accelerate_S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_accelerate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accelerate_S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accelerate_S) ProtoMessage() {}

func (x *Accelerate_S) ProtoReflect() protoreflect.Message {
	mi := &file_msg_accelerate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accelerate_S.ProtoReflect.Descriptor instead.
func (*Accelerate_S) Descriptor() ([]byte, []int) {
	return file_msg_accelerate_proto_rawDescGZIP(), []int{1}
}

func (x *Accelerate_S) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *Accelerate_S) GetMapId() int32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *Accelerate_S) GetTyp() int32 {
	if x != nil {
		return x.Typ
	}
	return 0
}

func (x *Accelerate_S) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_msg_accelerate_proto protoreflect.FileDescriptor

var file_msg_accelerate_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x73, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x64, 0x0a, 0x0c, 0x41,
	0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x43, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x61, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x79, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x74, 0x79, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x54, 0x79, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x54, 0x79, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x22, 0x62, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x53, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x70,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x79, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x79,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_accelerate_proto_rawDescOnce sync.Once
	file_msg_accelerate_proto_rawDescData = file_msg_accelerate_proto_rawDesc
)

func file_msg_accelerate_proto_rawDescGZIP() []byte {
	file_msg_accelerate_proto_rawDescOnce.Do(func() {
		file_msg_accelerate_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_accelerate_proto_rawDescData)
	})
	return file_msg_accelerate_proto_rawDescData
}

var file_msg_accelerate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msg_accelerate_proto_goTypes = []interface{}{
	(*Accelerate_C)(nil), // 0: msg.Accelerate_C
	(*Accelerate_S)(nil), // 1: msg.Accelerate_S
}
var file_msg_accelerate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msg_accelerate_proto_init() }
func file_msg_accelerate_proto_init() {
	if File_msg_accelerate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_accelerate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accelerate_C); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_accelerate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accelerate_S); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_accelerate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_accelerate_proto_goTypes,
		DependencyIndexes: file_msg_accelerate_proto_depIdxs,
		MessageInfos:      file_msg_accelerate_proto_msgTypes,
	}.Build()
	File_msg_accelerate_proto = out.File
	file_msg_accelerate_proto_rawDesc = nil
	file_msg_accelerate_proto_goTypes = nil
	file_msg_accelerate_proto_depIdxs = nil
}
