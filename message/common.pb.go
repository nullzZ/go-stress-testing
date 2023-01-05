// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: msg/common.proto

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

type Push struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId int32  `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"` //推送协议
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`    //推送数据
}

func (x *Push) Reset() {
	*x = Push{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Push) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Push) ProtoMessage() {}

func (x *Push) ProtoReflect() protoreflect.Message {
	mi := &file_msg_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Push.ProtoReflect.Descriptor instead.
func (*Push) Descriptor() ([]byte, []int) {
	return file_msg_common_proto_rawDescGZIP(), []int{0}
}

func (x *Push) GetMsgId() int32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *Push) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

//客户端请求的通用消息
type RequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId int32  `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RequestMessage) Reset() {
	*x = RequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMessage) ProtoMessage() {}

func (x *RequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_msg_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMessage.ProtoReflect.Descriptor instead.
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return file_msg_common_proto_rawDescGZIP(), []int{1}
}

func (x *RequestMessage) GetMsgId() int32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *RequestMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

//服务器响应的通用消息
type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`         //错误码
	ErrorParma []string `protobuf:"bytes,2,rep,name=errorParma,proto3" json:"errorParma,omitempty"`  //错误参数
	MsgId      int32    `protobuf:"varint,3,opt,name=msgId,proto3" json:"msgId,omitempty"`           //uri
	Data       []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`              //主消息
	Pushes     []*Push  `protobuf:"bytes,5,rep,name=pushes,proto3" json:"pushes,omitempty"`          //额外消息
	ServerTime int64    `protobuf:"varint,6,opt,name=serverTime,proto3" json:"serverTime,omitempty"` //服务器当前时间
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_msg_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_msg_common_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseMessage) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ResponseMessage) GetErrorParma() []string {
	if x != nil {
		return x.ErrorParma
	}
	return nil
}

func (x *ResponseMessage) GetMsgId() int32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *ResponseMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ResponseMessage) GetPushes() []*Push {
	if x != nil {
		return x.Pushes
	}
	return nil
}

func (x *ResponseMessage) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

type CDInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime int64 `protobuf:"varint,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	Cd        int32 `protobuf:"varint,2,opt,name=cd,proto3" json:"cd,omitempty"`
}

func (x *CDInfo) Reset() {
	*x = CDInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CDInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CDInfo) ProtoMessage() {}

func (x *CDInfo) ProtoReflect() protoreflect.Message {
	mi := &file_msg_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CDInfo.ProtoReflect.Descriptor instead.
func (*CDInfo) Descriptor() ([]byte, []int) {
	return file_msg_common_proto_rawDescGZIP(), []int{3}
}

func (x *CDInfo) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *CDInfo) GetCd() int32 {
	if x != nil {
		return x.Cd
	}
	return 0
}

var File_msg_common_proto protoreflect.FileDescriptor

var file_msg_common_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x73, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x30, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x0e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb6, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x6d, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x6d,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x06, 0x70,
	0x75, 0x73, 0x68, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x06, 0x70, 0x75, 0x73, 0x68, 0x65, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x36,
	0x0a, 0x06, 0x43, 0x44, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x63, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_common_proto_rawDescOnce sync.Once
	file_msg_common_proto_rawDescData = file_msg_common_proto_rawDesc
)

func file_msg_common_proto_rawDescGZIP() []byte {
	file_msg_common_proto_rawDescOnce.Do(func() {
		file_msg_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_common_proto_rawDescData)
	})
	return file_msg_common_proto_rawDescData
}

var file_msg_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_msg_common_proto_goTypes = []interface{}{
	(*Push)(nil),            // 0: msg.Push
	(*RequestMessage)(nil),  // 1: msg.RequestMessage
	(*ResponseMessage)(nil), // 2: msg.ResponseMessage
	(*CDInfo)(nil),          // 3: msg.CDInfo
}
var file_msg_common_proto_depIdxs = []int32{
	0, // 0: msg.ResponseMessage.pushes:type_name -> msg.Push
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_msg_common_proto_init() }
func file_msg_common_proto_init() {
	if File_msg_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Push); i {
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
		file_msg_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMessage); i {
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
		file_msg_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
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
		file_msg_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CDInfo); i {
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
			RawDescriptor: file_msg_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_common_proto_goTypes,
		DependencyIndexes: file_msg_common_proto_depIdxs,
		MessageInfos:      file_msg_common_proto_msgTypes,
	}.Build()
	File_msg_common_proto = out.File
	file_msg_common_proto_rawDesc = nil
	file_msg_common_proto_goTypes = nil
	file_msg_common_proto_depIdxs = nil
}
