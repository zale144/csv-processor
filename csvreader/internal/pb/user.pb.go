// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

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

type UserBatchReq struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBatchReq) Reset()         { *m = UserBatchReq{} }
func (m *UserBatchReq) String() string { return proto.CompactTextString(m) }
func (*UserBatchReq) ProtoMessage()    {}
func (*UserBatchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserBatchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBatchReq.Unmarshal(m, b)
}
func (m *UserBatchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBatchReq.Marshal(b, m, deterministic)
}
func (m *UserBatchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBatchReq.Merge(m, src)
}
func (m *UserBatchReq) XXX_Size() int {
	return xxx_messageInfo_UserBatchReq.Size(m)
}
func (m *UserBatchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBatchReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserBatchReq proto.InternalMessageInfo

func (m *UserBatchReq) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserBatchRsp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBatchRsp) Reset()         { *m = UserBatchRsp{} }
func (m *UserBatchRsp) String() string { return proto.CompactTextString(m) }
func (*UserBatchRsp) ProtoMessage()    {}
func (*UserBatchRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserBatchRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBatchRsp.Unmarshal(m, b)
}
func (m *UserBatchRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBatchRsp.Marshal(b, m, deterministic)
}
func (m *UserBatchRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBatchRsp.Merge(m, src)
}
func (m *UserBatchRsp) XXX_Size() int {
	return xxx_messageInfo_UserBatchRsp.Size(m)
}
func (m *UserBatchRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBatchRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserBatchRsp proto.InternalMessageInfo

func (m *UserBatchRsp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*UserBatchReq)(nil), "pb.UserBatchReq")
	proto.RegisterType((*UserBatchRsp)(nil), "pb.UserBatchRsp")
	proto.RegisterType((*User)(nil), "pb.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x49, 0xd2, 0x40, 0x7a, 0x45, 0xa8, 0x3a, 0x31, 0x58, 0x15, 0x42, 0x95, 0x27, 0x4f,
	0x1e, 0xca, 0xc2, 0xcc, 0xc2, 0x86, 0x90, 0x11, 0x0f, 0xe0, 0xb8, 0x87, 0x6a, 0xa9, 0xc5, 0xc6,
	0x97, 0x32, 0xf1, 0xf0, 0xc8, 0x0e, 0x29, 0x88, 0xf1, 0xff, 0xfe, 0xcf, 0xf2, 0xdd, 0x01, 0x1c,
	0x99, 0x92, 0x8e, 0x29, 0x0c, 0x01, 0xeb, 0xd8, 0x4b, 0x0d, 0x97, 0xaf, 0x4c, 0xe9, 0xc1, 0x0e,
	0x6e, 0x67, 0xe8, 0x03, 0x6f, 0xa1, 0xcd, 0x06, 0x8b, 0x6a, 0xdd, 0xa8, 0xc5, 0xa6, 0xd3, 0xb1,
	0xd7, 0x59, 0x30, 0x23, 0x96, 0xea, 0xaf, 0xcf, 0x11, 0x05, 0x5c, 0xf0, 0xd1, 0x39, 0xe2, 0xfc,
	0xa2, 0x52, 0x9d, 0x99, 0xa2, 0xfc, 0x82, 0x59, 0x36, 0xf1, 0x0a, 0x6a, 0xbf, 0x2d, 0x65, 0x63,
	0x6a, 0xbf, 0xc5, 0x1b, 0x98, 0xbf, 0xf9, 0xc4, 0xc3, 0x93, 0x3d, 0x90, 0xa8, 0xd7, 0x95, 0x9a,
	0x9b, 0x5f, 0x80, 0x2b, 0xe8, 0xf6, 0xf6, 0xa7, 0x6c, 0x4a, 0x79, 0xca, 0x78, 0x0d, 0x2d, 0x1d,
	0xac, 0xdf, 0x8b, 0x59, 0x29, 0xc6, 0x90, 0x69, 0xdc, 0x85, 0x77, 0x12, 0xed, 0x48, 0x4b, 0xd8,
	0x3c, 0xc2, 0x22, 0xff, 0xfe, 0x42, 0xe9, 0xd3, 0x3b, 0xc2, 0x7b, 0x58, 0x3e, 0xa7, 0x90, 0xe7,
	0x3a, 0x4d, 0x8f, 0xcb, 0x69, 0xb7, 0x69, 0xf9, 0xd5, 0x3f, 0xc2, 0x51, 0x9e, 0xf5, 0xe7, 0xe5,
	0x56, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xcc, 0x0f, 0x60, 0x39, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	ProcessUserBatch(ctx context.Context, in *UserBatchReq, opts ...grpc.CallOption) (*UserBatchRsp, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ProcessUserBatch(ctx context.Context, in *UserBatchReq, opts ...grpc.CallOption) (*UserBatchRsp, error) {
	out := new(UserBatchRsp)
	err := c.cc.Invoke(ctx, "/pb.UserService/ProcessUserBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	ProcessUserBatch(context.Context, *UserBatchReq) (*UserBatchRsp, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) ProcessUserBatch(ctx context.Context, req *UserBatchReq) (*UserBatchRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessUserBatch not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_ProcessUserBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ProcessUserBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/ProcessUserBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ProcessUserBatch(ctx, req.(*UserBatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessUserBatch",
			Handler:    _UserService_ProcessUserBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
