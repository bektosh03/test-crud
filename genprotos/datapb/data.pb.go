// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: protos/datapb/data.proto

package datapb

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

type DownloadPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DownloadPostsRequest) Reset() {
	*x = DownloadPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_datapb_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadPostsRequest) ProtoMessage() {}

func (x *DownloadPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_datapb_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadPostsRequest.ProtoReflect.Descriptor instead.
func (*DownloadPostsRequest) Descriptor() ([]byte, []int) {
	return file_protos_datapb_data_proto_rawDescGZIP(), []int{0}
}

type GetDownloadStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDownloadStatusRequest) Reset() {
	*x = GetDownloadStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_datapb_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDownloadStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadStatusRequest) ProtoMessage() {}

func (x *GetDownloadStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_datapb_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadStatusRequest.ProtoReflect.Descriptor instead.
func (*GetDownloadStatusRequest) Descriptor() ([]byte, []int) {
	return file_protos_datapb_data_proto_rawDescGZIP(), []int{1}
}

type GetDownloadStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (x *GetDownloadStatusResponse) Reset() {
	*x = GetDownloadStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_datapb_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDownloadStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadStatusResponse) ProtoMessage() {}

func (x *GetDownloadStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_datapb_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadStatusResponse.ProtoReflect.Descriptor instead.
func (*GetDownloadStatusResponse) Descriptor() ([]byte, []int) {
	return file_protos_datapb_data_proto_rawDescGZIP(), []int{2}
}

func (x *GetDownloadStatusResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetDownloadStatusResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

var File_protos_datapb_data_proto protoreflect.FileDescriptor

var file_protos_datapb_data_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x62, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6b, 0x74, 0x6f, 0x73, 0x68, 0x30, 0x33, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2d, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_datapb_data_proto_rawDescOnce sync.Once
	file_protos_datapb_data_proto_rawDescData = file_protos_datapb_data_proto_rawDesc
)

func file_protos_datapb_data_proto_rawDescGZIP() []byte {
	file_protos_datapb_data_proto_rawDescOnce.Do(func() {
		file_protos_datapb_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_datapb_data_proto_rawDescData)
	})
	return file_protos_datapb_data_proto_rawDescData
}

var file_protos_datapb_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_datapb_data_proto_goTypes = []interface{}{
	(*DownloadPostsRequest)(nil),      // 0: datapb.DownloadPostsRequest
	(*GetDownloadStatusRequest)(nil),  // 1: datapb.GetDownloadStatusRequest
	(*GetDownloadStatusResponse)(nil), // 2: datapb.GetDownloadStatusResponse
}
var file_protos_datapb_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_datapb_data_proto_init() }
func file_protos_datapb_data_proto_init() {
	if File_protos_datapb_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_datapb_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadPostsRequest); i {
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
		file_protos_datapb_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDownloadStatusRequest); i {
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
		file_protos_datapb_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDownloadStatusResponse); i {
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
			RawDescriptor: file_protos_datapb_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_datapb_data_proto_goTypes,
		DependencyIndexes: file_protos_datapb_data_proto_depIdxs,
		MessageInfos:      file_protos_datapb_data_proto_msgTypes,
	}.Build()
	File_protos_datapb_data_proto = out.File
	file_protos_datapb_data_proto_rawDesc = nil
	file_protos_datapb_data_proto_goTypes = nil
	file_protos_datapb_data_proto_depIdxs = nil
}
