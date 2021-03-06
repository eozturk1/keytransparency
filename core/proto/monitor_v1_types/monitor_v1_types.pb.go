// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor_v1_types.proto

/*
Package monitor_v1_types is a generated protocol buffer package.

Key Transparency Monitor Service


It is generated from these files:
	monitor_v1_types.proto

It has these top-level messages:
	GetMonitoringRequest
	GetMonitoringResponse
*/
package monitor_v1_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keytransparency_v1_types "github.com/google/keytransparency/core/proto/keytransparency_v1_types"
import trillian "github.com/google/trillian"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// GetMonitoringRequest contains the input parameters of the GetMonitoring APIs.
type GetMonitoringRequest struct {
	// epoch specifies the for which the monitoring results will
	// be returned (epochs start at 1).
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	// ktURL is the URL of the keytransparency server for which the monitoring
	// result will be returned.
	Kt_URL string `protobuf:"bytes,2,opt,name=kt_URL,json=ktURL" json:"kt_URL,omitempty"`
}

func (m *GetMonitoringRequest) Reset()                    { *m = GetMonitoringRequest{} }
func (m *GetMonitoringRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMonitoringRequest) ProtoMessage()               {}
func (*GetMonitoringRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetMonitoringRequest) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *GetMonitoringRequest) GetKt_URL() string {
	if m != nil {
		return m.Kt_URL
	}
	return ""
}

type GetMonitoringResponse struct {
	// smr contains the map root for the sparse Merkle Tree signed with the
	// monitor's key on success. If the checks were not successful the
	// smr will be empty. The epochs are encoded into the smr map_revision.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,1,opt,name=smr" json:"smr,omitempty"`
	// seen_timestamp_nanos contains the time in nanoseconds where this particular
	// signed map root was retrieved and processed. The actual timestamp of the
	// smr returned by the server is contained in above smr field.
	SeenTimestampNanos int64 `protobuf:"varint,2,opt,name=seen_timestamp_nanos,json=seenTimestampNanos" json:"seen_timestamp_nanos,omitempty"`
	// errors contains a list of error string representing the verification checks
	// that failed while monitoring the key-transparency server.
	Errors []string `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	// error_data contains the original response from the mutations API if and
	// only if at least one verification step failed. It can be used to re-run the
	// verification steps.
	ErrorData *keytransparency_v1_types.GetMutationsResponse `protobuf:"bytes,4,opt,name=error_data,json=errorData" json:"error_data,omitempty"`
}

func (m *GetMonitoringResponse) Reset()                    { *m = GetMonitoringResponse{} }
func (m *GetMonitoringResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMonitoringResponse) ProtoMessage()               {}
func (*GetMonitoringResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetMonitoringResponse) GetSmr() *trillian.SignedMapRoot {
	if m != nil {
		return m.Smr
	}
	return nil
}

func (m *GetMonitoringResponse) GetSeenTimestampNanos() int64 {
	if m != nil {
		return m.SeenTimestampNanos
	}
	return 0
}

func (m *GetMonitoringResponse) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *GetMonitoringResponse) GetErrorData() *keytransparency_v1_types.GetMutationsResponse {
	if m != nil {
		return m.ErrorData
	}
	return nil
}

func init() {
	proto.RegisterType((*GetMonitoringRequest)(nil), "monitor.v1.types.GetMonitoringRequest")
	proto.RegisterType((*GetMonitoringResponse)(nil), "monitor.v1.types.GetMonitoringResponse")
}

func init() { proto.RegisterFile("monitor_v1_types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x6b, 0xc2, 0x30,
	0x18, 0xc5, 0xe9, 0x3a, 0x05, 0x33, 0x18, 0x23, 0xa8, 0x2b, 0x9e, 0xc4, 0x93, 0xbb, 0xa4, 0x73,
	0xfb, 0x13, 0x36, 0xd8, 0x45, 0x77, 0xc8, 0xe6, 0x39, 0xc4, 0xfa, 0x51, 0x83, 0x36, 0x5f, 0x96,
	0x7c, 0x0a, 0xfe, 0xad, 0xfb, 0x67, 0x46, 0x53, 0xed, 0xa1, 0xb0, 0x5b, 0xdf, 0xf7, 0x6b, 0xde,
	0x7b, 0x3c, 0x36, 0xae, 0xd0, 0x1a, 0x42, 0xaf, 0x4e, 0x0b, 0x45, 0x67, 0x07, 0x41, 0x38, 0x8f,
	0x84, 0xfc, 0xe1, 0x72, 0x17, 0xa7, 0x85, 0x88, 0xf7, 0xc9, 0xb6, 0x34, 0xb4, 0x3b, 0x6e, 0x44,
	0x81, 0x55, 0x5e, 0x22, 0x96, 0x07, 0xc8, 0xf7, 0x70, 0x26, 0xaf, 0x6d, 0x70, 0xda, 0x83, 0x2d,
	0xce, 0x79, 0x81, 0x1e, 0xf2, 0xf8, 0xbe, 0x8b, 0x5a, 0xfb, 0x7f, 0x41, 0x93, 0x3b, 0xb9, 0x27,
	0x6f, 0x0e, 0x07, 0xa3, 0x6d, 0xa3, 0x67, 0x6f, 0x6c, 0xf8, 0x01, 0xb4, 0x6a, 0xca, 0x18, 0x5b,
	0x4a, 0xf8, 0x39, 0x42, 0x20, 0x3e, 0x64, 0x3d, 0x70, 0x58, 0xec, 0xb2, 0x64, 0x9a, 0xcc, 0x53,
	0xd9, 0x08, 0x3e, 0x62, 0xfd, 0x3d, 0xa9, 0xb5, 0x5c, 0x66, 0x37, 0xd3, 0x64, 0x3e, 0x90, 0xbd,
	0x3d, 0xad, 0xe5, 0x72, 0xf6, 0x9b, 0xb0, 0x51, 0xc7, 0x25, 0x38, 0xb4, 0x01, 0xf8, 0x13, 0x4b,
	0x43, 0xe5, 0xa3, 0xc9, 0xdd, 0xcb, 0xa3, 0x68, 0xc3, 0xbf, 0x4c, 0x69, 0x61, 0xbb, 0xd2, 0x4e,
	0x22, 0x92, 0xac, 0xff, 0xe1, 0xcf, 0x6c, 0x18, 0x00, 0xac, 0x22, 0x53, 0x41, 0x20, 0x5d, 0x39,
	0x65, 0xb5, 0xc5, 0x10, 0x93, 0x52, 0xc9, 0x6b, 0xf6, 0x7d, 0x45, 0x9f, 0x35, 0xe1, 0x63, 0xd6,
	0x07, 0xef, 0xd1, 0x87, 0x2c, 0x9d, 0xa6, 0xf3, 0x81, 0xbc, 0x28, 0xbe, 0x62, 0x2c, 0x7e, 0xa9,
	0xad, 0x26, 0x9d, 0xdd, 0xc6, 0x6c, 0x21, 0x3a, 0xc3, 0xb4, 0xc3, 0x8b, 0xba, 0xf9, 0x91, 0x34,
	0x19, 0xb4, 0xe1, 0x5a, 0x5c, 0x0e, 0xa2, 0xc3, 0xbb, 0x26, 0xbd, 0xe9, 0xc7, 0xa5, 0x5e, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xdc, 0x29, 0x97, 0xcb, 0x01, 0x00, 0x00,
}
