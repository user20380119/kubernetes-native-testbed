// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cart.proto

package cartpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Cart struct {
	UserUUID             string           `protobuf:"bytes,1,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	CartProducts         map[string]int32 `protobuf:"bytes,2,rep,name=cartProducts,proto3" json:"cartProducts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{0}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetUserUUID() string {
	if m != nil {
		return m.UserUUID
	}
	return ""
}

func (m *Cart) GetCartProducts() map[string]int32 {
	if m != nil {
		return m.CartProducts
	}
	return nil
}

type ShowRequest struct {
	UserUUID             string   `protobuf:"bytes,1,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowRequest) Reset()         { *m = ShowRequest{} }
func (m *ShowRequest) String() string { return proto.CompactTextString(m) }
func (*ShowRequest) ProtoMessage()    {}
func (*ShowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{1}
}

func (m *ShowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowRequest.Unmarshal(m, b)
}
func (m *ShowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowRequest.Marshal(b, m, deterministic)
}
func (m *ShowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowRequest.Merge(m, src)
}
func (m *ShowRequest) XXX_Size() int {
	return xxx_messageInfo_ShowRequest.Size(m)
}
func (m *ShowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShowRequest proto.InternalMessageInfo

func (m *ShowRequest) GetUserUUID() string {
	if m != nil {
		return m.UserUUID
	}
	return ""
}

type ShowResponse struct {
	Cart                 *Cart    `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowResponse) Reset()         { *m = ShowResponse{} }
func (m *ShowResponse) String() string { return proto.CompactTextString(m) }
func (*ShowResponse) ProtoMessage()    {}
func (*ShowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{2}
}

