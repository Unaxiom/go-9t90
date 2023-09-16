// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: buf/validate/priv/private.proto

package priv

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Do not use. Internal to protovalidate library
type FieldConstraints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cel []*Constraint `protobuf:"bytes,1,rep,name=cel,proto3" json:"cel,omitempty"`
}

func (x *FieldConstraints) Reset() {
	*x = FieldConstraints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_priv_private_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldConstraints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldConstraints) ProtoMessage() {}

func (x *FieldConstraints) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_priv_private_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldConstraints.ProtoReflect.Descriptor instead.
func (*FieldConstraints) Descriptor() ([]byte, []int) {
	return file_buf_validate_priv_private_proto_rawDescGZIP(), []int{0}
}

func (x *FieldConstraints) GetCel() []*Constraint {
	if x != nil {
		return x.Cel
	}
	return nil
}

// Do not use. Internal to protovalidate library
type Constraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Expression string `protobuf:"bytes,3,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *Constraint) Reset() {
	*x = Constraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_priv_private_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Constraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Constraint) ProtoMessage() {}

func (x *Constraint) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_priv_private_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Constraint.ProtoReflect.Descriptor instead.
func (*Constraint) Descriptor() ([]byte, []int) {
	return file_buf_validate_priv_private_proto_rawDescGZIP(), []int{1}
}

func (x *Constraint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Constraint) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Constraint) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

var file_buf_validate_priv_private_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldConstraints)(nil),
		Field:         1160,
		Name:          "buf.validate.priv.field",
		Tag:           "bytes,1160,opt,name=field",
		Filename:      "buf/validate/priv/private.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// Do not use. Internal to protovalidate library
	//
	// optional buf.validate.priv.FieldConstraints field = 1160;
	E_Field = &file_buf_validate_priv_private_proto_extTypes[0]
)

var File_buf_validate_priv_private_proto protoreflect.FileDescriptor

var file_buf_validate_priv_private_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x10, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x03, 0x63, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x03, 0x63, 0x65, 0x6c, 0x22, 0x56, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x3a, 0x5c, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x88, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x77, 0x0a, 0x17, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x42, 0x0c, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x62, 0x75,
	0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62,
	0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_buf_validate_priv_private_proto_rawDescOnce sync.Once
	file_buf_validate_priv_private_proto_rawDescData = file_buf_validate_priv_private_proto_rawDesc
)

func file_buf_validate_priv_private_proto_rawDescGZIP() []byte {
	file_buf_validate_priv_private_proto_rawDescOnce.Do(func() {
		file_buf_validate_priv_private_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_priv_private_proto_rawDescData)
	})
	return file_buf_validate_priv_private_proto_rawDescData
}

var file_buf_validate_priv_private_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_validate_priv_private_proto_goTypes = []interface{}{
	(*FieldConstraints)(nil),          // 0: buf.validate.priv.FieldConstraints
	(*Constraint)(nil),                // 1: buf.validate.priv.Constraint
	(*descriptorpb.FieldOptions)(nil), // 2: google.protobuf.FieldOptions
}
var file_buf_validate_priv_private_proto_depIdxs = []int32{
	1, // 0: buf.validate.priv.FieldConstraints.cel:type_name -> buf.validate.priv.Constraint
	2, // 1: buf.validate.priv.field:extendee -> google.protobuf.FieldOptions
	0, // 2: buf.validate.priv.field:type_name -> buf.validate.priv.FieldConstraints
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buf_validate_priv_private_proto_init() }
func file_buf_validate_priv_private_proto_init() {
	if File_buf_validate_priv_private_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_priv_private_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldConstraints); i {
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
		file_buf_validate_priv_private_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Constraint); i {
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
			RawDescriptor: file_buf_validate_priv_private_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_priv_private_proto_goTypes,
		DependencyIndexes: file_buf_validate_priv_private_proto_depIdxs,
		MessageInfos:      file_buf_validate_priv_private_proto_msgTypes,
		ExtensionInfos:    file_buf_validate_priv_private_proto_extTypes,
	}.Build()
	File_buf_validate_priv_private_proto = out.File
	file_buf_validate_priv_private_proto_rawDesc = nil
	file_buf_validate_priv_private_proto_goTypes = nil
	file_buf_validate_priv_private_proto_depIdxs = nil
}
