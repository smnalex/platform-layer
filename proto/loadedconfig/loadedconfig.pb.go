// Code generated by protoc-gen-go.
// source: github.com/hailocab/platform-layer/proto/loadedconfig/loadedconfig.proto
// DO NOT EDIT!

/*
Package com_hailocab_kernel_platform_loadedconfig is a generated protocol buffer package.

It is generated from these files:
	github.com/hailocab/platform-layer/proto/loadedconfig/loadedconfig.proto

It has these top-level messages:
	Request
	Response
*/
package com_hailocab_kernel_platform_loadedconfig

import proto "github.com/hailocab/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

type Response struct {
	Config           *string `protobuf:"bytes,1,req,name=config" json:"config,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetConfig() string {
	if m != nil && m.Config != nil {
		return *m.Config
	}
	return ""
}

func init() {
}
