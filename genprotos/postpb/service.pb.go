// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: protos/postpb/service.proto

package postpb

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

var File_protos_postpb_service_proto protoreflect.FileDescriptor

var file_protos_postpb_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x70, 0x62, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x70, 0x62, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf5, 0x01, 0x0a,
	0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x40, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0c, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6b, 0x74, 0x6f, 0x73, 0x68, 0x30, 0x33, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2d, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_postpb_service_proto_goTypes = []interface{}{
	(*GetPostRequest)(nil),    // 0: postpb.GetPostRequest
	(*ListPostsRequest)(nil),  // 1: postpb.ListPostsRequest
	(*Post)(nil),              // 2: postpb.Post
	(*DeletePostRequest)(nil), // 3: postpb.DeletePostRequest
	(*ListPostsResponse)(nil), // 4: postpb.ListPostsResponse
	(*emptypb.Empty)(nil),     // 5: google.protobuf.Empty
}
var file_protos_postpb_service_proto_depIdxs = []int32{
	0, // 0: postpb.PostService.GetPost:input_type -> postpb.GetPostRequest
	1, // 1: postpb.PostService.ListPosts:input_type -> postpb.ListPostsRequest
	2, // 2: postpb.PostService.UpdatePost:input_type -> postpb.Post
	3, // 3: postpb.PostService.DeletePost:input_type -> postpb.DeletePostRequest
	2, // 4: postpb.PostService.GetPost:output_type -> postpb.Post
	4, // 5: postpb.PostService.ListPosts:output_type -> postpb.ListPostsResponse
	5, // 6: postpb.PostService.UpdatePost:output_type -> google.protobuf.Empty
	5, // 7: postpb.PostService.DeletePost:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_postpb_service_proto_init() }
func file_protos_postpb_service_proto_init() {
	if File_protos_postpb_service_proto != nil {
		return
	}
	file_protos_postpb_post_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_postpb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_postpb_service_proto_goTypes,
		DependencyIndexes: file_protos_postpb_service_proto_depIdxs,
	}.Build()
	File_protos_postpb_service_proto = out.File
	file_protos_postpb_service_proto_rawDesc = nil
	file_protos_postpb_service_proto_goTypes = nil
	file_protos_postpb_service_proto_depIdxs = nil
}
