// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: myorg/todo/v1/update_todo.proto

package v1

import (
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

type UpdateTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FieldMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	Title     string                `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Done      bool                  `protobuf:"varint,4,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *UpdateTodoRequest) Reset() {
	*x = UpdateTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myorg_todo_v1_update_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoRequest) ProtoMessage() {}

func (x *UpdateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_myorg_todo_v1_update_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoRequest.ProtoReflect.Descriptor instead.
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return file_myorg_todo_v1_update_todo_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTodoRequest) GetFieldMask() *field_mask.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *UpdateTodoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTodoRequest) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type UpdateTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *UpdateTodoResponse) Reset() {
	*x = UpdateTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myorg_todo_v1_update_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoResponse) ProtoMessage() {}

func (x *UpdateTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_myorg_todo_v1_update_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoResponse.ProtoReflect.Descriptor instead.
func (*UpdateTodoResponse) Descriptor() ([]byte, []int) {
	return file_myorg_todo_v1_update_todo_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTodoResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

var File_myorg_todo_v1_update_todo_proto protoreflect.FileDescriptor

var file_myorg_todo_v1_update_todo_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x79, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6d, 0x79, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31,
	0x1a, 0x18, 0x6d, 0x79, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x3d, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x79,
	0x6f, 0x72, 0x67, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x42, 0x7a, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79,
	0x6f, 0x72, 0x67, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x54, 0x6f, 0x64,
	0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x79, 0x4f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6d,
	0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6d,
	0x79, 0x6f, 0x72, 0x67, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d,
	0x54, 0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x79, 0x4f, 0x72, 0x67, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0d, 0x4d, 0x79, 0x4f, 0x72, 0x67, 0x5c, 0x54, 0x6f, 0x64, 0x6f, 0x5c,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_myorg_todo_v1_update_todo_proto_rawDescOnce sync.Once
	file_myorg_todo_v1_update_todo_proto_rawDescData = file_myorg_todo_v1_update_todo_proto_rawDesc
)

func file_myorg_todo_v1_update_todo_proto_rawDescGZIP() []byte {
	file_myorg_todo_v1_update_todo_proto_rawDescOnce.Do(func() {
		file_myorg_todo_v1_update_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_myorg_todo_v1_update_todo_proto_rawDescData)
	})
	return file_myorg_todo_v1_update_todo_proto_rawDescData
}

var file_myorg_todo_v1_update_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_myorg_todo_v1_update_todo_proto_goTypes = []interface{}{
	(*UpdateTodoRequest)(nil),    // 0: myorg.todo.v1.UpdateTodoRequest
	(*UpdateTodoResponse)(nil),   // 1: myorg.todo.v1.UpdateTodoResponse
	(*field_mask.FieldMask)(nil), // 2: google.protobuf.FieldMask
	(*Todo)(nil),                 // 3: myorg.todo.v1.Todo
}
var file_myorg_todo_v1_update_todo_proto_depIdxs = []int32{
	2, // 0: myorg.todo.v1.UpdateTodoRequest.field_mask:type_name -> google.protobuf.FieldMask
	3, // 1: myorg.todo.v1.UpdateTodoResponse.todo:type_name -> myorg.todo.v1.Todo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_myorg_todo_v1_update_todo_proto_init() }
func file_myorg_todo_v1_update_todo_proto_init() {
	if File_myorg_todo_v1_update_todo_proto != nil {
		return
	}
	file_myorg_todo_v1_todo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_myorg_todo_v1_update_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoRequest); i {
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
		file_myorg_todo_v1_update_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoResponse); i {
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
			RawDescriptor: file_myorg_todo_v1_update_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_myorg_todo_v1_update_todo_proto_goTypes,
		DependencyIndexes: file_myorg_todo_v1_update_todo_proto_depIdxs,
		MessageInfos:      file_myorg_todo_v1_update_todo_proto_msgTypes,
	}.Build()
	File_myorg_todo_v1_update_todo_proto = out.File
	file_myorg_todo_v1_update_todo_proto_rawDesc = nil
	file_myorg_todo_v1_update_todo_proto_goTypes = nil
	file_myorg_todo_v1_update_todo_proto_depIdxs = nil
}