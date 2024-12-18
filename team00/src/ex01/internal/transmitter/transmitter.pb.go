// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: transmitter.proto

package transmitter

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

type FrequencyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string  `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Frequency float64 `protobuf:"fixed64,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Timestamp string  `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *FrequencyMessage) Reset() {
	*x = FrequencyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transmitter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrequencyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrequencyMessage) ProtoMessage() {}

func (x *FrequencyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transmitter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrequencyMessage.ProtoReflect.Descriptor instead.
func (*FrequencyMessage) Descriptor() ([]byte, []int) {
	return file_transmitter_proto_rawDescGZIP(), []int{0}
}

func (x *FrequencyMessage) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *FrequencyMessage) GetFrequency() float64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *FrequencyMessage) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type FrequencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FrequencyRequest) Reset() {
	*x = FrequencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transmitter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrequencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrequencyRequest) ProtoMessage() {}

func (x *FrequencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transmitter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrequencyRequest.ProtoReflect.Descriptor instead.
func (*FrequencyRequest) Descriptor() ([]byte, []int) {
	return file_transmitter_proto_rawDescGZIP(), []int{1}
}

var File_transmitter_proto protoreflect.FileDescriptor

var file_transmitter_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x22, 0x6d, 0x0a, 0x10, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x12, 0x0a, 0x10, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x32, 0x69, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x11, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1d,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transmitter_proto_rawDescOnce sync.Once
	file_transmitter_proto_rawDescData = file_transmitter_proto_rawDesc
)

func file_transmitter_proto_rawDescGZIP() []byte {
	file_transmitter_proto_rawDescOnce.Do(func() {
		file_transmitter_proto_rawDescData = protoimpl.X.CompressGZIP(file_transmitter_proto_rawDescData)
	})
	return file_transmitter_proto_rawDescData
}

var file_transmitter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transmitter_proto_goTypes = []any{
	(*FrequencyMessage)(nil), // 0: transmitter.FrequencyMessage
	(*FrequencyRequest)(nil), // 1: transmitter.FrequencyRequest
}
var file_transmitter_proto_depIdxs = []int32{
	1, // 0: transmitter.TransmitterService.StreamFrequencies:input_type -> transmitter.FrequencyRequest
	0, // 1: transmitter.TransmitterService.StreamFrequencies:output_type -> transmitter.FrequencyMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transmitter_proto_init() }
func file_transmitter_proto_init() {
	if File_transmitter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transmitter_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FrequencyMessage); i {
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
		file_transmitter_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FrequencyRequest); i {
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
			RawDescriptor: file_transmitter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transmitter_proto_goTypes,
		DependencyIndexes: file_transmitter_proto_depIdxs,
		MessageInfos:      file_transmitter_proto_msgTypes,
	}.Build()
	File_transmitter_proto = out.File
	file_transmitter_proto_rawDesc = nil
	file_transmitter_proto_goTypes = nil
	file_transmitter_proto_depIdxs = nil
}
