// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/booking/booking.proto

package booking

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	sdk "user-gateway/proto/sdk"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgGetBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId int64 `protobuf:"varint,1,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
}

func (x *MsgGetBookingRequest) Reset() {
	*x = MsgGetBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgGetBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgGetBookingRequest) ProtoMessage() {}

func (x *MsgGetBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgGetBookingRequest.ProtoReflect.Descriptor instead.
func (*MsgGetBookingRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_booking_proto_rawDescGZIP(), []int{0}
}

func (x *MsgGetBookingRequest) GetBookingId() int64 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

type MsgCreatePropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WardId       string `protobuf:"bytes,1,opt,name=wardId,proto3" json:"wardId,omitempty"`
	HostId       string `protobuf:"bytes,2,opt,name=hostId,proto3" json:"hostId,omitempty"`
	NumGuests    int64  `protobuf:"varint,6,opt,name=numGuests,proto3" json:"numGuests,omitempty"`
	NumBed       int64  `protobuf:"varint,3,opt,name=numBed,proto3" json:"numBed,omitempty"`
	NumBath      int64  `protobuf:"varint,4,opt,name=numBath,proto3" json:"numBath,omitempty"`
	NumBedroom   int64  `protobuf:"varint,5,opt,name=numBedroom,proto3" json:"numBedroom,omitempty"`
	IsGuestFavor bool   `protobuf:"varint,7,opt,name=isGuestFavor,proto3" json:"isGuestFavor,omitempty"`
	Body         string `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	Title        string `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	PropertyType int64  `protobuf:"varint,10,opt,name=propertyType,proto3" json:"propertyType,omitempty"`
}

func (x *MsgCreatePropertyRequest) Reset() {
	*x = MsgCreatePropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreatePropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreatePropertyRequest) ProtoMessage() {}

func (x *MsgCreatePropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreatePropertyRequest.ProtoReflect.Descriptor instead.
func (*MsgCreatePropertyRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_booking_proto_rawDescGZIP(), []int{1}
}

func (x *MsgCreatePropertyRequest) GetWardId() string {
	if x != nil {
		return x.WardId
	}
	return ""
}

func (x *MsgCreatePropertyRequest) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *MsgCreatePropertyRequest) GetNumGuests() int64 {
	if x != nil {
		return x.NumGuests
	}
	return 0
}

func (x *MsgCreatePropertyRequest) GetNumBed() int64 {
	if x != nil {
		return x.NumBed
	}
	return 0
}

func (x *MsgCreatePropertyRequest) GetNumBath() int64 {
	if x != nil {
		return x.NumBath
	}
	return 0
}

func (x *MsgCreatePropertyRequest) GetNumBedroom() int64 {
	if x != nil {
		return x.NumBedroom
	}
	return 0
}

func (x *MsgCreatePropertyRequest) GetIsGuestFavor() bool {
	if x != nil {
		return x.IsGuestFavor
	}
	return false
}

func (x *MsgCreatePropertyRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *MsgCreatePropertyRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MsgCreatePropertyRequest) GetPropertyType() int64 {
	if x != nil {
		return x.PropertyType
	}
	return 0
}

type MsgUpdatePropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId   string `protobuf:"bytes,10,opt,name=propertyId,proto3" json:"propertyId,omitempty"`
	NumGuests    int64  `protobuf:"varint,6,opt,name=numGuests,proto3" json:"numGuests,omitempty"`
	NumBed       int64  `protobuf:"varint,3,opt,name=numBed,proto3" json:"numBed,omitempty"`
	NumBath      int64  `protobuf:"varint,4,opt,name=numBath,proto3" json:"numBath,omitempty"`
	NumBedroom   int64  `protobuf:"varint,5,opt,name=numBedroom,proto3" json:"numBedroom,omitempty"`
	IsGuestFavor bool   `protobuf:"varint,7,opt,name=isGuestFavor,proto3" json:"isGuestFavor,omitempty"`
	Body         string `protobuf:"bytes,8,opt,name=body,proto3" json:"body,omitempty"`
	Title        string `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
	PropertyType int64  `protobuf:"varint,11,opt,name=propertyType,proto3" json:"propertyType,omitempty"`
}

