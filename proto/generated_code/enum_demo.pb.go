// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/enum_demo.proto

package generated_code

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

type EnumAllowingAlias int32

const (
	EnumAllowingAlias_EAA_UNSPECIFIED EnumAllowingAlias = 0
	EnumAllowingAlias_EAA_STARTED     EnumAllowingAlias = 1
	EnumAllowingAlias_EAA_RUNNING     EnumAllowingAlias = 1
	EnumAllowingAlias_EAA_FINISHED    EnumAllowingAlias = 2
)

// Enum value maps for EnumAllowingAlias.
var (
	EnumAllowingAlias_name = map[int32]string{
		0: "EAA_UNSPECIFIED",
		1: "EAA_STARTED",
		// Duplicate value: 1: "EAA_RUNNING",
		2: "EAA_FINISHED",
	}
	EnumAllowingAlias_value = map[string]int32{
		"EAA_UNSPECIFIED": 0,
		"EAA_STARTED":     1,
		"EAA_RUNNING":     1,
		"EAA_FINISHED":    2,
	}
)

func (x EnumAllowingAlias) Enum() *EnumAllowingAlias {
	p := new(EnumAllowingAlias)
	*p = x
	return p
}

func (x EnumAllowingAlias) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumAllowingAlias) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enum_demo_proto_enumTypes[0].Descriptor()
}

func (EnumAllowingAlias) Type() protoreflect.EnumType {
	return &file_proto_enum_demo_proto_enumTypes[0]
}

func (x EnumAllowingAlias) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumAllowingAlias.Descriptor instead.
func (EnumAllowingAlias) EnumDescriptor() ([]byte, []int) {
	return file_proto_enum_demo_proto_rawDescGZIP(), []int{0}
}

type EnumNotAllowingAlias int32

const (
	EnumNotAllowingAlias_ENAA_UNSPECIFIED EnumNotAllowingAlias = 0
	EnumNotAllowingAlias_ENAA_STARTED     EnumNotAllowingAlias = 1
	// ENAA_RUNNING = 1;  // Uncommenting this line will cause a warning message.
	EnumNotAllowingAlias_ENAA_FINISHED EnumNotAllowingAlias = 2
)

// Enum value maps for EnumNotAllowingAlias.
var (
	EnumNotAllowingAlias_name = map[int32]string{
		0: "ENAA_UNSPECIFIED",
		1: "ENAA_STARTED",
		2: "ENAA_FINISHED",
	}
	EnumNotAllowingAlias_value = map[string]int32{
		"ENAA_UNSPECIFIED": 0,
		"ENAA_STARTED":     1,
		"ENAA_FINISHED":    2,
	}
)

func (x EnumNotAllowingAlias) Enum() *EnumNotAllowingAlias {
	p := new(EnumNotAllowingAlias)
	*p = x
	return p
}

func (x EnumNotAllowingAlias) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumNotAllowingAlias) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enum_demo_proto_enumTypes[1].Descriptor()
}

func (EnumNotAllowingAlias) Type() protoreflect.EnumType {
	return &file_proto_enum_demo_proto_enumTypes[1]
}

func (x EnumNotAllowingAlias) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnumNotAllowingAlias.Descriptor instead.
func (EnumNotAllowingAlias) EnumDescriptor() ([]byte, []int) {
	return file_proto_enum_demo_proto_rawDescGZIP(), []int{1}
}

type DESTINATION int32

const (
	DESTINATION_DESTINATION_UNSPECIFIED DESTINATION = 0
	DESTINATION_DESTINATION_VARANASI    DESTINATION = 1
	DESTINATION_DESTINATION_MP          DESTINATION = 2
	DESTINATION_DESTINATION_INDORE      DESTINATION = 2
	DESTINATION_DESTINATION_HYD         DESTINATION = 3
)

// Enum value maps for DESTINATION.
var (
	DESTINATION_name = map[int32]string{
		0: "DESTINATION_UNSPECIFIED",
		1: "DESTINATION_VARANASI",
		2: "DESTINATION_MP",
		// Duplicate value: 2: "DESTINATION_INDORE",
		3: "DESTINATION_HYD",
	}
	DESTINATION_value = map[string]int32{
		"DESTINATION_UNSPECIFIED": 0,
		"DESTINATION_VARANASI":    1,
		"DESTINATION_MP":          2,
		"DESTINATION_INDORE":      2,
		"DESTINATION_HYD":         3,
	}
)

func (x DESTINATION) Enum() *DESTINATION {
	p := new(DESTINATION)
	*p = x
	return p
}

func (x DESTINATION) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESTINATION) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_enum_demo_proto_enumTypes[2].Descriptor()
}

