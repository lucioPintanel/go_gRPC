// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.1
// source: proto/service.proto

package pb

import (
	context "context"
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

type MeuRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mensagem      string                 `protobuf:"bytes,1,opt,name=mensagem,proto3" json:"mensagem,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MeuRequest) Reset() {
	*x = MeuRequest{}
	mi := &file_proto_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MeuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeuRequest) ProtoMessage() {}

func (x *MeuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeuRequest.ProtoReflect.Descriptor instead.
func (*MeuRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *MeuRequest) GetMensagem() string {
	if x != nil {
		return x.Mensagem
	}
	return ""
}

type MeuResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Resposta      string                 `protobuf:"bytes,1,opt,name=resposta,proto3" json:"resposta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MeuResponse) Reset() {
	*x = MeuResponse{}
	mi := &file_proto_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MeuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeuResponse) ProtoMessage() {}

func (x *MeuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeuResponse.ProtoReflect.Descriptor instead.
func (*MeuResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *MeuResponse) GetResposta() string {
	if x != nil {
		return x.Resposta
	}
	return ""
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x69, 0x5f, 0x73,
	0x72, 0x76, 0x22, 0x28, 0x0a, 0x0a, 0x4d, 0x65, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x67, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x67, 0x65, 0x6d, 0x22, 0x29, 0x0a, 0x0b,
	0x4d, 0x65, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x32, 0x4c, 0x0a, 0x0a, 0x4d, 0x65, 0x75, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x6f, 0x12, 0x3e, 0x0a, 0x09, 0x4d, 0x65, 0x75, 0x4d, 0x65, 0x74, 0x6f,
	0x64, 0x6f, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x6c, 0x69, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x4d, 0x65, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x6c, 0x69, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x4d, 0x65, 0x75, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_service_proto_goTypes = []any{
	(*MeuRequest)(nil),  // 0: grpccli_srv.MeuRequest
	(*MeuResponse)(nil), // 1: grpccli_srv.MeuResponse
}
var file_proto_service_proto_depIdxs = []int32{
	0, // 0: grpccli_srv.MeuServico.MeuMetodo:input_type -> grpccli_srv.MeuRequest
	1, // 1: grpccli_srv.MeuServico.MeuMetodo:output_type -> grpccli_srv.MeuResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MeuServicoClient is the client API for MeuServico service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeuServicoClient interface {
	MeuMetodo(ctx context.Context, in *MeuRequest, opts ...grpc.CallOption) (*MeuResponse, error)
}

type meuServicoClient struct {
	cc grpc.ClientConnInterface
}

func NewMeuServicoClient(cc grpc.ClientConnInterface) MeuServicoClient {
	return &meuServicoClient{cc}
}

func (c *meuServicoClient) MeuMetodo(ctx context.Context, in *MeuRequest, opts ...grpc.CallOption) (*MeuResponse, error) {
	out := new(MeuResponse)
	err := c.cc.Invoke(ctx, "/grpccli_srv.MeuServico/MeuMetodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeuServicoServer is the server API for MeuServico service.
type MeuServicoServer interface {
	MeuMetodo(context.Context, *MeuRequest) (*MeuResponse, error)
}

// UnimplementedMeuServicoServer can be embedded to have forward compatible implementations.
type UnimplementedMeuServicoServer struct {
}

func (*UnimplementedMeuServicoServer) MeuMetodo(context.Context, *MeuRequest) (*MeuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MeuMetodo not implemented")
}

func RegisterMeuServicoServer(s *grpc.Server, srv MeuServicoServer) {
	s.RegisterService(&_MeuServico_serviceDesc, srv)
}

func _MeuServico_MeuMetodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeuServicoServer).MeuMetodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpccli_srv.MeuServico/MeuMetodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeuServicoServer).MeuMetodo(ctx, req.(*MeuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MeuServico_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpccli_srv.MeuServico",
	HandlerType: (*MeuServicoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MeuMetodo",
			Handler:    _MeuServico_MeuMetodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
