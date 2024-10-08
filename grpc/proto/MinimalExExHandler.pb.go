// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: MinimalExExHandler.proto

package proto

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

// Request messages
type ChainCommittedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartBlock uint64 `protobuf:"varint,1,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	EndBlock   uint64 `protobuf:"varint,2,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
}

func (x *ChainCommittedRequest) Reset() {
	*x = ChainCommittedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MinimalExExHandler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainCommittedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainCommittedRequest) ProtoMessage() {}

func (x *ChainCommittedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MinimalExExHandler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainCommittedRequest.ProtoReflect.Descriptor instead.
func (*ChainCommittedRequest) Descriptor() ([]byte, []int) {
	return file_MinimalExExHandler_proto_rawDescGZIP(), []int{0}
}

func (x *ChainCommittedRequest) GetStartBlock() uint64 {
	if x != nil {
		return x.StartBlock
	}
	return 0
}

func (x *ChainCommittedRequest) GetEndBlock() uint64 {
	if x != nil {
		return x.EndBlock
	}
	return 0
}

type ChainReorgedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldStartBlock uint64 `protobuf:"varint,1,opt,name=old_start_block,json=oldStartBlock,proto3" json:"old_start_block,omitempty"`
	OldEndBlock   uint64 `protobuf:"varint,2,opt,name=old_end_block,json=oldEndBlock,proto3" json:"old_end_block,omitempty"`
	NewStartBlock uint64 `protobuf:"varint,3,opt,name=new_start_block,json=newStartBlock,proto3" json:"new_start_block,omitempty"`
	NewEndBlock   uint64 `protobuf:"varint,4,opt,name=new_end_block,json=newEndBlock,proto3" json:"new_end_block,omitempty"`
}

func (x *ChainReorgedRequest) Reset() {
	*x = ChainReorgedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MinimalExExHandler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainReorgedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainReorgedRequest) ProtoMessage() {}

func (x *ChainReorgedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MinimalExExHandler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainReorgedRequest.ProtoReflect.Descriptor instead.
func (*ChainReorgedRequest) Descriptor() ([]byte, []int) {
	return file_MinimalExExHandler_proto_rawDescGZIP(), []int{1}
}

func (x *ChainReorgedRequest) GetOldStartBlock() uint64 {
	if x != nil {
		return x.OldStartBlock
	}
	return 0
}

func (x *ChainReorgedRequest) GetOldEndBlock() uint64 {
	if x != nil {
		return x.OldEndBlock
	}
	return 0
}

func (x *ChainReorgedRequest) GetNewStartBlock() uint64 {
	if x != nil {
		return x.NewStartBlock
	}
	return 0
}

func (x *ChainReorgedRequest) GetNewEndBlock() uint64 {
	if x != nil {
		return x.NewEndBlock
	}
	return 0
}

type ChainRevertedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartBlock uint64 `protobuf:"varint,1,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	EndBlock   uint64 `protobuf:"varint,2,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
}

func (x *ChainRevertedRequest) Reset() {
	*x = ChainRevertedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MinimalExExHandler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainRevertedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainRevertedRequest) ProtoMessage() {}

func (x *ChainRevertedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MinimalExExHandler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainRevertedRequest.ProtoReflect.Descriptor instead.
func (*ChainRevertedRequest) Descriptor() ([]byte, []int) {
	return file_MinimalExExHandler_proto_rawDescGZIP(), []int{2}
}

func (x *ChainRevertedRequest) GetStartBlock() uint64 {
	if x != nil {
		return x.StartBlock
	}
	return 0
}

func (x *ChainRevertedRequest) GetEndBlock() uint64 {
	if x != nil {
		return x.EndBlock
	}
	return 0
}

// Response message
type HandlerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendFinishedHeightEvent bool   `protobuf:"varint,1,opt,name=send_finished_height_event,json=sendFinishedHeightEvent,proto3" json:"send_finished_height_event,omitempty"`
	FinishedHeight          uint64 `protobuf:"varint,2,opt,name=finished_height,json=finishedHeight,proto3" json:"finished_height,omitempty"`
}

func (x *HandlerResponse) Reset() {
	*x = HandlerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MinimalExExHandler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandlerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandlerResponse) ProtoMessage() {}

