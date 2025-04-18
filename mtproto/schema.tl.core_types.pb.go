//
// WARNING! All changes made in this file will be lost!
// Created from 'scheme.tl' by 'codegen_proto.py'
//
//  Copyright (c) 2017, https://github.com/xxxxx
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: schema.tl.core_types.proto

package mtproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// /////////////////////////////////////////////////////////////////////////////
// Error <--
//   - TL_error
type Error_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error_Data) Reset() {
	*x = Error_Data{}
	mi := &file_schema_tl_core_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error_Data) ProtoMessage() {}

func (x *Error_Data) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error_Data.ProtoReflect.Descriptor instead.
func (*Error_Data) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{0}
}

func (x *Error_Data) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error_Data) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Constructor   TLConstructor          `protobuf:"varint,1,opt,name=constructor,proto3,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data          *Error_Data            `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_schema_tl_core_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetConstructor() TLConstructor {
	if x != nil {
		return x.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (x *Error) GetData() *Error_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// error#c4b9f9bb code:int text:string = Error;
type TLError struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *Error_Data            `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TLError) Reset() {
	*x = TLError{}
	mi := &file_schema_tl_core_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TLError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLError) ProtoMessage() {}

func (x *TLError) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLError.ProtoReflect.Descriptor instead.
func (*TLError) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{2}
}

func (x *TLError) GetData() *Error_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// Null <--
//   - TL_null
type Null_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Null_Data) Reset() {
	*x = Null_Data{}
	mi := &file_schema_tl_core_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Null_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Null_Data) ProtoMessage() {}

func (x *Null_Data) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Null_Data.ProtoReflect.Descriptor instead.
func (*Null_Data) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{3}
}

type Null struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Constructor   TLConstructor          `protobuf:"varint,1,opt,name=constructor,proto3,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data          *Null_Data             `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Null) Reset() {
	*x = Null{}
	mi := &file_schema_tl_core_types_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Null) ProtoMessage() {}

func (x *Null) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Null.ProtoReflect.Descriptor instead.
func (*Null) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{4}
}

func (x *Null) GetConstructor() TLConstructor {
	if x != nil {
		return x.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (x *Null) GetData() *Null_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// null#56730bcc = Null;
type TLNull struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *Null_Data             `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TLNull) Reset() {
	*x = TLNull{}
	mi := &file_schema_tl_core_types_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TLNull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLNull) ProtoMessage() {}

func (x *TLNull) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLNull.ProtoReflect.Descriptor instead.
func (*TLNull) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{5}
}

func (x *TLNull) GetData() *Null_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// Bool <--
//   - TL_boolFalse
//   - TL_boolTrue
type Bool_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bool_Data) Reset() {
	*x = Bool_Data{}
	mi := &file_schema_tl_core_types_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bool_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bool_Data) ProtoMessage() {}

func (x *Bool_Data) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bool_Data.ProtoReflect.Descriptor instead.
func (*Bool_Data) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{6}
}

type Bool struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Constructor   TLConstructor          `protobuf:"varint,1,opt,name=constructor,proto3,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data          *Bool_Data             `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bool) Reset() {
	*x = Bool{}
	mi := &file_schema_tl_core_types_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bool) ProtoMessage() {}

func (x *Bool) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bool.ProtoReflect.Descriptor instead.
func (*Bool) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{7}
}

func (x *Bool) GetConstructor() TLConstructor {
	if x != nil {
		return x.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (x *Bool) GetData() *Bool_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// boolFalse#bc799737 = Bool;
type TLBoolFalse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *Bool_Data             `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TLBoolFalse) Reset() {
	*x = TLBoolFalse{}
	mi := &file_schema_tl_core_types_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TLBoolFalse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLBoolFalse) ProtoMessage() {}

func (x *TLBoolFalse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLBoolFalse.ProtoReflect.Descriptor instead.
func (*TLBoolFalse) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{8}
}

func (x *TLBoolFalse) GetData() *Bool_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// boolTrue#997275b5 = Bool;
type TLBoolTrue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *Bool_Data             `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TLBoolTrue) Reset() {
	*x = TLBoolTrue{}
	mi := &file_schema_tl_core_types_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TLBoolTrue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLBoolTrue) ProtoMessage() {}

func (x *TLBoolTrue) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLBoolTrue.ProtoReflect.Descriptor instead.
func (*TLBoolTrue) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{9}
}

