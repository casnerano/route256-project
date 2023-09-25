// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: api/v1/cart_service.proto

package cart

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

type ListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint64 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price uint32 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ListItem) Reset() {
	*x = ListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItem) ProtoMessage() {}

func (x *ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItem.ProtoReflect.Descriptor instead.
func (*ListItem) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListItem) GetSku() uint64 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *ListItem) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListItem) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User uint64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*ListItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	TotalPrice uint64      `protobuf:"varint,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListResponse) GetItems() []*ListItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListResponse) GetTotalPrice() uint64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type ClearRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User uint64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ClearRequest) Reset() {
	*x = ClearRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearRequest) ProtoMessage() {}

func (x *ClearRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearRequest.ProtoReflect.Descriptor instead.
func (*ClearRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{3}
}

func (x *ClearRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

type ClearResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearResponse) Reset() {
	*x = ClearResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearResponse) ProtoMessage() {}

func (x *ClearResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearResponse.ProtoReflect.Descriptor instead.
func (*ClearResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{4}
}

type ItemAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  uint64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Sku   uint64 `protobuf:"varint,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ItemAddRequest) Reset() {
	*x = ItemAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemAddRequest) ProtoMessage() {}

func (x *ItemAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemAddRequest.ProtoReflect.Descriptor instead.
func (*ItemAddRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{5}
}

func (x *ItemAddRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *ItemAddRequest) GetSku() uint64 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *ItemAddRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ItemAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ItemAddResponse) Reset() {
	*x = ItemAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemAddResponse) ProtoMessage() {}