func (x *HandlerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MinimalExExHandler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandlerResponse.ProtoReflect.Descriptor instead.
func (*HandlerResponse) Descriptor() ([]byte, []int) {
	return file_MinimalExExHandler_proto_rawDescGZIP(), []int{3}
}

func (x *HandlerResponse) GetSendFinishedHeightEvent() bool {
	if x != nil {
		return x.SendFinishedHeightEvent
	}
	return false
}

func (x *HandlerResponse) GetFinishedHeight() uint64 {
	if x != nil {
		return x.FinishedHeight
	}
	return 0
}

var File_MinimalExExHandler_proto protoreflect.FileDescriptor

var file_MinimalExExHandler_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x45, 0x78, 0x45, 0x78, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x78, 0x65, 0x78,
	0x22, 0x55, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e,
	0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0xad, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x6f, 0x72, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6f, 0x6c, 0x64, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x6c, 0x64, 0x5f, 0x65,
	0x6e, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x6f, 0x6c, 0x64, 0x45, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x45,
	0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x54, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x77, 0x0a,
	0x0f, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x1a, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x73, 0x65, 0x6e, 0x64, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0xeb, 0x01, 0x0a, 0x0b, 0x45, 0x78, 0x45, 0x78, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x14, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x12, 0x1b,
	0x2e, 0x65, 0x78, 0x65, 0x78, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x78,
	0x65, 0x78, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x6f, 0x72, 0x67, 0x65, 0x64, 0x12, 0x19, 0x2e, 0x65, 0x78, 0x65, 0x78, 0x2e,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x6f, 0x72, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x78, 0x65, 0x78, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x13, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x64, 0x12, 0x1a, 0x2e, 0x65, 0x78, 0x65, 0x78, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x65, 0x78, 0x65, 0x78, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x70, 0x32, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x72, 0x65, 0x74, 0x68,
	0x2d, 0x65, 0x78, 0x65, 0x78, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MinimalExExHandler_proto_rawDescOnce sync.Once
	file_MinimalExExHandler_proto_rawDescData = file_MinimalExExHandler_proto_rawDesc
)

func file_MinimalExExHandler_proto_rawDescGZIP() []byte {
	file_MinimalExExHandler_proto_rawDescOnce.Do(func() {
		file_MinimalExExHandler_proto_rawDescData = protoimpl.X.CompressGZIP(file_MinimalExExHandler_proto_rawDescData)
	})
	return file_MinimalExExHandler_proto_rawDescData
}

var file_MinimalExExHandler_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_MinimalExExHandler_proto_goTypes = []interface{}{
	(*ChainCommittedRequest)(nil), // 0: exex.ChainCommittedRequest
	(*ChainReorgedRequest)(nil),   // 1: exex.ChainReorgedRequest
	(*ChainRevertedRequest)(nil),  // 2: exex.ChainRevertedRequest
	(*HandlerResponse)(nil),       // 3: exex.HandlerResponse
}
var file_MinimalExExHandler_proto_depIdxs = []int32{
	0, // 0: exex.ExExHandler.HandleChainCommitted:input_type -> exex.ChainCommittedRequest
	1, // 1: exex.ExExHandler.HandleChainReorged:input_type -> exex.ChainReorgedRequest
	2, // 2: exex.ExExHandler.HandleChainReverted:input_type -> exex.ChainRevertedRequest
	3, // 3: exex.ExExHandler.HandleChainCommitted:output_type -> exex.HandlerResponse
	3, // 4: exex.ExExHandler.HandleChainReorged:output_type -> exex.HandlerResponse
	3, // 5: exex.ExExHandler.HandleChainReverted:output_type -> exex.HandlerResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MinimalExExHandler_proto_init() }
func file_MinimalExExHandler_proto_init() {
	if File_MinimalExExHandler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MinimalExExHandler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainCommittedRequest); i {
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
		file_MinimalExExHandler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainReorgedRequest); i {
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
		file_MinimalExExHandler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainRevertedRequest); i {
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
		file_MinimalExExHandler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandlerResponse); i {
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
			RawDescriptor: file_MinimalExExHandler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MinimalExExHandler_proto_goTypes,
		DependencyIndexes: file_MinimalExExHandler_proto_depIdxs,
		MessageInfos:      file_MinimalExExHandler_proto_msgTypes,
	}.Build()
	File_MinimalExExHandler_proto = out.File
	file_MinimalExExHandler_proto_rawDesc = nil
	file_MinimalExExHandler_proto_goTypes = nil
	file_MinimalExExHandler_proto_depIdxs = nil
}
