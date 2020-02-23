// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment_info.proto

package paymentinfopb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type PaymentInfo struct {
	UUID                 string               `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	UserUUID             string               `protobuf:"bytes,2,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	CardNumber           string               `protobuf:"bytes,3,opt,name=cardNumber,proto3" json:"cardNumber,omitempty"`
	CommentUUID          string               `protobuf:"bytes,4,opt,name=commentUUID,proto3" json:"commentUUID,omitempty"`
	ProductUUID          string               `protobuf:"bytes,5,opt,name=productUUID,proto3" json:"productUUID,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PaymentInfo) Reset()         { *m = PaymentInfo{} }
func (m *PaymentInfo) String() string { return proto.CompactTextString(m) }
func (*PaymentInfo) ProtoMessage()    {}
func (*PaymentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_60ad0eadf34cbae7, []int{0}
}

func (m *PaymentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentInfo.Unmarshal(m, b)
}
func (m *PaymentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentInfo.Marshal(b, m, deterministic)
}
func (m *PaymentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentInfo.Merge(m, src)
}
func (m *PaymentInfo) XXX_Size() int {
	return xxx_messageInfo_PaymentInfo.Size(m)
}
func (m *PaymentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentInfo proto.InternalMessageInfo

func (m *PaymentInfo) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *PaymentInfo) GetUserUUID() string {
	if m != nil {
		return m.UserUUID
	}
	return ""
}

func (m *PaymentInfo) GetCardNumber() string {
	if m != nil {
		return m.CardNumber
	}
	return ""
}

func (m *PaymentInfo) GetCommentUUID() string {
	if m != nil {
		return m.CommentUUID
	}
	return ""
}

func (m *PaymentInfo) GetProductUUID() string {
	if m != nil {
		return m.ProductUUID
	}
	return ""
}

func (m *PaymentInfo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *PaymentInfo) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *PaymentInfo) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type GetRequest struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60ad0eadf34cbae7, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type GetResponse struct {
	PaymentInfo          *PaymentInfo `protobuf:"bytes,1,opt,name=paymentInfo,proto3" json:"paymentInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60ad0eadf34cbae7, []int{2}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetPaymentInfo() *PaymentInfo {
	if m != nil {
		return m.PaymentInfo
	}
	return nil
}

type SetRequest struct {
	PaymentInfo          *PaymentInfo `protobuf:"bytes,1,opt,name=paymentInfo,proto3" json:"paymentInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60ad0eadf34cbae7, []int{3}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetPaymentInfo() *PaymentInfo {
	if m != nil {
		return m.PaymentInfo
	}
	return nil
}

type SetResponse struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60ad0eadf34cbae7, []int{4}
}

func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (m *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(m, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type UpdateRequest struct {
	PaymentInfo          *PaymentInfo `protobuf:"bytes,1,opt,name=paymentInfo,proto3" json:"paymentInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60ad0eadf34cbae7, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetPaymentInfo() *PaymentInfo {
	if m != nil {
		return m.PaymentInfo
	}
	return nil
}

