// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: todo.proto

package todo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_todo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

type Todo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Completed     bool                   `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Todo) Reset() {
	*x = Todo{}
	mi := &file_todo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *Todo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type TodoList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Todos         []*Todo                `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TodoList) Reset() {
	*x = TodoList{}
	mi := &file_todo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TodoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoList) ProtoMessage() {}

func (x *TodoList) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoList.ProtoReflect.Descriptor instead.
func (*TodoList) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *TodoList) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type CreateTodoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	mi := &file_todo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTodoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateTodoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Completed     bool                   `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTodoRequest) Reset() {
	*x = UpdateTodoRequest{}
	mi := &file_todo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoRequest) ProtoMessage() {}

func (x *UpdateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[4]
	if x != nil {
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
	return file_todo_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTodoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTodoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTodoRequest) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type DeleteTodoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTodoRequest) Reset() {
	*x = DeleteTodoRequest{}
	mi := &file_todo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoRequest) ProtoMessage() {}

func (x *DeleteTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoRequest.ProtoReflect.Descriptor instead.
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTodoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_todo_proto protoreflect.FileDescriptor

const file_todo_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"todo.proto\x12\x04todo\"\a\n" +
	"\x05Empty\"J\n" +
	"\x04Todo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x1c\n" +
	"\tcompleted\x18\x03 \x01(\bR\tcompleted\",\n" +
	"\bTodoList\x12 \n" +
	"\x05todos\x18\x01 \x03(\v2\n" +
	".todo.TodoR\x05todos\")\n" +
	"\x11CreateTodoRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\"W\n" +
	"\x11UpdateTodoRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x1c\n" +
	"\tcompleted\x18\x03 \x01(\bR\tcompleted\"#\n" +
	"\x11DeleteTodoRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id2\xd0\x01\n" +
	"\vTodoService\x121\n" +
	"\n" +
	"CreateTodo\x12\x17.todo.CreateTodoRequest\x1a\n" +
	".todo.Todo\x12'\n" +
	"\bGetTodos\x12\v.todo.Empty\x1a\x0e.todo.TodoList\x121\n" +
	"\n" +
	"UpdateTodo\x12\x17.todo.UpdateTodoRequest\x1a\n" +
	".todo.Todo\x122\n" +
	"\n" +
	"DeleteTodo\x12\x17.todo.DeleteTodoRequest\x1a\v.todo.EmptyB(Z&github.com/Deza415/toDoList/proto;todob\x06proto3"

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData []byte
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_todo_proto_rawDesc), len(file_todo_proto_rawDesc)))
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_todo_proto_goTypes = []any{
	(*Empty)(nil),             // 0: todo.Empty
	(*Todo)(nil),              // 1: todo.Todo
	(*TodoList)(nil),          // 2: todo.TodoList
	(*CreateTodoRequest)(nil), // 3: todo.CreateTodoRequest
	(*UpdateTodoRequest)(nil), // 4: todo.UpdateTodoRequest
	(*DeleteTodoRequest)(nil), // 5: todo.DeleteTodoRequest
}
var file_todo_proto_depIdxs = []int32{
	1, // 0: todo.TodoList.todos:type_name -> todo.Todo
	3, // 1: todo.TodoService.CreateTodo:input_type -> todo.CreateTodoRequest
	0, // 2: todo.TodoService.GetTodos:input_type -> todo.Empty
	4, // 3: todo.TodoService.UpdateTodo:input_type -> todo.UpdateTodoRequest
	5, // 4: todo.TodoService.DeleteTodo:input_type -> todo.DeleteTodoRequest
	1, // 5: todo.TodoService.CreateTodo:output_type -> todo.Todo
	2, // 6: todo.TodoService.GetTodos:output_type -> todo.TodoList
	1, // 7: todo.TodoService.UpdateTodo:output_type -> todo.Todo
	0, // 8: todo.TodoService.DeleteTodo:output_type -> todo.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_todo_proto_rawDesc), len(file_todo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
