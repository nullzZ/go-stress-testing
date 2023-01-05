// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: msg/reward.proto

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

type Reward_S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//通用奖励弹框
	Result      int32         `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	RewardItems []*RewardItem `protobuf:"bytes,2,rep,name=rewardItems,proto3" json:"rewardItems,omitempty"` //参数
	Source      int32         `protobuf:"varint,3,opt,name=source,proto3" json:"source,omitempty"`          //奖励来源
}

func (x *Reward_S) Reset() {
	*x = Reward_S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_reward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reward_S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reward_S) ProtoMessage() {}

func (x *Reward_S) ProtoReflect() protoreflect.Message {
	mi := &file_msg_reward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reward_S.ProtoReflect.Descriptor instead.
func (*Reward_S) Descriptor() ([]byte, []int) {
	return file_msg_reward_proto_rawDescGZIP(), []int{0}
}

func (x *Reward_S) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *Reward_S) GetRewardItems() []*RewardItem {
	if x != nil {
		return x.RewardItems
	}
	return nil
}

func (x *Reward_S) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

type RewardItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid int32 `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
	Num int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *RewardItem) Reset() {
	*x = RewardItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_reward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardItem) ProtoMessage() {}

func (x *RewardItem) ProtoReflect() protoreflect.Message {
	mi := &file_msg_reward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardItem.ProtoReflect.Descriptor instead.
func (*RewardItem) Descriptor() ([]byte, []int) {
	return file_msg_reward_proto_rawDescGZIP(), []int{1}
}

func (x *RewardItem) GetTid() int32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *RewardItem) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_msg_reward_proto protoreflect.FileDescriptor

var file_msg_reward_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x73, 0x67, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x6d, 0x0a, 0x08, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x30, 0x0a, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_reward_proto_rawDescOnce sync.Once
	file_msg_reward_proto_rawDescData = file_msg_reward_proto_rawDesc
)

func file_msg_reward_proto_rawDescGZIP() []byte {
	file_msg_reward_proto_rawDescOnce.Do(func() {
		file_msg_reward_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_reward_proto_rawDescData)
	})
	return file_msg_reward_proto_rawDescData
}

var file_msg_reward_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msg_reward_proto_goTypes = []interface{}{
	(*Reward_S)(nil),   // 0: msg.Reward_S
	(*RewardItem)(nil), // 1: msg.RewardItem
}
var file_msg_reward_proto_depIdxs = []int32{
	1, // 0: msg.Reward_S.rewardItems:type_name -> msg.RewardItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_msg_reward_proto_init() }
func file_msg_reward_proto_init() {
	if File_msg_reward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_reward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reward_S); i {
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
		file_msg_reward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardItem); i {
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
			RawDescriptor: file_msg_reward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_reward_proto_goTypes,
		DependencyIndexes: file_msg_reward_proto_depIdxs,
		MessageInfos:      file_msg_reward_proto_msgTypes,
	}.Build()
	File_msg_reward_proto = out.File
	file_msg_reward_proto_rawDesc = nil
	file_msg_reward_proto_goTypes = nil
	file_msg_reward_proto_depIdxs = nil
}
