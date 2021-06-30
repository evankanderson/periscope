//
//Copyright © 2021 Evan Anderson <Evan.K.Anderson@gmail.com>
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: reverseproxy.proto

package periscope

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used by Out to correlate requests and responses.
	// Does not need to be set for In().
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The next two items correspond to the Request-Line in RFC7230
	// The HTTP verb (method) of the request.
	Verb string `protobuf:"bytes,2,opt,name=verb,proto3" json:"verb,omitempty"`
	// The request-target part of the HTTP request.
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	// Host is extracted out specifically from general headers and/or request_path.
	Host string `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	// The request headers, verbatim, including items like Content-Length
	// (which may be duplicated by the length of body), and the Host header.
	Headers map[string]string `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The request body. This assumes that the request body is consumed
	// as a single read and not streamed over time.
	Body []byte `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ProxyRequest) Reset() {
	*x = ProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reverseproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRequest) ProtoMessage() {}

func (x *ProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reverseproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRequest.ProtoReflect.Descriptor instead.
func (*ProxyRequest) Descriptor() ([]byte, []int) {
	return file_reverseproxy_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProxyRequest) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *ProxyRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *ProxyRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ProxyRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *ProxyRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used by Out to correlate requests and responses.
	// Does not need to be set for In().
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The next two items correspond to the Status-Line in RFC7230
	// The HTTP status code for the response.
	Status int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// The "reason-phrase" of the HTTP response
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	// The request headers, verbatim, including items like Content-Length
	// (which may be duplicated by the length of body), and the Host header.
	Headers map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The request body. This assumes that the request body is consumed
	// as a single read and not streamed over time.
	Body []byte `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ProxyResponse) Reset() {
	*x = ProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reverseproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyResponse) ProtoMessage() {}

func (x *ProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reverseproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyResponse.ProtoReflect.Descriptor instead.
func (*ProxyResponse) Descriptor() ([]byte, []int) {
	return file_reverseproxy_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProxyResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ProxyResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ProxyResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *ProxyResponse) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_reverseproxy_proto protoreflect.FileDescriptor

var file_reverseproxy_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x65, 0x72, 0x69, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22,
	0xee, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x76, 0x65, 0x72, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x3e, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x70, 0x65, 0x72, 0x69, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xe0, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x65, 0x72, 0x69, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0x86, 0x01, 0x0a, 0x09, 0x50, 0x65, 0x72, 0x69, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x12, 0x39, 0x0a, 0x02, 0x49, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x65, 0x72, 0x69, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x65, 0x72, 0x69, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x03,
	0x4f, 0x75, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x65, 0x72, 0x69, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x17, 0x2e,
	0x70, 0x65, 0x72, 0x69, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x76, 0x61, 0x6e, 0x6b,
	0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x69, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x65, 0x72, 0x69, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reverseproxy_proto_rawDescOnce sync.Once
	file_reverseproxy_proto_rawDescData = file_reverseproxy_proto_rawDesc
)

func file_reverseproxy_proto_rawDescGZIP() []byte {
	file_reverseproxy_proto_rawDescOnce.Do(func() {
		file_reverseproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_reverseproxy_proto_rawDescData)
	})
	return file_reverseproxy_proto_rawDescData
}

var file_reverseproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_reverseproxy_proto_goTypes = []interface{}{
	(*ProxyRequest)(nil),  // 0: periscope.ProxyRequest
	(*ProxyResponse)(nil), // 1: periscope.ProxyResponse
	nil,                   // 2: periscope.ProxyRequest.HeadersEntry
	nil,                   // 3: periscope.ProxyResponse.HeadersEntry
}
var file_reverseproxy_proto_depIdxs = []int32{
	2, // 0: periscope.ProxyRequest.headers:type_name -> periscope.ProxyRequest.HeadersEntry
	3, // 1: periscope.ProxyResponse.headers:type_name -> periscope.ProxyResponse.HeadersEntry
	0, // 2: periscope.Periscope.In:input_type -> periscope.ProxyRequest
	1, // 3: periscope.Periscope.Out:input_type -> periscope.ProxyResponse
	1, // 4: periscope.Periscope.In:output_type -> periscope.ProxyResponse
	0, // 5: periscope.Periscope.Out:output_type -> periscope.ProxyRequest
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_reverseproxy_proto_init() }
func file_reverseproxy_proto_init() {
	if File_reverseproxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reverseproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reverseproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reverseproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reverseproxy_proto_goTypes,
		DependencyIndexes: file_reverseproxy_proto_depIdxs,
		MessageInfos:      file_reverseproxy_proto_msgTypes,
	}.Build()
	File_reverseproxy_proto = out.File
	file_reverseproxy_proto_rawDesc = nil
	file_reverseproxy_proto_goTypes = nil
	file_reverseproxy_proto_depIdxs = nil
}