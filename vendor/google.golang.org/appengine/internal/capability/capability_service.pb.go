// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google.golang.org/appengine/internal/capability/capability_service.proto

/*
Package capability is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/appengine/internal/capability/capability_service.proto

It has these top-level messages:
	IsEnabledRequest
	IsEnabledResponse
*/
package capability

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

type IsEnabledResponse_SummaryStatus int32

const (
	IsEnabledResponse_DEFAULT          IsEnabledResponse_SummaryStatus = 0
	IsEnabledResponse_ENABLED          IsEnabledResponse_SummaryStatus = 1
	IsEnabledResponse_SCHEDULED_FUTURE IsEnabledResponse_SummaryStatus = 2
	IsEnabledResponse_SCHEDULED_NOW    IsEnabledResponse_SummaryStatus = 3
	IsEnabledResponse_DISABLED         IsEnabledResponse_SummaryStatus = 4
	IsEnabledResponse_UNKNOWN          IsEnabledResponse_SummaryStatus = 5
)

var IsEnabledResponse_SummaryStatus_name = map[int32]string{
	0: "DEFAULT",
	1: "ENABLED",
	2: "SCHEDULED_FUTURE",
	3: "SCHEDULED_NOW",
	4: "DISABLED",
	5: "UNKNOWN",
}
var IsEnabledResponse_SummaryStatus_value = map[string]int32{
	"DEFAULT":          0,
	"ENABLED":          1,
	"SCHEDULED_FUTURE": 2,
	"SCHEDULED_NOW":    3,
	"DISABLED":         4,
	"UNKNOWN":          5,
}

