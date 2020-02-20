// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/common/fault/v3/fault.proto

package envoy_extensions_filters_common_fault_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/datawire/ambassador/pkg/api/envoy/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FaultDelay_FaultDelayType int32

const (
	// Unused and deprecated.
	FaultDelay_FIXED FaultDelay_FaultDelayType = 0
)

var FaultDelay_FaultDelayType_name = map[int32]string{
	0: "FIXED",
}

var FaultDelay_FaultDelayType_value = map[string]int32{
	"FIXED": 0,
}

func (x FaultDelay_FaultDelayType) String() string {
	return proto.EnumName(FaultDelay_FaultDelayType_name, int32(x))
}

func (FaultDelay_FaultDelayType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{0, 0}
}

// Delay specification is used to inject latency into the
// HTTP/gRPC/Mongo/Redis operation or delay proxying of TCP connections.
// [#next-free-field: 6]
type FaultDelay struct {
	// Types that are valid to be assigned to FaultDelaySecifier:
	//	*FaultDelay_FixedDelay
	//	*FaultDelay_HeaderDelay_
	FaultDelaySecifier isFaultDelay_FaultDelaySecifier `protobuf_oneof:"fault_delay_secifier"`
	// The percentage of operations/connections/requests on which the delay will be injected.
	Percentage           *v3.FractionalPercent `protobuf:"bytes,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FaultDelay) Reset()         { *m = FaultDelay{} }
func (m *FaultDelay) String() string { return proto.CompactTextString(m) }
func (*FaultDelay) ProtoMessage()    {}
func (*FaultDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{0}
}
func (m *FaultDelay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FaultDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FaultDelay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FaultDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultDelay.Merge(m, src)
}
func (m *FaultDelay) XXX_Size() int {
	return m.Size()
}
func (m *FaultDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultDelay.DiscardUnknown(m)
}

var xxx_messageInfo_FaultDelay proto.InternalMessageInfo

type isFaultDelay_FaultDelaySecifier interface {
	isFaultDelay_FaultDelaySecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FaultDelay_FixedDelay struct {
	FixedDelay *types.Duration `protobuf:"bytes,3,opt,name=fixed_delay,json=fixedDelay,proto3,oneof"`
}
type FaultDelay_HeaderDelay_ struct {
	HeaderDelay *FaultDelay_HeaderDelay `protobuf:"bytes,5,opt,name=header_delay,json=headerDelay,proto3,oneof"`
}

func (*FaultDelay_FixedDelay) isFaultDelay_FaultDelaySecifier()   {}
func (*FaultDelay_HeaderDelay_) isFaultDelay_FaultDelaySecifier() {}

func (m *FaultDelay) GetFaultDelaySecifier() isFaultDelay_FaultDelaySecifier {
	if m != nil {
		return m.FaultDelaySecifier
	}
	return nil
}

func (m *FaultDelay) GetFixedDelay() *types.Duration {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

func (m *FaultDelay) GetHeaderDelay() *FaultDelay_HeaderDelay {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_HeaderDelay_); ok {
		return x.HeaderDelay
	}
	return nil
}

func (m *FaultDelay) GetPercentage() *v3.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultDelay) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultDelay_FixedDelay)(nil),
		(*FaultDelay_HeaderDelay_)(nil),
	}
}

// Fault delays are controlled via an HTTP header (if applicable). See the
// :ref:`http fault filter <config_http_filters_fault_injection_http_header>` documentation for
// more information.
type FaultDelay_HeaderDelay struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultDelay_HeaderDelay) Reset()         { *m = FaultDelay_HeaderDelay{} }
func (m *FaultDelay_HeaderDelay) String() string { return proto.CompactTextString(m) }
func (*FaultDelay_HeaderDelay) ProtoMessage()    {}
func (*FaultDelay_HeaderDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{0, 0}
}
func (m *FaultDelay_HeaderDelay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FaultDelay_HeaderDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FaultDelay_HeaderDelay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FaultDelay_HeaderDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultDelay_HeaderDelay.Merge(m, src)
}
func (m *FaultDelay_HeaderDelay) XXX_Size() int {
	return m.Size()
}
func (m *FaultDelay_HeaderDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultDelay_HeaderDelay.DiscardUnknown(m)
}

var xxx_messageInfo_FaultDelay_HeaderDelay proto.InternalMessageInfo

// Describes a rate limit to be applied.
type FaultRateLimit struct {
	// Types that are valid to be assigned to LimitType:
	//	*FaultRateLimit_FixedLimit_
	//	*FaultRateLimit_HeaderLimit_
	LimitType isFaultRateLimit_LimitType `protobuf_oneof:"limit_type"`
	// The percentage of operations/connections/requests on which the rate limit will be injected.
	Percentage           *v3.FractionalPercent `protobuf:"bytes,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FaultRateLimit) Reset()         { *m = FaultRateLimit{} }
func (m *FaultRateLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit) ProtoMessage()    {}
func (*FaultRateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{1}
}
func (m *FaultRateLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FaultRateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FaultRateLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FaultRateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit.Merge(m, src)
}
func (m *FaultRateLimit) XXX_Size() int {
	return m.Size()
}
func (m *FaultRateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit proto.InternalMessageInfo

type isFaultRateLimit_LimitType interface {
	isFaultRateLimit_LimitType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FaultRateLimit_FixedLimit_ struct {
	FixedLimit *FaultRateLimit_FixedLimit `protobuf:"bytes,1,opt,name=fixed_limit,json=fixedLimit,proto3,oneof"`
}
type FaultRateLimit_HeaderLimit_ struct {
	HeaderLimit *FaultRateLimit_HeaderLimit `protobuf:"bytes,3,opt,name=header_limit,json=headerLimit,proto3,oneof"`
}

func (*FaultRateLimit_FixedLimit_) isFaultRateLimit_LimitType()  {}
func (*FaultRateLimit_HeaderLimit_) isFaultRateLimit_LimitType() {}

func (m *FaultRateLimit) GetLimitType() isFaultRateLimit_LimitType {
	if m != nil {
		return m.LimitType
	}
	return nil
}

func (m *FaultRateLimit) GetFixedLimit() *FaultRateLimit_FixedLimit {
	if x, ok := m.GetLimitType().(*FaultRateLimit_FixedLimit_); ok {
		return x.FixedLimit
	}
	return nil
}

func (m *FaultRateLimit) GetHeaderLimit() *FaultRateLimit_HeaderLimit {
	if x, ok := m.GetLimitType().(*FaultRateLimit_HeaderLimit_); ok {
		return x.HeaderLimit
	}
	return nil
}

func (m *FaultRateLimit) GetPercentage() *v3.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultRateLimit) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultRateLimit_FixedLimit_)(nil),
		(*FaultRateLimit_HeaderLimit_)(nil),
	}
}

