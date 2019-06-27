// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: webhookd.proto

package webhookd

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HttpMethod int32

const (
	HttpMethod_UNKNOWN HttpMethod = 0
	HttpMethod_GET     HttpMethod = 1
	HttpMethod_HEAD    HttpMethod = 2
	HttpMethod_POST    HttpMethod = 3
	HttpMethod_PUT     HttpMethod = 4
	HttpMethod_DELETE  HttpMethod = 5
	HttpMethod_CONNECT HttpMethod = 6
	HttpMethod_OPTIONS HttpMethod = 7
	HttpMethod_TRACE   HttpMethod = 8
)

var HttpMethod_name = map[int32]string{
	0: "UNKNOWN",
	1: "GET",
	2: "HEAD",
	3: "POST",
	4: "PUT",
	5: "DELETE",
	6: "CONNECT",
	7: "OPTIONS",
	8: "TRACE",
}

var HttpMethod_value = map[string]int32{
	"UNKNOWN": 0,
	"GET":     1,
	"HEAD":    2,
	"POST":    3,
	"PUT":     4,
	"DELETE":  5,
	"CONNECT": 6,
	"OPTIONS": 7,
	"TRACE":   8,
}

func (x HttpMethod) String() string {
	return proto.EnumName(HttpMethod_name, int32(x))
}

func (HttpMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{0}
}

type SubscriptionCreateRequest struct {
	Provider             *Provider `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Key                  string    `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SubscriptionCreateRequest) Reset()         { *m = SubscriptionCreateRequest{} }
