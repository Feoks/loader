// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: echo-service.proto

package echopb

import (
	context "context"
	_ "git.repo.services.lenvendo.ru/grade-factor/api/service"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_echo_service_proto protoreflect.FileDescriptor

var file_echo_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x63, 0x68, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6c, 0x65, 0x6e, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x65, 0x63, 0x68, 0x6f, 0x2d, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x65, 0x63, 0x68, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x85, 0x02, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x08, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65,
	0x73, 0x73, 0x12, 0x17, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x76, 0x65,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x63,
	0x68, 0x6f, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f,
	0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x54, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x18, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x4c,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x65, 0x63, 0x68, 0x6f,
	0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x60, 0x0a, 0x0b,
	0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x1a, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x63, 0x68, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x42, 0x11,
	0x5a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_echo_service_proto_goTypes = []interface{}{
	(*LivenessRequest)(nil),     // 0: echopb.LivenessRequest
	(*ReadinessRequest)(nil),    // 1: echopb.ReadinessRequest
	(*VersionRequest)(nil),      // 2: echopb.VersionRequest
	(*GetEchoListRequest)(nil),  // 3: echopb.GetEchoListRequest
	(*LivenessResponse)(nil),    // 4: echopb.LivenessResponse
	(*ReadinessResponse)(nil),   // 5: echopb.ReadinessResponse
	(*VersionResponse)(nil),     // 6: echopb.VersionResponse
	(*GetEchoListResponse)(nil), // 7: echopb.GetEchoListResponse
}
var file_echo_service_proto_depIdxs = []int32{
	0, // 0: echopb.HealthService.Liveness:input_type -> echopb.LivenessRequest
	1, // 1: echopb.HealthService.Readiness:input_type -> echopb.ReadinessRequest
	2, // 2: echopb.HealthService.Version:input_type -> echopb.VersionRequest
	3, // 3: echopb.EchoService.GetEcho:input_type -> echopb.GetEchoListRequest
	4, // 4: echopb.HealthService.Liveness:output_type -> echopb.LivenessResponse
	5, // 5: echopb.HealthService.Readiness:output_type -> echopb.ReadinessResponse
	6, // 6: echopb.HealthService.Version:output_type -> echopb.VersionResponse
	7, // 7: echopb.EchoService.GetEcho:output_type -> echopb.GetEchoListResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_echo_service_proto_init() }
func file_echo_service_proto_init() {
	if File_echo_service_proto != nil {
		return
	}
	file_echo_proto_init()
	file_echo_health_proto_init()
	file_echo_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_echo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_echo_service_proto_goTypes,
		DependencyIndexes: file_echo_service_proto_depIdxs,
	}.Build()
	File_echo_service_proto = out.File
	file_echo_service_proto_rawDesc = nil
	file_echo_service_proto_goTypes = nil
	file_echo_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthServiceClient interface {
	// returns a error if service doesn`t live.
	// The kubelet uses liveness probes to know when to restart a Container.
	Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error)
	// returns a error if service doesn`t ready.
	// Service doesn`t ready by default.
	// The kubelet uses readiness probes to know when a Container is ready to start accepting traffic.
	Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error)
	// returns buid time, last commit and version app
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error) {
	out := new(LivenessResponse)
	err := c.cc.Invoke(ctx, "/echopb.HealthService/Liveness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error) {
	out := new(ReadinessResponse)
	err := c.cc.Invoke(ctx, "/echopb.HealthService/Readiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/echopb.HealthService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
type HealthServiceServer interface {
	// returns a error if service doesn`t live.
	// The kubelet uses liveness probes to know when to restart a Container.
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)
	// returns a error if service doesn`t ready.
	// Service doesn`t ready by default.
	// The kubelet uses readiness probes to know when a Container is ready to start accepting traffic.
	Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error)
	// returns buid time, last commit and version app
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedHealthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (*UnimplementedHealthServiceServer) Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Liveness not implemented")
}
func (*UnimplementedHealthServiceServer) Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Readiness not implemented")
}
func (*UnimplementedHealthServiceServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}

func RegisterHealthServiceServer(s *grpc.Server, srv HealthServiceServer) {
	s.RegisterService(&_HealthService_serviceDesc, srv)
}

func _HealthService_Liveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Liveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echopb.HealthService/Liveness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Liveness(ctx, req.(*LivenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_Readiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Readiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echopb.HealthService/Readiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Readiness(ctx, req.(*ReadinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echopb.HealthService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echopb.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Liveness",
			Handler:    _HealthService_Liveness_Handler,
		},
		{
			MethodName: "Readiness",
			Handler:    _HealthService_Readiness_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _HealthService_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo-service.proto",
}

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	// returns buid time, last commit and version app
	GetEcho(ctx context.Context, in *GetEchoListRequest, opts ...grpc.CallOption) (*GetEchoListResponse, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) GetEcho(ctx context.Context, in *GetEchoListRequest, opts ...grpc.CallOption) (*GetEchoListResponse, error) {
	out := new(GetEchoListResponse)
	err := c.cc.Invoke(ctx, "/echopb.EchoService/GetEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	// returns buid time, last commit and version app
	GetEcho(context.Context, *GetEchoListRequest) (*GetEchoListResponse, error)
}

// UnimplementedEchoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (*UnimplementedEchoServiceServer) GetEcho(context.Context, *GetEchoListRequest) (*GetEchoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEcho not implemented")
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_GetEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEchoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).GetEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echopb.EchoService/GetEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).GetEcho(ctx, req.(*GetEchoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echopb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEcho",
			Handler:    _EchoService_GetEcho_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo-service.proto",
}