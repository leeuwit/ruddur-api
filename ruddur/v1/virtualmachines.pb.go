// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: ruddur/v1/virtualmachines.proto

package ruddurv1

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

type VirtualMachine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status  string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Cpu     string   `protobuf:"bytes,3,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory  string   `protobuf:"bytes,4,opt,name=memory,proto3" json:"memory,omitempty"`
	Storage []string `protobuf:"bytes,5,rep,name=storage,proto3" json:"storage,omitempty"`
	Boot    string   `protobuf:"bytes,6,opt,name=boot,proto3" json:"boot,omitempty"`
}

func (x *VirtualMachine) Reset() {
	*x = VirtualMachine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualMachine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachine) ProtoMessage() {}

func (x *VirtualMachine) ProtoReflect() protoreflect.Message {
	mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachine.ProtoReflect.Descriptor instead.
func (*VirtualMachine) Descriptor() ([]byte, []int) {
	return file_ruddur_v1_virtualmachines_proto_rawDescGZIP(), []int{0}
}

func (x *VirtualMachine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VirtualMachine) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VirtualMachine) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *VirtualMachine) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *VirtualMachine) GetStorage() []string {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *VirtualMachine) GetBoot() string {
	if x != nil {
		return x.Boot
	}
	return ""
}

type VirtualMachineServiceStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *VirtualMachineServiceStatusRequest) Reset() {
	*x = VirtualMachineServiceStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualMachineServiceStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachineServiceStatusRequest) ProtoMessage() {}

func (x *VirtualMachineServiceStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachineServiceStatusRequest.ProtoReflect.Descriptor instead.
func (*VirtualMachineServiceStatusRequest) Descriptor() ([]byte, []int) {
	return file_ruddur_v1_virtualmachines_proto_rawDescGZIP(), []int{1}
}

func (x *VirtualMachineServiceStatusRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *VirtualMachineServiceStatusRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type VirtualMachineServiceStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VirtualMachineServiceStatusResponse) Reset() {
	*x = VirtualMachineServiceStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualMachineServiceStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachineServiceStatusResponse) ProtoMessage() {}

func (x *VirtualMachineServiceStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachineServiceStatusResponse.ProtoReflect.Descriptor instead.
func (*VirtualMachineServiceStatusResponse) Descriptor() ([]byte, []int) {
	return file_ruddur_v1_virtualmachines_proto_rawDescGZIP(), []int{2}
}

func (x *VirtualMachineServiceStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type VirtualMachineServiceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *VirtualMachineServiceListRequest) Reset() {
	*x = VirtualMachineServiceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualMachineServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachineServiceListRequest) ProtoMessage() {}

func (x *VirtualMachineServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachineServiceListRequest.ProtoReflect.Descriptor instead.
func (*VirtualMachineServiceListRequest) Descriptor() ([]byte, []int) {
	return file_ruddur_v1_virtualmachines_proto_rawDescGZIP(), []int{3}
}

func (x *VirtualMachineServiceListRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

type VirtualMachineServiceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vms []*VirtualMachine `protobuf:"bytes,1,rep,name=vms,proto3" json:"vms,omitempty"`
}

func (x *VirtualMachineServiceListResponse) Reset() {
	*x = VirtualMachineServiceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualMachineServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachineServiceListResponse) ProtoMessage() {}

func (x *VirtualMachineServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachineServiceListResponse.ProtoReflect.Descriptor instead.
func (*VirtualMachineServiceListResponse) Descriptor() ([]byte, []int) {
	return file_ruddur_v1_virtualmachines_proto_rawDescGZIP(), []int{4}
}

func (x *VirtualMachineServiceListResponse) GetVms() []*VirtualMachine {
	if x != nil {
		return x.Vms
	}
	return nil
}

type VirtualMachineServiceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string   `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cpu     string   `protobuf:"bytes,3,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory  string   `protobuf:"bytes,4,opt,name=memory,proto3" json:"memory,omitempty"`
	Storage []string `protobuf:"bytes,5,rep,name=storage,proto3" json:"storage,omitempty"`
	Boot    string   `protobuf:"bytes,6,opt,name=boot,proto3" json:"boot,omitempty"`
}

func (x *VirtualMachineServiceCreateRequest) Reset() {
	*x = VirtualMachineServiceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualMachineServiceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachineServiceCreateRequest) ProtoMessage() {}

func (x *VirtualMachineServiceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachineServiceCreateRequest.ProtoReflect.Descriptor instead.
func (*VirtualMachineServiceCreateRequest) Descriptor() ([]byte, []int) {
	return file_ruddur_v1_virtualmachines_proto_rawDescGZIP(), []int{5}
}

func (x *VirtualMachineServiceCreateRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *VirtualMachineServiceCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VirtualMachineServiceCreateRequest) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *VirtualMachineServiceCreateRequest) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *VirtualMachineServiceCreateRequest) GetStorage() []string {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *VirtualMachineServiceCreateRequest) GetBoot() string {
	if x != nil {
		return x.Boot
	}
	return ""
}

type VirtualMachineServiceCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VirtualMachineServiceCreateResponse) Reset() {
	*x = VirtualMachineServiceCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualMachineServiceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachineServiceCreateResponse) ProtoMessage() {}

func (x *VirtualMachineServiceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachineServiceCreateResponse.ProtoReflect.Descriptor instead.
func (*VirtualMachineServiceCreateResponse) Descriptor() ([]byte, []int) {
	return file_ruddur_v1_virtualmachines_proto_rawDescGZIP(), []int{6}
}

func (x *VirtualMachineServiceCreateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VirtualMachineServiceCreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type VirtualMachineServiceDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *VirtualMachineServiceDeleteRequest) Reset() {
	*x = VirtualMachineServiceDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualMachineServiceDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachineServiceDeleteRequest) ProtoMessage() {}

func (x *VirtualMachineServiceDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachineServiceDeleteRequest.ProtoReflect.Descriptor instead.
func (*VirtualMachineServiceDeleteRequest) Descriptor() ([]byte, []int) {
	return file_ruddur_v1_virtualmachines_proto_rawDescGZIP(), []int{7}
}

func (x *VirtualMachineServiceDeleteRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *VirtualMachineServiceDeleteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type VirtualMachineServiceDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VirtualMachineServiceDeleteResponse) Reset() {
	*x = VirtualMachineServiceDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualMachineServiceDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualMachineServiceDeleteResponse) ProtoMessage() {}

func (x *VirtualMachineServiceDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ruddur_v1_virtualmachines_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualMachineServiceDeleteResponse.ProtoReflect.Descriptor instead.
func (*VirtualMachineServiceDeleteResponse) Descriptor() ([]byte, []int) {
	return file_ruddur_v1_virtualmachines_proto_rawDescGZIP(), []int{8}
}

func (x *VirtualMachineServiceDeleteResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VirtualMachineServiceDeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_ruddur_v1_virtualmachines_proto protoreflect.FileDescriptor

var file_ruddur_v1_virtualmachines_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x6f,
	0x74, 0x22, 0x52, 0x0a, 0x22, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x23, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x3c, 0x0a, 0x20, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x50, 0x0a, 0x21, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x76, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52,
	0x03, 0x76, 0x6d, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x22, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x6f,
	0x74, 0x22, 0x57, 0x0a, 0x23, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x52, 0x0a, 0x22, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x57,
	0x0a, 0x23, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x89, 0x05, 0x0a, 0x15, 0x56, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x9d, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x2e, 0x72,
	0x75, 0x64, 0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x72, 0x75,
	0x64, 0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x30,
	0x01, 0x12, 0x93, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x72, 0x75, 0x64,
	0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a,
	0x22, 0x25, 0x2f, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x9b, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x2d, 0x2e, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x72,
	0x75, 0x64, 0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x9b, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x2d, 0x2e, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x72, 0x75, 0x64,
	0x64, 0x75, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x9c, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x75, 0x64, 0x64,
	0x75, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x65, 0x75, 0x77, 0x69,
	0x74, 0x2f, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x75, 0x64,
	0x64, 0x75, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x75, 0x64, 0x64, 0x75, 0x72, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x52, 0x75, 0x64, 0x64, 0x75, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x09, 0x52, 0x75, 0x64, 0x64, 0x75, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15,
	0x52, 0x75, 0x64, 0x64, 0x75, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x52, 0x75, 0x64, 0x64, 0x75, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ruddur_v1_virtualmachines_proto_rawDescOnce sync.Once
	file_ruddur_v1_virtualmachines_proto_rawDescData = file_ruddur_v1_virtualmachines_proto_rawDesc
)

func file_ruddur_v1_virtualmachines_proto_rawDescGZIP() []byte {
	file_ruddur_v1_virtualmachines_proto_rawDescOnce.Do(func() {
		file_ruddur_v1_virtualmachines_proto_rawDescData = protoimpl.X.CompressGZIP(file_ruddur_v1_virtualmachines_proto_rawDescData)
	})
	return file_ruddur_v1_virtualmachines_proto_rawDescData
}

var file_ruddur_v1_virtualmachines_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ruddur_v1_virtualmachines_proto_goTypes = []interface{}{
	(*VirtualMachine)(nil),                      // 0: ruddur.v1.VirtualMachine
	(*VirtualMachineServiceStatusRequest)(nil),  // 1: ruddur.v1.VirtualMachineServiceStatusRequest
	(*VirtualMachineServiceStatusResponse)(nil), // 2: ruddur.v1.VirtualMachineServiceStatusResponse
	(*VirtualMachineServiceListRequest)(nil),    // 3: ruddur.v1.VirtualMachineServiceListRequest
	(*VirtualMachineServiceListResponse)(nil),   // 4: ruddur.v1.VirtualMachineServiceListResponse
	(*VirtualMachineServiceCreateRequest)(nil),  // 5: ruddur.v1.VirtualMachineServiceCreateRequest
	(*VirtualMachineServiceCreateResponse)(nil), // 6: ruddur.v1.VirtualMachineServiceCreateResponse
	(*VirtualMachineServiceDeleteRequest)(nil),  // 7: ruddur.v1.VirtualMachineServiceDeleteRequest
	(*VirtualMachineServiceDeleteResponse)(nil), // 8: ruddur.v1.VirtualMachineServiceDeleteResponse
}
var file_ruddur_v1_virtualmachines_proto_depIdxs = []int32{
	0, // 0: ruddur.v1.VirtualMachineServiceListResponse.vms:type_name -> ruddur.v1.VirtualMachine
	1, // 1: ruddur.v1.VirtualMachineService.Status:input_type -> ruddur.v1.VirtualMachineServiceStatusRequest
	3, // 2: ruddur.v1.VirtualMachineService.List:input_type -> ruddur.v1.VirtualMachineServiceListRequest
	5, // 3: ruddur.v1.VirtualMachineService.Create:input_type -> ruddur.v1.VirtualMachineServiceCreateRequest
	7, // 4: ruddur.v1.VirtualMachineService.Delete:input_type -> ruddur.v1.VirtualMachineServiceDeleteRequest
	2, // 5: ruddur.v1.VirtualMachineService.Status:output_type -> ruddur.v1.VirtualMachineServiceStatusResponse
	4, // 6: ruddur.v1.VirtualMachineService.List:output_type -> ruddur.v1.VirtualMachineServiceListResponse
	6, // 7: ruddur.v1.VirtualMachineService.Create:output_type -> ruddur.v1.VirtualMachineServiceCreateResponse
	8, // 8: ruddur.v1.VirtualMachineService.Delete:output_type -> ruddur.v1.VirtualMachineServiceDeleteResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ruddur_v1_virtualmachines_proto_init() }
func file_ruddur_v1_virtualmachines_proto_init() {
	if File_ruddur_v1_virtualmachines_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ruddur_v1_virtualmachines_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualMachine); i {
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
		file_ruddur_v1_virtualmachines_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualMachineServiceStatusRequest); i {
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
		file_ruddur_v1_virtualmachines_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualMachineServiceStatusResponse); i {
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
		file_ruddur_v1_virtualmachines_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualMachineServiceListRequest); i {
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
		file_ruddur_v1_virtualmachines_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualMachineServiceListResponse); i {
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
		file_ruddur_v1_virtualmachines_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualMachineServiceCreateRequest); i {
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
		file_ruddur_v1_virtualmachines_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualMachineServiceCreateResponse); i {
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
		file_ruddur_v1_virtualmachines_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualMachineServiceDeleteRequest); i {
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
		file_ruddur_v1_virtualmachines_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualMachineServiceDeleteResponse); i {
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
			RawDescriptor: file_ruddur_v1_virtualmachines_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ruddur_v1_virtualmachines_proto_goTypes,
		DependencyIndexes: file_ruddur_v1_virtualmachines_proto_depIdxs,
		MessageInfos:      file_ruddur_v1_virtualmachines_proto_msgTypes,
	}.Build()
	File_ruddur_v1_virtualmachines_proto = out.File
	file_ruddur_v1_virtualmachines_proto_rawDesc = nil
	file_ruddur_v1_virtualmachines_proto_goTypes = nil
	file_ruddur_v1_virtualmachines_proto_depIdxs = nil
}
