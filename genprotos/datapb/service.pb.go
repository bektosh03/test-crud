// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: protos/datapb/service.proto

package datapb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protos_datapb_service_proto protoreflect.FileDescriptor

var file_protos_datapb_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64,
	0x61, 0x74, 0x61, 0x64, 0x62, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xae, 0x01, 0x0a,
	0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0d,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1c, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x64, 0x62, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x58, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6b, 0x74,
	0x6f, 0x73, 0x68, 0x30, 0x33, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x63, 0x72, 0x75, 0x64, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_datapb_service_proto_goTypes = []interface{}{
	(*DownloadPostsRequest)(nil),      // 0: datadb.DownloadPostsRequest
	(*GetDownloadStatusRequest)(nil),  // 1: datadb.GetDownloadStatusRequest
	(*emptypb.Empty)(nil),             // 2: google.protobuf.Empty
	(*GetDownloadStatusResponse)(nil), // 3: datadb.GetDownloadStatusResponse
}
var file_protos_datapb_service_proto_depIdxs = []int32{
	0, // 0: datadb.DataService.DownloadPosts:input_type -> datadb.DownloadPostsRequest
	1, // 1: datadb.DataService.GetDownloadStatus:input_type -> datadb.GetDownloadStatusRequest
	2, // 2: datadb.DataService.DownloadPosts:output_type -> google.protobuf.Empty
	3, // 3: datadb.DataService.GetDownloadStatus:output_type -> datadb.GetDownloadStatusResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_datapb_service_proto_init() }
func file_protos_datapb_service_proto_init() {
	if File_protos_datapb_service_proto != nil {
		return
	}
	file_protos_datapb_data_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_datapb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_datapb_service_proto_goTypes,
		DependencyIndexes: file_protos_datapb_service_proto_depIdxs,
	}.Build()
	File_protos_datapb_service_proto = out.File
	file_protos_datapb_service_proto_rawDesc = nil
	file_protos_datapb_service_proto_goTypes = nil
	file_protos_datapb_service_proto_depIdxs = nil
}
