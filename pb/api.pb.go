// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientHello struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PublicId             []byte   `protobuf:"bytes,2,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientHello) Reset()         { *m = ClientHello{} }
func (m *ClientHello) String() string { return proto.CompactTextString(m) }
func (*ClientHello) ProtoMessage()    {}
func (*ClientHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_3fc349b9ce835f03, []int{0}
}
func (m *ClientHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientHello.Unmarshal(m, b)
}
func (m *ClientHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientHello.Marshal(b, m, deterministic)
}
func (dst *ClientHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientHello.Merge(dst, src)
}
func (m *ClientHello) XXX_Size() int {
	return xxx_messageInfo_ClientHello.Size(m)
}
func (m *ClientHello) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientHello.DiscardUnknown(m)
}

var xxx_messageInfo_ClientHello proto.InternalMessageInfo

func (m *ClientHello) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ClientHello) GetPublicId() []byte {
	if m != nil {
		return m.PublicId
	}
	return nil
}

type ClientNumber struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// double number = 2;
	Number               []byte   `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientNumber) Reset()         { *m = ClientNumber{} }
func (m *ClientNumber) String() string { return proto.CompactTextString(m) }
func (*ClientNumber) ProtoMessage()    {}
func (*ClientNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_3fc349b9ce835f03, []int{1}
}
func (m *ClientNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientNumber.Unmarshal(m, b)
}
func (m *ClientNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientNumber.Marshal(b, m, deterministic)
}
func (dst *ClientNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientNumber.Merge(dst, src)
}
func (m *ClientNumber) XXX_Size() int {
	return xxx_messageInfo_ClientNumber.Size(m)
}
func (m *ClientNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientNumber.DiscardUnknown(m)
}

var xxx_messageInfo_ClientNumber proto.InternalMessageInfo

func (m *ClientNumber) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ClientNumber) GetNumber() []byte {
	if m != nil {
		return m.Number
	}
	return nil
}

type ClientCurrentMaxNumber struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientCurrentMaxNumber) Reset()         { *m = ClientCurrentMaxNumber{} }
func (m *ClientCurrentMaxNumber) String() string { return proto.CompactTextString(m) }
func (*ClientCurrentMaxNumber) ProtoMessage()    {}
func (*ClientCurrentMaxNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_3fc349b9ce835f03, []int{2}
}
func (m *ClientCurrentMaxNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientCurrentMaxNumber.Unmarshal(m, b)
}
func (m *ClientCurrentMaxNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientCurrentMaxNumber.Marshal(b, m, deterministic)
}
func (dst *ClientCurrentMaxNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientCurrentMaxNumber.Merge(dst, src)
}
func (m *ClientCurrentMaxNumber) XXX_Size() int {
	return xxx_messageInfo_ClientCurrentMaxNumber.Size(m)
}
func (m *ClientCurrentMaxNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientCurrentMaxNumber.DiscardUnknown(m)
}

var xxx_messageInfo_ClientCurrentMaxNumber proto.InternalMessageInfo

func (m *ClientCurrentMaxNumber) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ClientResetCurrentMaxNumber struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientResetCurrentMaxNumber) Reset()         { *m = ClientResetCurrentMaxNumber{} }
func (m *ClientResetCurrentMaxNumber) String() string { return proto.CompactTextString(m) }
func (*ClientResetCurrentMaxNumber) ProtoMessage()    {}
func (*ClientResetCurrentMaxNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_3fc349b9ce835f03, []int{3}
}
func (m *ClientResetCurrentMaxNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientResetCurrentMaxNumber.Unmarshal(m, b)
}
func (m *ClientResetCurrentMaxNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientResetCurrentMaxNumber.Marshal(b, m, deterministic)
}
func (dst *ClientResetCurrentMaxNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientResetCurrentMaxNumber.Merge(dst, src)
}
func (m *ClientResetCurrentMaxNumber) XXX_Size() int {
	return xxx_messageInfo_ClientResetCurrentMaxNumber.Size(m)
}
func (m *ClientResetCurrentMaxNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientResetCurrentMaxNumber.DiscardUnknown(m)
}

var xxx_messageInfo_ClientResetCurrentMaxNumber proto.InternalMessageInfo

func (m *ClientResetCurrentMaxNumber) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ClientMessage struct {
	// Types that are valid to be assigned to Message:
	//	*ClientMessage_Hello
	//	*ClientMessage_Number
	//	*ClientMessage_CurrentMaxNumber
	//	*ClientMessage_ResetCurrentMaxNumber
	Message              isClientMessage_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientMessage) Reset()         { *m = ClientMessage{} }
func (m *ClientMessage) String() string { return proto.CompactTextString(m) }
func (*ClientMessage) ProtoMessage()    {}
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_3fc349b9ce835f03, []int{4}
}
func (m *ClientMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientMessage.Unmarshal(m, b)
}
func (m *ClientMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientMessage.Marshal(b, m, deterministic)
}
func (dst *ClientMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientMessage.Merge(dst, src)
}
func (m *ClientMessage) XXX_Size() int {
	return xxx_messageInfo_ClientMessage.Size(m)
}
func (m *ClientMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ClientMessage proto.InternalMessageInfo

type isClientMessage_Message interface {
	isClientMessage_Message()
}

type ClientMessage_Hello struct {
	Hello *ClientHello `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}
