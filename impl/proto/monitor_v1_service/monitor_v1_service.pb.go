// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor_v1_service.proto

/*
Package monitor_v1_service is a generated protocol buffer package.

Monitor Service

The Key Transparency monitor server service consists of APIs to fetch
monitor results queried using the mutations API.

It is generated from these files:
	monitor_v1_service.proto

It has these top-level messages:
*/
package monitor_v1_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import monitor_v1_types "github.com/google/keytransparency/core/proto/monitor_v1_types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MonitorService service

type MonitorServiceClient interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains additional data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest epoch the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current epoch it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetSignedMapRoot(ctx context.Context, in *monitor_v1_types.GetMonitoringRequest, opts ...grpc.CallOption) (*monitor_v1_types.GetMonitoringResponse, error)
	// GetSignedMapRootByRevision works similar to GetSignedMapRoot but returns
	// the monitor's result for a specific map revision.
	//
	// Returns the signed map root for the specified epoch the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current epoch it won't sign the map root
	// and additional data will be provided to reproduce the failure.
	GetSignedMapRootByRevision(ctx context.Context, in *monitor_v1_types.GetMonitoringRequest, opts ...grpc.CallOption) (*monitor_v1_types.GetMonitoringResponse, error)
}

type monitorServiceClient struct {
	cc *grpc.ClientConn
}

func NewMonitorServiceClient(cc *grpc.ClientConn) MonitorServiceClient {
	return &monitorServiceClient{cc}
}

func (c *monitorServiceClient) GetSignedMapRoot(ctx context.Context, in *monitor_v1_types.GetMonitoringRequest, opts ...grpc.CallOption) (*monitor_v1_types.GetMonitoringResponse, error) {
	out := new(monitor_v1_types.GetMonitoringResponse)
	err := grpc.Invoke(ctx, "/monitor.v1.service.MonitorService/GetSignedMapRoot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) GetSignedMapRootByRevision(ctx context.Context, in *monitor_v1_types.GetMonitoringRequest, opts ...grpc.CallOption) (*monitor_v1_types.GetMonitoringResponse, error) {
	out := new(monitor_v1_types.GetMonitoringResponse)
	err := grpc.Invoke(ctx, "/monitor.v1.service.MonitorService/GetSignedMapRootByRevision", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MonitorService service

type MonitorServiceServer interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains additional data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest epoch the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current epoch it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetSignedMapRoot(context.Context, *monitor_v1_types.GetMonitoringRequest) (*monitor_v1_types.GetMonitoringResponse, error)
	// GetSignedMapRootByRevision works similar to GetSignedMapRoot but returns
	// the monitor's result for a specific map revision.
	//
	// Returns the signed map root for the specified epoch the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current epoch it won't sign the map root
	// and additional data will be provided to reproduce the failure.
	GetSignedMapRootByRevision(context.Context, *monitor_v1_types.GetMonitoringRequest) (*monitor_v1_types.GetMonitoringResponse, error)
}

func RegisterMonitorServiceServer(s *grpc.Server, srv MonitorServiceServer) {
	s.RegisterService(&_MonitorService_serviceDesc, srv)
}

func _MonitorService_GetSignedMapRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(monitor_v1_types.GetMonitoringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).GetSignedMapRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.service.MonitorService/GetSignedMapRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).GetSignedMapRoot(ctx, req.(*monitor_v1_types.GetMonitoringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_GetSignedMapRootByRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(monitor_v1_types.GetMonitoringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).GetSignedMapRootByRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.service.MonitorService/GetSignedMapRootByRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).GetSignedMapRootByRevision(ctx, req.(*monitor_v1_types.GetMonitoringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitor.v1.service.MonitorService",
	HandlerType: (*MonitorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSignedMapRoot",
			Handler:    _MonitorService_GetSignedMapRoot_Handler,
		},
		{
			MethodName: "GetSignedMapRootByRevision",
			Handler:    _MonitorService_GetSignedMapRootByRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor_v1_service.proto",
}

func init() { proto.RegisterFile("monitor_v1_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xb1, 0x4e, 0xc3, 0x40,
	0x0c, 0x86, 0xd5, 0x0e, 0x0c, 0x19, 0x10, 0xba, 0x09, 0x45, 0x4c, 0x0c, 0x14, 0x18, 0x62, 0x0a,
	0x1b, 0x23, 0x4b, 0x17, 0xba, 0xa4, 0x30, 0x47, 0xd7, 0x60, 0xa5, 0xa7, 0x26, 0xf6, 0x71, 0x76,
	0x22, 0x45, 0x55, 0x17, 0x5e, 0x01, 0x76, 0x9e, 0x87, 0x99, 0x57, 0xe0, 0x41, 0x90, 0x92, 0x20,
	0x55, 0x30, 0x74, 0x61, 0xf5, 0xef, 0xff, 0xf3, 0x27, 0x47, 0xc7, 0x15, 0x93, 0x53, 0x0e, 0x59,
	0x33, 0xcd, 0x04, 0x43, 0xe3, 0x72, 0x4c, 0x7c, 0x60, 0x65, 0x63, 0x86, 0x24, 0x69, 0xa6, 0xc9,
	0x90, 0xc4, 0x0f, 0x85, 0xd3, 0x55, 0xbd, 0x4c, 0x72, 0xae, 0xa0, 0x60, 0x2e, 0x4a, 0x84, 0x35,
	0xb6, 0x1a, 0x2c, 0x89, 0xb7, 0x01, 0x29, 0x6f, 0x21, 0xe7, 0x80, 0xd0, 0x11, 0x60, 0x07, 0xad,
	0xad, 0x47, 0xf9, 0x33, 0xe8, 0x2f, 0xc5, 0x27, 0x03, 0xca, 0x7a, 0x07, 0x96, 0x88, 0xd5, 0xaa,
	0x63, 0x1a, 0xd2, 0xeb, 0x8f, 0x71, 0x74, 0x38, 0xef, 0x8b, 0x8b, 0x5e, 0xc3, 0xbc, 0x8d, 0xa2,
	0xa3, 0x19, 0xea, 0xc2, 0x15, 0x84, 0x4f, 0x73, 0xeb, 0x53, 0x66, 0x35, 0x67, 0xc9, 0x8e, 0x70,
	0x8f, 0x9f, 0xa1, 0x0e, 0x4d, 0x47, 0x45, 0x8a, 0xcf, 0x35, 0x8a, 0xc6, 0x93, 0xbd, 0x7b, 0xe2,
	0x99, 0x04, 0x4f, 0xe1, 0xe5, 0xf3, 0xeb, 0x75, 0x7c, 0x61, 0x26, 0xd0, 0x4c, 0x7f, 0xd4, 0x61,
	0xb3, 0xd6, 0xec, 0x31, 0xbd, 0xdf, 0x42, 0x65, 0x3d, 0x04, 0x94, 0xba, 0x54, 0xb9, 0x2d, 0xad,
	0xa2, 0xa8, 0x79, 0x1f, 0x45, 0xf1, 0x6f, 0xad, 0xbb, 0x36, 0xc5, 0xc6, 0x89, 0x63, 0xfa, 0x7f,
	0xc1, 0xab, 0x4e, 0xf0, 0xd2, 0x9c, 0xef, 0x13, 0x84, 0x0d, 0x7a, 0xce, 0x57, 0xdb, 0xe5, 0x41,
	0xf7, 0xd2, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x96, 0x6d, 0xf3, 0xf6, 0x01, 0x00,
	0x00,
}
