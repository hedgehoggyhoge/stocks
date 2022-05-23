// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: grpc_utils/proto/PortfolioStorage/portfolio_storage.proto

package portfolioStorage

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

type PortfolioItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity int64  `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PortfolioItem) Reset() {
	*x = PortfolioItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortfolioItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortfolioItem) ProtoMessage() {}

func (x *PortfolioItem) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*PortfolioItem) Descriptor() ([]byte, []int) {
	return file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescGZIP(), []int{0}
}

func (x *PortfolioItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *PortfolioItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdatePortfolioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         int64            `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	PortfolioItems []*PortfolioItem `protobuf:"bytes,2,rep,name=portfolioItems,proto3" json:"portfolioItems,omitempty"`
}

func (x *UpdatePortfolioRequest) Reset() {
	*x = UpdatePortfolioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePortfolioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePortfolioRequest) ProtoMessage() {}

func (x *UpdatePortfolioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*UpdatePortfolioRequest) Descriptor() ([]byte, []int) {
	return file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePortfolioRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdatePortfolioRequest) GetPortfolioItems() []*PortfolioItem {
	if x != nil {
		return x.PortfolioItems
	}
	return nil
}

type UpdatePortfolioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePortfolioResponse) Reset() {
	*x = UpdatePortfolioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePortfolioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePortfolioResponse) ProtoMessage() {}

func (x *UpdatePortfolioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*UpdatePortfolioResponse) Descriptor() ([]byte, []int) {
	return file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescGZIP(), []int{2}
}

type GetPortfolioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetPortfolioRequest) Reset() {
	*x = GetPortfolioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPortfolioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPortfolioRequest) ProtoMessage() {}

func (x *GetPortfolioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPortfolioRequest.ProtoReflect.Descriptor instead.
func (*GetPortfolioRequest) Descriptor() ([]byte, []int) {
	return file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescGZIP(), []int{3}
}

func (x *GetPortfolioRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetPortfolioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortfolioItems []*PortfolioItem `protobuf:"bytes,1,rep,name=portfolioItems,proto3" json:"portfolioItems,omitempty"`
	Price          float32          `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GetPortfolioResponse) Reset() {
	*x = GetPortfolioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPortfolioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPortfolioResponse) ProtoMessage() {}

func (x *GetPortfolioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPortfolioResponse.ProtoReflect.Descriptor instead.
func (*GetPortfolioResponse) Descriptor() ([]byte, []int) {
	return file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescGZIP(), []int{4}
}

func (x *GetPortfolioResponse) GetPortfolioItems() []*PortfolioItem {
	if x != nil {
		return x.PortfolioItems
	}
	return nil
}

func (x *GetPortfolioResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto protoreflect.FileDescriptor

var file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x50, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a,
	0x0d, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x79,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x47, 0x0a, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x66,
	0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x66,
	0x6f, 0x6c, 0x69, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66,
	0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f,
	0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x70,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0xdd, 0x01, 0x0a, 0x10, 0x50,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x68, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c,
	0x69, 0x6f, 0x12, 0x28, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x25, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescOnce sync.Once
	file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescData = file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDesc
)

func file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescGZIP() []byte {
	file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescOnce.Do(func() {
		file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescData)
	})
	return file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDescData
}

var file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_goTypes = []interface{}{
	(*PortfolioItem)(nil),           // 0: PortfolioStorage.PortfolioItem
	(*UpdatePortfolioRequest)(nil),  // 1: PortfolioStorage.UpdatePortfolioRequest
	(*UpdatePortfolioResponse)(nil), // 2: PortfolioStorage.UpdatePortfolioResponse
	(*GetPortfolioRequest)(nil),     // 3: PortfolioStorage.GetPortfolioRequest
	(*GetPortfolioResponse)(nil),    // 4: PortfolioStorage.GetPortfolioResponse
}
var file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_depIdxs = []int32{
	0, // 0: PortfolioStorage.UpdatePortfolioRequest.portfolioItems:type_name -> PortfolioStorage.PortfolioItem
	0, // 1: PortfolioStorage.GetPortfolioResponse.portfolioItems:type_name -> PortfolioStorage.PortfolioItem
	1, // 2: PortfolioStorage.PortfolioStorage.UpdatePortfolio:input_type -> PortfolioStorage.UpdatePortfolioRequest
	3, // 3: PortfolioStorage.PortfolioStorage.GetPortfolio:input_type -> PortfolioStorage.GetPortfolioRequest
	2, // 4: PortfolioStorage.PortfolioStorage.UpdatePortfolio:output_type -> PortfolioStorage.UpdatePortfolioResponse
	4, // 5: PortfolioStorage.PortfolioStorage.GetPortfolio:output_type -> PortfolioStorage.GetPortfolioResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_init() }
func file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_init() {
	if File_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortfolioItem); i {
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
		file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePortfolioRequest); i {
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
		file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePortfolioResponse); i {
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
		file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPortfolioRequest); i {
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
		file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPortfolioResponse); i {
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
			RawDescriptor: file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_goTypes,
		DependencyIndexes: file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_depIdxs,
		MessageInfos:      file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_msgTypes,
	}.Build()
	File_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto = out.File
	file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_rawDesc = nil
	file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_goTypes = nil
	file_grpc_utils_proto_PortfolioStorage_portfolio_storage_proto_depIdxs = nil
}
