// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Req struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Resp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Resp) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "rpc.Req")
	proto.RegisterType((*Resp)(nil), "rpc.Resp")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcf, 0x41, 0x4a, 0x43, 0x31,
	0x10, 0xc6, 0xf1, 0xa6, 0x7d, 0xaf, 0xd8, 0x11, 0x41, 0x66, 0xa5, 0x45, 0x50, 0xb3, 0x2a, 0x82,
	0x79, 0x45, 0x77, 0x9e, 0xc1, 0x55, 0x7a, 0x82, 0x98, 0x0c, 0x25, 0x50, 0x33, 0x69, 0x32, 0x8a,
	0x6b, 0xaf, 0xe0, 0x15, 0xbc, 0x91, 0x57, 0xf0, 0x20, 0x62, 0x14, 0xdd, 0xb4, 0xbb, 0x3f, 0x7c,
	0x3f, 0x06, 0x06, 0x8e, 0xe8, 0xc5, 0x3d, 0xe6, 0x0d, 0x99, 0x5c, 0x58, 0x18, 0x27, 0x25, 0xfb,
	0xf9, 0xd9, 0x9a, 0x79, 0xbd, 0xa1, 0xc1, 0xe5, 0x38, 0xb8, 0x94, 0x58, 0x9c, 0x44, 0x4e, 0xf5,
	0x87, 0xe8, 0x53, 0x98, 0x58, 0xda, 0x22, 0x42, 0x17, 0x9c, 0xb8, 0x13, 0x75, 0xa1, 0x16, 0x33,
	0xdb, 0x5a, 0x1b, 0xe8, 0x2c, 0xd5, 0xfc, 0xbd, 0x79, 0x0e, 0xd4, 0xb6, 0xde, 0xb6, 0xfe, 0xf3,
	0xe3, 0x7f, 0x7f, 0xf3, 0xae, 0xe0, 0x50, 0xa8, 0xca, 0x8a, 0xca, 0x73, 0xf4, 0x84, 0xd7, 0xd0,
	0xad, 0x28, 0x05, 0x3c, 0x30, 0x25, 0x7b, 0x63, 0x69, 0x3b, 0x9f, 0xfd, 0x56, 0xcd, 0xfa, 0xf8,
	0xf5, 0xe3, 0xf3, 0x6d, 0x0c, 0xba, 0x1f, 0x2a, 0xa5, 0x70, 0xa7, 0xae, 0xf0, 0x1c, 0xba, 0xfb,
	0x58, 0x65, 0x37, 0x1f, 0x2d, 0x15, 0x5e, 0xc2, 0xd4, 0x92, 0xe7, 0xb2, 0xe7, 0xe2, 0x68, 0xa1,
	0x50, 0x43, 0x6f, 0xf9, 0x49, 0x68, 0xaf, 0x58, 0xaa, 0x87, 0x69, 0x7b, 0xfc, 0xf6, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x94, 0xc3, 0x46, 0x78, 0x2c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	//发送信息
	Send(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	//服务端流式RPC
	List(ctx context.Context, in *Req, opts ...grpc.CallOption) (TestService_ListClient, error)
	//客户端流式RPC
	Record(ctx context.Context, opts ...grpc.CallOption) (TestService_RecordClient, error)
	//双向流式RPC
	Route(ctx context.Context, opts ...grpc.CallOption) (TestService_RouteClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Send(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/rpc.testService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) List(ctx context.Context, in *Req, opts ...grpc.CallOption) (TestService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/rpc.testService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_ListClient interface {
	Recv() (*Resp, error)
	grpc.ClientStream
}

type testServiceListClient struct {
	grpc.ClientStream
}

func (x *testServiceListClient) Recv() (*Resp, error) {
	m := new(Resp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) Record(ctx context.Context, opts ...grpc.CallOption) (TestService_RecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[1], "/rpc.testService/Record", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceRecordClient{stream}
	return x, nil
}

type TestService_RecordClient interface {
	Send(*Req) error
	CloseAndRecv() (*Resp, error)
	grpc.ClientStream
}

type testServiceRecordClient struct {
	grpc.ClientStream
}

func (x *testServiceRecordClient) Send(m *Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceRecordClient) CloseAndRecv() (*Resp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Resp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) Route(ctx context.Context, opts ...grpc.CallOption) (TestService_RouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[2], "/rpc.testService/Route", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceRouteClient{stream}
	return x, nil
}

type TestService_RouteClient interface {
	Send(*Req) error
	Recv() (*Resp, error)
	grpc.ClientStream
}

type testServiceRouteClient struct {
	grpc.ClientStream
}

func (x *testServiceRouteClient) Send(m *Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceRouteClient) Recv() (*Resp, error) {
	m := new(Resp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	//发送信息
	Send(context.Context, *Req) (*Resp, error)
	//服务端流式RPC
	List(*Req, TestService_ListServer) error
	//客户端流式RPC
	Record(TestService_RecordServer) error
	//双向流式RPC
	Route(TestService_RouteServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.testService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Send(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Req)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).List(m, &testServiceListServer{stream})
}

type TestService_ListServer interface {
	Send(*Resp) error
	grpc.ServerStream
}

type testServiceListServer struct {
	grpc.ServerStream
}

func (x *testServiceListServer) Send(m *Resp) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_Record_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).Record(&testServiceRecordServer{stream})
}

type TestService_RecordServer interface {
	SendAndClose(*Resp) error
	Recv() (*Req, error)
	grpc.ServerStream
}

type testServiceRecordServer struct {
	grpc.ServerStream
}

func (x *testServiceRecordServer) SendAndClose(m *Resp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceRecordServer) Recv() (*Req, error) {
	m := new(Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_Route_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).Route(&testServiceRouteServer{stream})
}

type TestService_RouteServer interface {
	Send(*Resp) error
	Recv() (*Req, error)
	grpc.ServerStream
}

type testServiceRouteServer struct {
	grpc.ServerStream
}

func (x *testServiceRouteServer) Send(m *Resp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceRouteServer) Recv() (*Req, error) {
	m := new(Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.testService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _TestService_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _TestService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Record",
			Handler:       _TestService_Record_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Route",
			Handler:       _TestService_Route_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "example.proto",
}