func (m *ShowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowResponse.Unmarshal(m, b)
}
func (m *ShowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowResponse.Marshal(b, m, deterministic)
}
func (m *ShowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowResponse.Merge(m, src)
}
func (m *ShowResponse) XXX_Size() int {
	return xxx_messageInfo_ShowResponse.Size(m)
}
func (m *ShowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowResponse proto.InternalMessageInfo

func (m *ShowResponse) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

type AddRequest struct {
	Cart                 *Cart    `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{3}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

type RemoveRequest struct {
	Cart                 *Cart    `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{4}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

func init() {
	proto.RegisterType((*Cart)(nil), "cartpb.Cart")
	proto.RegisterMapType((map[string]int32)(nil), "cartpb.Cart.CartProductsEntry")
	proto.RegisterType((*ShowRequest)(nil), "cartpb.ShowRequest")
	proto.RegisterType((*ShowResponse)(nil), "cartpb.ShowResponse")
	proto.RegisterType((*AddRequest)(nil), "cartpb.AddRequest")
	proto.RegisterType((*RemoveRequest)(nil), "cartpb.RemoveRequest")
}

func init() { proto.RegisterFile("cart.proto", fileDescriptor_bf731a5c8f9a516f) }

var fileDescriptor_bf731a5c8f9a516f = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4e, 0x02, 0x31,
	0x10, 0xc7, 0x29, 0x5f, 0xea, 0x80, 0x89, 0x8e, 0x68, 0xc8, 0x9a, 0x98, 0x4d, 0x4f, 0x78, 0x29,
	0x0a, 0x17, 0xf5, 0x62, 0x56, 0xe5, 0xc0, 0x8d, 0xd4, 0xf0, 0x00, 0xbb, 0x6c, 0xc5, 0x44, 0xa0,
	0x6b, 0xb7, 0x8b, 0xd9, 0x27, 0xf2, 0xe6, 0x33, 0x9a, 0x6e, 0xad, 0xb2, 0x31, 0x7e, 0x5c, 0x9a,
	0xce, 0xcc, 0xff, 0x3f, 0xf3, 0x9b, 0x0c, 0xc0, 0x2c, 0x54, 0x9a, 0x25, 0x4a, 0x6a, 0x89, 0x4d,
	0xf3, 0x4f, 0x22, 0xef, 0x78, 0x2e, 0xe5, 0x7c, 0x21, 0xfa, 0x45, 0x36, 0xca, 0x1e, 0xfa, 0x62,
	0x99, 0xe8, 0xdc, 0x8a, 0xe8, 0x2b, 0x81, 0xfa, 0x6d, 0xa8, 0x34, 0x7a, 0xb0, 0x9d, 0xa5, 0x42,
	0x4d, 0xa7, 0xe3, 0xbb, 0x2e, 0xf1, 0x49, 0x6f, 0x87, 0x7f, 0xc6, 0x78, 0x03, 0x6d, 0xd3, 0x6b,
	0xa2, 0x64, 0x9c, 0xcd, 0x74, 0xda, 0xad, 0xfa, 0xb5, 0x5e, 0x6b, 0x70, 0xc2, 0xec, 0x00, 0x66,
	0xfc, 0xc5, 0xe3, 0x04, 0xa3, 0x95, 0x56, 0x39, 0x2f, 0x79, 0xbc, 0x6b, 0xd8, 0xff, 0x26, 0xc1,
	0x3d, 0xa8, 0x3d, 0x89, 0xfc, 0x63, 0x9e, 0xf9, 0x62, 0x07, 0x1a, 0xeb, 0x70, 0x91, 0x89, 0x6e,
	0xd5, 0x27, 0xbd, 0x06, 0xb7, 0xc1, 0x55, 0xf5, 0x82, 0xd0, 0x53, 0x68, 0xdd, 0x3f, 0xca, 0x17,
	0x2e, 0x9e, 0x33, 0x91, 0xfe, 0xca, 0x4b, 0xcf, 0xa0, 0x6d, 0xa5, 0x69, 0x22, 0x57, 0xa9, 0x40,
	0x1f, 0xea, 0x86, 0xa5, 0xd0, 0xb5, 0x06, 0xed, 0x4d, 0x6e, 0x5e, 0x54, 0x28, 0x03, 0x08, 0xe2,
	0xd8, 0xf5, 0xfe, 0x5b, 0x7f, 0x0e, 0xbb, 0x5c, 0x2c, 0xe5, 0x5a, 0xfc, 0xdb, 0x32, 0x78, 0x23,
	0xb0, 0x65, 0xc2, 0x60, 0x32, 0xc6, 0x21, 0xd4, 0x0d, 0x20, 0x1e, 0x38, 0xdd, 0xc6, 0x66, 0x5e,
	0xa7, 0x9c, 0xb4, 0x3b, 0xd0, 0x0a, 0x0e, 0xa1, 0x16, 0xc4, 0x31, 0xa2, 0x2b, 0x7f, 0x01, 0x7b,
	0x47, 0xcc, 0xde, 0x98, 0xb9, 0x1b, 0xb3, 0x91, 0xb9, 0x31, 0xad, 0xe0, 0x25, 0x34, 0x2d, 0x28,
	0x1e, 0x3a, 0x5f, 0x09, 0xfc, 0x67, 0x6b, 0xd4, 0x2c, 0x32, 0xc3, 0xf7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb8, 0xb5, 0xb8, 0x67, 0x54, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CartAPIClient is the client API for CartAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartAPIClient interface {
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type cartAPIClient struct {
	cc *grpc.ClientConn
}

func NewCartAPIClient(cc *grpc.ClientConn) CartAPIClient {
	return &cartAPIClient{cc}
}

func (c *cartAPIClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, "/cartpb.CartAPI/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartAPIClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cartpb.CartAPI/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartAPIClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cartpb.CartAPI/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartAPIServer is the server API for CartAPI service.
type CartAPIServer interface {
	Show(context.Context, *ShowRequest) (*ShowResponse, error)
	Add(context.Context, *AddRequest) (*empty.Empty, error)
	Remove(context.Context, *RemoveRequest) (*empty.Empty, error)
}

// UnimplementedCartAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCartAPIServer struct {
}

func (*UnimplementedCartAPIServer) Show(ctx context.Context, req *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (*UnimplementedCartAPIServer) Add(ctx context.Context, req *AddRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCartAPIServer) Remove(ctx context.Context, req *RemoveRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func RegisterCartAPIServer(s *grpc.Server, srv CartAPIServer) {
	s.RegisterService(&_CartAPI_serviceDesc, srv)
}

func _CartAPI_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartAPIServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cartpb.CartAPI/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartAPIServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartAPI_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartAPIServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cartpb.CartAPI/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartAPIServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartAPI_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartAPIServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cartpb.CartAPI/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartAPIServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CartAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cartpb.CartAPI",
	HandlerType: (*CartAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Show",
			Handler:    _CartAPI_Show_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _CartAPI_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _CartAPI_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart.proto",
}