func (DESTINATION) Type() protoreflect.EnumType {
	return &file_proto_enum_demo_proto_enumTypes[2]
}

func (x DESTINATION) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESTINATION.Descriptor instead.
func (DESTINATION) EnumDescriptor() ([]byte, []int) {
	return file_proto_enum_demo_proto_rawDescGZIP(), []int{2}
}

type EnumContainer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Destination   *DESTINATION           `protobuf:"varint,1,opt,name=destination,proto3,enum=DESTINATION,oneof" json:"destination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnumContainer) Reset() {
	*x = EnumContainer{}
	mi := &file_proto_enum_demo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnumContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumContainer) ProtoMessage() {}

func (x *EnumContainer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_enum_demo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumContainer.ProtoReflect.Descriptor instead.
func (*EnumContainer) Descriptor() ([]byte, []int) {
	return file_proto_enum_demo_proto_rawDescGZIP(), []int{0}
}

func (x *EnumContainer) GetDestination() DESTINATION {
	if x != nil && x.Destination != nil {
		return *x.Destination
	}
	return DESTINATION_DESTINATION_UNSPECIFIED
}

var File_proto_enum_demo_proto protoreflect.FileDescriptor

const file_proto_enum_demo_proto_rawDesc = "" +
	"\n" +
	"\x15proto/enum_demo.proto\"T\n" +
	"\rEnumContainer\x123\n" +
	"\vdestination\x18\x01 \x01(\x0e2\f.DESTINATIONH\x00R\vdestination\x88\x01\x01B\x0e\n" +
	"\f_destination*`\n" +
	"\x11EnumAllowingAlias\x12\x13\n" +
	"\x0fEAA_UNSPECIFIED\x10\x00\x12\x0f\n" +
	"\vEAA_STARTED\x10\x01\x12\x0f\n" +
	"\vEAA_RUNNING\x10\x01\x12\x10\n" +
	"\fEAA_FINISHED\x10\x02\x1a\x02\x10\x01*Q\n" +
	"\x14EnumNotAllowingAlias\x12\x14\n" +
	"\x10ENAA_UNSPECIFIED\x10\x00\x12\x10\n" +
	"\fENAA_STARTED\x10\x01\x12\x11\n" +
	"\rENAA_FINISHED\x10\x02*\x89\x01\n" +
	"\vDESTINATION\x12\x1b\n" +
	"\x17DESTINATION_UNSPECIFIED\x10\x00\x12\x18\n" +
	"\x14DESTINATION_VARANASI\x10\x01\x12\x12\n" +
	"\x0eDESTINATION_MP\x10\x02\x12\x16\n" +
	"\x12DESTINATION_INDORE\x10\x02\x12\x13\n" +
	"\x0fDESTINATION_HYD\x10\x03\x1a\x02\x10\x01B\x16Z\x14proto/generated_codeb\x06proto3"

var (
	file_proto_enum_demo_proto_rawDescOnce sync.Once
	file_proto_enum_demo_proto_rawDescData []byte
)

func file_proto_enum_demo_proto_rawDescGZIP() []byte {
	file_proto_enum_demo_proto_rawDescOnce.Do(func() {
		file_proto_enum_demo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_enum_demo_proto_rawDesc), len(file_proto_enum_demo_proto_rawDesc)))
	})
	return file_proto_enum_demo_proto_rawDescData
}

var file_proto_enum_demo_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_enum_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_enum_demo_proto_goTypes = []any{
	(EnumAllowingAlias)(0),    // 0: EnumAllowingAlias
	(EnumNotAllowingAlias)(0), // 1: EnumNotAllowingAlias
	(DESTINATION)(0),          // 2: DESTINATION
	(*EnumContainer)(nil),     // 3: EnumContainer
}
var file_proto_enum_demo_proto_depIdxs = []int32{
	2, // 0: EnumContainer.destination:type_name -> DESTINATION
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_enum_demo_proto_init() }
func file_proto_enum_demo_proto_init() {
	if File_proto_enum_demo_proto != nil {
		return
	}
	file_proto_enum_demo_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_enum_demo_proto_rawDesc), len(file_proto_enum_demo_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_enum_demo_proto_goTypes,
		DependencyIndexes: file_proto_enum_demo_proto_depIdxs,
		EnumInfos:         file_proto_enum_demo_proto_enumTypes,
		MessageInfos:      file_proto_enum_demo_proto_msgTypes,
	}.Build()
	File_proto_enum_demo_proto = out.File
	file_proto_enum_demo_proto_goTypes = nil
	file_proto_enum_demo_proto_depIdxs = nil
}
