// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: pkg/api/proto/matching_engine.proto

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

type AddOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  string  `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Side     string  `protobuf:"bytes,2,opt,name=side,proto3" json:"side,omitempty"` // "buy" or "sell"
	Quantity float64 `protobuf:"fixed64,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price    float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *AddOrderRequest) Reset() {
	*x = AddOrderRequest{}
	mi := &file_pkg_api_proto_matching_engine_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderRequest) ProtoMessage() {}

func (x *AddOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_proto_matching_engine_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderRequest.ProtoReflect.Descriptor instead.
func (*AddOrderRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_proto_matching_engine_proto_rawDescGZIP(), []int{0}
}

func (x *AddOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *AddOrderRequest) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}

func (x *AddOrderRequest) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *AddOrderRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type AddOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddOrderResponse) Reset() {
	*x = AddOrderResponse{}
	mi := &file_pkg_api_proto_matching_engine_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderResponse) ProtoMessage() {}

func (x *AddOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_proto_matching_engine_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderResponse.ProtoReflect.Descriptor instead.
func (*AddOrderResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_proto_matching_engine_proto_rawDescGZIP(), []int{1}
}

func (x *AddOrderResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddOrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_api_proto_matching_engine_proto protoreflect.FileDescriptor

var file_pkg_api_proto_matching_engine_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x22, 0x72, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x46, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x68, 0x0a, 0x15, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x43, 0x5a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x6d, 0x61, 0x6e,
	0x4b, 0x75, 0x7a, 0x6f, 0x76, 0x6c, 0x65, 0x76, 0x2f, 0x61, 0x73, 0x6d, 0x65, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_proto_matching_engine_proto_rawDescOnce sync.Once
	file_pkg_api_proto_matching_engine_proto_rawDescData = file_pkg_api_proto_matching_engine_proto_rawDesc
)

func file_pkg_api_proto_matching_engine_proto_rawDescGZIP() []byte {
	file_pkg_api_proto_matching_engine_proto_rawDescOnce.Do(func() {
		file_pkg_api_proto_matching_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_proto_matching_engine_proto_rawDescData)
	})
	return file_pkg_api_proto_matching_engine_proto_rawDescData
}

var file_pkg_api_proto_matching_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_api_proto_matching_engine_proto_goTypes = []any{
	(*AddOrderRequest)(nil),  // 0: matching_engine.AddOrderRequest
	(*AddOrderResponse)(nil), // 1: matching_engine.AddOrderResponse
}
var file_pkg_api_proto_matching_engine_proto_depIdxs = []int32{
	0, // 0: matching_engine.MatchingEngineService.AddOrder:input_type -> matching_engine.AddOrderRequest
	1, // 1: matching_engine.MatchingEngineService.AddOrder:output_type -> matching_engine.AddOrderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_api_proto_matching_engine_proto_init() }
func file_pkg_api_proto_matching_engine_proto_init() {
	if File_pkg_api_proto_matching_engine_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_api_proto_matching_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_proto_matching_engine_proto_goTypes,
		DependencyIndexes: file_pkg_api_proto_matching_engine_proto_depIdxs,
		MessageInfos:      file_pkg_api_proto_matching_engine_proto_msgTypes,
	}.Build()
	File_pkg_api_proto_matching_engine_proto = out.File
	file_pkg_api_proto_matching_engine_proto_rawDesc = nil
	file_pkg_api_proto_matching_engine_proto_goTypes = nil
	file_pkg_api_proto_matching_engine_proto_depIdxs = nil
}