func (x *MsgUpdatePropertyRequest) Reset() {
	*x = MsgUpdatePropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUpdatePropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUpdatePropertyRequest) ProtoMessage() {}

func (x *MsgUpdatePropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUpdatePropertyRequest.ProtoReflect.Descriptor instead.
func (*MsgUpdatePropertyRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_booking_proto_rawDescGZIP(), []int{2}
}

func (x *MsgUpdatePropertyRequest) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

func (x *MsgUpdatePropertyRequest) GetNumGuests() int64 {
	if x != nil {
		return x.NumGuests
	}
	return 0
}

func (x *MsgUpdatePropertyRequest) GetNumBed() int64 {
	if x != nil {
		return x.NumBed
	}
	return 0
}

func (x *MsgUpdatePropertyRequest) GetNumBath() int64 {
	if x != nil {
		return x.NumBath
	}
	return 0
}

func (x *MsgUpdatePropertyRequest) GetNumBedroom() int64 {
	if x != nil {
		return x.NumBedroom
	}
	return 0
}

func (x *MsgUpdatePropertyRequest) GetIsGuestFavor() bool {
	if x != nil {
		return x.IsGuestFavor
	}
	return false
}

func (x *MsgUpdatePropertyRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *MsgUpdatePropertyRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MsgUpdatePropertyRequest) GetPropertyType() int64 {
	if x != nil {
		return x.PropertyType
	}
	return 0
}

type MsgDeletePropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId string `protobuf:"bytes,1,opt,name=propertyId,proto3" json:"propertyId,omitempty"`
}

func (x *MsgDeletePropertyRequest) Reset() {
	*x = MsgDeletePropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeletePropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeletePropertyRequest) ProtoMessage() {}

func (x *MsgDeletePropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeletePropertyRequest.ProtoReflect.Descriptor instead.
func (*MsgDeletePropertyRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_booking_proto_rawDescGZIP(), []int{3}
}

func (x *MsgDeletePropertyRequest) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

type MsgGetPropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId string `protobuf:"bytes,1,opt,name=propertyId,proto3" json:"propertyId,omitempty"`
}

func (x *MsgGetPropertyRequest) Reset() {
	*x = MsgGetPropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgGetPropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgGetPropertyRequest) ProtoMessage() {}

func (x *MsgGetPropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgGetPropertyRequest.ProtoReflect.Descriptor instead.
func (*MsgGetPropertyRequest) Descriptor() ([]byte, []int) {
	return file_proto_booking_booking_proto_rawDescGZIP(), []int{4}
}

func (x *MsgGetPropertyRequest) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

type MessageQueryRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginate    *sdk.Pagination `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty"`
	TimeQuery   *sdk.TimeQuery  `protobuf:"bytes,2,opt,name=timeQuery,proto3,oneof" json:"timeQuery,omitempty"`
	QueryFields *QueryRoom      `protobuf:"bytes,3,opt,name=queryFields,proto3,oneof" json:"queryFields,omitempty"`
	OrderFields *OrderRoom      `protobuf:"bytes,4,opt,name=orderFields,proto3,oneof" json:"orderFields,omitempty"`
}

func (x *MessageQueryRoom) Reset() {
	*x = MessageQueryRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageQueryRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageQueryRoom) ProtoMessage() {}

func (x *MessageQueryRoom) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageQueryRoom.ProtoReflect.Descriptor instead.
func (*MessageQueryRoom) Descriptor() ([]byte, []int) {
	return file_proto_booking_booking_proto_rawDescGZIP(), []int{5}
}

