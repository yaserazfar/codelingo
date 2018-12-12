// Code generated by protoc-gen-go.
// source: github.com/codelingo/codelingo/flows/codelingo/rewrite/rpc/rewrite.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	github.com/codelingo/codelingo/flows/codelingo/rewrite/rpc/rewrite.proto

It has these top-level messages:
	Request
	Hunk
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the files or directories to review.
type Request struct {
	Host             string   `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Hostname         string   `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	GenerateComments bool     `protobuf:"varint,3,opt,name=generateComments" json:"generateComments,omitempty"`
	Repo             string   `protobuf:"bytes,4,opt,name=repo" json:"repo,omitempty"`
	Sha              string   `protobuf:"bytes,5,opt,name=sha" json:"sha,omitempty"`
	Patches          []string `protobuf:"bytes,6,rep,name=Patches" json:"Patches,omitempty"`
	IsPullRequest    bool     `protobuf:"varint,7,opt,name=isPullRequest" json:"isPullRequest,omitempty"`
	PullRequestID    int64    `protobuf:"varint,8,opt,name=pullRequestID" json:"pullRequestID,omitempty"`
	Vcs              string   `protobuf:"bytes,9,opt,name=vcs" json:"vcs,omitempty"`
	Dotlingo         string   `protobuf:"bytes,10,opt,name=dotlingo" json:"dotlingo,omitempty"`
	Dir              string   `protobuf:"bytes,11,opt,name=dir" json:"dir,omitempty"`
	// Types that are valid to be assigned to OwnerOrDepot:
	//	*Request_Owner
	//	*Request_Depot
	OwnerOrDepot isRequest_OwnerOrDepot `protobuf_oneof:"ownerOrDepot"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isRequest_OwnerOrDepot interface {
	isRequest_OwnerOrDepot()
}

type Request_Owner struct {
	Owner string `protobuf:"bytes,12,opt,name=owner,oneof"`
}
type Request_Depot struct {
	Depot string `protobuf:"bytes,13,opt,name=depot,oneof"`
}

func (*Request_Owner) isRequest_OwnerOrDepot() {}
func (*Request_Depot) isRequest_OwnerOrDepot() {}

func (m *Request) GetOwnerOrDepot() isRequest_OwnerOrDepot {
	if m != nil {
		return m.OwnerOrDepot
	}
	return nil
}

func (m *Request) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Request) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Request) GetGenerateComments() bool {
	if m != nil {
		return m.GenerateComments
	}
	return false
}

func (m *Request) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *Request) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func (m *Request) GetPatches() []string {
	if m != nil {
		return m.Patches
	}
	return nil
}

func (m *Request) GetIsPullRequest() bool {
	if m != nil {
		return m.IsPullRequest
	}
	return false
}

func (m *Request) GetPullRequestID() int64 {
	if m != nil {
		return m.PullRequestID
	}
	return 0
}

func (m *Request) GetVcs() string {
	if m != nil {
		return m.Vcs
	}
	return ""
}

func (m *Request) GetDotlingo() string {
	if m != nil {
		return m.Dotlingo
	}
	return ""
}

func (m *Request) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

func (m *Request) GetOwner() string {
	if x, ok := m.GetOwnerOrDepot().(*Request_Owner); ok {
		return x.Owner
	}
	return ""
}

func (m *Request) GetDepot() string {
	if x, ok := m.GetOwnerOrDepot().(*Request_Depot); ok {
		return x.Depot
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Request) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Request_OneofMarshaler, _Request_OneofUnmarshaler, _Request_OneofSizer, []interface{}{
		(*Request_Owner)(nil),
		(*Request_Depot)(nil),
	}
}

func _Request_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Request)
	// ownerOrDepot
	switch x := m.OwnerOrDepot.(type) {
	case *Request_Owner:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Owner)
	case *Request_Depot:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Depot)
	case nil:
	default:
		return fmt.Errorf("Request.OwnerOrDepot has unexpected type %T", x)
	}
	return nil
}

func _Request_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Request)
	switch tag {
	case 12: // ownerOrDepot.owner
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OwnerOrDepot = &Request_Owner{x}
		return true, err
	case 13: // ownerOrDepot.depot
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OwnerOrDepot = &Request_Depot{x}
		return true, err
	default:
		return false, nil
	}
}

func _Request_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Request)
	// ownerOrDepot
	switch x := m.OwnerOrDepot.(type) {
	case *Request_Owner:
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Owner)))
		n += len(x.Owner)
	case *Request_Depot:
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Depot)))
		n += len(x.Depot)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Hunk struct {
	Filename         string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	StartOffset      int32  `protobuf:"varint,2,opt,name=startOffset" json:"startOffset,omitempty"`
	EndOffset        int32  `protobuf:"varint,3,opt,name=endOffset" json:"endOffset,omitempty"`
	SRC              string `protobuf:"bytes,4,opt,name=SRC" json:"SRC,omitempty"`
	DecoratorOptions string `protobuf:"bytes,5,opt,name=decoratorOptions" json:"decoratorOptions,omitempty"`
}

func (m *Hunk) Reset()                    { *m = Hunk{} }
func (m *Hunk) String() string            { return proto.CompactTextString(m) }
func (*Hunk) ProtoMessage()               {}
func (*Hunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Hunk) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Hunk) GetStartOffset() int32 {
	if m != nil {
		return m.StartOffset
	}
	return 0
}

func (m *Hunk) GetEndOffset() int32 {
	if m != nil {
		return m.EndOffset
	}
	return 0
}

func (m *Hunk) GetSRC() string {
	if m != nil {
		return m.SRC
	}
	return ""
}

func (m *Hunk) GetDecoratorOptions() string {
	if m != nil {
		return m.DecoratorOptions
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "rpc.Request")
	proto.RegisterType((*Hunk)(nil), "rpc.Hunk")
}

func init() {
	proto.RegisterFile("github.com/codelingo/codelingo/flows/codelingo/rewrite/rpc/rewrite.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0x4d, 0x6e, 0xdb, 0x30,
	0x10, 0x85, 0x2b, 0xcb, 0x8e, 0x6d, 0xe6, 0x07, 0x01, 0x51, 0x14, 0x44, 0xd1, 0x85, 0x10, 0x74,
	0x61, 0x74, 0xa1, 0x2c, 0x7a, 0x03, 0x3b, 0x0b, 0x17, 0x08, 0x60, 0x83, 0x5d, 0x74, 0xad, 0x90,
	0x63, 0x9b, 0xa8, 0xc4, 0x61, 0x49, 0x2a, 0xbe, 0x4e, 0xd1, 0x1b, 0xf6, 0x06, 0xc1, 0x50, 0x92,
	0xe3, 0x20, 0x2b, 0xbd, 0xf7, 0x71, 0xc0, 0xe1, 0x3c, 0x0d, 0x5b, 0xef, 0x4d, 0x3c, 0xb4, 0x4f,
	0xa5, 0xc2, 0xe6, 0x5e, 0xa1, 0x86, 0xda, 0xd8, 0x3d, 0x9e, 0xa9, 0x5d, 0x8d, 0xc7, 0x70, 0xe6,
	0x3d, 0x1c, 0xbd, 0x89, 0x70, 0xef, 0x9d, 0x1a, 0x74, 0xe9, 0x3c, 0x46, 0xe4, 0xb9, 0x77, 0xea,
	0xee, 0xff, 0x88, 0x4d, 0x25, 0xfc, 0x69, 0x21, 0x44, 0xce, 0xd9, 0xf8, 0x80, 0x21, 0x8a, 0xac,
	0xc8, 0x16, 0x73, 0x99, 0x34, 0xff, 0xcc, 0x66, 0xf4, 0xb5, 0x55, 0x03, 0x62, 0x94, 0xf8, 0xc9,
	0xf3, 0x6f, 0xec, 0x76, 0x0f, 0x16, 0x7c, 0x15, 0x61, 0x85, 0x4d, 0x03, 0x36, 0x06, 0x91, 0x17,
	0xd9, 0x62, 0x26, 0xdf, 0x71, 0xba, 0xdb, 0x83, 0x43, 0x31, 0xee, 0xee, 0x26, 0xcd, 0x6f, 0x59,
	0x1e, 0x0e, 0x95, 0x98, 0x24, 0x44, 0x92, 0x0b, 0x36, 0xdd, 0x56, 0x51, 0x1d, 0x20, 0x88, 0x8b,
	0x22, 0x5f, 0xcc, 0xe5, 0x60, 0xf9, 0x57, 0x76, 0x6d, 0xc2, 0xb6, 0xad, 0xeb, 0xfe, 0xb1, 0x62,
	0x9a, 0x1a, 0xbd, 0x85, 0x54, 0xe5, 0x5e, 0xed, 0x8f, 0x07, 0x31, 0x2b, 0xb2, 0x45, 0x2e, 0xdf,
	0x42, 0xea, 0xfb, 0xac, 0x82, 0x98, 0x77, 0x7d, 0x9f, 0x55, 0xa0, 0x29, 0x35, 0xc6, 0x14, 0x98,
	0x60, 0xdd, 0x94, 0x83, 0xa7, 0x6a, 0x6d, 0xbc, 0xb8, 0xec, 0xaa, 0xb5, 0xf1, 0xfc, 0x13, 0x9b,
	0xe0, 0xd1, 0x82, 0x17, 0x57, 0xc4, 0xd6, 0x1f, 0x64, 0x67, 0x89, 0x6b, 0x70, 0x18, 0xc5, 0xf5,
	0xc0, 0x93, 0x5d, 0xde, 0xb0, 0xab, 0x54, 0xb0, 0xf1, 0x0f, 0xe4, 0xef, 0xfe, 0x66, 0x6c, 0xbc,
	0x6e, 0xed, 0x6f, 0x6a, 0xbb, 0x33, 0x35, 0xa4, 0x70, 0xbb, 0xd0, 0x4f, 0x9e, 0x17, 0xec, 0x32,
	0xc4, 0xca, 0xc7, 0xcd, 0x6e, 0x17, 0x20, 0xa6, 0xec, 0x27, 0xf2, 0x1c, 0xf1, 0x2f, 0x6c, 0x0e,
	0x56, 0xf7, 0xe7, 0x79, 0x3a, 0x7f, 0x05, 0xf4, 0xec, 0x9f, 0x72, 0xd5, 0xe7, 0x4d, 0x92, 0x7e,
	0x97, 0x06, 0x85, 0xbe, 0x8a, 0xe8, 0x37, 0x2e, 0x1a, 0xb4, 0xa1, 0xcf, 0xfe, 0x1d, 0x5f, 0x96,
	0xec, 0xa3, 0xc1, 0xf2, 0xb4, 0x45, 0xe5, 0xde, 0x3b, 0x55, 0xaa, 0x7a, 0x79, 0xb3, 0x42, 0x0d,
	0x8f, 0x84, 0xb6, 0xb4, 0x43, 0xdb, 0xec, 0xdf, 0x28, 0x5f, 0x3d, 0xfe, 0x7a, 0xba, 0x48, 0x2b,
	0xf5, 0xfd, 0x25, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xc9, 0xd2, 0x93, 0x9e, 0x02, 0x00, 0x00,
}
