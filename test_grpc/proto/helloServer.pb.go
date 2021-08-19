// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: proto/helloServer.proto

package my_grpc_proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_helloServer_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReplay) Reset() {
	*x = HelloReplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReplay) ProtoMessage() {}

func (x *HelloReplay) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReplay.ProtoReflect.Descriptor instead.
func (*HelloReplay) Descriptor() ([]byte, []int) {
	return file_proto_helloServer_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReplay) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HelloMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloMessage) Reset() {
	*x = HelloMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloServer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloMessage) ProtoMessage() {}

func (x *HelloMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloServer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloMessage.ProtoReflect.Descriptor instead.
func (*HelloMessage) Descriptor() ([]byte, []int) {
	return file_proto_helloServer_proto_rawDescGZIP(), []int{2}
}

func (x *HelloMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_proto_helloServer_proto protoreflect.FileDescriptor

var file_proto_helloServer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x79, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0b,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x9f, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x1b, 0x2e, 0x6d, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x6d, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x73, 0x67, 0x12, 0x1b, 0x2e,
	0x6d, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x79, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_helloServer_proto_rawDescOnce sync.Once
	file_proto_helloServer_proto_rawDescData = file_proto_helloServer_proto_rawDesc
)

func file_proto_helloServer_proto_rawDescGZIP() []byte {
	file_proto_helloServer_proto_rawDescOnce.Do(func() {
		file_proto_helloServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_helloServer_proto_rawDescData)
	})
	return file_proto_helloServer_proto_rawDescData
}

var file_proto_helloServer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_helloServer_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: my_grpc_proto.HelloRequest
	(*HelloReplay)(nil),  // 1: my_grpc_proto.HelloReplay
	(*HelloMessage)(nil), // 2: my_grpc_proto.HelloMessage
}
var file_proto_helloServer_proto_depIdxs = []int32{
	0, // 0: my_grpc_proto.HelloServer.SayHello:input_type -> my_grpc_proto.HelloRequest
	0, // 1: my_grpc_proto.HelloServer.GetHelloMsg:input_type -> my_grpc_proto.HelloRequest
	1, // 2: my_grpc_proto.HelloServer.SayHello:output_type -> my_grpc_proto.HelloReplay
	2, // 3: my_grpc_proto.HelloServer.GetHelloMsg:output_type -> my_grpc_proto.HelloMessage
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_helloServer_proto_init() }
func file_proto_helloServer_proto_init() {
	if File_proto_helloServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_helloServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_proto_helloServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReplay); i {
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
		file_proto_helloServer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloMessage); i {
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
			RawDescriptor: file_proto_helloServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_helloServer_proto_goTypes,
		DependencyIndexes: file_proto_helloServer_proto_depIdxs,
		MessageInfos:      file_proto_helloServer_proto_msgTypes,
	}.Build()
	File_proto_helloServer_proto = out.File
	file_proto_helloServer_proto_rawDesc = nil
	file_proto_helloServer_proto_goTypes = nil
	file_proto_helloServer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloServerClient is the client API for HelloServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServerClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReplay, error)
	GetHelloMsg(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloMessage, error)
}

type helloServerClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServerClient(cc grpc.ClientConnInterface) HelloServerClient {
	return &helloServerClient{cc}
}

func (c *helloServerClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReplay, error) {
	out := new(HelloReplay)
	err := c.cc.Invoke(ctx, "/my_grpc_proto.HelloServer/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServerClient) GetHelloMsg(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloMessage, error) {
	out := new(HelloMessage)
	err := c.cc.Invoke(ctx, "/my_grpc_proto.HelloServer/GetHelloMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServerServer is the server API for HelloServer service.
type HelloServerServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReplay, error)
	GetHelloMsg(context.Context, *HelloRequest) (*HelloMessage, error)
}

// UnimplementedHelloServerServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServerServer struct {
}

func (*UnimplementedHelloServerServer) SayHello(context.Context, *HelloRequest) (*HelloReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloServerServer) GetHelloMsg(context.Context, *HelloRequest) (*HelloMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHelloMsg not implemented")
}

func RegisterHelloServerServer(s *grpc.Server, srv HelloServerServer) {
	s.RegisterService(&_HelloServer_serviceDesc, srv)
}

func _HelloServer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/my_grpc_proto.HelloServer/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloServer_GetHelloMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).GetHelloMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/my_grpc_proto.HelloServer/GetHelloMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).GetHelloMsg(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "my_grpc_proto.HelloServer",
	HandlerType: (*HelloServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloServer_SayHello_Handler,
		},
		{
			MethodName: "GetHelloMsg",
			Handler:    _HelloServer_GetHelloMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/helloServer.proto",
}
