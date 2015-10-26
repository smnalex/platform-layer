// Code generated by protoc-gen-go.
// source: github.com/hailocab/platform-layer/proto/endpointsample/endpointmetric.proto
// DO NOT EDIT!

/*
Package com_hailocab_kernel_platform_endpointmetrics is a generated protocol buffer package.

It is generated from these files:
	github.com/hailocab/platform-layer/proto/endpointsample/endpointmetric.proto

It has these top-level messages:
	EndpointMetrics
	EndpointSample
*/
package com_hailocab_kernel_platform_endpointmetrics

import proto "github.com/hailocab/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type EndpointMetrics struct {
	ServiceName      *string         `protobuf:"bytes,1,req,name=serviceName" json:"serviceName,omitempty"`
	EndpointName     *string         `protobuf:"bytes,2,req,name=endpointName" json:"endpointName,omitempty"`
	InstanceId       *string         `protobuf:"bytes,3,req,name=instanceId" json:"instanceId,omitempty"`
	Timestamp        *int64          `protobuf:"varint,4,req,name=timestamp" json:"timestamp,omitempty"`
	SuccessSample    *EndpointSample `protobuf:"bytes,5,req,name=successSample" json:"successSample,omitempty"`
	ErrorSample      *EndpointSample `protobuf:"bytes,6,req,name=errorSample" json:"errorSample,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *EndpointMetrics) Reset()         { *m = EndpointMetrics{} }
func (m *EndpointMetrics) String() string { return proto.CompactTextString(m) }
func (*EndpointMetrics) ProtoMessage()    {}

func (m *EndpointMetrics) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *EndpointMetrics) GetEndpointName() string {
	if m != nil && m.EndpointName != nil {
		return *m.EndpointName
	}
	return ""
}

func (m *EndpointMetrics) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

func (m *EndpointMetrics) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *EndpointMetrics) GetSuccessSample() *EndpointSample {
	if m != nil {
		return m.SuccessSample
	}
	return nil
}

func (m *EndpointMetrics) GetErrorSample() *EndpointSample {
	if m != nil {
		return m.ErrorSample
	}
	return nil
}

type EndpointSample struct {
	Rate1            *float32 `protobuf:"fixed32,1,req,name=rate1" json:"rate1,omitempty"`
	Rate5            *float32 `protobuf:"fixed32,2,req,name=rate5" json:"rate5,omitempty"`
	Rate15           *float32 `protobuf:"fixed32,3,req,name=rate15" json:"rate15,omitempty"`
	Mean             *float32 `protobuf:"fixed32,4,req,name=mean" json:"mean,omitempty"`
	StdDev           *float32 `protobuf:"fixed32,5,req,name=stdDev" json:"stdDev,omitempty"`
	Upper95          *float32 `protobuf:"fixed32,6,req,name=upper95" json:"upper95,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EndpointSample) Reset()         { *m = EndpointSample{} }
func (m *EndpointSample) String() string { return proto.CompactTextString(m) }
func (*EndpointSample) ProtoMessage()    {}

func (m *EndpointSample) GetRate1() float32 {
	if m != nil && m.Rate1 != nil {
		return *m.Rate1
	}
	return 0
}

func (m *EndpointSample) GetRate5() float32 {
	if m != nil && m.Rate5 != nil {
		return *m.Rate5
	}
	return 0
}

func (m *EndpointSample) GetRate15() float32 {
	if m != nil && m.Rate15 != nil {
		return *m.Rate15
	}
	return 0
}

func (m *EndpointSample) GetMean() float32 {
	if m != nil && m.Mean != nil {
		return *m.Mean
	}
	return 0
}

func (m *EndpointSample) GetStdDev() float32 {
	if m != nil && m.StdDev != nil {
		return *m.StdDev
	}
	return 0
}

func (m *EndpointSample) GetUpper95() float32 {
	if m != nil && m.Upper95 != nil {
		return *m.Upper95
	}
	return 0
}

func init() {
}
