// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.5
// source: levels.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Levels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http    *Enabled `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc    *Enabled `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	Logging *Enabled `protobuf:"bytes,3,opt,name=logging,proto3" json:"logging,omitempty"`
	Metric  *Enabled `protobuf:"bytes,4,opt,name=metric,proto3" json:"metric,omitempty"`
	Sentry  *Enabled `protobuf:"bytes,5,opt,name=sentry,proto3" json:"sentry,omitempty"`
	Tracing *Enabled `protobuf:"bytes,6,opt,name=tracing,proto3" json:"tracing,omitempty"`
	Queue   *Enabled `protobuf:"bytes,7,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *Levels) Reset() {
	*x = Levels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_levels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Levels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Levels) ProtoMessage() {}

func (x *Levels) ProtoReflect() protoreflect.Message {
	mi := &file_levels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Levels.ProtoReflect.Descriptor instead.
func (*Levels) Descriptor() ([]byte, []int) {
	return file_levels_proto_rawDescGZIP(), []int{0}
}

func (x *Levels) GetHttp() *Enabled {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Levels) GetGrpc() *Enabled {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Levels) GetLogging() *Enabled {
	if x != nil {
		return x.Logging
	}
	return nil
}

func (x *Levels) GetMetric() *Enabled {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *Levels) GetSentry() *Enabled {
	if x != nil {
		return x.Sentry
	}
	return nil
}

func (x *Levels) GetTracing() *Enabled {
	if x != nil {
		return x.Tracing
	}
	return nil
}

func (x *Levels) GetQueue() *Enabled {
	if x != nil {
		return x.Queue
	}
	return nil
}

type Enabled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*Enabled_Null
	//	*Enabled_Enabled
	Kind isEnabled_Kind `protobuf_oneof:"kind"`
}

func (x *Enabled) Reset() {
	*x = Enabled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_levels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enabled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enabled) ProtoMessage() {}

func (x *Enabled) ProtoReflect() protoreflect.Message {
	mi := &file_levels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enabled.ProtoReflect.Descriptor instead.
func (*Enabled) Descriptor() ([]byte, []int) {
	return file_levels_proto_rawDescGZIP(), []int{1}
}

func (m *Enabled) GetKind() isEnabled_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Enabled) GetNull() structpb.NullValue {
	if x, ok := x.GetKind().(*Enabled_Null); ok {
		return x.Null
	}
	return structpb.NullValue(0)
}

func (x *Enabled) GetEnabled() bool {
	if x, ok := x.GetKind().(*Enabled_Enabled); ok {
		return x.Enabled
	}
	return false
}

type isEnabled_Kind interface {
	isEnabled_Kind()
}

type Enabled_Null struct {
	Null structpb.NullValue `protobuf:"varint,1,opt,name=null,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Enabled_Enabled struct {
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3,oneof"`
}

func (*Enabled_Null) isEnabled_Kind() {}

func (*Enabled_Enabled) isEnabled_Kind() {}

var File_levels_proto protoreflect.FileDescriptor

var file_levels_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x65, 0x6c, 0x64, 0x6f, 0x72, 0x61, 0x64, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7,
	0x02, 0x0a, 0x06, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x68, 0x74, 0x74,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6c, 0x64, 0x6f, 0x72, 0x61,
	0x64, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x2d, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6c, 0x64, 0x6f, 0x72, 0x61, 0x64,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x33, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6c, 0x64, 0x6f, 0x72,
	0x61, 0x64, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65,
	0x6c, 0x64, 0x6f, 0x72, 0x61, 0x64, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12,
	0x31, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x65, 0x6c, 0x64, 0x6f, 0x72, 0x61, 0x64, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6c, 0x64, 0x6f, 0x72, 0x61, 0x64, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6c, 0x64, 0x6f, 0x72, 0x61, 0x64,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0x5f, 0x0a, 0x07, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x65, 0x6c, 0x64, 0x6f, 0x72, 0x61, 0x64, 0x6f, 0x2e, 0x72, 0x75, 0x2f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_levels_proto_rawDescOnce sync.Once
	file_levels_proto_rawDescData = file_levels_proto_rawDesc
)

func file_levels_proto_rawDescGZIP() []byte {
	file_levels_proto_rawDescOnce.Do(func() {
		file_levels_proto_rawDescData = protoimpl.X.CompressGZIP(file_levels_proto_rawDescData)
	})
	return file_levels_proto_rawDescData
}

var file_levels_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_levels_proto_goTypes = []interface{}{
	(*Levels)(nil),          // 0: eldorado.service.Levels
	(*Enabled)(nil),         // 1: eldorado.service.Enabled
	(structpb.NullValue)(0), // 2: google.protobuf.NullValue
}
var file_levels_proto_depIdxs = []int32{
	1, // 0: eldorado.service.Levels.http:type_name -> eldorado.service.Enabled
	1, // 1: eldorado.service.Levels.grpc:type_name -> eldorado.service.Enabled
	1, // 2: eldorado.service.Levels.logging:type_name -> eldorado.service.Enabled
	1, // 3: eldorado.service.Levels.metric:type_name -> eldorado.service.Enabled
	1, // 4: eldorado.service.Levels.sentry:type_name -> eldorado.service.Enabled
	1, // 5: eldorado.service.Levels.tracing:type_name -> eldorado.service.Enabled
	1, // 6: eldorado.service.Levels.queue:type_name -> eldorado.service.Enabled
	2, // 7: eldorado.service.Enabled.null:type_name -> google.protobuf.NullValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_levels_proto_init() }
func file_levels_proto_init() {
	if File_levels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_levels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Levels); i {
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
		file_levels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enabled); i {
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
	file_levels_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Enabled_Null)(nil),
		(*Enabled_Enabled)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_levels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_levels_proto_goTypes,
		DependencyIndexes: file_levels_proto_depIdxs,
		MessageInfos:      file_levels_proto_msgTypes,
	}.Build()
	File_levels_proto = out.File
	file_levels_proto_rawDesc = nil
	file_levels_proto_goTypes = nil
	file_levels_proto_depIdxs = nil
}