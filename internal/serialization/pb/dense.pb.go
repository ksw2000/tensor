// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: dense.proto

package pb

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

type Triangle int32

const (
	Triangle_NOTTRIANGLE Triangle = 0
	Triangle_UPPER       Triangle = 1
	Triangle_LOWER       Triangle = 2
	Triangle_SYMMETRIC   Triangle = 3
)

// Enum value maps for Triangle.
var (
	Triangle_name = map[int32]string{
		0: "NOTTRIANGLE",
		1: "UPPER",
		2: "LOWER",
		3: "SYMMETRIC",
	}
	Triangle_value = map[string]int32{
		"NOTTRIANGLE": 0,
		"UPPER":       1,
		"LOWER":       2,
		"SYMMETRIC":   3,
	}
)

func (x Triangle) Enum() *Triangle {
	p := new(Triangle)
	*p = x
	return p
}

func (x Triangle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Triangle) Descriptor() protoreflect.EnumDescriptor {
	return file_dense_proto_enumTypes[0].Descriptor()
}

func (Triangle) Type() protoreflect.EnumType {
	return &file_dense_proto_enumTypes[0]
}

func (x Triangle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Triangle.Descriptor instead.
func (Triangle) EnumDescriptor() ([]byte, []int) {
	return file_dense_proto_rawDescGZIP(), []int{0}
}

type AP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shape   []int32  `protobuf:"varint,1,rep,packed,name=shape,proto3" json:"shape,omitempty"`
	Strides []int32  `protobuf:"varint,2,rep,packed,name=strides,proto3" json:"strides,omitempty"`
	O       uint32   `protobuf:"varint,3,opt,name=o,proto3" json:"o,omitempty"`
	T       Triangle `protobuf:"varint,4,opt,name=t,proto3,enum=dense.Triangle" json:"t,omitempty"`
}

func (x *AP) Reset() {
	*x = AP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dense_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AP) ProtoMessage() {}

func (x *AP) ProtoReflect() protoreflect.Message {
	mi := &file_dense_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AP.ProtoReflect.Descriptor instead.
func (*AP) Descriptor() ([]byte, []int) {
	return file_dense_proto_rawDescGZIP(), []int{0}
}

func (x *AP) GetShape() []int32 {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *AP) GetStrides() []int32 {
	if x != nil {
		return x.Strides
	}
	return nil
}

func (x *AP) GetO() uint32 {
	if x != nil {
		return x.O
	}
	return 0
}

func (x *AP) GetT() Triangle {
	if x != nil {
		return x.T
	}
	return Triangle_NOTTRIANGLE
}

type Dense struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ap             *AP     `protobuf:"bytes,1,opt,name=ap,proto3" json:"ap,omitempty"`
	Data           []byte  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	TransposedWith []int32 `protobuf:"varint,3,rep,packed,name=transposedWith,proto3" json:"transposedWith,omitempty"`
	Memoryflag     uint32  `protobuf:"varint,4,opt,name=memoryflag,proto3" json:"memoryflag,omitempty"`
	Type           string  `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Dense) Reset() {
	*x = Dense{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dense_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dense) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dense) ProtoMessage() {}

func (x *Dense) ProtoReflect() protoreflect.Message {
	mi := &file_dense_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dense.ProtoReflect.Descriptor instead.
func (*Dense) Descriptor() ([]byte, []int) {
	return file_dense_proto_rawDescGZIP(), []int{1}
}

func (x *Dense) GetAp() *AP {
	if x != nil {
		return x.Ap
	}
	return nil
}

func (x *Dense) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Dense) GetTransposedWith() []int32 {
	if x != nil {
		return x.TransposedWith
	}
	return nil
}

func (x *Dense) GetMemoryflag() uint32 {
	if x != nil {
		return x.Memoryflag
	}
	return 0
}

func (x *Dense) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_dense_proto protoreflect.FileDescriptor

var file_dense_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64,
	0x65, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x0a, 0x02, 0x41, 0x50, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x6f, 0x12, 0x1d, 0x0a, 0x01, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x69, 0x61,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x01, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x05, 0x44, 0x65, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x02, 0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x64, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x50, 0x52, 0x02, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x57, 0x69,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x73, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x40, 0x0a, 0x08,
	0x54, 0x72, 0x69, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x54,
	0x52, 0x49, 0x41, 0x4e, 0x47, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x50, 0x50,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x59, 0x4d, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x10, 0x03, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dense_proto_rawDescOnce sync.Once
	file_dense_proto_rawDescData = file_dense_proto_rawDesc
)

func file_dense_proto_rawDescGZIP() []byte {
	file_dense_proto_rawDescOnce.Do(func() {
		file_dense_proto_rawDescData = protoimpl.X.CompressGZIP(file_dense_proto_rawDescData)
	})
	return file_dense_proto_rawDescData
}

var file_dense_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dense_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dense_proto_goTypes = []interface{}{
	(Triangle)(0), // 0: dense.Triangle
	(*AP)(nil),    // 1: dense.AP
	(*Dense)(nil), // 2: dense.Dense
}
var file_dense_proto_depIdxs = []int32{
	0, // 0: dense.AP.t:type_name -> dense.Triangle
	1, // 1: dense.Dense.ap:type_name -> dense.AP
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dense_proto_init() }
func file_dense_proto_init() {
	if File_dense_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dense_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AP); i {
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
		file_dense_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dense); i {
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
			RawDescriptor: file_dense_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dense_proto_goTypes,
		DependencyIndexes: file_dense_proto_depIdxs,
		EnumInfos:         file_dense_proto_enumTypes,
		MessageInfos:      file_dense_proto_msgTypes,
	}.Build()
	File_dense_proto = out.File
	file_dense_proto_rawDesc = nil
	file_dense_proto_goTypes = nil
	file_dense_proto_depIdxs = nil
}
