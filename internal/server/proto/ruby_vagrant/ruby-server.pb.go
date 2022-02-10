// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/ruby_vagrant/ruby-server.proto

package ruby_vagrant

import (
	vagrant_plugin_sdk "github.com/hashicorp/vagrant-plugin-sdk/proto/vagrant_plugin_sdk"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Supported plugin types, the values here MUST match the enum values
// in the Go sdk/component package exactly. A test in internal/server
// validates this.
type Plugin_Type int32

const (
	Plugin_UNKNOWN       Plugin_Type = 0
	Plugin_COMMAND       Plugin_Type = 1
	Plugin_COMMUNICATOR  Plugin_Type = 2
	Plugin_GUEST         Plugin_Type = 3
	Plugin_HOST          Plugin_Type = 4
	Plugin_PROVIDER      Plugin_Type = 5
	Plugin_PROVISIONER   Plugin_Type = 6
	Plugin_SYNCEDFOLDER  Plugin_Type = 7
	Plugin_AUTHENTICATOR Plugin_Type = 8
	Plugin_LOGPLATFORM   Plugin_Type = 9
	Plugin_LOGVIEWER     Plugin_Type = 10
	Plugin_MAPPER        Plugin_Type = 11
	Plugin_CONFIG        Plugin_Type = 12
	Plugin_PLUGININFO    Plugin_Type = 13
	Plugin_PUSH          Plugin_Type = 14
)

// Enum value maps for Plugin_Type.
var (
	Plugin_Type_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "COMMAND",
		2:  "COMMUNICATOR",
		3:  "GUEST",
		4:  "HOST",
		5:  "PROVIDER",
		6:  "PROVISIONER",
		7:  "SYNCEDFOLDER",
		8:  "AUTHENTICATOR",
		9:  "LOGPLATFORM",
		10: "LOGVIEWER",
		11: "MAPPER",
		12: "CONFIG",
		13: "PLUGININFO",
		14: "PUSH",
	}
	Plugin_Type_value = map[string]int32{
		"UNKNOWN":       0,
		"COMMAND":       1,
		"COMMUNICATOR":  2,
		"GUEST":         3,
		"HOST":          4,
		"PROVIDER":      5,
		"PROVISIONER":   6,
		"SYNCEDFOLDER":  7,
		"AUTHENTICATOR": 8,
		"LOGPLATFORM":   9,
		"LOGVIEWER":     10,
		"MAPPER":        11,
		"CONFIG":        12,
		"PLUGININFO":    13,
		"PUSH":          14,
	}
)

func (x Plugin_Type) Enum() *Plugin_Type {
	p := new(Plugin_Type)
	*p = x
	return p
}

func (x Plugin_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Plugin_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ruby_vagrant_ruby_server_proto_enumTypes[0].Descriptor()
}

func (Plugin_Type) Type() protoreflect.EnumType {
	return &file_proto_ruby_vagrant_ruby_server_proto_enumTypes[0]
}

func (x Plugin_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Plugin_Type.Descriptor instead.
func (Plugin_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_ruby_vagrant_ruby_server_proto_rawDescGZIP(), []int{1, 0}
}

type GetPluginsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugins []*Plugin `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty"`
}

func (x *GetPluginsResponse) Reset() {
	*x = GetPluginsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ruby_vagrant_ruby_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPluginsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginsResponse) ProtoMessage() {}

func (x *GetPluginsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ruby_vagrant_ruby_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginsResponse.ProtoReflect.Descriptor instead.
func (*GetPluginsResponse) Descriptor() ([]byte, []int) {
	return file_proto_ruby_vagrant_ruby_server_proto_rawDescGZIP(), []int{0}
}

func (x *GetPluginsResponse) GetPlugins() []*Plugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the plugin
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// type of the plugin
	Type Plugin_Type `protobuf:"varint,2,opt,name=type,proto3,enum=hashicorp.vagrant.Plugin_Type" json:"type,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ruby_vagrant_ruby_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ruby_vagrant_ruby_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_proto_ruby_vagrant_ruby_server_proto_rawDescGZIP(), []int{1}
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetType() Plugin_Type {
	if x != nil {
		return x.Type
	}
	return Plugin_UNKNOWN
}

type ParseVagrantfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the Vagrantfile
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"` // TODO: might be good to add an option for passing cmd line args
}

func (x *ParseVagrantfileRequest) Reset() {
	*x = ParseVagrantfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ruby_vagrant_ruby_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseVagrantfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseVagrantfileRequest) ProtoMessage() {}