func (x *MessageQueryRoom) GetPaginate() *sdk.Pagination {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *MessageQueryRoom) GetTimeQuery() *sdk.TimeQuery {
	if x != nil {
		return x.TimeQuery
	}
	return nil
}

func (x *MessageQueryRoom) GetQueryFields() *QueryRoom {
	if x != nil {
		return x.QueryFields
	}
	return nil
}

func (x *MessageQueryRoom) GetOrderFields() *OrderRoom {
	if x != nil {
		return x.OrderFields
	}
	return nil
}

type QueryRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Id     *string `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name   *string `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *QueryRoom) Reset() {
	*x = QueryRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRoom) ProtoMessage() {}

func (x *QueryRoom) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRoom.ProtoReflect.Descriptor instead.
func (*QueryRoom) Descriptor() ([]byte, []int) {
	return file_proto_booking_booking_proto_rawDescGZIP(), []int{6}
}

func (x *QueryRoom) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryRoom) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *QueryRoom) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type OrderRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt *sdk.Sort `protobuf:"varint,1,opt,name=createdAt,proto3,enum=sdk.Sort,oneof" json:"createdAt,omitempty"`
}

func (x *OrderRoom) Reset() {
	*x = OrderRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_booking_booking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRoom) ProtoMessage() {}

func (x *OrderRoom) ProtoReflect() protoreflect.Message {
	mi := &file_proto_booking_booking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRoom.ProtoReflect.Descriptor instead.
func (*OrderRoom) Descriptor() ([]byte, []int) {
	return file_proto_booking_booking_proto_rawDescGZIP(), []int{7}
}

func (x *OrderRoom) GetCreatedAt() sdk.Sort {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return sdk.Sort(0)
}

var File_proto_booking_booking_proto protoreflect.FileDescriptor

var file_proto_booking_booking_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x13, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x34, 0x0a, 0x14, 0x4d, 0x73, 0x67, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0xac, 0x02, 0x0a, 0x18, 0x4d, 0x73, 0x67,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x75, 0x6d, 0x42, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x75,
	0x6d, 0x42, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x72,
	0x6f, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x42, 0x65,
	0x64, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x9c, 0x02, 0x0a, 0x18, 0x4d, 0x73, 0x67, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75,
	0x6d, 0x42, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x75, 0x6d,
	0x42, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x72, 0x6f,
	0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64,
	0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x47, 0x75, 0x65, 0x73, 0x74, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x47, 0x75,
	0x65, 0x73, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a, 0x18, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x49, 0x64, 0x22, 0x37, 0x0a, 0x15, 0x4d, 0x73, 0x67, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x22, 0xa4, 0x02, 0x0a, 0x10,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6f, 0x6d,
	0x12, 0x2b, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x48, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x40, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6f, 0x6d,
	0x48, 0x01, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x6f,
	0x6f, 0x6d, 0x48, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x22, 0x61, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x2c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x6f, 0x72, 0x74,
	0x48, 0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xe0,
	0x03, 0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x25, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12,
	0x20, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6f,
	0x6d, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x28, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x28, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x12, 0x28, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_booking_booking_proto_rawDescOnce sync.Once
	file_proto_booking_booking_proto_rawDescData = file_proto_booking_booking_proto_rawDesc
)

func file_proto_booking_booking_proto_rawDescGZIP() []byte {
	file_proto_booking_booking_proto_rawDescOnce.Do(func() {
		file_proto_booking_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_booking_booking_proto_rawDescData)
	})
	return file_proto_booking_booking_proto_rawDescData
}

