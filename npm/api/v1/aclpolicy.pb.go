// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aclpolicy.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Verdict int32

const (
	Verdict_ALLOW Verdict = 0
	Verdict_DENY  Verdict = 1
)

var Verdict_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}

var Verdict_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x Verdict) String() string {
	return proto.EnumName(Verdict_name, int32(x))
}

func (Verdict) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_867e1063938b2946, []int{0}
}

type Direction int32

const (
	Direction_INGRESS Direction = 0
	Direction_EGRESS  Direction = 1
	Direction_BOTH    Direction = 2
)

var Direction_name = map[int32]string{
	0: "INGRESS",
	1: "EGRESS",
	2: "BOTH",
}

var Direction_value = map[string]int32{
	"INGRESS": 0,
	"EGRESS":  1,
	"BOTH":    2,
}

func (x Direction) String() string {
	return proto.EnumName(Direction_name, int32(x))
}

func (Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_867e1063938b2946, []int{1}
}

type AclPolicy struct {
	PolicyID             string     `protobuf:"bytes,1,opt,name=PolicyID,proto3" json:"PolicyID,omitempty"`
	PolicyName           string     `protobuf:"bytes,2,opt,name=PolicyName,proto3" json:"PolicyName,omitempty"`
	Comment              string     `protobuf:"bytes,3,opt,name=Comment,proto3" json:"Comment,omitempty"`
	SrcList              []*SetInfo `protobuf:"bytes,4,rep,name=SrcList,proto3" json:"SrcList,omitempty"`
	DstList              []*SetInfo `protobuf:"bytes,5,rep,name=DstList,proto3" json:"DstList,omitempty"`
	Target               Verdict    `protobuf:"varint,6,opt,name=Target,proto3,enum=api.Verdict" json:"Target,omitempty"`
	Direction            Direction  `protobuf:"varint,7,opt,name=Direction,proto3,enum=api.Direction" json:"Direction,omitempty"`
	Protocol             string     `protobuf:"bytes,8,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	SrcPortList          []*Ports   `protobuf:"bytes,9,rep,name=SrcPortList,proto3" json:"SrcPortList,omitempty"`
	DstPortList          []*Ports   `protobuf:"bytes,10,rep,name=DstPortList,proto3" json:"DstPortList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AclPolicy) Reset()         { *m = AclPolicy{} }
func (m *AclPolicy) String() string { return proto.CompactTextString(m) }
func (*AclPolicy) ProtoMessage()    {}
func (*AclPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_867e1063938b2946, []int{0}
}

func (m *AclPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AclPolicy.Unmarshal(m, b)
}
func (m *AclPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AclPolicy.Marshal(b, m, deterministic)
}
func (m *AclPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AclPolicy.Merge(m, src)
}
func (m *AclPolicy) XXX_Size() int {
	return xxx_messageInfo_AclPolicy.Size(m)
}
func (m *AclPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_AclPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_AclPolicy proto.InternalMessageInfo

func (m *AclPolicy) GetPolicyID() string {
	if m != nil {
		return m.PolicyID
	}
	return ""
}

func (m *AclPolicy) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *AclPolicy) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AclPolicy) GetSrcList() []*SetInfo {
	if m != nil {
		return m.SrcList
	}
	return nil
}

func (m *AclPolicy) GetDstList() []*SetInfo {
	if m != nil {
		return m.DstList
	}
	return nil
}

func (m *AclPolicy) GetTarget() Verdict {
	if m != nil {
		return m.Target
	}
	return Verdict_ALLOW
}

func (m *AclPolicy) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return Direction_INGRESS
}

func (m *AclPolicy) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *AclPolicy) GetSrcPortList() []*Ports {
	if m != nil {
		return m.SrcPortList
	}
	return nil
}

func (m *AclPolicy) GetDstPortList() []*Ports {
	if m != nil {
		return m.DstPortList
	}
	return nil
}

