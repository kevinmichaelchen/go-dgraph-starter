// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: myorg/todo/v1/pagination.proto

package v1

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

type PaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*PaginationRequest_ForwardPaginationInfo
	//	*PaginationRequest_BackwardPaginationInfo
	Request isPaginationRequest_Request `protobuf_oneof:"request"`
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myorg_todo_v1_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_myorg_todo_v1_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_myorg_todo_v1_pagination_proto_rawDescGZIP(), []int{0}
}

func (m *PaginationRequest) GetRequest() isPaginationRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *PaginationRequest) GetForwardPaginationInfo() *ForwardPaginationRequest {
	if x, ok := x.GetRequest().(*PaginationRequest_ForwardPaginationInfo); ok {
		return x.ForwardPaginationInfo
	}
	return nil
}

func (x *PaginationRequest) GetBackwardPaginationInfo() *BackwardPaginationRequest {
	if x, ok := x.GetRequest().(*PaginationRequest_BackwardPaginationInfo); ok {
		return x.BackwardPaginationInfo
	}
	return nil
}

type isPaginationRequest_Request interface {
	isPaginationRequest_Request()
}

type PaginationRequest_ForwardPaginationInfo struct {
	ForwardPaginationInfo *ForwardPaginationRequest `protobuf:"bytes,1,opt,name=forward_pagination_info,json=forwardPaginationInfo,proto3,oneof"`
}

type PaginationRequest_BackwardPaginationInfo struct {
	BackwardPaginationInfo *BackwardPaginationRequest `protobuf:"bytes,2,opt,name=backward_pagination_info,json=backwardPaginationInfo,proto3,oneof"`
}

func (*PaginationRequest_ForwardPaginationInfo) isPaginationRequest_Request() {}

func (*PaginationRequest_BackwardPaginationInfo) isPaginationRequest_Request() {}

type ForwardPaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First int32  `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	After string `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *ForwardPaginationRequest) Reset() {
	*x = ForwardPaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myorg_todo_v1_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardPaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardPaginationRequest) ProtoMessage() {}

func (x *ForwardPaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_myorg_todo_v1_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardPaginationRequest.ProtoReflect.Descriptor instead.
func (*ForwardPaginationRequest) Descriptor() ([]byte, []int) {
	return file_myorg_todo_v1_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *ForwardPaginationRequest) GetFirst() int32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *ForwardPaginationRequest) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

type BackwardPaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Last   int32  `protobuf:"varint,1,opt,name=last,proto3" json:"last,omitempty"`
	Before string `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
}

func (x *BackwardPaginationRequest) Reset() {
	*x = BackwardPaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myorg_todo_v1_pagination_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackwardPaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackwardPaginationRequest) ProtoMessage() {}

func (x *BackwardPaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_myorg_todo_v1_pagination_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackwardPaginationRequest.ProtoReflect.Descriptor instead.
func (*BackwardPaginationRequest) Descriptor() ([]byte, []int) {
	return file_myorg_todo_v1_pagination_proto_rawDescGZIP(), []int{2}
}

func (x *BackwardPaginationRequest) GetLast() int32 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *BackwardPaginationRequest) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cursor of first item in page
	StartCursor string `protobuf:"bytes,1,opt,name=start_cursor,json=startCursor,proto3" json:"start_cursor,omitempty"`
	// cursor of last item in page
	EndCursor   string `protobuf:"bytes,2,opt,name=end_cursor,json=endCursor,proto3" json:"end_cursor,omitempty"`
	HasNextPage bool   `protobuf:"varint,3,opt,name=has_next_page,json=hasNextPage,proto3" json:"has_next_page,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myorg_todo_v1_pagination_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_myorg_todo_v1_pagination_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_myorg_todo_v1_pagination_proto_rawDescGZIP(), []int{3}
}

func (x *PageInfo) GetStartCursor() string {
	if x != nil {
		return x.StartCursor
	}
	return ""
}

func (x *PageInfo) GetEndCursor() string {
	if x != nil {
		return x.EndCursor
	}
	return ""
}