func (x *TLBoolTrue) GetData() *Bool_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// True <--
//   - TL_true
type True_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *True_Data) Reset() {
	*x = True_Data{}
	mi := &file_schema_tl_core_types_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *True_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*True_Data) ProtoMessage() {}

func (x *True_Data) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use True_Data.ProtoReflect.Descriptor instead.
func (*True_Data) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{10}
}

type True struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Constructor   TLConstructor          `protobuf:"varint,1,opt,name=constructor,proto3,enum=mtproto.TLConstructor" json:"constructor,omitempty"`
	Data          *True_Data             `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *True) Reset() {
	*x = True{}
	mi := &file_schema_tl_core_types_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *True) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*True) ProtoMessage() {}

func (x *True) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use True.ProtoReflect.Descriptor instead.
func (*True) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{11}
}

func (x *True) GetConstructor() TLConstructor {
	if x != nil {
		return x.Constructor
	}
	return TLConstructor_CRC32_UNKNOWN
}

func (x *True) GetData() *True_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// true#3fedd339 = True;
type TLTrue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *True_Data             `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TLTrue) Reset() {
	*x = TLTrue{}
	mi := &file_schema_tl_core_types_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TLTrue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLTrue) ProtoMessage() {}

func (x *TLTrue) ProtoReflect() protoreflect.Message {
	mi := &file_schema_tl_core_types_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLTrue.ProtoReflect.Descriptor instead.
func (*TLTrue) Descriptor() ([]byte, []int) {
	return file_schema_tl_core_types_proto_rawDescGZIP(), []int{12}
}

func (x *TLTrue) GetData() *True_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_schema_tl_core_types_proto protoreflect.FileDescriptor

const file_schema_tl_core_types_proto_rawDesc = "" +
	"\n" +
	"\x1aschema.tl.core_types.proto\x12\amtproto\x1a\x15schema.tl.crc32.proto\"4\n" +
	"\n" +
	"Error_Data\x12\x12\n" +
	"\x04code\x18\x01 \x01(\x05R\x04code\x12\x12\n" +
	"\x04text\x18\x02 \x01(\tR\x04text\"j\n" +
	"\x05Error\x128\n" +
	"\vconstructor\x18\x01 \x01(\x0e2\x16.mtproto.TLConstructorR\vconstructor\x12'\n" +
	"\x04Data\x18\x02 \x01(\v2\x13.mtproto.Error_DataR\x04Data\"3\n" +
	"\bTL_error\x12'\n" +
	"\x04Data\x18\x02 \x01(\v2\x13.mtproto.Error_DataR\x04Data\"\v\n" +
	"\tNull_Data\"h\n" +
	"\x04Null\x128\n" +
	"\vconstructor\x18\x01 \x01(\x0e2\x16.mtproto.TLConstructorR\vconstructor\x12&\n" +
	"\x04Data\x18\x02 \x01(\v2\x12.mtproto.Null_DataR\x04Data\"1\n" +
	"\aTL_null\x12&\n" +
	"\x04Data\x18\x02 \x01(\v2\x12.mtproto.Null_DataR\x04Data\"\v\n" +
	"\tBool_Data\"h\n" +
	"\x04Bool\x128\n" +
	"\vconstructor\x18\x01 \x01(\x0e2\x16.mtproto.TLConstructorR\vconstructor\x12&\n" +
	"\x04Data\x18\x02 \x01(\v2\x12.mtproto.Bool_DataR\x04Data\"6\n" +
	"\fTL_boolFalse\x12&\n" +
	"\x04Data\x18\x02 \x01(\v2\x12.mtproto.Bool_DataR\x04Data\"5\n" +
	"\vTL_boolTrue\x12&\n" +
	"\x04Data\x18\x02 \x01(\v2\x12.mtproto.Bool_DataR\x04Data\"\v\n" +
	"\tTrue_Data\"h\n" +
	"\x04True\x128\n" +
	"\vconstructor\x18\x01 \x01(\x0e2\x16.mtproto.TLConstructorR\vconstructor\x12&\n" +
	"\x04Data\x18\x02 \x01(\v2\x12.mtproto.True_DataR\x04Data\"1\n" +
	"\aTL_true\x12&\n" +
	"\x04Data\x18\x02 \x01(\v2\x12.mtproto.True_DataR\x04DataB[\n" +
	"\x1fcom.takemakefake.engine.mtprotoB\aMTProtoH\x02Z-github.com/takemakefake/proto/mtproto;mtprotob\x06proto3"

var (
	file_schema_tl_core_types_proto_rawDescOnce sync.Once
	file_schema_tl_core_types_proto_rawDescData []byte
)

func file_schema_tl_core_types_proto_rawDescGZIP() []byte {
	file_schema_tl_core_types_proto_rawDescOnce.Do(func() {
		file_schema_tl_core_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_schema_tl_core_types_proto_rawDesc), len(file_schema_tl_core_types_proto_rawDesc)))
	})
	return file_schema_tl_core_types_proto_rawDescData
}

