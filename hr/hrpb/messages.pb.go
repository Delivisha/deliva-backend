// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: hrpb/messages.proto

package hrpb

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

type DepartmentCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	HeadOfDepartment  string               `protobuf:"bytes,3,opt,name=head_of_department,json=headOfDepartment,proto3" json:"head_of_department,omitempty"`
	NumberOfEmployees int32                `protobuf:"varint,4,opt,name=number_of_employees,json=numberOfEmployees,proto3" json:"number_of_employees,omitempty"`
	Projects          []*DepartmentProject `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *DepartmentCreated) Reset() {
	*x = DepartmentCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrpb_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartmentCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartmentCreated) ProtoMessage() {}

func (x *DepartmentCreated) ProtoReflect() protoreflect.Message {
	mi := &file_hrpb_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartmentCreated.ProtoReflect.Descriptor instead.
func (*DepartmentCreated) Descriptor() ([]byte, []int) {
	return file_hrpb_messages_proto_rawDescGZIP(), []int{0}
}

func (x *DepartmentCreated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DepartmentCreated) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DepartmentCreated) GetHeadOfDepartment() string {
	if x != nil {
		return x.HeadOfDepartment
	}
	return ""
}

func (x *DepartmentCreated) GetNumberOfEmployees() int32 {
	if x != nil {
		return x.NumberOfEmployees
	}
	return 0
}

func (x *DepartmentCreated) GetProjects() []*DepartmentProject {
	if x != nil {
		return x.Projects
	}
	return nil
}

type ProjectCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *DepartmentProject `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ProjectCreated) Reset() {
	*x = ProjectCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrpb_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectCreated) ProtoMessage() {}

func (x *ProjectCreated) ProtoReflect() protoreflect.Message {
	mi := &file_hrpb_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectCreated.ProtoReflect.Descriptor instead.
func (*ProjectCreated) Descriptor() ([]byte, []int) {
	return file_hrpb_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectCreated) GetProject() *DepartmentProject {
	if x != nil {
		return x.Project
	}
	return nil
}

type DepartmentProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Duration          string `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	NumberOfEmployees int32  `protobuf:"varint,4,opt,name=number_of_employees,json=numberOfEmployees,proto3" json:"number_of_employees,omitempty"`
	Budget            int32  `protobuf:"varint,5,opt,name=budget,proto3" json:"budget,omitempty"`
	Resources         string `protobuf:"bytes,6,opt,name=resources,proto3" json:"resources,omitempty"`
}

func (x *DepartmentProject) Reset() {
	*x = DepartmentProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrpb_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartmentProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartmentProject) ProtoMessage() {}

func (x *DepartmentProject) ProtoReflect() protoreflect.Message {
	mi := &file_hrpb_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartmentProject.ProtoReflect.Descriptor instead.
func (*DepartmentProject) Descriptor() ([]byte, []int) {
	return file_hrpb_messages_proto_rawDescGZIP(), []int{2}
}

func (x *DepartmentProject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DepartmentProject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DepartmentProject) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *DepartmentProject) GetNumberOfEmployees() int32 {
	if x != nil {
		return x.NumberOfEmployees
	}
	return 0
}

func (x *DepartmentProject) GetBudget() int32 {
	if x != nil {
		return x.Budget
	}
	return 0
}

func (x *DepartmentProject) GetResources() string {
	if x != nil {
		return x.Resources
	}
	return ""
}

type EmployeeCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName     string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName    string `protobuf:"bytes,3,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName      string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	DateOfBirth   string `protobuf:"bytes,5,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	HomeAddress   string `protobuf:"bytes,6,opt,name=home_address,json=homeAddress,proto3" json:"home_address,omitempty"`
	PhoneNumber   string `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Bank          string `protobuf:"bytes,8,opt,name=bank,proto3" json:"bank,omitempty"`
	AccountNumber string `protobuf:"bytes,9,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	GrossSalary   int32  `protobuf:"varint,10,opt,name=gross_salary,json=grossSalary,proto3" json:"gross_salary,omitempty"`
	Department    string `protobuf:"bytes,11,opt,name=Department,proto3" json:"Department,omitempty"`
	Country       string `protobuf:"bytes,12,opt,name=Country,proto3" json:"Country,omitempty"`
	Suspended     bool   `protobuf:"varint,13,opt,name=suspended,proto3" json:"suspended,omitempty"`
	Sacked        bool   `protobuf:"varint,14,opt,name=sacked,proto3" json:"sacked,omitempty"`
}

func (x *EmployeeCreated) Reset() {
	*x = EmployeeCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrpb_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeCreated) ProtoMessage() {}

func (x *EmployeeCreated) ProtoReflect() protoreflect.Message {
	mi := &file_hrpb_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeCreated.ProtoReflect.Descriptor instead.
func (*EmployeeCreated) Descriptor() ([]byte, []int) {
	return file_hrpb_messages_proto_rawDescGZIP(), []int{3}
}

func (x *EmployeeCreated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EmployeeCreated) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *EmployeeCreated) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *EmployeeCreated) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *EmployeeCreated) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *EmployeeCreated) GetHomeAddress() string {
	if x != nil {
		return x.HomeAddress
	}
	return ""
}

func (x *EmployeeCreated) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *EmployeeCreated) GetBank() string {
	if x != nil {
		return x.Bank
	}
	return ""
}

func (x *EmployeeCreated) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *EmployeeCreated) GetGrossSalary() int32 {
	if x != nil {
		return x.GrossSalary
	}
	return 0
}