type SetInfo struct {
	IPSet                *IPSet   `protobuf:"bytes,1,opt,name=IPSet,proto3" json:"IPSet,omitempty"`
	Included             bool     `protobuf:"varint,2,opt,name=Included,proto3" json:"Included,omitempty"`
	MatchType            string   `protobuf:"bytes,3,opt,name=MatchType,proto3" json:"MatchType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetInfo) Reset()         { *m = SetInfo{} }
func (m *SetInfo) String() string { return proto.CompactTextString(m) }
func (*SetInfo) ProtoMessage()    {}
func (*SetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_867e1063938b2946, []int{1}
}

func (m *SetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetInfo.Unmarshal(m, b)
}
func (m *SetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetInfo.Marshal(b, m, deterministic)
}
func (m *SetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetInfo.Merge(m, src)
}
func (m *SetInfo) XXX_Size() int {
	return xxx_messageInfo_SetInfo.Size(m)
}
func (m *SetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SetInfo proto.InternalMessageInfo

func (m *SetInfo) GetIPSet() *IPSet {
	if m != nil {
		return m.IPSet
	}
	return nil
}

func (m *SetInfo) GetIncluded() bool {
	if m != nil {
		return m.Included
	}
	return false
}

func (m *SetInfo) GetMatchType() string {
	if m != nil {
		return m.MatchType
	}
	return ""
}

type Ports struct {
	Port                 int64    `protobuf:"varint,1,opt,name=Port,proto3" json:"Port,omitempty"`
	EndPort              int64    `protobuf:"varint,2,opt,name=EndPort,proto3" json:"EndPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ports) Reset()         { *m = Ports{} }
func (m *Ports) String() string { return proto.CompactTextString(m) }
func (*Ports) ProtoMessage()    {}
func (*Ports) Descriptor() ([]byte, []int) {
	return fileDescriptor_867e1063938b2946, []int{2}
}

func (m *Ports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ports.Unmarshal(m, b)
}
func (m *Ports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ports.Marshal(b, m, deterministic)
}
func (m *Ports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ports.Merge(m, src)
}
func (m *Ports) XXX_Size() int {
	return xxx_messageInfo_Ports.Size(m)
}
func (m *Ports) XXX_DiscardUnknown() {
	xxx_messageInfo_Ports.DiscardUnknown(m)
}

var xxx_messageInfo_Ports proto.InternalMessageInfo

func (m *Ports) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Ports) GetEndPort() int64 {
	if m != nil {
		return m.EndPort
	}
	return 0
}

func init() {
	proto.RegisterEnum("api.Verdict", Verdict_name, Verdict_value)
	proto.RegisterEnum("api.Direction", Direction_name, Direction_value)
	proto.RegisterType((*AclPolicy)(nil), "api.AclPolicy")
	proto.RegisterType((*SetInfo)(nil), "api.SetInfo")
	proto.RegisterType((*Ports)(nil), "api.Ports")
}

func init() { proto.RegisterFile("aclpolicy.proto", fileDescriptor_867e1063938b2946) }

