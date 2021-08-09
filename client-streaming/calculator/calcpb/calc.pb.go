// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: calculator/calcpb/calc.proto

package calcpb

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

type Calculate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Calculate) Reset() {
	*x = Calculate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Calculate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Calculate) ProtoMessage() {}

func (x *Calculate) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Calculate.ProtoReflect.Descriptor instead.
func (*Calculate) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{0}
}

func (x *Calculate) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type CalculateLongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calculate *Calculate `protobuf:"bytes,1,opt,name=calculate,proto3" json:"calculate,omitempty"`
}

func (x *CalculateLongRequest) Reset() {
	*x = CalculateLongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateLongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateLongRequest) ProtoMessage() {}

func (x *CalculateLongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateLongRequest.ProtoReflect.Descriptor instead.
func (*CalculateLongRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateLongRequest) GetCalculate() *Calculate {
	if x != nil {
		return x.Calculate
	}
	return nil
}

type CalculateLongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalculateLongResponse) Reset() {
	*x = CalculateLongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateLongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateLongResponse) ProtoMessage() {}

func (x *CalculateLongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateLongResponse.ProtoReflect.Descriptor instead.
func (*CalculateLongResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{2}
}

func (x *CalculateLongResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_calcpb_calc_proto protoreflect.FileDescriptor

var file_calculator_calcpb_calc_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x63, 0x61, 0x6c, 0x63, 0x22, 0x23, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x14, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2d, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x09, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x22, 0x2f, 0x0a, 0x15, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x60, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_calculator_calcpb_calc_proto_rawDescOnce sync.Once
	file_calculator_calcpb_calc_proto_rawDescData = file_calculator_calcpb_calc_proto_rawDesc
)

func file_calculator_calcpb_calc_proto_rawDescGZIP() []byte {
	file_calculator_calcpb_calc_proto_rawDescOnce.Do(func() {
		file_calculator_calcpb_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calcpb_calc_proto_rawDescData)
	})
	return file_calculator_calcpb_calc_proto_rawDescData
}

var file_calculator_calcpb_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_calculator_calcpb_calc_proto_goTypes = []interface{}{
	(*Calculate)(nil),             // 0: calc.Calculate
	(*CalculateLongRequest)(nil),  // 1: calc.CalculateLongRequest
	(*CalculateLongResponse)(nil), // 2: calc.CalculateLongResponse
}
var file_calculator_calcpb_calc_proto_depIdxs = []int32{
	0, // 0: calc.CalculateLongRequest.calculate:type_name -> calc.Calculate
	1, // 1: calc.CalculateService.LongCalculate:input_type -> calc.CalculateLongRequest
	2, // 2: calc.CalculateService.LongCalculate:output_type -> calc.CalculateLongResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_calculator_calcpb_calc_proto_init() }
func file_calculator_calcpb_calc_proto_init() {
	if File_calculator_calcpb_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calcpb_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Calculate); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateLongRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateLongResponse); i {
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
			RawDescriptor: file_calculator_calcpb_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calcpb_calc_proto_goTypes,
		DependencyIndexes: file_calculator_calcpb_calc_proto_depIdxs,
		MessageInfos:      file_calculator_calcpb_calc_proto_msgTypes,
	}.Build()
	File_calculator_calcpb_calc_proto = out.File
	file_calculator_calcpb_calc_proto_rawDesc = nil
	file_calculator_calcpb_calc_proto_goTypes = nil
	file_calculator_calcpb_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateServiceClient interface {
	LongCalculate(ctx context.Context, opts ...grpc.CallOption) (CalculateService_LongCalculateClient, error)
}

type calculateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateServiceClient(cc grpc.ClientConnInterface) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) LongCalculate(ctx context.Context, opts ...grpc.CallOption) (CalculateService_LongCalculateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[0], "/calc.CalculateService/LongCalculate", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceLongCalculateClient{stream}
	return x, nil
}

type CalculateService_LongCalculateClient interface {
	Send(*CalculateLongRequest) error
	CloseAndRecv() (*CalculateLongResponse, error)
	grpc.ClientStream
}

type calculateServiceLongCalculateClient struct {
	grpc.ClientStream
}

func (x *calculateServiceLongCalculateClient) Send(m *CalculateLongRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceLongCalculateClient) CloseAndRecv() (*CalculateLongResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalculateLongResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculateServiceServer is the server API for CalculateService service.
type CalculateServiceServer interface {
	LongCalculate(CalculateService_LongCalculateServer) error
}

// UnimplementedCalculateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (*UnimplementedCalculateServiceServer) LongCalculate(CalculateService_LongCalculateServer) error {
	return status.Errorf(codes.Unimplemented, "method LongCalculate not implemented")
}

func RegisterCalculateServiceServer(s *grpc.Server, srv CalculateServiceServer) {
	s.RegisterService(&_CalculateService_serviceDesc, srv)
}

func _CalculateService_LongCalculate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).LongCalculate(&calculateServiceLongCalculateServer{stream})
}

type CalculateService_LongCalculateServer interface {
	SendAndClose(*CalculateLongResponse) error
	Recv() (*CalculateLongRequest, error)
	grpc.ServerStream
}

type calculateServiceLongCalculateServer struct {
	grpc.ServerStream
}

func (x *calculateServiceLongCalculateServer) SendAndClose(m *CalculateLongResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceLongCalculateServer) Recv() (*CalculateLongRequest, error) {
	m := new(CalculateLongRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LongCalculate",
			Handler:       _CalculateService_LongCalculate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}