func (x *EmployeeCreated) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *EmployeeCreated) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *EmployeeCreated) GetSuspended() bool {
	if x != nil {
		return x.Suspended
	}
	return false
}

func (x *EmployeeCreated) GetSacked() bool {
	if x != nil {
		return x.Sacked
	}
	return false
}

type EmployeeSuspended struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Reason   string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Duration string `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *EmployeeSuspended) Reset() {
	*x = EmployeeSuspended{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrpb_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeSuspended) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeSuspended) ProtoMessage() {}

func (x *EmployeeSuspended) ProtoReflect() protoreflect.Message {
	mi := &file_hrpb_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeSuspended.ProtoReflect.Descriptor instead.
func (*EmployeeSuspended) Descriptor() ([]byte, []int) {
	return file_hrpb_messages_proto_rawDescGZIP(), []int{4}
}

func (x *EmployeeSuspended) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EmployeeSuspended) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *EmployeeSuspended) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

type EmployeeUnsuspended struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EmployeeUnsuspended) Reset() {
	*x = EmployeeUnsuspended{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrpb_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeUnsuspended) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeUnsuspended) ProtoMessage() {}

func (x *EmployeeUnsuspended) ProtoReflect() protoreflect.Message {
	mi := &file_hrpb_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeUnsuspended.ProtoReflect.Descriptor instead.
func (*EmployeeUnsuspended) Descriptor() ([]byte, []int) {
	return file_hrpb_messages_proto_rawDescGZIP(), []int{5}
}

func (x *EmployeeUnsuspended) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmployeeSacked struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *EmployeeSacked) Reset() {
	*x = EmployeeSacked{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hrpb_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeSacked) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeSacked) ProtoMessage() {}

func (x *EmployeeSacked) ProtoReflect() protoreflect.Message {
	mi := &file_hrpb_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeSacked.ProtoReflect.Descriptor instead.
func (*EmployeeSacked) Descriptor() ([]byte, []int) {
	return file_hrpb_messages_proto_rawDescGZIP(), []int{6}
}

func (x *EmployeeSacked) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EmployeeSacked) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_hrpb_messages_proto protoreflect.FileDescriptor

var file_hrpb_messages_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x72, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x72, 0x70, 0x62, 0x22, 0xca, 0x01, 0x0a, 0x11,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x66,
	0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x68, 0x65, 0x61, 0x64, 0x4f, 0x66, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66,
	0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x72, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x43, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x72,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0xb9, 0x01,
	0x0a, 0x11, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66,
	0x5f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0xb6, 0x03, 0x0a, 0x0f, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x21,
	0x0a, 0x0c, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x53, 0x61, 0x6c, 0x61,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61,
	0x63, 0x6b, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x61, 0x63, 0x6b,
	0x65, 0x64, 0x22, 0x57, 0x0a, 0x11, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x75,
	0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x13, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x55, 0x6e, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x38, 0x0a, 0x0e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x61,
	0x63, 0x6b, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x65, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x72, 0x70, 0x62, 0x42, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x61, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x72, 0x70, 0x62,
	0x2f, 0x68, 0x72, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x48, 0x72,
	0x70, 0x62, 0xca, 0x02, 0x04, 0x48, 0x72, 0x70, 0x62, 0xe2, 0x02, 0x10, 0x48, 0x72, 0x70, 0x62,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x48,
	0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hrpb_messages_proto_rawDescOnce sync.Once
	file_hrpb_messages_proto_rawDescData = file_hrpb_messages_proto_rawDesc
)

func file_hrpb_messages_proto_rawDescGZIP() []byte {
	file_hrpb_messages_proto_rawDescOnce.Do(func() {
		file_hrpb_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_hrpb_messages_proto_rawDescData)
	})
	return file_hrpb_messages_proto_rawDescData
}

var file_hrpb_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_hrpb_messages_proto_goTypes = []interface{}{
	(*DepartmentCreated)(nil),   // 0: hrpb.DepartmentCreated
	(*ProjectCreated)(nil),      // 1: hrpb.ProjectCreated
	(*DepartmentProject)(nil),   // 2: hrpb.DepartmentProject
	(*EmployeeCreated)(nil),     // 3: hrpb.EmployeeCreated
	(*EmployeeSuspended)(nil),   // 4: hrpb.EmployeeSuspended
	(*EmployeeUnsuspended)(nil), // 5: hrpb.EmployeeUnsuspended
	(*EmployeeSacked)(nil),      // 6: hrpb.EmployeeSacked
}
var file_hrpb_messages_proto_depIdxs = []int32{
	2, // 0: hrpb.DepartmentCreated.projects:type_name -> hrpb.DepartmentProject
	2, // 1: hrpb.ProjectCreated.project:type_name -> hrpb.DepartmentProject
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hrpb_messages_proto_init() }
func file_hrpb_messages_proto_init() {
	if File_hrpb_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hrpb_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartmentCreated); i {
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
		file_hrpb_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectCreated); i {
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
		file_hrpb_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartmentProject); i {
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
		file_hrpb_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeCreated); i {
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
		file_hrpb_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeSuspended); i {
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
		file_hrpb_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeUnsuspended); i {
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
		file_hrpb_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeSacked); i {
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
			RawDescriptor: file_hrpb_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hrpb_messages_proto_goTypes,
		DependencyIndexes: file_hrpb_messages_proto_depIdxs,
		MessageInfos:      file_hrpb_messages_proto_msgTypes,
	}.Build()
	File_hrpb_messages_proto = out.File
	file_hrpb_messages_proto_rawDesc = nil
	file_hrpb_messages_proto_goTypes = nil
	file_hrpb_messages_proto_depIdxs = nil
}