// Describes a fixed/constant rate limit.
type FaultRateLimit_FixedLimit struct {
	// The limit supplied in KiB/s.
	LimitKbps            uint64   `protobuf:"varint,1,opt,name=limit_kbps,json=limitKbps,proto3" json:"limit_kbps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultRateLimit_FixedLimit) Reset()         { *m = FaultRateLimit_FixedLimit{} }
func (m *FaultRateLimit_FixedLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit_FixedLimit) ProtoMessage()    {}
func (*FaultRateLimit_FixedLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{1, 0}
}
func (m *FaultRateLimit_FixedLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FaultRateLimit_FixedLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FaultRateLimit_FixedLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FaultRateLimit_FixedLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit_FixedLimit.Merge(m, src)
}
func (m *FaultRateLimit_FixedLimit) XXX_Size() int {
	return m.Size()
}
func (m *FaultRateLimit_FixedLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit_FixedLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit_FixedLimit proto.InternalMessageInfo

func (m *FaultRateLimit_FixedLimit) GetLimitKbps() uint64 {
	if m != nil {
		return m.LimitKbps
	}
	return 0
}

// Rate limits are controlled via an HTTP header (if applicable). See the
// :ref:`http fault filter <config_http_filters_fault_injection_http_header>` documentation for
// more information.
type FaultRateLimit_HeaderLimit struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultRateLimit_HeaderLimit) Reset()         { *m = FaultRateLimit_HeaderLimit{} }
func (m *FaultRateLimit_HeaderLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit_HeaderLimit) ProtoMessage()    {}
func (*FaultRateLimit_HeaderLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_99c2ca93a0e6a448, []int{1, 1}
}
func (m *FaultRateLimit_HeaderLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FaultRateLimit_HeaderLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FaultRateLimit_HeaderLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit_HeaderLimit.Merge(m, src)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Size() int {
	return m.Size()
}
func (m *FaultRateLimit_HeaderLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit_HeaderLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit_HeaderLimit proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.extensions.filters.common.fault.v3.FaultDelay_FaultDelayType", FaultDelay_FaultDelayType_name, FaultDelay_FaultDelayType_value)
	proto.RegisterType((*FaultDelay)(nil), "envoy.extensions.filters.common.fault.v3.FaultDelay")
	proto.RegisterType((*FaultDelay_HeaderDelay)(nil), "envoy.extensions.filters.common.fault.v3.FaultDelay.HeaderDelay")
	proto.RegisterType((*FaultRateLimit)(nil), "envoy.extensions.filters.common.fault.v3.FaultRateLimit")
	proto.RegisterType((*FaultRateLimit_FixedLimit)(nil), "envoy.extensions.filters.common.fault.v3.FaultRateLimit.FixedLimit")
	proto.RegisterType((*FaultRateLimit_HeaderLimit)(nil), "envoy.extensions.filters.common.fault.v3.FaultRateLimit.HeaderLimit")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/common/fault/v3/fault.proto", fileDescriptor_99c2ca93a0e6a448)
}

var fileDescriptor_99c2ca93a0e6a448 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x9b, 0x36, 0xd5, 0xee, 0xa9, 0x48, 0x0d, 0x82, 0xb5, 0x85, 0xb2, 0x56, 0xd0, 0xaa,
	0x38, 0x81, 0x56, 0x57, 0x2d, 0x28, 0x4b, 0xac, 0xa5, 0xbb, 0x2a, 0x94, 0xe0, 0x85, 0x77, 0x65,
	0x9a, 0x9c, 0x74, 0x07, 0xd3, 0x4c, 0x48, 0xd2, 0xd2, 0xbe, 0x81, 0xcf, 0xe0, 0x23, 0x78, 0xe5,
	0x43, 0x28, 0x78, 0xe9, 0x23, 0x48, 0x9f, 0x42, 0xf6, 0x4a, 0x66, 0x26, 0x69, 0x2b, 0x28, 0xb6,
	0x7b, 0x37, 0x27, 0x33, 0xe7, 0x3b, 0x7f, 0xfe, 0x7f, 0x18, 0x78, 0x84, 0xc1, 0x9c, 0x2f, 0x4d,
	0x5c, 0x24, 0x18, 0xc4, 0x8c, 0x07, 0xb1, 0xe9, 0x31, 0x3f, 0xc1, 0x28, 0x36, 0x1d, 0x3e, 0x9d,
	0xf2, 0xc0, 0xf4, 0xe8, 0xcc, 0x4f, 0xcc, 0x79, 0x47, 0x2d, 0x48, 0x18, 0xf1, 0x84, 0x1b, 0x2d,
	0xd9, 0x45, 0x36, 0x5d, 0x24, 0xed, 0x22, 0xaa, 0x8b, 0xa8, 0xc3, 0xf3, 0x4e, 0xad, 0xae, 0xf8,
	0xc9, 0x32, 0x44, 0x01, 0x09, 0x31, 0x72, 0x30, 0x48, 0x31, 0xb5, 0xc6, 0x84, 0xf3, 0x89, 0x8f,
	0xa6, 0xac, 0xc6, 0x33, 0xcf, 0x74, 0x67, 0x11, 0x4d, 0x18, 0x0f, 0xd2, 0xfd, 0x5b, 0x33, 0x37,
	0xa4, 0x26, 0x0d, 0x02, 0x9e, 0xc8, 0xcf, 0xb1, 0x39, 0xc7, 0x48, 0xcc, 0x63, 0xc1, 0x24, 0x3d,
	0x72, 0x5b, 0xf1, 0xb7, 0xcf, 0xb8, 0x18, 0x46, 0xe8, 0x6c, 0x73, 0x6e, 0xcc, 0xa9, 0xcf, 0x5c,
	0x9a, 0xa0, 0x99, 0x2d, 0xd4, 0x46, 0xf3, 0x6b, 0x01, 0xa0, 0x2f, 0xa4, 0xf6, 0xd0, 0xa7, 0x4b,
	0xa3, 0x0f, 0x65, 0x8f, 0x2d, 0xd0, 0x1d, 0xb9, 0xa2, 0xac, 0x16, 0x0e, 0xb5, 0x56, 0xb9, 0x7d,
	0x93, 0x28, 0x95, 0x24, 0x53, 0x49, 0x7a, 0xa9, 0x4a, 0xab, 0x74, 0x6e, 0x15, 0x3f, 0x6b, 0xf9,
	0xfb, 0xb9, 0x41, 0xce, 0x06, 0xd9, 0xa9, 0x38, 0x08, 0x57, 0xce, 0x90, 0xba, 0x18, 0xa5, 0xa0,
	0xa2, 0x04, 0x1d, 0x93, 0x5d, 0x5d, 0x23, 0x1b, 0x4d, 0x64, 0x20, 0x41, 0x72, 0x3d, 0xc8, 0xd9,
	0xe5, 0xb3, 0x4d, 0x69, 0x1c, 0x03, 0xa4, 0x7e, 0xd2, 0x09, 0x56, 0x75, 0x39, 0xe4, 0x30, 0x1d,
	0x22, 0x0c, 0x97, 0xa4, 0x88, 0x3a, 0x42, 0x2b, 0xf5, 0x87, 0xea, 0xa8, 0xbd, 0xd5, 0x53, 0x3b,
	0x81, 0xf2, 0x16, 0xbf, 0xdb, 0xfd, 0xf4, 0xed, 0x63, 0xe3, 0x31, 0x74, 0x14, 0xc2, 0xe1, 0x81,
	0xc7, 0x26, 0xa9, 0xc6, 0x4c, 0x5b, 0xfb, 0x1f, 0xda, 0x9a, 0x75, 0xb8, 0xba, 0xd9, 0x79, 0xb7,
	0x0c, 0xd1, 0x38, 0x80, 0x62, 0xff, 0xe4, 0xfd, 0xab, 0x5e, 0x25, 0xd7, 0x25, 0x02, 0x7c, 0x0f,
	0xee, 0xee, 0x08, 0xb6, 0xea, 0x70, 0x5d, 0x7e, 0x56, 0xfe, 0x8d, 0x62, 0x74, 0x98, 0xc7, 0x30,
	0x32, 0x0a, 0xbf, 0x2c, 0xed, 0x54, 0x2f, 0xe5, 0x2b, 0x85, 0x53, 0xbd, 0xa4, 0x55, 0xf2, 0xb6,
	0x2e, 0xfe, 0xb3, 0xf9, 0x45, 0x4f, 0x87, 0xdb, 0x34, 0xc1, 0x37, 0x6c, 0xca, 0x12, 0xc3, 0xcb,
	0xa2, 0xf4, 0x45, 0x59, 0xd5, 0xa4, 0x39, 0x2f, 0xf7, 0x4c, 0x60, 0x8d, 0x23, 0x7d, 0xc1, 0x92,
	0xcb, 0x75, 0xd4, 0x6a, 0x0e, 0x5b, 0x47, 0xad, 0x06, 0xa9, 0x3b, 0xd3, 0xbb, 0xf0, 0x20, 0x65,
	0x69, 0x36, 0x29, 0x8d, 0x5b, 0x8d, 0xfa, 0x33, 0xee, 0xfc, 0x05, 0xe2, 0x8e, 0x01, 0x36, 0x3f,
	0x62, 0xdc, 0x01, 0x90, 0x9a, 0x47, 0x1f, 0xc6, 0x61, 0x2c, 0x1d, 0xd2, 0xad, 0xcb, 0xe7, 0x96,
	0xde, 0xce, 0xb7, 0x34, 0xfb, 0x40, 0x6e, 0xbd, 0x1e, 0x87, 0x71, 0xf7, 0xb9, 0x08, 0xef, 0x29,
	0x1c, 0xfd, 0x3f, 0xbc, 0xbf, 0xf9, 0x55, 0x7b, 0x9b, 0xdd, 0x31, 0x59, 0x76, 0x5f, 0x08, 0xda,
	0x33, 0x78, 0xb2, 0x0f, 0x6d, 0xbb, 0xbf, 0x2d, 0xfa, 0x1f, 0xc2, 0x83, 0x3d, 0xfa, 0xad, 0x6b,
	0xd9, 0x9f, 0x0a, 0x9b, 0xe4, 0x25, 0xb2, 0xec, 0xef, 0xab, 0x86, 0xf6, 0x63, 0xd5, 0xd0, 0x7e,
	0xae, 0x1a, 0x1a, 0x1c, 0x31, 0xae, 0x8c, 0x0c, 0x23, 0xbe, 0x58, 0xee, 0x1c, 0x9e, 0xa5, 0x1e,
	0x8f, 0xa1, 0x78, 0x18, 0x86, 0xda, 0xf8, 0x92, 0x7c, 0x21, 0x3a, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xa9, 0x50, 0x8e, 0x14, 0x54, 0x05, 0x00, 0x00,
}

func (m *FaultDelay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultDelay) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FaultDelay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.FaultDelaySecifier != nil {
		{
			size := m.FaultDelaySecifier.Size()
			i -= size
			if _, err := m.FaultDelaySecifier.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Percentage != nil {
		{
			size, err := m.Percentage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}

func (m *FaultDelay_FixedDelay) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *FaultDelay_FixedDelay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FixedDelay != nil {
		{
			size, err := m.FixedDelay.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *FaultDelay_HeaderDelay_) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *FaultDelay_HeaderDelay_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HeaderDelay != nil {
		{
			size, err := m.HeaderDelay.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *FaultDelay_HeaderDelay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultDelay_HeaderDelay) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FaultDelay_HeaderDelay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *FaultRateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultRateLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FaultRateLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LimitType != nil {
		{
			size := m.LimitType.Size()
			i -= size
			if _, err := m.LimitType.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Percentage != nil {
		{
			size, err := m.Percentage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *FaultRateLimit_FixedLimit_) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *FaultRateLimit_FixedLimit_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FixedLimit != nil {
		{
			size, err := m.FixedLimit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *FaultRateLimit_HeaderLimit_) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *FaultRateLimit_HeaderLimit_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HeaderLimit != nil {
		{
			size, err := m.HeaderLimit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFault(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *FaultRateLimit_FixedLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultRateLimit_FixedLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FaultRateLimit_FixedLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LimitKbps != 0 {
		i = encodeVarintFault(dAtA, i, uint64(m.LimitKbps))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FaultRateLimit_HeaderLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultRateLimit_HeaderLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FaultRateLimit_HeaderLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintFault(dAtA []byte, offset int, v uint64) int {
	offset -= sovFault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FaultDelay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FaultDelaySecifier != nil {
		n += m.FaultDelaySecifier.Size()
	}
	if m.Percentage != nil {
		l = m.Percentage.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FaultDelay_FixedDelay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FixedDelay != nil {
		l = m.FixedDelay.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	return n
}
func (m *FaultDelay_HeaderDelay_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HeaderDelay != nil {
		l = m.HeaderDelay.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	return n
}
func (m *FaultDelay_HeaderDelay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FaultRateLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LimitType != nil {
		n += m.LimitType.Size()
	}
	if m.Percentage != nil {
		l = m.Percentage.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FaultRateLimit_FixedLimit_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FixedLimit != nil {
		l = m.FixedLimit.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	return n
}
func (m *FaultRateLimit_HeaderLimit_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HeaderLimit != nil {
		l = m.HeaderLimit.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	return n
}
func (m *FaultRateLimit_FixedLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LimitKbps != 0 {
		n += 1 + sovFault(uint64(m.LimitKbps))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FaultRateLimit_HeaderLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFault(x uint64) (n int) {
	return sovFault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FaultDelay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FaultDelay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FaultDelay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedDelay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.Duration{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.FaultDelaySecifier = &FaultDelay_FixedDelay{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percentage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Percentage == nil {
				m.Percentage = &v3.FractionalPercent{}
			}
			if err := m.Percentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderDelay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FaultDelay_HeaderDelay{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.FaultDelaySecifier = &FaultDelay_HeaderDelay_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FaultDelay_HeaderDelay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeaderDelay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeaderDelay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FaultRateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FaultRateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FaultRateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FaultRateLimit_FixedLimit{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.LimitType = &FaultRateLimit_FixedLimit_{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percentage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Percentage == nil {
				m.Percentage = &v3.FractionalPercent{}
			}
			if err := m.Percentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FaultRateLimit_HeaderLimit{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.LimitType = &FaultRateLimit_HeaderLimit_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FaultRateLimit_FixedLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FixedLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FixedLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitKbps", wireType)
			}
			m.LimitKbps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LimitKbps |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FaultRateLimit_HeaderLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeaderLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeaderLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFault
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFault
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFault
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFault
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFault
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFault
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFault(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFault
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFault = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFault   = fmt.Errorf("proto: integer overflow")
)