func (x *ItemAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemAddResponse.ProtoReflect.Descriptor instead.
func (*ItemAddResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{6}
}

type ItemDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User uint64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Sku  uint64 `protobuf:"varint,2,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *ItemDeleteRequest) Reset() {
	*x = ItemDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemDeleteRequest) ProtoMessage() {}

func (x *ItemDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemDeleteRequest.ProtoReflect.Descriptor instead.
func (*ItemDeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{7}
}

func (x *ItemDeleteRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *ItemDeleteRequest) GetSku() uint64 {
	if x != nil {
		return x.Sku
	}
	return 0
}

type ItemDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ItemDeleteResponse) Reset() {
	*x = ItemDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemDeleteResponse) ProtoMessage() {}

func (x *ItemDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemDeleteResponse.ProtoReflect.Descriptor instead.
func (*ItemDeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{8}
}

type CheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User uint64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CheckoutRequest) Reset() {
	*x = CheckoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequest) ProtoMessage() {}

func (x *CheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequest.ProtoReflect.Descriptor instead.
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{9}
}

func (x *CheckoutRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

type CheckoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *CheckoutResponse) Reset() {
	*x = CheckoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_cart_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutResponse) ProtoMessage() {}

func (x *CheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cart_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutResponse.ProtoReflect.Descriptor instead.
func (*CheckoutResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cart_service_proto_rawDescGZIP(), []int{10}
}

func (x *CheckoutResponse) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_api_v1_cart_service_proto protoreflect.FileDescriptor

var file_api_v1_cart_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x22, 0x5c, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x21, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x65, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x0f, 0x0a,
	0x0d, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c,
	0x0a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x11, 0x0a, 0x0f,
	0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x39, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x14, 0x0a, 0x12, 0x49, 0x74,
	0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x25, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x32, 0xbb, 0x03, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x4d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32,
	0x35, 0x36, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50,
	0x0a, 0x05, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x22, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32,
	0x35, 0x36, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x56, 0x0a, 0x07, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x12, 0x24, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35,
	0x36, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x08, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x25, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_cart_service_proto_rawDescOnce sync.Once
	file_api_v1_cart_service_proto_rawDescData = file_api_v1_cart_service_proto_rawDesc
)

func file_api_v1_cart_service_proto_rawDescGZIP() []byte {
	file_api_v1_cart_service_proto_rawDescOnce.Do(func() {
		file_api_v1_cart_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_cart_service_proto_rawDescData)
	})
	return file_api_v1_cart_service_proto_rawDescData
}

var file_api_v1_cart_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v1_cart_service_proto_goTypes = []interface{}{
	(*ListItem)(nil),           // 0: route256.cart.api.v1.ListItem
	(*ListRequest)(nil),        // 1: route256.cart.api.v1.ListRequest
	(*ListResponse)(nil),       // 2: route256.cart.api.v1.ListResponse
	(*ClearRequest)(nil),       // 3: route256.cart.api.v1.ClearRequest
	(*ClearResponse)(nil),      // 4: route256.cart.api.v1.ClearResponse
	(*ItemAddRequest)(nil),     // 5: route256.cart.api.v1.ItemAddRequest
	(*ItemAddResponse)(nil),    // 6: route256.cart.api.v1.ItemAddResponse
	(*ItemDeleteRequest)(nil),  // 7: route256.cart.api.v1.ItemDeleteRequest
	(*ItemDeleteResponse)(nil), // 8: route256.cart.api.v1.ItemDeleteResponse
	(*CheckoutRequest)(nil),    // 9: route256.cart.api.v1.CheckoutRequest
	(*CheckoutResponse)(nil),   // 10: route256.cart.api.v1.CheckoutResponse
}
var file_api_v1_cart_service_proto_depIdxs = []int32{
	0,  // 0: route256.cart.api.v1.ListResponse.items:type_name -> route256.cart.api.v1.ListItem
	1,  // 1: route256.cart.api.v1.Cart.List:input_type -> route256.cart.api.v1.ListRequest
	3,  // 2: route256.cart.api.v1.Cart.Clear:input_type -> route256.cart.api.v1.ClearRequest
	5,  // 3: route256.cart.api.v1.Cart.ItemAdd:input_type -> route256.cart.api.v1.ItemAddRequest
	7,  // 4: route256.cart.api.v1.Cart.ItemDelete:input_type -> route256.cart.api.v1.ItemDeleteRequest
	9,  // 5: route256.cart.api.v1.Cart.Checkout:input_type -> route256.cart.api.v1.CheckoutRequest
	2,  // 6: route256.cart.api.v1.Cart.List:output_type -> route256.cart.api.v1.ListResponse
	4,  // 7: route256.cart.api.v1.Cart.Clear:output_type -> route256.cart.api.v1.ClearResponse
	6,  // 8: route256.cart.api.v1.Cart.ItemAdd:output_type -> route256.cart.api.v1.ItemAddResponse
	8,  // 9: route256.cart.api.v1.Cart.ItemDelete:output_type -> route256.cart.api.v1.ItemDeleteResponse
	10, // 10: route256.cart.api.v1.Cart.Checkout:output_type -> route256.cart.api.v1.CheckoutResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_cart_service_proto_init() }
func file_api_v1_cart_service_proto_init() {
	if File_api_v1_cart_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_cart_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItem); i {
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
		file_api_v1_cart_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_api_v1_cart_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_api_v1_cart_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearRequest); i {
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
		file_api_v1_cart_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearResponse); i {
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
		file_api_v1_cart_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemAddRequest); i {
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
		file_api_v1_cart_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemAddResponse); i {
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
		file_api_v1_cart_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemDeleteRequest); i {
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
		file_api_v1_cart_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemDeleteResponse); i {
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
		file_api_v1_cart_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutRequest); i {
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
		file_api_v1_cart_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutResponse); i {
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
			RawDescriptor: file_api_v1_cart_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_cart_service_proto_goTypes,
		DependencyIndexes: file_api_v1_cart_service_proto_depIdxs,
		MessageInfos:      file_api_v1_cart_service_proto_msgTypes,
	}.Build()
	File_api_v1_cart_service_proto = out.File
	file_api_v1_cart_service_proto_rawDesc = nil
	file_api_v1_cart_service_proto_goTypes = nil
	file_api_v1_cart_service_proto_depIdxs = nil
}