var file_schema_tl_core_types_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_schema_tl_core_types_proto_goTypes = []any{
	(*Error_Data)(nil),  // 0: mtproto.Error_Data
	(*Error)(nil),       // 1: mtproto.Error
	(*TLError)(nil),     // 2: mtproto.TL_error
	(*Null_Data)(nil),   // 3: mtproto.Null_Data
	(*Null)(nil),        // 4: mtproto.Null
	(*TLNull)(nil),      // 5: mtproto.TL_null
	(*Bool_Data)(nil),   // 6: mtproto.Bool_Data
	(*Bool)(nil),        // 7: mtproto.Bool
	(*TLBoolFalse)(nil), // 8: mtproto.TL_boolFalse
	(*TLBoolTrue)(nil),  // 9: mtproto.TL_boolTrue
	(*True_Data)(nil),   // 10: mtproto.True_Data
	(*True)(nil),        // 11: mtproto.True
	(*TLTrue)(nil),      // 12: mtproto.TL_true
	(TLConstructor)(0),  // 13: mtproto.TLConstructor
}
var file_schema_tl_core_types_proto_depIdxs = []int32{
	13, // 0: mtproto.Error.constructor:type_name -> mtproto.TLConstructor
	0,  // 1: mtproto.Error.Data:type_name -> mtproto.Error_Data
	0,  // 2: mtproto.TL_error.Data:type_name -> mtproto.Error_Data
	13, // 3: mtproto.Null.constructor:type_name -> mtproto.TLConstructor
	3,  // 4: mtproto.Null.Data:type_name -> mtproto.Null_Data
	3,  // 5: mtproto.TL_null.Data:type_name -> mtproto.Null_Data
	13, // 6: mtproto.Bool.constructor:type_name -> mtproto.TLConstructor
	6,  // 7: mtproto.Bool.Data:type_name -> mtproto.Bool_Data
	6,  // 8: mtproto.TL_boolFalse.Data:type_name -> mtproto.Bool_Data
	6,  // 9: mtproto.TL_boolTrue.Data:type_name -> mtproto.Bool_Data
	13, // 10: mtproto.True.constructor:type_name -> mtproto.TLConstructor
	10, // 11: mtproto.True.Data:type_name -> mtproto.True_Data
	10, // 12: mtproto.TL_true.Data:type_name -> mtproto.True_Data
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_schema_tl_core_types_proto_init() }
func file_schema_tl_core_types_proto_init() {
	if File_schema_tl_core_types_proto != nil {
		return
	}
	file_schema_tl_crc32_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_schema_tl_core_types_proto_rawDesc), len(file_schema_tl_core_types_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_tl_core_types_proto_goTypes,
		DependencyIndexes: file_schema_tl_core_types_proto_depIdxs,
		MessageInfos:      file_schema_tl_core_types_proto_msgTypes,
	}.Build()
	File_schema_tl_core_types_proto = out.File
	file_schema_tl_core_types_proto_goTypes = nil
	file_schema_tl_core_types_proto_depIdxs = nil
}