func (x *PageInfo) GetHasNextPage() bool {
	if x != nil {
		return x.HasNextPage
	}
	return false
}

var File_myorg_todo_v1_pagination_proto protoreflect.FileDescriptor

var file_myorg_todo_v1_pagination_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x79, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6d, 0x79, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x22,
	0xe7, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x61, 0x0a, 0x17, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x79, 0x6f, 0x72, 0x67, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x15, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x64, 0x0a, 0x18, 0x62, 0x61, 0x63, 0x6b,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x79, 0x6f,
	0x72, 0x67, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x77,
	0x61, 0x72, 0x64, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x16, 0x62, 0x61, 0x63, 0x6b, 0x77, 0x61, 0x72, 0x64,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x09,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x18, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x22, 0x47, 0x0a, 0x19, 0x42, 0x61, 0x63, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x61,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0x70, 0x0a, 0x08, 0x50, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64,
	0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x6e, 0x64, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x5f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x42, 0x7a, 0x0a, 0x11,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76,
	0x31, 0x42, 0x09, 0x54, 0x6f, 0x64, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x79, 0x4f, 0x72, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x79, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x54, 0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x79, 0x4f, 0x72, 0x67,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x4d, 0x79, 0x4f, 0x72, 0x67,
	0x5c, 0x54, 0x6f, 0x64, 0x6f, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_myorg_todo_v1_pagination_proto_rawDescOnce sync.Once
	file_myorg_todo_v1_pagination_proto_rawDescData = file_myorg_todo_v1_pagination_proto_rawDesc
)

func file_myorg_todo_v1_pagination_proto_rawDescGZIP() []byte {
	file_myorg_todo_v1_pagination_proto_rawDescOnce.Do(func() {
		file_myorg_todo_v1_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_myorg_todo_v1_pagination_proto_rawDescData)
	})
	return file_myorg_todo_v1_pagination_proto_rawDescData
}

var file_myorg_todo_v1_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_myorg_todo_v1_pagination_proto_goTypes = []interface{}{
	(*PaginationRequest)(nil),         // 0: myorg.todo.v1.PaginationRequest
	(*ForwardPaginationRequest)(nil),  // 1: myorg.todo.v1.ForwardPaginationRequest
	(*BackwardPaginationRequest)(nil), // 2: myorg.todo.v1.BackwardPaginationRequest
	(*PageInfo)(nil),                  // 3: myorg.todo.v1.PageInfo
}
var file_myorg_todo_v1_pagination_proto_depIdxs = []int32{
	1, // 0: myorg.todo.v1.PaginationRequest.forward_pagination_info:type_name -> myorg.todo.v1.ForwardPaginationRequest
	2, // 1: myorg.todo.v1.PaginationRequest.backward_pagination_info:type_name -> myorg.todo.v1.BackwardPaginationRequest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_myorg_todo_v1_pagination_proto_init() }
func file_myorg_todo_v1_pagination_proto_init() {
	if File_myorg_todo_v1_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_myorg_todo_v1_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationRequest); i {
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
		file_myorg_todo_v1_pagination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardPaginationRequest); i {
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
		file_myorg_todo_v1_pagination_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackwardPaginationRequest); i {
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
		file_myorg_todo_v1_pagination_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
	file_myorg_todo_v1_pagination_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PaginationRequest_ForwardPaginationInfo)(nil),
		(*PaginationRequest_BackwardPaginationInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_myorg_todo_v1_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_myorg_todo_v1_pagination_proto_goTypes,
		DependencyIndexes: file_myorg_todo_v1_pagination_proto_depIdxs,
		MessageInfos:      file_myorg_todo_v1_pagination_proto_msgTypes,
	}.Build()
	File_myorg_todo_v1_pagination_proto = out.File
	file_myorg_todo_v1_pagination_proto_rawDesc = nil
	file_myorg_todo_v1_pagination_proto_goTypes = nil
	file_myorg_todo_v1_pagination_proto_depIdxs = nil
}
