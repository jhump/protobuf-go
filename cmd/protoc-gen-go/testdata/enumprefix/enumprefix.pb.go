// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/enumprefix/enumprefix.proto

package enumprefix

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
	sync "sync"
)

type Strip int32

const (
	Strip_ZERO Strip = 0
	Strip_ONE  Strip = 1
)

// Enum value maps for Strip.
var (
	Strip_name = map[int32]string{
		0: "STRIP_ZERO",
		1: "STRIP_ONE",
	}
	Strip_value = map[string]int32{
		"STRIP_ZERO": 0,
		"STRIP_ONE":  1,
	}
)

func (x Strip) Enum() *Strip {
	p := new(Strip)
	*p = x
	return p
}

func (x Strip) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Strip) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_enumTypes[0].Descriptor()
}

func (Strip) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_enumTypes[0]
}

func (x Strip) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Strip.Descriptor instead.
func (Strip) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescGZIP(), []int{0}
}

type Both int32

const (
	Both_ZERO Both = 0
	Both_ONE  Both = 1
)

// Old (prefixed) names for Both enum values.
const (
	Both_BOTH_ZERO Both = Both_ZERO
	Both_BOTH_ONE  Both = Both_ONE
)

// Enum value maps for Both.
var (
	Both_name = map[int32]string{
		0: "BOTH_ZERO",
		1: "BOTH_ONE",
	}
	Both_value = map[string]int32{
		"BOTH_ZERO": 0,
		"BOTH_ONE":  1,
	}
)

func (x Both) Enum() *Both {
	p := new(Both)
	*p = x
	return p
}

func (x Both) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Both) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_enumTypes[1].Descriptor()
}

func (Both) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_enumTypes[1]
}

func (x Both) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Both.Descriptor instead.
func (Both) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescGZIP(), []int{1}
}

type BothNoPrefix int32

const (
	BothNoPrefix_ZERO BothNoPrefix = 0
	BothNoPrefix_ONE  BothNoPrefix = 1
)

// Enum value maps for BothNoPrefix.
var (
	BothNoPrefix_name = map[int32]string{
		0: "ZERO",
		1: "ONE",
	}
	BothNoPrefix_value = map[string]int32{
		"ZERO": 0,
		"ONE":  1,
	}
)

func (x BothNoPrefix) Enum() *BothNoPrefix {
	p := new(BothNoPrefix)
	*p = x
	return p
}

func (x BothNoPrefix) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BothNoPrefix) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_enumTypes[2].Descriptor()
}

func (BothNoPrefix) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_enumTypes[2]
}

func (x BothNoPrefix) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BothNoPrefix.Descriptor instead.
func (BothNoPrefix) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescGZIP(), []int{2}
}

type BothButOne int32

const (
	BothButOne_ZERO             BothButOne = 0
	BothButOne_BOTH_BUT_ONE_ONE BothButOne = 1
)

// Old (prefixed) names for BothButOne enum values.
const (
	BothButOne_BOTH_BUT_ONE_ZERO BothButOne = BothButOne_ZERO
)

// Enum value maps for BothButOne.
var (
	BothButOne_name = map[int32]string{
		0: "BOTH_BUT_ONE_ZERO",
		1: "BOTH_BUT_ONE_ONE",
	}
	BothButOne_value = map[string]int32{
		"BOTH_BUT_ONE_ZERO": 0,
		"BOTH_BUT_ONE_ONE":  1,
	}
)

func (x BothButOne) Enum() *BothButOne {
	p := new(BothButOne)
	*p = x
	return p
}

func (x BothButOne) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BothButOne) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_enumTypes[3].Descriptor()
}

func (BothButOne) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_enumTypes[3]
}

func (x BothButOne) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BothButOne.Descriptor instead.
func (BothButOne) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescGZIP(), []int{3}
}

var File_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x26, 0x0a, 0x05, 0x53, 0x74, 0x72, 0x69, 0x70, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x52, 0x49, 0x50, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x54, 0x52, 0x49, 0x50, 0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x2a, 0x2c,
	0x0a, 0x04, 0x42, 0x6f, 0x74, 0x68, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x4f, 0x54, 0x48, 0x5f, 0x5a,
	0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x4f, 0x54, 0x48, 0x5f, 0x4f, 0x4e,
	0x45, 0x10, 0x01, 0x1a, 0x07, 0x3a, 0x05, 0xd2, 0x3e, 0x02, 0x18, 0x02, 0x2a, 0x2a, 0x0a, 0x0c,
	0x42, 0x6f, 0x74, 0x68, 0x4e, 0x6f, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x08, 0x0a, 0x04,
	0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x1a,
	0x07, 0x3a, 0x05, 0xd2, 0x3e, 0x02, 0x18, 0x02, 0x2a, 0x4b, 0x0a, 0x0a, 0x42, 0x6f, 0x74, 0x68,
	0x42, 0x75, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x4f, 0x54, 0x48, 0x5f, 0x42,
	0x55, 0x54, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x1d, 0x0a,
	0x10, 0x42, 0x4f, 0x54, 0x48, 0x5f, 0x42, 0x55, 0x54, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x4f, 0x4e,
	0x45, 0x10, 0x01, 0x1a, 0x07, 0x12, 0x05, 0xd2, 0x3e, 0x02, 0x18, 0x01, 0x1a, 0x07, 0x3a, 0x05,
	0xd2, 0x3e, 0x02, 0x18, 0x02, 0x42, 0x4a, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x92, 0x03, 0x05, 0xd2, 0x3e, 0x02, 0x18,
	0x03, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe9, 0x07,
}

var (
	file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescData = file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_goTypes = []any{
	(Strip)(0),        // 0: goproto.protoc.enumprefix.Strip
	(Both)(0),         // 1: goproto.protoc.enumprefix.Both
	(BothNoPrefix)(0), // 2: goproto.protoc.enumprefix.BothNoPrefix
	(BothButOne)(0),   // 3: goproto.protoc.enumprefix.BothButOne
}
var file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_init() }
func file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_init() {
	if File_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_enumTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto = out.File
	file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_enumprefix_enumprefix_proto_depIdxs = nil
}