type ClientMessage_Number struct {
	Number *ClientNumber `protobuf:"bytes,2,opt,name=number,proto3,oneof"`
}
type ClientMessage_CurrentMaxNumber struct {
	CurrentMaxNumber *ClientCurrentMaxNumber `protobuf:"bytes,3,opt,name=current_max_number,json=currentMaxNumber,proto3,oneof"`
}
type ClientMessage_ResetCurrentMaxNumber struct {
	ResetCurrentMaxNumber *ClientResetCurrentMaxNumber `protobuf:"bytes,4,opt,name=reset_current_max_number,json=resetCurrentMaxNumber,proto3,oneof"`
}

func (*ClientMessage_Hello) isClientMessage_Message()                 {}
func (*ClientMessage_Number) isClientMessage_Message()                {}
func (*ClientMessage_CurrentMaxNumber) isClientMessage_Message()      {}
func (*ClientMessage_ResetCurrentMaxNumber) isClientMessage_Message() {}

func (m *ClientMessage) GetMessage() isClientMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ClientMessage) GetHello() *ClientHello {
	if x, ok := m.GetMessage().(*ClientMessage_Hello); ok {
		return x.Hello
	}
	return nil
}

func (m *ClientMessage) GetNumber() *ClientNumber {
	if x, ok := m.GetMessage().(*ClientMessage_Number); ok {
		return x.Number
	}
	return nil
}

func (m *ClientMessage) GetCurrentMaxNumber() *ClientCurrentMaxNumber {
	if x, ok := m.GetMessage().(*ClientMessage_CurrentMaxNumber); ok {
		return x.CurrentMaxNumber
	}
	return nil
}

func (m *ClientMessage) GetResetCurrentMaxNumber() *ClientResetCurrentMaxNumber {
	if x, ok := m.GetMessage().(*ClientMessage_ResetCurrentMaxNumber); ok {
		return x.ResetCurrentMaxNumber
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ClientMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ClientMessage_OneofMarshaler, _ClientMessage_OneofUnmarshaler, _ClientMessage_OneofSizer, []interface{}{
		(*ClientMessage_Hello)(nil),
		(*ClientMessage_Number)(nil),
		(*ClientMessage_CurrentMaxNumber)(nil),
		(*ClientMessage_ResetCurrentMaxNumber)(nil),
	}
}

func _ClientMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ClientMessage)
	// Message
	switch x := m.Message.(type) {
	case *ClientMessage_Hello:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Hello); err != nil {
			return err
		}
	case *ClientMessage_Number:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Number); err != nil {
			return err
		}
	case *ClientMessage_CurrentMaxNumber:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CurrentMaxNumber); err != nil {
			return err
		}
	case *ClientMessage_ResetCurrentMaxNumber:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResetCurrentMaxNumber); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ClientMessage.Message has unexpected type %T", x)
	}
	return nil
}