func (x *ParseVagrantfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ruby_vagrant_ruby_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseVagrantfileRequest.ProtoReflect.Descriptor instead.
func (*ParseVagrantfileRequest) Descriptor() ([]byte, []int) {
	return file_proto_ruby_vagrant_ruby_server_proto_rawDescGZIP(), []int{2}
}

func (x *ParseVagrantfileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ParseVagrantfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Vagrantfile representation
	Vagrantfile *vagrant_plugin_sdk.Vagrantfile_Vagrantfile `protobuf:"bytes,1,opt,name=vagrantfile,proto3" json:"vagrantfile,omitempty"`
}

func (x *ParseVagrantfileResponse) Reset() {
	*x = ParseVagrantfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ruby_vagrant_ruby_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseVagrantfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseVagrantfileResponse) ProtoMessage() {}

func (x *ParseVagrantfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ruby_vagrant_ruby_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseVagrantfileResponse.ProtoReflect.Descriptor instead.
func (*ParseVagrantfileResponse) Descriptor() ([]byte, []int) {
	return file_proto_ruby_vagrant_ruby_server_proto_rawDescGZIP(), []int{3}
}

func (x *ParseVagrantfileResponse) GetVagrantfile() *vagrant_plugin_sdk.Vagrantfile_Vagrantfile {
	if x != nil {
		return x.Vagrantfile
	}
	return nil
}

var File_proto_ruby_vagrant_ruby_server_proto protoreflect.FileDescriptor

var file_proto_ruby_vagrant_ruby_server_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x75, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x2f, 0x72, 0x75, 0x62, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x22, 0xb0, 0x02, 0x0a, 0x06, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xdd, 0x01,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4d, 0x4d, 0x55, 0x4e, 0x49, 0x43, 0x41, 0x54, 0x4f, 0x52,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x56, 0x49,
	0x44, 0x45, 0x52, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49,
	0x4f, 0x4e, 0x45, 0x52, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x59, 0x4e, 0x43, 0x45, 0x44,
	0x46, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x55, 0x54, 0x48,
	0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x4c,
	0x4f, 0x47, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09,
	0x4c, 0x4f, 0x47, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x41, 0x50, 0x50, 0x45, 0x52, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x10, 0x0c, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e, 0x49, 0x4e, 0x46,
	0x4f, 0x10, 0x0d, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55, 0x53, 0x48, 0x10, 0x0e, 0x22, 0x2d, 0x0a,
	0x17, 0x50, 0x61, 0x72, 0x73, 0x65, 0x56, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x6c, 0x0a, 0x18,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x56, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x76, 0x61, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x56, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x56, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0b, 0x76,
	0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x32, 0xc7, 0x01, 0x0a, 0x0b, 0x52,
	0x75, 0x62, 0x79, 0x56, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x25, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x76, 0x61, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x73, 0x65,
	0x56, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x56, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x56, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x76, 0x61, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x75, 0x62, 0x79, 0x5f,
	0x76, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ruby_vagrant_ruby_server_proto_rawDescOnce sync.Once
	file_proto_ruby_vagrant_ruby_server_proto_rawDescData = file_proto_ruby_vagrant_ruby_server_proto_rawDesc
)

func file_proto_ruby_vagrant_ruby_server_proto_rawDescGZIP() []byte {
	file_proto_ruby_vagrant_ruby_server_proto_rawDescOnce.Do(func() {
		file_proto_ruby_vagrant_ruby_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ruby_vagrant_ruby_server_proto_rawDescData)
	})
	return file_proto_ruby_vagrant_ruby_server_proto_rawDescData
}

var file_proto_ruby_vagrant_ruby_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_ruby_vagrant_ruby_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_ruby_vagrant_ruby_server_proto_goTypes = []interface{}{
	(Plugin_Type)(0),                                   // 0: hashicorp.vagrant.Plugin.Type
	(*GetPluginsResponse)(nil),                         // 1: hashicorp.vagrant.GetPluginsResponse
	(*Plugin)(nil),                                     // 2: hashicorp.vagrant.Plugin
	(*ParseVagrantfileRequest)(nil),                    // 3: hashicorp.vagrant.ParseVagrantfileRequest
	(*ParseVagrantfileResponse)(nil),                   // 4: hashicorp.vagrant.ParseVagrantfileResponse
	(*vagrant_plugin_sdk.Vagrantfile_Vagrantfile)(nil), // 5: hashicorp.vagrant.sdk.Vagrantfile.Vagrantfile
	(*emptypb.Empty)(nil),                              // 6: google.protobuf.Empty
}
var file_proto_ruby_vagrant_ruby_server_proto_depIdxs = []int32{
	2, // 0: hashicorp.vagrant.GetPluginsResponse.plugins:type_name -> hashicorp.vagrant.Plugin
	0, // 1: hashicorp.vagrant.Plugin.type:type_name -> hashicorp.vagrant.Plugin.Type
	5, // 2: hashicorp.vagrant.ParseVagrantfileResponse.vagrantfile:type_name -> hashicorp.vagrant.sdk.Vagrantfile.Vagrantfile
	6, // 3: hashicorp.vagrant.RubyVagrant.GetPlugins:input_type -> google.protobuf.Empty
	3, // 4: hashicorp.vagrant.RubyVagrant.ParseVagrantfile:input_type -> hashicorp.vagrant.ParseVagrantfileRequest
	1, // 5: hashicorp.vagrant.RubyVagrant.GetPlugins:output_type -> hashicorp.vagrant.GetPluginsResponse
	4, // 6: hashicorp.vagrant.RubyVagrant.ParseVagrantfile:output_type -> hashicorp.vagrant.ParseVagrantfileResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_ruby_vagrant_ruby_server_proto_init() }
func file_proto_ruby_vagrant_ruby_server_proto_init() {
	if File_proto_ruby_vagrant_ruby_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ruby_vagrant_ruby_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPluginsResponse); i {
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
		file_proto_ruby_vagrant_ruby_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_proto_ruby_vagrant_ruby_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseVagrantfileRequest); i {
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
		file_proto_ruby_vagrant_ruby_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseVagrantfileResponse); i {
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
			RawDescriptor: file_proto_ruby_vagrant_ruby_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ruby_vagrant_ruby_server_proto_goTypes,
		DependencyIndexes: file_proto_ruby_vagrant_ruby_server_proto_depIdxs,
		EnumInfos:         file_proto_ruby_vagrant_ruby_server_proto_enumTypes,
		MessageInfos:      file_proto_ruby_vagrant_ruby_server_proto_msgTypes,
	}.Build()
	File_proto_ruby_vagrant_ruby_server_proto = out.File
	file_proto_ruby_vagrant_ruby_server_proto_rawDesc = nil
	file_proto_ruby_vagrant_ruby_server_proto_goTypes = nil
	file_proto_ruby_vagrant_ruby_server_proto_depIdxs = nil
}