var fileDescriptor_867e1063938b2946 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x37, 0xfd, 0x97, 0xe6, 0x56, 0xd6, 0x70, 0x9f, 0x86, 0x45, 0x96, 0x50, 0x44, 0xca,
	0xb2, 0xe4, 0x41, 0xf1, 0x03, 0xac, 0x26, 0x68, 0xa0, 0x76, 0xcb, 0xa4, 0x28, 0x3e, 0xc6, 0xc9,
	0xa8, 0x03, 0x69, 0x26, 0x24, 0xe3, 0xc3, 0x7e, 0x78, 0x41, 0xe6, 0x66, 0x92, 0x06, 0xdc, 0xb7,
	0x7b, 0xce, 0xf9, 0xb5, 0x73, 0xff, 0x04, 0x5e, 0x16, 0xa2, 0x6a, 0x74, 0xa5, 0xc4, 0x53, 0xdc,
	0xb4, 0xda, 0x68, 0x9c, 0x17, 0x8d, 0xba, 0xd9, 0xa8, 0xa6, 0x93, 0xa6, 0x77, 0xb6, 0x7f, 0x67,
	0x10, 0x3c, 0x88, 0xea, 0x48, 0x14, 0xde, 0xc0, 0xba, 0xaf, 0xb2, 0x84, 0x79, 0x91, 0xb7, 0x0b,
	0xf8, 0xa8, 0xf1, 0x16, 0xa0, 0xaf, 0x0f, 0xc5, 0x59, 0xb2, 0x19, 0xa5, 0x13, 0x07, 0x19, 0xf8,
	0x1f, 0xf5, 0xf9, 0x2c, 0x6b, 0xc3, 0xe6, 0x14, 0x0e, 0x12, 0xdf, 0x80, 0x9f, 0xb7, 0x62, 0xaf,
	0x3a, 0xc3, 0x16, 0xd1, 0x7c, 0xb7, 0x79, 0xfb, 0x22, 0x2e, 0x1a, 0x15, 0xe7, 0xd2, 0x64, 0xf5,
	0x4f, 0xcd, 0x87, 0xd0, 0x72, 0x49, 0x67, 0x88, 0x5b, 0x3e, 0xc7, 0xb9, 0x10, 0x5f, 0xc3, 0xea,
	0x54, 0xb4, 0xbf, 0xa4, 0x61, 0xab, 0xc8, 0xdb, 0x5d, 0x3b, 0xec, 0xab, 0x6c, 0x4b, 0x25, 0x0c,
	0x77, 0x19, 0xde, 0x43, 0x90, 0xa8, 0x56, 0x0a, 0xa3, 0x74, 0xcd, 0x7c, 0x02, 0xaf, 0x09, 0x1c,
	0x5d, 0x7e, 0x01, 0x68, 0x72, 0xbb, 0x10, 0xa1, 0x2b, 0xb6, 0x76, 0x93, 0x3b, 0x8d, 0xf7, 0xb0,
	0xc9, 0x5b, 0x71, 0xd4, 0x6d, 0xdf, 0x5b, 0x40, 0xbd, 0x01, 0xfd, 0x97, 0x35, 0x3b, 0x3e, 0x8d,
	0x2d, 0x9d, 0x74, 0x66, 0xa4, 0xe1, 0x7f, 0x7a, 0x12, 0x6f, 0x25, 0xf8, 0x6e, 0x3e, 0x8c, 0x60,
	0x99, 0x1d, 0x73, 0x69, 0x68, 0xf3, 0xc3, 0x4f, 0xc8, 0xe1, 0x7d, 0x60, 0x9b, 0xcc, 0x6a, 0x51,
	0xfd, 0x29, 0x65, 0x49, 0x07, 0x58, 0xf3, 0x51, 0xe3, 0x2b, 0x08, 0xbe, 0x14, 0x46, 0xfc, 0x3e,
	0x3d, 0x35, 0xd2, 0x1d, 0xe0, 0x62, 0x6c, 0xdf, 0xc3, 0x92, 0x1e, 0x47, 0x84, 0x85, 0x2d, 0xe8,
	0x8d, 0x39, 0xa7, 0xda, 0x5e, 0x2e, 0xad, 0x4b, 0xb2, 0x67, 0x64, 0x0f, 0xf2, 0xee, 0x16, 0x7c,
	0xb7, 0x56, 0x0c, 0x60, 0xf9, 0xb0, 0xdf, 0x3f, 0x7e, 0x0b, 0xaf, 0x70, 0x0d, 0x8b, 0x24, 0x3d,
	0x7c, 0x0f, 0xbd, 0xbb, 0x78, 0xb2, 0x63, 0xdc, 0x80, 0x9f, 0x1d, 0x3e, 0xf1, 0x34, 0xcf, 0xc3,
	0x2b, 0x04, 0x58, 0xa5, 0x7d, 0xed, 0x59, 0xfe, 0xc3, 0xe3, 0xe9, 0x73, 0x38, 0xfb, 0xb1, 0xa2,
	0x8f, 0xee, 0xdd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x44, 0x6e, 0x0c, 0x99, 0x02, 0x00,
	0x00,
}
