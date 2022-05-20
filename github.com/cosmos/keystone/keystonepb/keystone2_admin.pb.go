// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: keystone2_admin.proto

package keystonepb

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

type KeyringSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *KeyringSpec) Reset() {
	*x = KeyringSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone2_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyringSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyringSpec) ProtoMessage() {}

func (x *KeyringSpec) ProtoReflect() protoreflect.Message {
	mi := &file_keystone2_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyringSpec.ProtoReflect.Descriptor instead.
func (*KeyringSpec) Descriptor() ([]byte, []int) {
	return file_keystone2_admin_proto_rawDescGZIP(), []int{0}
}

func (x *KeyringSpec) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KeyringSpec) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type KeyringRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InResponseTo uint64 `protobuf:"varint,2,opt,name=inResponseTo,proto3" json:"inResponseTo,omitempty"`
	Label        string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	IssuerUrl    string `protobuf:"bytes,4,opt,name=issuerUrl,proto3" json:"issuerUrl,omitempty"`
	Expires      uint64 `protobuf:"varint,5,opt,name=expires,proto3" json:"expires,omitempty"`
	IssuerIdUrl  string `protobuf:"bytes,6,opt,name=issuerIdUrl,proto3" json:"issuerIdUrl,omitempty"`
}

func (x *KeyringRef) Reset() {
	*x = KeyringRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone2_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyringRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyringRef) ProtoMessage() {}

func (x *KeyringRef) ProtoReflect() protoreflect.Message {
	mi := &file_keystone2_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyringRef.ProtoReflect.Descriptor instead.
func (*KeyringRef) Descriptor() ([]byte, []int) {
	return file_keystone2_admin_proto_rawDescGZIP(), []int{1}
}

func (x *KeyringRef) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KeyringRef) GetInResponseTo() uint64 {
	if x != nil {
		return x.InResponseTo
	}
	return 0
}

func (x *KeyringRef) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *KeyringRef) GetIssuerUrl() string {
	if x != nil {
		return x.IssuerUrl
	}
	return ""
}

func (x *KeyringRef) GetExpires() uint64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

func (x *KeyringRef) GetIssuerIdUrl() string {
	if x != nil {
		return x.IssuerIdUrl
	}
	return ""
}

type KeyrefList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InResponseTo uint64    `protobuf:"varint,2,opt,name=inResponseTo,proto3" json:"inResponseTo,omitempty"`
	Label        []*KeyRef `protobuf:"bytes,3,rep,name=label,proto3" json:"label,omitempty"`
}

func (x *KeyrefList) Reset() {
	*x = KeyrefList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone2_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyrefList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyrefList) ProtoMessage() {}

func (x *KeyrefList) ProtoReflect() protoreflect.Message {
	mi := &file_keystone2_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyrefList.ProtoReflect.Descriptor instead.
func (*KeyrefList) Descriptor() ([]byte, []int) {
	return file_keystone2_admin_proto_rawDescGZIP(), []int{2}
}

func (x *KeyrefList) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KeyrefList) GetInResponseTo() uint64 {
	if x != nil {
		return x.InResponseTo
	}
	return 0
}

func (x *KeyrefList) GetLabel() []*KeyRef {
	if x != nil {
		return x.Label
	}
	return nil
}

type KeyringLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InResponseTo uint64 `protobuf:"varint,2,opt,name=inResponseTo,proto3" json:"inResponseTo,omitempty"`
	Label        string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *KeyringLabel) Reset() {
	*x = KeyringLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keystone2_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyringLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyringLabel) ProtoMessage() {}

func (x *KeyringLabel) ProtoReflect() protoreflect.Message {
	mi := &file_keystone2_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyringLabel.ProtoReflect.Descriptor instead.
func (*KeyringLabel) Descriptor() ([]byte, []int) {
	return file_keystone2_admin_proto_rawDescGZIP(), []int{3}
}

func (x *KeyringLabel) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KeyringLabel) GetInResponseTo() uint64 {
	if x != nil {
		return x.InResponseTo
	}
	return 0
}