func (x IsEnabledResponse_SummaryStatus) Enum() *IsEnabledResponse_SummaryStatus {
	p := new(IsEnabledResponse_SummaryStatus)
	*p = x
	return p
}
func (x IsEnabledResponse_SummaryStatus) String() string {
	return proto.EnumName(IsEnabledResponse_SummaryStatus_name, int32(x))
}
func (x *IsEnabledResponse_SummaryStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IsEnabledResponse_SummaryStatus_value, data, "IsEnabledResponse_SummaryStatus")
	if err != nil {
		return err
	}
	*x = IsEnabledResponse_SummaryStatus(value)
	return nil
}
func (IsEnabledResponse_SummaryStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type IsEnabledRequest struct {
	Package          *string  `protobuf:"bytes,1,req,name=package" json:"package,omitempty"`
	Capability       []string `protobuf:"bytes,2,rep,name=capability" json:"capability,omitempty"`
	Call             []string `protobuf:"bytes,3,rep,name=call" json:"call,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *IsEnabledRequest) Reset()                    { *m = IsEnabledRequest{} }
func (m *IsEnabledRequest) String() string            { return proto.CompactTextString(m) }
func (*IsEnabledRequest) ProtoMessage()               {}
func (*IsEnabledRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsEnabledRequest) GetPackage() string {
	if m != nil && m.Package != nil {
		return *m.Package
	}
	return ""
}

func (m *IsEnabledRequest) GetCapability() []string {
	if m != nil {
		return m.Capability
	}
	return nil
}

func (m *IsEnabledRequest) GetCall() []string {
	if m != nil {
		return m.Call
	}
	return nil
}

type IsEnabledResponse struct {
	SummaryStatus      *IsEnabledResponse_SummaryStatus `protobuf:"varint,1,opt,name=summary_status,json=summaryStatus,enum=appengine.IsEnabledResponse_SummaryStatus" json:"summary_status,omitempty"`
	TimeUntilScheduled *int64                           `protobuf:"varint,2,opt,name=time_until_scheduled,json=timeUntilScheduled" json:"time_until_scheduled,omitempty"`
	XXX_unrecognized   []byte                           `json:"-"`
}

func (m *IsEnabledResponse) Reset()                    { *m = IsEnabledResponse{} }
func (m *IsEnabledResponse) String() string            { return proto.CompactTextString(m) }
func (*IsEnabledResponse) ProtoMessage()               {}
func (*IsEnabledResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsEnabledResponse) GetSummaryStatus() IsEnabledResponse_SummaryStatus {
	if m != nil && m.SummaryStatus != nil {
		return *m.SummaryStatus
	}
	return IsEnabledResponse_DEFAULT
}

func (m *IsEnabledResponse) GetTimeUntilScheduled() int64 {
	if m != nil && m.TimeUntilScheduled != nil {
		return *m.TimeUntilScheduled
	}
	return 0
}

func init() {
	proto.RegisterType((*IsEnabledRequest)(nil), "appengine.IsEnabledRequest")
	proto.RegisterType((*IsEnabledResponse)(nil), "appengine.IsEnabledResponse")
}

func init() {
	proto.RegisterFile("google.golang.org/appengine/internal/capability/capability_service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x8a, 0x9b, 0x40,
	0x14, 0x86, 0xa3, 0xa6, 0xa4, 0x9e, 0x26, 0xc1, 0x0c, 0xb9, 0x90, 0xb6, 0x14, 0xf1, 0x4a, 0x7a,
	0x61, 0x4a, 0xde, 0x20, 0x89, 0x86, 0x84, 0x06, 0x43, 0x35, 0x12, 0x28, 0x14, 0x3b, 0x31, 0x83,
	0x95, 0x8e, 0xa3, 0xeb, 0x8c, 0x0b, 0x79, 0x82, 0x7d, 0xed, 0x45, 0x43, 0x8c, 0xcb, 0x2e, 0x7b,
	0x77, 0xce, 0xf9, 0xf9, 0xfe, 0x99, 0x73, 0x7e, 0xd8, 0x24, 0x79, 0x9e, 0x50, 0x62, 0x27, 0x39,
	0xc5, 0x2c, 0xb1, 0xf3, 0x32, 0x99, 0xe1, 0xa2, 0x20, 0x2c, 0x49, 0x19, 0x99, 0xa5, 0x4c, 0x90,
	0x92, 0x61, 0x3a, 0x8b, 0x71, 0x81, 0x4f, 0x29, 0x4d, 0xc5, 0xa5, 0x53, 0x46, 0x9c, 0x94, 0x8f,
	0x69, 0x4c, 0xec, 0xa2, 0xcc, 0x45, 0x8e, 0xd4, 0x96, 0x33, 0xff, 0x82, 0xb6, 0xe5, 0x2e, 0xc3,
	0x27, 0x4a, 0xce, 0x3e, 0x79, 0xa8, 0x08, 0x17, 0x48, 0x87, 0x41, 0x81, 0xe3, 0xff, 0x38, 0x21,
	0xba, 0x64, 0xc8, 0x96, 0xea, 0xdf, 0x5a, 0xf4, 0x0d, 0xe0, 0x6e, 0xaa, 0xcb, 0x86, 0x62, 0xa9,
	0x7e, 0x67, 0x82, 0x10, 0xf4, 0x63, 0x4c, 0xa9, 0xae, 0x34, 0x4a, 0x53, 0x9b, 0x4f, 0x32, 0x4c,
	0x3a, 0x4f, 0xf0, 0x22, 0x67, 0x9c, 0xa0, 0x5f, 0x30, 0xe6, 0x55, 0x96, 0xe1, 0xf2, 0x12, 0x71,
	0x81, 0x45, 0xc5, 0x75, 0xc9, 0x90, 0xac, 0xf1, 0xfc, 0xbb, 0xdd, 0xfe, 0xcd, 0x7e, 0x45, 0xd9,
	0xc1, 0x15, 0x09, 0x1a, 0xc2, 0x1f, 0xf1, 0x6e, 0x8b, 0x7e, 0xc0, 0x54, 0xa4, 0x19, 0x89, 0x2a,
	0x26, 0x52, 0x1a, 0xf1, 0xf8, 0x1f, 0x39, 0x57, 0x94, 0x9c, 0x75, 0xd9, 0x90, 0x2c, 0xc5, 0x47,
	0xb5, 0x16, 0xd6, 0x52, 0x70, 0x53, 0xcc, 0x0c, 0x46, 0x2f, 0x1c, 0xd1, 0x27, 0x18, 0x38, 0xee,
	0x7a, 0x11, 0xee, 0x0e, 0x5a, 0xaf, 0x6e, 0x5c, 0x6f, 0xb1, 0xdc, 0xb9, 0x8e, 0x26, 0xa1, 0x29,
	0x68, 0xc1, 0x6a, 0xe3, 0x3a, 0xe1, 0xce, 0x75, 0xa2, 0x75, 0x78, 0x08, 0x7d, 0x57, 0x93, 0xd1,
	0x04, 0x46, 0xf7, 0xa9, 0xb7, 0x3f, 0x6a, 0x0a, 0x1a, 0xc2, 0x47, 0x67, 0x1b, 0x5c, 0xb1, 0x7e,
	0xed, 0x11, 0x7a, 0x3f, 0xbd, 0xfd, 0xd1, 0xd3, 0x3e, 0xcc, 0xff, 0xc0, 0x64, 0xd5, 0xde, 0x2a,
	0xb8, 0x26, 0x82, 0x36, 0xa0, 0xb6, 0x7b, 0xa2, 0x2f, 0x6f, 0x6f, 0xdf, 0xc4, 0xf2, 0xf9, 0xeb,
	0x7b, 0xa7, 0x31, 0x7b, 0xcb, 0xe1, 0xef, 0x4e, 0x14, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0x03, 0x26, 0x25, 0x2e, 0x02, 0x00, 0x00,
}