type DeleteRequest struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60ad0eadf34cbae7, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func init() {
	proto.RegisterType((*PaymentInfo)(nil), "paymentinfopb.PaymentInfo")
	proto.RegisterType((*GetRequest)(nil), "paymentinfopb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "paymentinfopb.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "paymentinfopb.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "paymentinfopb.SetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "paymentinfopb.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "paymentinfopb.DeleteRequest")
}

func init() { proto.RegisterFile("payment_info.proto", fileDescriptor_60ad0eadf34cbae7) }

var fileDescriptor_60ad0eadf34cbae7 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0x2f, 0x7f, 0x2e, 0x17, 0x4e, 0xc3, 0x5d, 0xcc, 0xe2, 0x86, 0x5b, 0x8d, 0xd6, 0xba,
	0x61, 0x55, 0x12, 0xdc, 0xb8, 0x30, 0x46, 0x12, 0x0c, 0x41, 0xa3, 0x21, 0xad, 0xac, 0x4d, 0xff,
	0x1c, 0x08, 0x09, 0xed, 0x8c, 0xed, 0x74, 0xc1, 0x7b, 0xf8, 0x5e, 0xbe, 0x92, 0x99, 0x19, 0xa0,
	0x2d, 0x14, 0x59, 0xe8, 0xae, 0xfd, 0xce, 0xef, 0x3b, 0x67, 0xce, 0x37, 0x2d, 0x10, 0xe6, 0xae,
	0x42, 0x8c, 0xf8, 0xeb, 0x22, 0x9a, 0x51, 0x8b, 0xc5, 0x94, 0x53, 0xd2, 0x5e, 0x6b, 0x42, 0x62,
	0x9e, 0x7e, 0x32, 0xa7, 0x74, 0xbe, 0xc4, 0x9e, 0x2c, 0x7a, 0xe9, 0xac, 0x87, 0x21, 0xe3, 0x2b,
	0xc5, 0xea, 0xe7, 0xbb, 0x45, 0xbe, 0x08, 0x31, 0xe1, 0x6e, 0xc8, 0x14, 0x60, 0x7e, 0x54, 0x41,
	0x9b, 0xa8, 0x7e, 0xe3, 0x68, 0x46, 0x09, 0x81, 0xfa, 0x74, 0x3a, 0x1e, 0x76, 0x2a, 0x46, 0xa5,
	0xdb, 0xb2, 0xe5, 0x33, 0xd1, 0xa1, 0x99, 0x26, 0x18, 0x4b, 0xbd, 0x2a, 0xf5, 0xed, 0x3b, 0x39,
	0x03, 0xf0, 0xdd, 0x38, 0x78, 0x4e, 0x43, 0x0f, 0xe3, 0x4e, 0x4d, 0x56, 0x73, 0x0a, 0x31, 0x40,
	0xf3, 0x69, 0x28, 0xda, 0x4b, 0x7b, 0x5d, 0x02, 0x79, 0x49, 0x10, 0x2c, 0xa6, 0x41, 0xea, 0x2b,
	0xe2, 0xb7, 0x22, 0x72, 0x12, 0xb9, 0x86, 0x96, 0x1f, 0xa3, 0xcb, 0x31, 0x18, 0xf0, 0x4e, 0xc3,
	0xa8, 0x74, 0xb5, 0xbe, 0x6e, 0xa9, 0xc5, 0xac, 0xcd, 0x62, 0xd6, 0xcb, 0x66, 0x31, 0x3b, 0x83,
	0x85, 0x33, 0x65, 0xc1, 0xda, 0xf9, 0xe7, 0xb8, 0x73, 0x0b, 0x0b, 0x67, 0x80, 0x4b, 0x54, 0xce,
	0xe6, 0x71, 0xe7, 0x16, 0x36, 0x0d, 0x80, 0x11, 0x72, 0x1b, 0xdf, 0x52, 0x4c, 0x78, 0x59, 0x9e,
	0xe6, 0x23, 0x68, 0x92, 0x48, 0x18, 0x8d, 0x12, 0x24, 0x37, 0xa0, 0xb1, 0xec, 0x06, 0x24, 0x29,
	0x86, 0x15, 0x6e, 0xd9, 0xca, 0xdd, 0x91, 0x9d, 0xc7, 0xcd, 0x07, 0x00, 0x27, 0x1b, 0xf7, 0xbd,
	0x5e, 0x17, 0xa0, 0x39, 0xb9, 0x83, 0x95, 0x9d, 0xfd, 0x09, 0xda, 0x53, 0x19, 0xd2, 0xcf, 0x4c,
	0xbc, 0x84, 0xf6, 0x50, 0x26, 0xf7, 0x45, 0x5e, 0xfd, 0xf7, 0x2a, 0xfc, 0xcd, 0x75, 0x18, 0x4c,
	0xc6, 0xe4, 0x16, 0x6a, 0x23, 0xe4, 0xe4, 0xff, 0xce, 0x9c, 0x2c, 0x78, 0x5d, 0x2f, 0x2b, 0xa9,
	0xc5, 0xcc, 0x5f, 0xc2, 0xef, 0x94, 0xf8, 0x9d, 0xc3, 0x7e, 0xa7, 0xe0, 0xbf, 0x83, 0x86, 0x8a,
	0x81, 0x9c, 0xee, 0x70, 0x85, 0x74, 0xf4, 0x7f, 0x7b, 0xdf, 0xcc, 0xbd, 0xf8, 0x3b, 0x55, 0x07,
	0xb5, 0xf9, 0x5e, 0x87, 0x42, 0x20, 0x87, 0x3b, 0x78, 0x0d, 0xa9, 0x5c, 0x7d, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x1d, 0xec, 0x6c, 0x1c, 0x24, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaymentInfoAPIClient is the client API for PaymentInfoAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentInfoAPIClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type paymentInfoAPIClient struct {
	cc *grpc.ClientConn
}

func NewPaymentInfoAPIClient(cc *grpc.ClientConn) PaymentInfoAPIClient {
	return &paymentInfoAPIClient{cc}
}

func (c *paymentInfoAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/paymentinfopb.PaymentInfoAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentInfoAPIClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/paymentinfopb.PaymentInfoAPI/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentInfoAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/paymentinfopb.PaymentInfoAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentInfoAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/paymentinfopb.PaymentInfoAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentInfoAPIServer is the server API for PaymentInfoAPI service.
type PaymentInfoAPIServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Update(context.Context, *UpdateRequest) (*empty.Empty, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
}

// UnimplementedPaymentInfoAPIServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentInfoAPIServer struct {
}

func (*UnimplementedPaymentInfoAPIServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedPaymentInfoAPIServer) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedPaymentInfoAPIServer) Update(ctx context.Context, req *UpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedPaymentInfoAPIServer) Delete(ctx context.Context, req *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterPaymentInfoAPIServer(s *grpc.Server, srv PaymentInfoAPIServer) {
	s.RegisterService(&_PaymentInfoAPI_serviceDesc, srv)
}

func _PaymentInfoAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentInfoAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentinfopb.PaymentInfoAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentInfoAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentInfoAPI_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentInfoAPIServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentinfopb.PaymentInfoAPI/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentInfoAPIServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentInfoAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentInfoAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentinfopb.PaymentInfoAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentInfoAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentInfoAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentInfoAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentinfopb.PaymentInfoAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentInfoAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentInfoAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "paymentinfopb.PaymentInfoAPI",
	HandlerType: (*PaymentInfoAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PaymentInfoAPI_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _PaymentInfoAPI_Set_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PaymentInfoAPI_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PaymentInfoAPI_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment_info.proto",
}