func (x *KeyringLabel) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_keystone2_admin_proto protoreflect.FileDescriptor

var file_keystone2_admin_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x32, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x1a, 0x13, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e,
	0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xb0, 0x01, 0x0a, 0x0a,
	0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x55,
	0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x68,
	0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x72, 0x65, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f,
	0x12, 0x26, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x52, 0x65,
	0x66, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x58, 0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x72,
	0x69, 0x6e, 0x67, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x32, 0xe7, 0x01, 0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x3b, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x4b, 0x65, 0x79, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x15, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x6b, 0x65, 0x79,
	0x72, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x14, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74,
	0x6f, 0x6e, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x66, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x15, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74,
	0x6f, 0x6e, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x65, 0x63, 0x1a,
	0x14, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x72, 0x65,
	0x66, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x0f, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x6b, 0x65, 0x79,
	0x72, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65,
	0x2e, 0x6b, 0x65, 0x79, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x10, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x74,
	0x6f, 0x6e, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x74,
	0x6f, 0x6e, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keystone2_admin_proto_rawDescOnce sync.Once
	file_keystone2_admin_proto_rawDescData = file_keystone2_admin_proto_rawDesc
)

func file_keystone2_admin_proto_rawDescGZIP() []byte {
	file_keystone2_admin_proto_rawDescOnce.Do(func() {
		file_keystone2_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_keystone2_admin_proto_rawDescData)
	})
	return file_keystone2_admin_proto_rawDescData
}

var file_keystone2_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_keystone2_admin_proto_goTypes = []interface{}{
	(*KeyringSpec)(nil),  // 0: keystone.keyringSpec
	(*KeyringRef)(nil),   // 1: keystone.keyringRef
	(*KeyrefList)(nil),   // 2: keystone.keyrefList
	(*KeyringLabel)(nil), // 3: keystone.keyringLabel
	(*KeyRef)(nil),       // 4: keystone.keyRef
	(*Empty)(nil),        // 5: keystone.empty
	(*KeySpec)(nil),      // 6: keystone.keySpec
	(*Result)(nil),       // 7: keystone.result
}
var file_keystone2_admin_proto_depIdxs = []int32{
	4, // 0: keystone.keyrefList.label:type_name -> keystone.keyRef
	0, // 1: keystone.keyringAdmin.newKeyring:input_type -> keystone.keyringSpec
	0, // 2: keystone.keyringAdmin.keys:input_type -> keystone.keyringSpec
	5, // 3: keystone.keyringAdmin.label:input_type -> keystone.empty
	6, // 4: keystone.keyringAdmin.remove:input_type -> keystone.keySpec
	1, // 5: keystone.keyringAdmin.newKeyring:output_type -> keystone.keyringRef
	2, // 6: keystone.keyringAdmin.keys:output_type -> keystone.keyrefList
	3, // 7: keystone.keyringAdmin.label:output_type -> keystone.keyringLabel
	7, // 8: keystone.keyringAdmin.remove:output_type -> keystone.result
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_keystone2_admin_proto_init() }
func file_keystone2_admin_proto_init() {
	if File_keystone2_admin_proto != nil {
		return
	}
	file_keystone_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_keystone2_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyringSpec); i {
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
		file_keystone2_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyringRef); i {
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
		file_keystone2_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyrefList); i {
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
		file_keystone2_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyringLabel); i {
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
			RawDescriptor: file_keystone2_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keystone2_admin_proto_goTypes,
		DependencyIndexes: file_keystone2_admin_proto_depIdxs,
		MessageInfos:      file_keystone2_admin_proto_msgTypes,
	}.Build()
	File_keystone2_admin_proto = out.File
	file_keystone2_admin_proto_rawDesc = nil
	file_keystone2_admin_proto_goTypes = nil
	file_keystone2_admin_proto_depIdxs = nil
}
