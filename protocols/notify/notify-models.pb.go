// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.0
// source: notify-models.proto

package notify

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

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteUuid      string `protobuf:"bytes,1,opt,name=note_uuid,json=noteUuid,proto3" json:"note_uuid,omitempty"`
	CandidateName string `protobuf:"bytes,2,opt,name=candidate_name,json=candidateName,proto3" json:"candidate_name,omitempty"`
	CanditateUrl  string `protobuf:"bytes,3,opt,name=canditate_url,json=canditateUrl,proto3" json:"canditate_url,omitempty"`
	UserUuid      string `protobuf:"bytes,4,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	VacancyUuid   string `protobuf:"bytes,5,opt,name=vacancy_uuid,json=vacancyUuid,proto3" json:"vacancy_uuid,omitempty"`
	IsReaded      bool   `protobuf:"varint,6,opt,name=is_readed,json=isReaded,proto3" json:"is_readed,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_notify_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notify_models_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notify_models_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetNoteUuid() string {
	if x != nil {
		return x.NoteUuid
	}
	return ""
}

func (x *Notification) GetCandidateName() string {
	if x != nil {
		return x.CandidateName
	}
	return ""
}

func (x *Notification) GetCanditateUrl() string {
	if x != nil {
		return x.CanditateUrl
	}
	return ""
}

func (x *Notification) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *Notification) GetVacancyUuid() string {
	if x != nil {
		return x.VacancyUuid
	}
	return ""
}

func (x *Notification) GetIsReaded() bool {
	if x != nil {
		return x.IsReaded
	}
	return false
}

type Notifications struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *Notifications) Reset() {
	*x = Notifications{}
	mi := &file_notify_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notifications) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notifications) ProtoMessage() {}

func (x *Notifications) ProtoReflect() protoreflect.Message {
	mi := &file_notify_models_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notifications.ProtoReflect.Descriptor instead.
func (*Notifications) Descriptor() ([]byte, []int) {
	return file_notify_models_proto_rawDescGZIP(), []int{1}
}

func (x *Notifications) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

var File_notify_models_proto protoreflect.FileDescriptor

var file_notify_models_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0xd4, 0x01,
	0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x74, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x55, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x61, 0x63, 0x61,
	0x6e, 0x63, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x72, 0x65,
	0x61, 0x64, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x52, 0x65,
	0x61, 0x64, 0x65, 0x64, 0x22, 0x4b, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notify_models_proto_rawDescOnce sync.Once
	file_notify_models_proto_rawDescData = file_notify_models_proto_rawDesc
)

func file_notify_models_proto_rawDescGZIP() []byte {
	file_notify_models_proto_rawDescOnce.Do(func() {
		file_notify_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_notify_models_proto_rawDescData)
	})
	return file_notify_models_proto_rawDescData
}

var file_notify_models_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notify_models_proto_goTypes = []any{
	(*Notification)(nil),  // 0: models.Notification
	(*Notifications)(nil), // 1: models.Notifications
}
var file_notify_models_proto_depIdxs = []int32{
	0, // 0: models.Notifications.notifications:type_name -> models.Notification
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_notify_models_proto_init() }
func file_notify_models_proto_init() {
	if File_notify_models_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notify_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notify_models_proto_goTypes,
		DependencyIndexes: file_notify_models_proto_depIdxs,
		MessageInfos:      file_notify_models_proto_msgTypes,
	}.Build()
	File_notify_models_proto = out.File
	file_notify_models_proto_rawDesc = nil
	file_notify_models_proto_goTypes = nil
	file_notify_models_proto_depIdxs = nil
}