func _ClientMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ClientMessage)
	switch tag {
	case 1: // Message.hello
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientHello)
		err := b.DecodeMessage(msg)
		m.Message = &ClientMessage_Hello{msg}
		return true, err
	case 2: // Message.number
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientNumber)
		err := b.DecodeMessage(msg)
		m.Message = &ClientMessage_Number{msg}
		return true, err
	case 3: // Message.current_max_number
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientCurrentMaxNumber)
		err := b.DecodeMessage(msg)
		m.Message = &ClientMessage_CurrentMaxNumber{msg}
		return true, err
	case 4: // Message.reset_current_max_number
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClientResetCurrentMaxNumber)
		err := b.DecodeMessage(msg)
		m.Message = &ClientMessage_ResetCurrentMaxNumber{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ClientMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ClientMessage)
	// Message
	switch x := m.Message.(type) {
	case *ClientMessage_Hello:
		s := proto.Size(x.Hello)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientMessage_Number:
		s := proto.Size(x.Number)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientMessage_CurrentMaxNumber:
		s := proto.Size(x.CurrentMaxNumber)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ClientMessage_ResetCurrentMaxNumber:
		s := proto.Size(x.ResetCurrentMaxNumber)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ServerHi struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PublicId             []byte   `protobuf:"bytes,2,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerHi) Reset()         { *m = ServerHi{} }
func (m *ServerHi) String() string { return proto.CompactTextString(m) }
func (*ServerHi) ProtoMessage()    {}
func (*ServerHi) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_3fc349b9ce835f03, []int{5}
}
func (m *ServerHi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerHi.Unmarshal(m, b)
}
func (m *ServerHi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerHi.Marshal(b, m, deterministic)
}
func (dst *ServerHi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerHi.Merge(dst, src)
}
func (m *ServerHi) XXX_Size() int {
	return xxx_messageInfo_ServerHi.Size(m)
}
func (m *ServerHi) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerHi.DiscardUnknown(m)
}

var xxx_messageInfo_ServerHi proto.InternalMessageInfo

func (m *ServerHi) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ServerHi) GetPublicId() []byte {
	if m != nil {
		return m.PublicId
	}
	return nil
}

type ServerNumber struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// double number = 2;
	Number               []byte   `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerNumber) Reset()         { *m = ServerNumber{} }
func (m *ServerNumber) String() string { return proto.CompactTextString(m) }
func (*ServerNumber) ProtoMessage()    {}
func (*ServerNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_3fc349b9ce835f03, []int{6}
}
func (m *ServerNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerNumber.Unmarshal(m, b)
}
func (m *ServerNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerNumber.Marshal(b, m, deterministic)
}
func (dst *ServerNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerNumber.Merge(dst, src)
}
func (m *ServerNumber) XXX_Size() int {
	return xxx_messageInfo_ServerNumber.Size(m)
}
func (m *ServerNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerNumber.DiscardUnknown(m)
}

var xxx_messageInfo_ServerNumber proto.InternalMessageInfo

func (m *ServerNumber) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ServerNumber) GetNumber() []byte {
	if m != nil {
		return m.Number
	}
	return nil
}

type ServerMessage struct {
	// Types that are valid to be assigned to Message:
	//	*ServerMessage_Hi
	//	*ServerMessage_Number
	Message              isServerMessage_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ServerMessage) Reset()         { *m = ServerMessage{} }
func (m *ServerMessage) String() string { return proto.CompactTextString(m) }
func (*ServerMessage) ProtoMessage()    {}
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_3fc349b9ce835f03, []int{7}
}
func (m *ServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerMessage.Unmarshal(m, b)
}
func (m *ServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerMessage.Marshal(b, m, deterministic)
}
func (dst *ServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerMessage.Merge(dst, src)
}
func (m *ServerMessage) XXX_Size() int {
	return xxx_messageInfo_ServerMessage.Size(m)
}
func (m *ServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ServerMessage proto.InternalMessageInfo

type isServerMessage_Message interface {
	isServerMessage_Message()
}

type ServerMessage_Hi struct {
	Hi *ServerHi `protobuf:"bytes,1,opt,name=hi,proto3,oneof"`
}
type ServerMessage_Number struct {
	Number *ServerNumber `protobuf:"bytes,2,opt,name=number,proto3,oneof"`
}

func (*ServerMessage_Hi) isServerMessage_Message()     {}
func (*ServerMessage_Number) isServerMessage_Message() {}

func (m *ServerMessage) GetMessage() isServerMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ServerMessage) GetHi() *ServerHi {
	if x, ok := m.GetMessage().(*ServerMessage_Hi); ok {
		return x.Hi
	}
	return nil
}

