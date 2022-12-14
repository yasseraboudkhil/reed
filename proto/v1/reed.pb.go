// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reed.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type EncryptPayload struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Plaintext            string   `protobuf:"bytes,2,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	FieldName            string   `protobuf:"bytes,3,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptPayload) Reset()         { *m = EncryptPayload{} }
func (m *EncryptPayload) String() string { return proto.CompactTextString(m) }
func (*EncryptPayload) ProtoMessage()    {}
func (*EncryptPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e85daebbf806b99, []int{0}
}

func (m *EncryptPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptPayload.Unmarshal(m, b)
}
func (m *EncryptPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptPayload.Marshal(b, m, deterministic)
}
func (m *EncryptPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptPayload.Merge(m, src)
}
func (m *EncryptPayload) XXX_Size() int {
	return xxx_messageInfo_EncryptPayload.Size(m)
}
func (m *EncryptPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptPayload.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptPayload proto.InternalMessageInfo

func (m *EncryptPayload) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *EncryptPayload) GetPlaintext() string {
	if m != nil {
		return m.Plaintext
	}
	return ""
}

func (m *EncryptPayload) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

type DecryptPayload struct {
	UserId               string        `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CipherAndKey         *CipherAndKey `protobuf:"bytes,2,opt,name=cipherAndKey,proto3" json:"cipherAndKey,omitempty"`
	FieldName            string        `protobuf:"bytes,3,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DecryptPayload) Reset()         { *m = DecryptPayload{} }
func (m *DecryptPayload) String() string { return proto.CompactTextString(m) }
func (*DecryptPayload) ProtoMessage()    {}
func (*DecryptPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e85daebbf806b99, []int{1}
}

func (m *DecryptPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptPayload.Unmarshal(m, b)
}
func (m *DecryptPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptPayload.Marshal(b, m, deterministic)
}
func (m *DecryptPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptPayload.Merge(m, src)
}
func (m *DecryptPayload) XXX_Size() int {
	return xxx_messageInfo_DecryptPayload.Size(m)
}
func (m *DecryptPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptPayload.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptPayload proto.InternalMessageInfo

func (m *DecryptPayload) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DecryptPayload) GetCipherAndKey() *CipherAndKey {
	if m != nil {
		return m.CipherAndKey
	}
	return nil
}

func (m *DecryptPayload) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

type EncryptRequest struct {
	Api                  string          `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	EncryptPayload       *EncryptPayload `protobuf:"bytes,2,opt,name=encryptPayload,proto3" json:"encryptPayload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EncryptRequest) Reset()         { *m = EncryptRequest{} }
func (m *EncryptRequest) String() string { return proto.CompactTextString(m) }
func (*EncryptRequest) ProtoMessage()    {}
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e85daebbf806b99, []int{2}
}

func (m *EncryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptRequest.Unmarshal(m, b)
}
func (m *EncryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptRequest.Marshal(b, m, deterministic)
}
func (m *EncryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptRequest.Merge(m, src)
}
func (m *EncryptRequest) XXX_Size() int {
	return xxx_messageInfo_EncryptRequest.Size(m)
}
func (m *EncryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptRequest proto.InternalMessageInfo

func (m *EncryptRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *EncryptRequest) GetEncryptPayload() *EncryptPayload {
	if m != nil {
		return m.EncryptPayload
	}
	return nil
}

type CipherAndKey struct {
	Ciphertext           string   `protobuf:"bytes,1,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	EncryptedDataKey     string   `protobuf:"bytes,2,opt,name=encryptedDataKey,proto3" json:"encryptedDataKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipherAndKey) Reset()         { *m = CipherAndKey{} }
func (m *CipherAndKey) String() string { return proto.CompactTextString(m) }
func (*CipherAndKey) ProtoMessage()    {}
func (*CipherAndKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e85daebbf806b99, []int{3}
}

func (m *CipherAndKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipherAndKey.Unmarshal(m, b)
}
func (m *CipherAndKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipherAndKey.Marshal(b, m, deterministic)
}
func (m *CipherAndKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipherAndKey.Merge(m, src)
}
func (m *CipherAndKey) XXX_Size() int {
	return xxx_messageInfo_CipherAndKey.Size(m)
}
func (m *CipherAndKey) XXX_DiscardUnknown() {
	xxx_messageInfo_CipherAndKey.DiscardUnknown(m)
}

var xxx_messageInfo_CipherAndKey proto.InternalMessageInfo

func (m *CipherAndKey) GetCiphertext() string {
	if m != nil {
		return m.Ciphertext
	}
	return ""
}

func (m *CipherAndKey) GetEncryptedDataKey() string {
	if m != nil {
		return m.EncryptedDataKey
	}
	return ""
}

type EncryptResponse struct {
	Api                  string        `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	CipherAndKey         *CipherAndKey `protobuf:"bytes,2,opt,name=cipherAndKey,proto3" json:"cipherAndKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EncryptResponse) Reset()         { *m = EncryptResponse{} }
func (m *EncryptResponse) String() string { return proto.CompactTextString(m) }
func (*EncryptResponse) ProtoMessage()    {}
func (*EncryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e85daebbf806b99, []int{4}
}

func (m *EncryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptResponse.Unmarshal(m, b)
}
func (m *EncryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptResponse.Marshal(b, m, deterministic)
}
func (m *EncryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptResponse.Merge(m, src)
}
func (m *EncryptResponse) XXX_Size() int {
	return xxx_messageInfo_EncryptResponse.Size(m)
}
func (m *EncryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptResponse proto.InternalMessageInfo

func (m *EncryptResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *EncryptResponse) GetCipherAndKey() *CipherAndKey {
	if m != nil {
		return m.CipherAndKey
	}
	return nil
}

type DecryptRequest struct {
	Api                  string          `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	DecryptPayload       *DecryptPayload `protobuf:"bytes,2,opt,name=decryptPayload,proto3" json:"decryptPayload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DecryptRequest) Reset()         { *m = DecryptRequest{} }
func (m *DecryptRequest) String() string { return proto.CompactTextString(m) }
func (*DecryptRequest) ProtoMessage()    {}
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e85daebbf806b99, []int{5}
}

func (m *DecryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptRequest.Unmarshal(m, b)
}
func (m *DecryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptRequest.Marshal(b, m, deterministic)
}
func (m *DecryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptRequest.Merge(m, src)
}
func (m *DecryptRequest) XXX_Size() int {
	return xxx_messageInfo_DecryptRequest.Size(m)
}
func (m *DecryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptRequest proto.InternalMessageInfo

func (m *DecryptRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DecryptRequest) GetDecryptPayload() *DecryptPayload {
	if m != nil {
		return m.DecryptPayload
	}
	return nil
}

type DecryptResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	DecryptedCipher      string   `protobuf:"bytes,2,opt,name=decryptedCipher,proto3" json:"decryptedCipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptResponse) Reset()         { *m = DecryptResponse{} }
func (m *DecryptResponse) String() string { return proto.CompactTextString(m) }
func (*DecryptResponse) ProtoMessage()    {}
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e85daebbf806b99, []int{6}
}

func (m *DecryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptResponse.Unmarshal(m, b)
}
func (m *DecryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptResponse.Marshal(b, m, deterministic)
}
func (m *DecryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptResponse.Merge(m, src)
}
func (m *DecryptResponse) XXX_Size() int {
	return xxx_messageInfo_DecryptResponse.Size(m)
}
func (m *DecryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptResponse proto.InternalMessageInfo

func (m *DecryptResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DecryptResponse) GetDecryptedCipher() string {
	if m != nil {
		return m.DecryptedCipher
	}
	return ""
}

func init() {
	proto.RegisterType((*EncryptPayload)(nil), "model.EncryptPayload")
	proto.RegisterType((*DecryptPayload)(nil), "model.DecryptPayload")
	proto.RegisterType((*EncryptRequest)(nil), "model.EncryptRequest")
	proto.RegisterType((*CipherAndKey)(nil), "model.CipherAndKey")
	proto.RegisterType((*EncryptResponse)(nil), "model.EncryptResponse")
	proto.RegisterType((*DecryptRequest)(nil), "model.DecryptRequest")
	proto.RegisterType((*DecryptResponse)(nil), "model.DecryptResponse")
}

func init() { proto.RegisterFile("reed.proto", fileDescriptor_2e85daebbf806b99) }

var fileDescriptor_2e85daebbf806b99 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xb1, 0x4b, 0xc3, 0x40,
	0x14, 0xc6, 0xa9, 0xc5, 0x96, 0xbe, 0x96, 0xa4, 0x9c, 0x58, 0x82, 0x88, 0x48, 0xa6, 0xe2, 0x90,
	0xa1, 0x0e, 0x76, 0x71, 0x10, 0xeb, 0x20, 0xa2, 0xc8, 0xb9, 0x89, 0xcb, 0x99, 0x7b, 0x62, 0x20,
	0x4d, 0xce, 0xcb, 0x55, 0xcc, 0x24, 0xfe, 0xe7, 0x92, 0xe4, 0x19, 0x93, 0x53, 0xab, 0xb8, 0x25,
	0xdf, 0xbb, 0xcb, 0xf7, 0xfb, 0xbe, 0x47, 0x00, 0x34, 0xa2, 0x0c, 0x94, 0x4e, 0x4d, 0xca, 0x36,
	0x97, 0xa9, 0xc4, 0xd8, 0x97, 0xe0, 0x9c, 0x25, 0xa1, 0xce, 0x95, 0xb9, 0x16, 0x79, 0x9c, 0x0a,
	0xc9, 0x26, 0xd0, 0x5b, 0x65, 0xa8, 0xcf, 0xa5, 0xd7, 0xd9, 0xef, 0x4c, 0x07, 0x9c, 0xde, 0xd8,
	0x2e, 0x0c, 0x54, 0x2c, 0xa2, 0xc4, 0xe0, 0x8b, 0xf1, 0x36, 0xca, 0xd1, 0xa7, 0x50, 0x4c, 0x1f,
	0x22, 0x8c, 0xe5, 0x95, 0x58, 0xa2, 0xd7, 0xad, 0xa6, 0xb5, 0xe0, 0xbf, 0x82, 0xb3, 0xc0, 0x3f,
	0xb9, 0x1c, 0xc1, 0x28, 0x8c, 0xd4, 0x23, 0xea, 0x93, 0x44, 0x5e, 0x60, 0x5e, 0x1a, 0x0d, 0x67,
	0x5b, 0x41, 0x49, 0x1b, 0x9c, 0x36, 0x46, 0xbc, 0x75, 0xf0, 0x17, 0x00, 0x51, 0xc7, 0xe4, 0xf8,
	0xb4, 0xc2, 0xcc, 0xb0, 0x31, 0x74, 0x85, 0x8a, 0xc8, 0xbd, 0x78, 0x64, 0xc7, 0xe0, 0x60, 0xab,
	0x0a, 0x32, 0xdf, 0x26, 0xf3, 0x76, 0x4f, 0xdc, 0x3a, 0xec, 0xdf, 0xc2, 0xa8, 0x89, 0xc7, 0xf6,
	0x00, 0x2a, 0xc0, 0xb2, 0xb0, 0xca, 0xa7, 0xa1, 0xb0, 0x03, 0x18, 0xd3, 0x17, 0x50, 0x2e, 0x84,
	0x11, 0x1f, 0x69, 0x07, 0xfc, 0x8b, 0xee, 0xdf, 0x81, 0x5b, 0xe3, 0x67, 0x2a, 0x4d, 0x32, 0xfc,
	0x86, 0xff, 0xbf, 0xd5, 0x15, 0xe5, 0xd0, 0x76, 0xd6, 0x96, 0x23, 0x71, 0x4d, 0x39, 0xed, 0xf5,
	0x72, 0xeb, 0xb0, 0x7f, 0x09, 0x6e, 0x6d, 0xf1, 0x63, 0x80, 0x29, 0xb8, 0x74, 0x0d, 0x65, 0x85,
	0x4b, 0x85, 0xd8, 0xf2, 0xec, 0xad, 0x03, 0x43, 0x8e, 0x28, 0x6f, 0x50, 0x3f, 0x47, 0x21, 0xb2,
	0x39, 0xf4, 0xa9, 0x1f, 0x66, 0x6d, 0x8b, 0x12, 0xed, 0x4c, 0x6c, 0x99, 0x28, 0xe6, 0xd0, 0x27,
	0x30, 0x66, 0x45, 0xb1, 0x6f, 0x5a, 0xfc, 0xf7, 0xbd, 0xf2, 0x3f, 0x3a, 0x7c, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x30, 0x98, 0x61, 0xc6, 0x55, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReedServiceClient is the client API for ReedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReedServiceClient interface {
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
}

type reedServiceClient struct {
	cc *grpc.ClientConn
}

func NewReedServiceClient(cc *grpc.ClientConn) ReedServiceClient {
	return &reedServiceClient{cc}
}

func (c *reedServiceClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/model.ReedService/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reedServiceClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/model.ReedService/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReedServiceServer is the server API for ReedService service.
type ReedServiceServer interface {
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
}

// UnimplementedReedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReedServiceServer struct {
}

func (*UnimplementedReedServiceServer) Encrypt(ctx context.Context, req *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (*UnimplementedReedServiceServer) Decrypt(ctx context.Context, req *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}

func RegisterReedServiceServer(s *grpc.Server, srv ReedServiceServer) {
	s.RegisterService(&_ReedService_serviceDesc, srv)
}

func _ReedService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReedServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.ReedService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReedServiceServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReedService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReedServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.ReedService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReedServiceServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.ReedService",
	HandlerType: (*ReedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _ReedService_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _ReedService_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reed.proto",
}