func (m *SubscriptionCreateRequest) String() string { return proto.CompactTextString(m) }
func (*SubscriptionCreateRequest) ProtoMessage()    {}
func (*SubscriptionCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{0}
}
func (m *SubscriptionCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionCreateRequest.Unmarshal(m, b)
}
func (m *SubscriptionCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionCreateRequest.Marshal(b, m, deterministic)
}
func (m *SubscriptionCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionCreateRequest.Merge(m, src)
}
func (m *SubscriptionCreateRequest) XXX_Size() int {
	return xxx_messageInfo_SubscriptionCreateRequest.Size(m)
}
func (m *SubscriptionCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionCreateRequest proto.InternalMessageInfo

func (m *SubscriptionCreateRequest) GetProvider() *Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *SubscriptionCreateRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SubscriptionInviteRequest struct {
	Subscription         *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SubscriptionInviteRequest) Reset()         { *m = SubscriptionInviteRequest{} }
func (m *SubscriptionInviteRequest) String() string { return proto.CompactTextString(m) }
func (*SubscriptionInviteRequest) ProtoMessage()    {}
func (*SubscriptionInviteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{1}
}
func (m *SubscriptionInviteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionInviteRequest.Unmarshal(m, b)
}
func (m *SubscriptionInviteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionInviteRequest.Marshal(b, m, deterministic)
}
func (m *SubscriptionInviteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionInviteRequest.Merge(m, src)
}
func (m *SubscriptionInviteRequest) XXX_Size() int {
	return xxx_messageInfo_SubscriptionInviteRequest.Size(m)
}
func (m *SubscriptionInviteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionInviteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionInviteRequest proto.InternalMessageInfo

func (m *SubscriptionInviteRequest) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *SubscriptionInviteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SubscriptionInviteResponse struct {
	Subscription         *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Token                string        `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	AcceptUrl            string        `protobuf:"bytes,4,opt,name=accept_url,json=acceptUrl,proto3" json:"accept_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SubscriptionInviteResponse) Reset()         { *m = SubscriptionInviteResponse{} }
func (m *SubscriptionInviteResponse) String() string { return proto.CompactTextString(m) }
func (*SubscriptionInviteResponse) ProtoMessage()    {}
func (*SubscriptionInviteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{2}
}
func (m *SubscriptionInviteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionInviteResponse.Unmarshal(m, b)
}
func (m *SubscriptionInviteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionInviteResponse.Marshal(b, m, deterministic)
}
func (m *SubscriptionInviteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionInviteResponse.Merge(m, src)
}
func (m *SubscriptionInviteResponse) XXX_Size() int {
	return xxx_messageInfo_SubscriptionInviteResponse.Size(m)
}
func (m *SubscriptionInviteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionInviteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionInviteResponse proto.InternalMessageInfo

func (m *SubscriptionInviteResponse) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *SubscriptionInviteResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SubscriptionInviteResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SubscriptionInviteResponse) GetAcceptUrl() string {
	if m != nil {
		return m.AcceptUrl
	}
	return ""
}

type SubscriptionAcceptRequest struct {
	Consumer             *Consumer     `protobuf:"bytes,1,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Subscription         *Subscription `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Token                string        `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SubscriptionAcceptRequest) Reset()         { *m = SubscriptionAcceptRequest{} }
func (m *SubscriptionAcceptRequest) String() string { return proto.CompactTextString(m) }
func (*SubscriptionAcceptRequest) ProtoMessage()    {}
func (*SubscriptionAcceptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{3}
}
func (m *SubscriptionAcceptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionAcceptRequest.Unmarshal(m, b)
}
func (m *SubscriptionAcceptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionAcceptRequest.Marshal(b, m, deterministic)
}
func (m *SubscriptionAcceptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionAcceptRequest.Merge(m, src)
}
func (m *SubscriptionAcceptRequest) XXX_Size() int {
	return xxx_messageInfo_SubscriptionAcceptRequest.Size(m)
}
func (m *SubscriptionAcceptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionAcceptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionAcceptRequest proto.InternalMessageInfo

func (m *SubscriptionAcceptRequest) GetConsumer() *Consumer {
	if m != nil {
		return m.Consumer
	}
	return nil
}

func (m *SubscriptionAcceptRequest) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *SubscriptionAcceptRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SubscriptionAcceptResponse struct {
	Received             bool     `protobuf:"varint,1,opt,name=received,proto3" json:"received,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriptionAcceptResponse) Reset()         { *m = SubscriptionAcceptResponse{} }
func (m *SubscriptionAcceptResponse) String() string { return proto.CompactTextString(m) }
func (*SubscriptionAcceptResponse) ProtoMessage()    {}
func (*SubscriptionAcceptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{4}
}
func (m *SubscriptionAcceptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionAcceptResponse.Unmarshal(m, b)
}
func (m *SubscriptionAcceptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionAcceptResponse.Marshal(b, m, deterministic)
}
func (m *SubscriptionAcceptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionAcceptResponse.Merge(m, src)
}
func (m *SubscriptionAcceptResponse) XXX_Size() int {
	return xxx_messageInfo_SubscriptionAcceptResponse.Size(m)
}
func (m *SubscriptionAcceptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionAcceptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionAcceptResponse proto.InternalMessageInfo

func (m *SubscriptionAcceptResponse) GetReceived() bool {
	if m != nil {
		return m.Received
	}
	return false
}

type TransmitRequest struct {
	Provider             *Provider     `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Subscription         *Subscription `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Message              *Message      `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TransmitRequest) Reset()         { *m = TransmitRequest{} }
func (m *TransmitRequest) String() string { return proto.CompactTextString(m) }
func (*TransmitRequest) ProtoMessage()    {}
func (*TransmitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{5}
}
func (m *TransmitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransmitRequest.Unmarshal(m, b)
}
func (m *TransmitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransmitRequest.Marshal(b, m, deterministic)
}
func (m *TransmitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransmitRequest.Merge(m, src)
}
func (m *TransmitRequest) XXX_Size() int {
	return xxx_messageInfo_TransmitRequest.Size(m)
}
func (m *TransmitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransmitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransmitRequest proto.InternalMessageInfo

func (m *TransmitRequest) GetProvider() *Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *TransmitRequest) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *TransmitRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type TransmitResponse struct {
	Received             bool     `protobuf:"varint,1,opt,name=received,proto3" json:"received,omitempty"`
	Message              *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransmitResponse) Reset()         { *m = TransmitResponse{} }
func (m *TransmitResponse) String() string { return proto.CompactTextString(m) }
func (*TransmitResponse) ProtoMessage()    {}
func (*TransmitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{6}
}
func (m *TransmitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransmitResponse.Unmarshal(m, b)
}
func (m *TransmitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransmitResponse.Marshal(b, m, deterministic)
}
func (m *TransmitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransmitResponse.Merge(m, src)
}
func (m *TransmitResponse) XXX_Size() int {
	return xxx_messageInfo_TransmitResponse.Size(m)
}
func (m *TransmitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransmitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransmitResponse proto.InternalMessageInfo

func (m *TransmitResponse) GetReceived() bool {
	if m != nil {
		return m.Received
	}
	return false
}

func (m *TransmitResponse) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type Message struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Method               HttpMethod        `protobuf:"varint,2,opt,name=method,proto3,enum=webhookd.HttpMethod" json:"method,omitempty"`
	Body                 string            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Query                map[string]string `protobuf:"bytes,4,rep,name=query,proto3" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Headers              map[string]string `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{7}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetMethod() HttpMethod {
	if m != nil {
		return m.Method
	}
	return HttpMethod_UNKNOWN
}

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Message) GetQuery() map[string]string {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *Message) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

type Provider struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{8}
}
func (m *Provider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Provider.Unmarshal(m, b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return xxx_messageInfo_Provider.Size(m)
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func (m *Provider) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Provider) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Consumer struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consumer) Reset()         { *m = Consumer{} }
func (m *Consumer) String() string { return proto.CompactTextString(m) }
func (*Consumer) ProtoMessage()    {}
func (*Consumer) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{9}
}
func (m *Consumer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consumer.Unmarshal(m, b)
}
func (m *Consumer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consumer.Marshal(b, m, deterministic)
}
func (m *Consumer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consumer.Merge(m, src)
}
func (m *Consumer) XXX_Size() int {
	return xxx_messageInfo_Consumer.Size(m)
}
func (m *Consumer) XXX_DiscardUnknown() {
	xxx_messageInfo_Consumer.DiscardUnknown(m)
}

var xxx_messageInfo_Consumer proto.InternalMessageInfo

func (m *Consumer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consumer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Subscription struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_f771f8a0984858e8, []int{10}
}
func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Subscription) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterEnum("webhookd.HttpMethod", HttpMethod_name, HttpMethod_value)
	proto.RegisterType((*SubscriptionCreateRequest)(nil), "webhookd.SubscriptionCreateRequest")
	proto.RegisterType((*SubscriptionInviteRequest)(nil), "webhookd.SubscriptionInviteRequest")
	proto.RegisterType((*SubscriptionInviteResponse)(nil), "webhookd.SubscriptionInviteResponse")
	proto.RegisterType((*SubscriptionAcceptRequest)(nil), "webhookd.SubscriptionAcceptRequest")
	proto.RegisterType((*SubscriptionAcceptResponse)(nil), "webhookd.SubscriptionAcceptResponse")
	proto.RegisterType((*TransmitRequest)(nil), "webhookd.TransmitRequest")
	proto.RegisterType((*TransmitResponse)(nil), "webhookd.TransmitResponse")
	proto.RegisterType((*Message)(nil), "webhookd.Message")
	proto.RegisterMapType((map[string]string)(nil), "webhookd.Message.HeadersEntry")
	proto.RegisterMapType((map[string]string)(nil), "webhookd.Message.QueryEntry")
	proto.RegisterType((*Provider)(nil), "webhookd.Provider")
	proto.RegisterType((*Consumer)(nil), "webhookd.Consumer")
	proto.RegisterType((*Subscription)(nil), "webhookd.Subscription")
}

func init() { proto.RegisterFile("webhookd.proto", fileDescriptor_f771f8a0984858e8) }

var fileDescriptor_f771f8a0984858e8 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xc5, 0xce, 0xcb, 0xb9, 0x89, 0x8a, 0x19, 0x55, 0x28, 0xb5, 0x00, 0x55, 0x86, 0x45, 0x79,
	0xd4, 0x45, 0x61, 0x13, 0x75, 0x97, 0xa6, 0x16, 0xad, 0xa0, 0x49, 0x70, 0x5d, 0x55, 0x02, 0xa1,
	0xca, 0xb1, 0x47, 0x8d, 0x95, 0xda, 0xe3, 0xfa, 0x11, 0x94, 0x8f, 0x61, 0xc1, 0x96, 0x9f, 0xe0,
	0xa3, 0xf8, 0x01, 0xe4, 0xf1, 0x3b, 0x76, 0x55, 0x5a, 0xd8, 0xcd, 0xcc, 0x39, 0xf7, 0x9c, 0x99,
	0x3b, 0x73, 0xef, 0xc0, 0xc6, 0x37, 0x3c, 0x9b, 0x13, 0xb2, 0x30, 0x24, 0xc7, 0x25, 0x3e, 0x41,
	0x5c, 0x32, 0x17, 0xbf, 0xc2, 0xd6, 0x69, 0x30, 0xf3, 0x74, 0xd7, 0x74, 0x7c, 0x93, 0xd8, 0x23,
	0x17, 0x6b, 0x3e, 0x56, 0xf0, 0x75, 0x80, 0x3d, 0x1f, 0x49, 0xc0, 0x39, 0x2e, 0x59, 0x9a, 0x06,
	0x76, 0x7b, 0xcc, 0x36, 0xb3, 0xd3, 0xe9, 0x23, 0x29, 0x55, 0x9a, 0xc6, 0x88, 0x92, 0x72, 0x10,
	0x0f, 0xb5, 0x05, 0x5e, 0xf5, 0xd8, 0x6d, 0x66, 0xa7, 0xad, 0x84, 0x43, 0x71, 0x51, 0x94, 0x3f,
	0xb6, 0x97, 0x66, 0x26, 0xbf, 0x0f, 0x5d, 0x2f, 0x07, 0xc6, 0x16, 0x8f, 0x33, 0x8b, 0x7c, 0xa8,
	0x52, 0xe0, 0x22, 0x04, 0x75, 0x5b, 0xb3, 0x70, 0xec, 0x45, 0xc7, 0xe2, 0x0f, 0x06, 0x84, 0x2a,
	0x37, 0xcf, 0x21, 0xb6, 0x87, 0xff, 0xb7, 0x1d, 0xda, 0x84, 0x86, 0x4f, 0x16, 0xd8, 0xee, 0xd5,
	0xe8, 0x62, 0x34, 0x41, 0x4f, 0x01, 0x34, 0x5d, 0xc7, 0x8e, 0x7f, 0x11, 0xb8, 0x57, 0xbd, 0x3a,
	0x85, 0xda, 0xd1, 0xca, 0x99, 0x7b, 0x25, 0x7e, 0x67, 0x8a, 0x19, 0x19, 0x52, 0x24, 0x97, 0x70,
	0x9d, 0xd8, 0x5e, 0x60, 0x55, 0x25, 0x7c, 0x14, 0x23, 0x4a, 0xca, 0x29, 0x1d, 0x89, 0xbd, 0xc3,
	0x91, 0x2a, 0xb7, 0x2f, 0x0e, 0x8a, 0x29, 0x4c, 0xb6, 0x17, 0xa7, 0x50, 0x00, 0xce, 0xc5, 0x3a,
	0x36, 0x97, 0xd8, 0xa0, 0xfb, 0xe3, 0x94, 0x74, 0x2e, 0xfe, 0x64, 0xe0, 0xa1, 0xea, 0x6a, 0xb6,
	0x67, 0x99, 0xfe, 0x7d, 0x1f, 0xd0, 0xbf, 0x9c, 0xe7, 0x35, 0xb4, 0x2c, 0xec, 0x79, 0xda, 0x25,
	0xa6, 0x27, 0xea, 0xf4, 0x1f, 0x65, 0x61, 0x27, 0x11, 0xa0, 0x24, 0x0c, 0xf1, 0x0b, 0xf0, 0xd9,
	0x5e, 0x6f, 0x3f, 0x5c, 0x5e, 0x9c, 0xbd, 0x55, 0xfc, 0x17, 0x0b, 0xad, 0x78, 0x11, 0x6d, 0x00,
	0x6b, 0x46, 0x72, 0x6d, 0x85, 0x35, 0x0d, 0xf4, 0x06, 0x9a, 0x16, 0xf6, 0xe7, 0xc4, 0xa0, 0x3a,
	0x1b, 0xfd, 0xcd, 0x4c, 0xe7, 0xc8, 0xf7, 0x9d, 0x13, 0x8a, 0x29, 0x31, 0x27, 0x7c, 0x76, 0x33,
	0x62, 0xac, 0xe2, 0x2b, 0xa2, 0x63, 0xd4, 0x87, 0xc6, 0x75, 0x80, 0xdd, 0x55, 0xaf, 0xbe, 0x5d,
	0xdb, 0xe9, 0xf4, 0x9f, 0x94, 0x36, 0x22, 0x7d, 0x0a, 0x61, 0xd9, 0xf6, 0xdd, 0x95, 0x12, 0x51,
	0xd1, 0x00, 0x5a, 0x73, 0xac, 0x19, 0xd8, 0xf5, 0x7a, 0x0d, 0x1a, 0xf5, 0xac, 0x1c, 0x75, 0x14,
	0x11, 0xa2, 0xb8, 0x84, 0x2e, 0x0c, 0x00, 0x32, 0xb9, 0xa4, 0xc0, 0x99, 0xb4, 0xc0, 0xc3, 0x57,
	0xb4, 0xd4, 0xae, 0x82, 0xa4, 0x32, 0xa2, 0xc9, 0x3e, 0x3b, 0x60, 0x84, 0x7d, 0xe8, 0xe6, 0x25,
	0xef, 0x12, 0x2b, 0x4a, 0xc0, 0x25, 0xaf, 0xa3, 0x94, 0xc1, 0xaa, 0xca, 0x97, 0x80, 0x4b, 0xaa,
	0xe3, 0xaf, 0xf8, 0x6f, 0xa1, 0x9b, 0x7f, 0x49, 0xa5, 0x98, 0x52, 0x23, 0x7b, 0xe5, 0x00, 0x64,
	0xf7, 0x83, 0x3a, 0xd0, 0x3a, 0x1b, 0x7f, 0x18, 0x4f, 0xce, 0xc7, 0xfc, 0x03, 0xd4, 0x82, 0xda,
	0x7b, 0x59, 0xe5, 0x19, 0xc4, 0x41, 0xfd, 0x48, 0x1e, 0x1e, 0xf2, 0x6c, 0x38, 0x9a, 0x4e, 0x4e,
	0x55, 0xbe, 0x16, 0x82, 0xd3, 0x33, 0x95, 0xaf, 0x23, 0x80, 0xe6, 0xa1, 0xfc, 0x51, 0x56, 0x65,
	0xbe, 0x11, 0x86, 0x8f, 0x26, 0xe3, 0xb1, 0x3c, 0x52, 0xf9, 0x66, 0x38, 0x99, 0x4c, 0xd5, 0xe3,
	0xc9, 0xf8, 0x94, 0x6f, 0xa1, 0x36, 0x34, 0x54, 0x65, 0x38, 0x92, 0x79, 0xae, 0xff, 0x9b, 0x85,
	0xce, 0x79, 0x7c, 0x49, 0x43, 0xc7, 0x44, 0x17, 0x80, 0xca, 0x9d, 0x1a, 0x3d, 0xaf, 0xae, 0x8d,
	0x42, 0x1f, 0x17, 0x5e, 0x54, 0x93, 0xd6, 0xfa, 0xe3, 0x9a, 0x41, 0x84, 0xde, 0x64, 0x50, 0xe8,
	0xe4, 0xf7, 0x33, 0x88, 0x7a, 0xcb, 0x4d, 0x06, 0x85, 0xc6, 0x78, 0x93, 0xc1, 0x5a, 0x7b, 0x1a,
	0x02, 0x97, 0x54, 0x35, 0xda, 0xca, 0x22, 0xd6, 0xba, 0x92, 0x20, 0x54, 0x41, 0x91, 0xc4, 0xc1,
	0x09, 0x74, 0x75, 0x62, 0xa5, 0x84, 0x03, 0x3e, 0x77, 0x05, 0xd3, 0xf0, 0xef, 0x9c, 0x32, 0x9f,
	0x5f, 0x5e, 0x9a, 0xfe, 0x3c, 0x98, 0x49, 0x3a, 0xb1, 0xf6, 0x12, 0xe2, 0x9e, 0xe6, 0x98, 0xbb,
	0xf4, 0x73, 0xdd, 0xbd, 0x24, 0xe9, 0xea, 0xac, 0x49, 0x97, 0xde, 0xfd, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0xa8, 0xed, 0xf0, 0x81, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WebhookdApiClient is the client API for WebhookdApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebhookdApiClient interface {
	SubscriptionCreate(ctx context.Context, in *SubscriptionCreateRequest, opts ...grpc.CallOption) (*SubscriptionInviteResponse, error)
	SubscriptionInvite(ctx context.Context, in *SubscriptionInviteRequest, opts ...grpc.CallOption) (*SubscriptionInviteResponse, error)
	SubscriptionAccept(ctx context.Context, in *SubscriptionAcceptRequest, opts ...grpc.CallOption) (*SubscriptionAcceptResponse, error)
	Transmit(ctx context.Context, in *TransmitRequest, opts ...grpc.CallOption) (*TransmitResponse, error)
}

type webhookdApiClient struct {
	cc *grpc.ClientConn
}

func NewWebhookdApiClient(cc *grpc.ClientConn) WebhookdApiClient {
	return &webhookdApiClient{cc}
}

func (c *webhookdApiClient) SubscriptionCreate(ctx context.Context, in *SubscriptionCreateRequest, opts ...grpc.CallOption) (*SubscriptionInviteResponse, error) {
	out := new(SubscriptionInviteResponse)
	err := c.cc.Invoke(ctx, "/webhookd.WebhookdApi/SubscriptionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookdApiClient) SubscriptionInvite(ctx context.Context, in *SubscriptionInviteRequest, opts ...grpc.CallOption) (*SubscriptionInviteResponse, error) {
	out := new(SubscriptionInviteResponse)
	err := c.cc.Invoke(ctx, "/webhookd.WebhookdApi/SubscriptionInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookdApiClient) SubscriptionAccept(ctx context.Context, in *SubscriptionAcceptRequest, opts ...grpc.CallOption) (*SubscriptionAcceptResponse, error) {
	out := new(SubscriptionAcceptResponse)
	err := c.cc.Invoke(ctx, "/webhookd.WebhookdApi/SubscriptionAccept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookdApiClient) Transmit(ctx context.Context, in *TransmitRequest, opts ...grpc.CallOption) (*TransmitResponse, error) {
	out := new(TransmitResponse)
	err := c.cc.Invoke(ctx, "/webhookd.WebhookdApi/Transmit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebhookdApiServer is the server API for WebhookdApi service.
type WebhookdApiServer interface {
	SubscriptionCreate(context.Context, *SubscriptionCreateRequest) (*SubscriptionInviteResponse, error)
	SubscriptionInvite(context.Context, *SubscriptionInviteRequest) (*SubscriptionInviteResponse, error)
	SubscriptionAccept(context.Context, *SubscriptionAcceptRequest) (*SubscriptionAcceptResponse, error)
	Transmit(context.Context, *TransmitRequest) (*TransmitResponse, error)
}

func RegisterWebhookdApiServer(s *grpc.Server, srv WebhookdApiServer) {
	s.RegisterService(&_WebhookdApi_serviceDesc, srv)
}

func _WebhookdApi_SubscriptionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookdApiServer).SubscriptionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webhookd.WebhookdApi/SubscriptionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookdApiServer).SubscriptionCreate(ctx, req.(*SubscriptionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebhookdApi_SubscriptionInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookdApiServer).SubscriptionInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webhookd.WebhookdApi/SubscriptionInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookdApiServer).SubscriptionInvite(ctx, req.(*SubscriptionInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebhookdApi_SubscriptionAccept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionAcceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookdApiServer).SubscriptionAccept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webhookd.WebhookdApi/SubscriptionAccept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookdApiServer).SubscriptionAccept(ctx, req.(*SubscriptionAcceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebhookdApi_Transmit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookdApiServer).Transmit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webhookd.WebhookdApi/Transmit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookdApiServer).Transmit(ctx, req.(*TransmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WebhookdApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "webhookd.WebhookdApi",
	HandlerType: (*WebhookdApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubscriptionCreate",
			Handler:    _WebhookdApi_SubscriptionCreate_Handler,
		},
		{
			MethodName: "SubscriptionInvite",
			Handler:    _WebhookdApi_SubscriptionInvite_Handler,
		},
		{
			MethodName: "SubscriptionAccept",
			Handler:    _WebhookdApi_SubscriptionAccept_Handler,
		},
		{
			MethodName: "Transmit",
			Handler:    _WebhookdApi_Transmit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webhookd.proto",
}