func (m *ServerMessage) GetNumber() *ServerNumber {
	if x, ok := m.GetMessage().(*ServerMessage_Number); ok {
		return x.Number
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServerMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServerMessage_OneofMarshaler, _ServerMessage_OneofUnmarshaler, _ServerMessage_OneofSizer, []interface{}{
		(*ServerMessage_Hi)(nil),
		(*ServerMessage_Number)(nil),
	}
}

func _ServerMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServerMessage)
	// Message
	switch x := m.Message.(type) {
	case *ServerMessage_Hi:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Hi); err != nil {
			return err
		}
	case *ServerMessage_Number:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Number); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ServerMessage.Message has unexpected type %T", x)
	}
	return nil
}

func _ServerMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServerMessage)
	switch tag {
	case 1: // Message.hi
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ServerHi)
		err := b.DecodeMessage(msg)
		m.Message = &ServerMessage_Hi{msg}
		return true, err
	case 2: // Message.number
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ServerNumber)
		err := b.DecodeMessage(msg)
		m.Message = &ServerMessage_Number{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ServerMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServerMessage)
	// Message
	switch x := m.Message.(type) {
	case *ServerMessage_Hi:
		s := proto.Size(x.Hi)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerMessage_Number:
		s := proto.Size(x.Number)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ClientHello)(nil), "pb.ClientHello")
	proto.RegisterType((*ClientNumber)(nil), "pb.ClientNumber")
	proto.RegisterType((*ClientCurrentMaxNumber)(nil), "pb.ClientCurrentMaxNumber")
	proto.RegisterType((*ClientResetCurrentMaxNumber)(nil), "pb.ClientResetCurrentMaxNumber")
	proto.RegisterType((*ClientMessage)(nil), "pb.ClientMessage")
	proto.RegisterType((*ServerHi)(nil), "pb.ServerHi")
	proto.RegisterType((*ServerNumber)(nil), "pb.ServerNumber")
	proto.RegisterType((*ServerMessage)(nil), "pb.ServerMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Message(ctx context.Context, opts ...grpc.CallOption) (Service_MessageClient, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Message(ctx context.Context, opts ...grpc.CallOption) (Service_MessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[0], "/pb.Service/Message", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceMessageClient{stream}
	return x, nil
}

type Service_MessageClient interface {
	Send(*ClientMessage) error
	Recv() (*ServerMessage, error)
	grpc.ClientStream
}

type serviceMessageClient struct {
	grpc.ClientStream
}

func (x *serviceMessageClient) Send(m *ClientMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceMessageClient) Recv() (*ServerMessage, error) {
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Message(Service_MessageServer) error
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Message_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).Message(&serviceMessageServer{stream})
}

type Service_MessageServer interface {
	Send(*ServerMessage) error
	Recv() (*ClientMessage, error)
	grpc.ServerStream
}

type serviceMessageServer struct {
	grpc.ServerStream
}

func (x *serviceMessageServer) Send(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceMessageServer) Recv() (*ClientMessage, error) {
	m := new(ClientMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Message",
			Handler:       _Service_Message_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_3fc349b9ce835f03) }