var file_proto_booking_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_booking_booking_proto_goTypes = []interface{}{
	(*MsgGetBookingRequest)(nil),     // 0: bookingService.MsgGetBookingRequest
	(*MsgCreatePropertyRequest)(nil), // 1: bookingService.MsgCreatePropertyRequest
	(*MsgUpdatePropertyRequest)(nil), // 2: bookingService.MsgUpdatePropertyRequest
	(*MsgDeletePropertyRequest)(nil), // 3: bookingService.MsgDeletePropertyRequest
	(*MsgGetPropertyRequest)(nil),    // 4: bookingService.MsgGetPropertyRequest
	(*MessageQueryRoom)(nil),         // 5: bookingService.MessageQueryRoom
	(*QueryRoom)(nil),                // 6: bookingService.QueryRoom
	(*OrderRoom)(nil),                // 7: bookingService.OrderRoom
	(*sdk.Pagination)(nil),           // 8: sdk.Pagination
	(*sdk.TimeQuery)(nil),            // 9: sdk.TimeQuery
	(sdk.Sort)(0),                    // 10: sdk.Sort
	(*sdk.BaseResponse)(nil),         // 11: sdk.BaseResponse
}
var file_proto_booking_booking_proto_depIdxs = []int32{
	8,  // 0: bookingService.MessageQueryRoom.paginate:type_name -> sdk.Pagination
	9,  // 1: bookingService.MessageQueryRoom.timeQuery:type_name -> sdk.TimeQuery
	6,  // 2: bookingService.MessageQueryRoom.queryFields:type_name -> bookingService.QueryRoom
	7,  // 3: bookingService.MessageQueryRoom.orderFields:type_name -> bookingService.OrderRoom
	10, // 4: bookingService.OrderRoom.createdAt:type_name -> sdk.Sort
	0,  // 5: bookingService.bookingService.GetBookingDetail:input_type -> bookingService.MsgGetBookingRequest
	4,  // 6: bookingService.bookingService.GetPropertyDetail:input_type -> bookingService.MsgGetPropertyRequest
	5,  // 7: bookingService.bookingService.GetAllProperty:input_type -> bookingService.MessageQueryRoom
	1,  // 8: bookingService.bookingService.CreateProperty:input_type -> bookingService.MsgCreatePropertyRequest
	2,  // 9: bookingService.bookingService.UpdateProperty:input_type -> bookingService.MsgUpdatePropertyRequest
	3,  // 10: bookingService.bookingService.DeleteProperty:input_type -> bookingService.MsgDeletePropertyRequest
	11, // 11: bookingService.bookingService.GetBookingDetail:output_type -> sdk.BaseResponse
	11, // 12: bookingService.bookingService.GetPropertyDetail:output_type -> sdk.BaseResponse
	11, // 13: bookingService.bookingService.GetAllProperty:output_type -> sdk.BaseResponse
	11, // 14: bookingService.bookingService.CreateProperty:output_type -> sdk.BaseResponse
	11, // 15: bookingService.bookingService.UpdateProperty:output_type -> sdk.BaseResponse
	11, // 16: bookingService.bookingService.DeleteProperty:output_type -> sdk.BaseResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_booking_booking_proto_init() }
func file_proto_booking_booking_proto_init() {
	if File_proto_booking_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_booking_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgGetBookingRequest); i {
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
		file_proto_booking_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreatePropertyRequest); i {
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
		file_proto_booking_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUpdatePropertyRequest); i {
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
		file_proto_booking_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeletePropertyRequest); i {
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
		file_proto_booking_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgGetPropertyRequest); i {
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
		file_proto_booking_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageQueryRoom); i {
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
		file_proto_booking_booking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRoom); i {
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
		file_proto_booking_booking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRoom); i {
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
	file_proto_booking_booking_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_proto_booking_booking_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_proto_booking_booking_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_booking_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_booking_booking_proto_goTypes,
		DependencyIndexes: file_proto_booking_booking_proto_depIdxs,
		MessageInfos:      file_proto_booking_booking_proto_msgTypes,
	}.Build()
	File_proto_booking_booking_proto = out.File
	file_proto_booking_booking_proto_rawDesc = nil
	file_proto_booking_booking_proto_goTypes = nil
	file_proto_booking_booking_proto_depIdxs = nil
}
