// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: base_9t90.proto

package sdk

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// Describes the order in which the records need to be returned
type SORT_ORDER int32

const (
	// Fetch results in the ascending order of the provided sort key
	SORT_ORDER_ASCENDING_UNSPECIFIED SORT_ORDER = 0
	// Fetch results in the descending order of the provided sort key
	SORT_ORDER_DESCENDING SORT_ORDER = 1
)

// Enum value maps for SORT_ORDER.
var (
	SORT_ORDER_name = map[int32]string{
		0: "ASCENDING_UNSPECIFIED",
		1: "DESCENDING",
	}
	SORT_ORDER_value = map[string]int32{
		"ASCENDING_UNSPECIFIED": 0,
		"DESCENDING":            1,
	}
)

func (x SORT_ORDER) Enum() *SORT_ORDER {
	p := new(SORT_ORDER)
	*p = x
	return p
}

func (x SORT_ORDER) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SORT_ORDER) Descriptor() protoreflect.EnumDescriptor {
	return file_base_9t90_proto_enumTypes[0].Descriptor()
}

func (SORT_ORDER) Type() protoreflect.EnumType {
	return &file_base_9t90_proto_enumTypes[0]
}

func (x SORT_ORDER) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SORT_ORDER.Descriptor instead.
func (SORT_ORDER) EnumDescriptor() ([]byte, []int) {
	return file_base_9t90_proto_rawDescGZIP(), []int{0}
}

// Describes an empty object
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_9t90_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_base_9t90_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_base_9t90_proto_rawDescGZIP(), []int{0}
}

// Describes the boolean response
type BooleanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stores if the value is true or false
	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BooleanResponse) Reset() {
	*x = BooleanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_9t90_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooleanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooleanResponse) ProtoMessage() {}

func (x *BooleanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_9t90_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooleanResponse.ProtoReflect.Descriptor instead.
func (*BooleanResponse) Descriptor() ([]byte, []int) {
	return file_base_9t90_proto_rawDescGZIP(), []int{1}
}

func (x *BooleanResponse) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

// Describes the payload for a request to determine the count of records
type CountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Denotes if only active records need to be returned
	IsActive bool `protobuf:"varint,1,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *CountRequest) Reset() {
	*x = CountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_9t90_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountRequest) ProtoMessage() {}

func (x *CountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_9t90_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountRequest.ProtoReflect.Descriptor instead.
func (*CountRequest) Descriptor() ([]byte, []int) {
	return file_base_9t90_proto_rawDescGZIP(), []int{2}
}

func (x *CountRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

// Describes the count response
type CountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of records
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountResponse) Reset() {
	*x = CountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_9t90_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountResponse) ProtoMessage() {}

func (x *CountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_9t90_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountResponse.ProtoReflect.Descriptor instead.
func (*CountResponse) Descriptor() ([]byte, []int) {
	return file_base_9t90_proto_rawDescGZIP(), []int{3}
}

func (x *CountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// Describes the URL response
type URLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The URL
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *URLResponse) Reset() {
	*x = URLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_9t90_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLResponse) ProtoMessage() {}

func (x *URLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_9t90_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLResponse.ProtoReflect.Descriptor instead.
func (*URLResponse) Descriptor() ([]byte, []int) {
	return file_base_9t90_proto_rawDescGZIP(), []int{4}
}

func (x *URLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Describes the metadata of each resource
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID of the resource
	Uuid string `protobuf:"bytes,5,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Represents if the resource is active
	IsActive bool `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Stores the timestamp of when the resource was created
	CreatedAt int64 `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Stores the timestamp of when the resource was last modified
	ModifiedAt int64 `protobuf:"varint,9,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_9t90_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_base_9t90_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_base_9t90_proto_rawDescGZIP(), []int{5}
}

func (x *Metadata) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Metadata) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Metadata) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Metadata) GetModifiedAt() int64 {
	if x != nil {
		return x.ModifiedAt
	}
	return 0
}

// Describes the UUID identifier
type Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID of the resource
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Identifier) Reset() {
	*x = Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_9t90_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identifier) ProtoMessage() {}

func (x *Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_base_9t90_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identifier.ProtoReflect.Descriptor instead.
func (*Identifier) Descriptor() ([]byte, []int) {
	return file_base_9t90_proto_rawDescGZIP(), []int{6}
}

func (x *Identifier) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// Describes a simple search key request
type SearchKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The search key
	SearchKey string `protobuf:"bytes,1,opt,name=search_key,json=searchKey,proto3" json:"search_key,omitempty"`
}

func (x *SearchKeyRequest) Reset() {
	*x = SearchKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_9t90_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchKeyRequest) ProtoMessage() {}

func (x *SearchKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_9t90_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchKeyRequest.ProtoReflect.Descriptor instead.
func (*SearchKeyRequest) Descriptor() ([]byte, []int) {
	return file_base_9t90_proto_rawDescGZIP(), []int{7}
}

func (x *SearchKeyRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

var File_base_9t90_proto protoreflect.FileDescriptor

var file_base_9t90_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x39, 0x74, 0x39, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x6e, 0x69, 0x6e, 0x65, 0x54, 0x39, 0x30, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x27, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2b, 0x0a, 0x0c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1f, 0x0a,
	0x0b, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x7b,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2a, 0x0a, 0x0a, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x2a, 0x37, 0x0a, 0x0a, 0x53, 0x4f,
	0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x53, 0x43, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x39, 0x74, 0x39, 0x30, 0x2f, 0x73, 0x64, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_9t90_proto_rawDescOnce sync.Once
	file_base_9t90_proto_rawDescData = file_base_9t90_proto_rawDesc
)

func file_base_9t90_proto_rawDescGZIP() []byte {
	file_base_9t90_proto_rawDescOnce.Do(func() {
		file_base_9t90_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_9t90_proto_rawDescData)
	})
	return file_base_9t90_proto_rawDescData
}

var file_base_9t90_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_base_9t90_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_base_9t90_proto_goTypes = []interface{}{
	(SORT_ORDER)(0),          // 0: nineT90.SORT_ORDER
	(*Empty)(nil),            // 1: nineT90.Empty
	(*BooleanResponse)(nil),  // 2: nineT90.BooleanResponse
	(*CountRequest)(nil),     // 3: nineT90.CountRequest
	(*CountResponse)(nil),    // 4: nineT90.CountResponse
	(*URLResponse)(nil),      // 5: nineT90.URLResponse
	(*Metadata)(nil),         // 6: nineT90.Metadata
	(*Identifier)(nil),       // 7: nineT90.Identifier
	(*SearchKeyRequest)(nil), // 8: nineT90.SearchKeyRequest
}
var file_base_9t90_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_9t90_proto_init() }
func file_base_9t90_proto_init() {
	if File_base_9t90_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_9t90_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_base_9t90_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooleanResponse); i {
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
		file_base_9t90_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountRequest); i {
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
		file_base_9t90_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountResponse); i {
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
		file_base_9t90_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLResponse); i {
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
		file_base_9t90_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_base_9t90_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identifier); i {
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
		file_base_9t90_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchKeyRequest); i {
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
			RawDescriptor: file_base_9t90_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_9t90_proto_goTypes,
		DependencyIndexes: file_base_9t90_proto_depIdxs,
		EnumInfos:         file_base_9t90_proto_enumTypes,
		MessageInfos:      file_base_9t90_proto_msgTypes,
	}.Build()
	File_base_9t90_proto = out.File
	file_base_9t90_proto_rawDesc = nil
	file_base_9t90_proto_goTypes = nil
	file_base_9t90_proto_depIdxs = nil
}