var fileDescriptor_api_3fc349b9ce835f03 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x4c, 0xb6, 0xda, 0x8f, 0xd7, 0x56, 0xdb, 0x05, 0x4b, 0x69, 0x41, 0x4b, 0x2e, 0x06, 0xc1,
	0x22, 0x2d, 0x28, 0x78, 0xf0, 0x60, 0x2f, 0xab, 0x50, 0x0f, 0xeb, 0xcd, 0x4b, 0xc8, 0xc7, 0x6a,
	0x17, 0xd2, 0x36, 0x6c, 0x12, 0xe9, 0x3f, 0xf1, 0xef, 0x4a, 0xb2, 0x1b, 0x37, 0xd6, 0x80, 0xf4,
	0x98, 0xf7, 0x66, 0x66, 0x67, 0xde, 0x10, 0x68, 0xb9, 0x11, 0x9f, 0x46, 0x62, 0x9b, 0x6c, 0x31,
	0x8a, 0x3c, 0xeb, 0x1e, 0xda, 0x8b, 0x90, 0xb3, 0x4d, 0x42, 0x58, 0x18, 0x6e, 0xf1, 0x09, 0x20,
	0x1e, 0x0c, 0xcd, 0x89, 0x69, 0xd7, 0x28, 0xe2, 0x01, 0x1e, 0x43, 0x2b, 0x4a, 0xbd, 0x90, 0xfb,
	0x0e, 0x0f, 0x86, 0x68, 0x62, 0xda, 0x1d, 0xda, 0x94, 0x83, 0xa7, 0xc0, 0xba, 0x85, 0x8e, 0xe4,
	0xbe, 0xa4, 0x6b, 0x8f, 0x89, 0x3f, 0xe4, 0x01, 0xd4, 0x37, 0xf9, 0x46, 0x31, 0xd5, 0x97, 0x65,
	0xc3, 0x40, 0xf2, 0x16, 0xa9, 0x10, 0x6c, 0x93, 0x2c, 0xdd, 0x5d, 0xb5, 0x82, 0x75, 0x0d, 0x63,
	0x89, 0xa4, 0x2c, 0x66, 0xff, 0xc3, 0xbf, 0x10, 0x74, 0x25, 0x7e, 0xc9, 0xe2, 0xd8, 0xfd, 0x60,
	0xf8, 0x12, 0x8e, 0x57, 0x59, 0xb0, 0x1c, 0xd4, 0x9e, 0x9d, 0x4e, 0x23, 0x6f, 0x5a, 0xca, 0x4b,
	0x0c, 0x2a, 0xf7, 0xf8, 0xea, 0x97, 0xd7, 0xf6, 0xac, 0xa7, 0x91, 0xf2, 0x31, 0x62, 0x14, 0xfe,
	0xf1, 0x33, 0x60, 0x5f, 0x5a, 0x71, 0xd6, 0xee, 0xce, 0x51, 0xbc, 0x5a, 0xce, 0x1b, 0x69, 0xde,
	0xbe, 0x5d, 0x62, 0xd0, 0x9e, 0xbf, 0x1f, 0xe1, 0x0d, 0x86, 0x22, 0xcb, 0xe6, 0x54, 0x28, 0x1e,
	0xe5, 0x8a, 0x17, 0x5a, 0xb1, 0xf2, 0x0a, 0xc4, 0xa0, 0x67, 0xa2, 0x6a, 0xf1, 0xd8, 0x82, 0x86,
	0xba, 0x83, 0x75, 0x07, 0xcd, 0x57, 0x26, 0x3e, 0x99, 0x20, 0xfc, 0xe0, 0x8e, 0x25, 0xf1, 0xc0,
	0x8e, 0xdf, 0xa1, 0x2b, 0x79, 0x45, 0x13, 0xe7, 0x80, 0x56, 0x5c, 0xd5, 0xd0, 0xc9, 0x22, 0x15,
	0x7e, 0x88, 0x41, 0xd1, 0x8a, 0x57, 0x17, 0x50, 0x7e, 0x5a, 0x17, 0x50, 0x0a, 0x36, 0x7b, 0x80,
	0x46, 0x06, 0xe2, 0x3e, 0xc3, 0xf3, 0x9f, 0x29, 0xee, 0xeb, 0x9b, 0xa9, 0xd1, 0xa8, 0xaf, 0xf5,
	0xd4, 0xc8, 0x36, 0x6f, 0x4c, 0xaf, 0x9e, 0xff, 0x0a, 0xf3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4c, 0xeb, 0xee, 0xe5, 0x17, 0x03, 0x00, 0x00,
}
