// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: program/program.proto

package program

import (
	common "github.com/dot-coaching/gen/go/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pricing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Currency string  `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount   float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Type     string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Pricing) Reset() {
	*x = Pricing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_program_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pricing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pricing) ProtoMessage() {}

func (x *Pricing) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pricing.ProtoReflect.Descriptor instead.
func (*Pricing) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{0}
}

func (x *Pricing) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pricing) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Pricing) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Pricing) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Coach struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Role   string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Coach) Reset() {
	*x = Coach{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_program_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coach) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coach) ProtoMessage() {}

func (x *Coach) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coach.ProtoReflect.Descriptor instead.
func (*Coach) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{1}
}

func (x *Coach) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Coach) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type Program struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StartDate    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Status       string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	CreatedBy    uint32                 `protobuf:"varint,7,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	Participants []string               `protobuf:"bytes,8,rep,name=participants,proto3" json:"participants,omitempty"`
	Coaches      []string               `protobuf:"bytes,9,rep,name=coaches,proto3" json:"coaches,omitempty"`
	Tags         []string               `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	Pricings     []*Pricing             `protobuf:"bytes,11,rep,name=pricings,proto3" json:"pricings,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Program) Reset() {
	*x = Program{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_program_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Program) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Program) ProtoMessage() {}

func (x *Program) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Program.ProtoReflect.Descriptor instead.
func (*Program) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{2}
}

func (x *Program) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Program) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Program) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Program) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Program) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Program) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Program) GetCreatedBy() uint32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *Program) GetParticipants() []string {
	if x != nil {
		return x.Participants
	}
	return nil
}

func (x *Program) GetCoaches() []string {
	if x != nil {
		return x.Coaches
	}
	return nil
}

func (x *Program) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Program) GetPricings() []*Pricing {
	if x != nil {
		return x.Pricings
	}
	return nil
}

func (x *Program) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Program) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateProgramRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	StartDate    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Status       string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	CreatedBy    uint32                 `protobuf:"varint,6,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	Participants []string               `protobuf:"bytes,7,rep,name=participants,proto3" json:"participants,omitempty"`
	Coaches      []string               `protobuf:"bytes,8,rep,name=coaches,proto3" json:"coaches,omitempty"`
	Tags         []string               `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	Pricings     []*Pricing             `protobuf:"bytes,10,rep,name=pricings,proto3" json:"pricings,omitempty"`
}

func (x *CreateProgramRequest) Reset() {
	*x = CreateProgramRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_program_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProgramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProgramRequest) ProtoMessage() {}

func (x *CreateProgramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProgramRequest.ProtoReflect.Descriptor instead.
func (*CreateProgramRequest) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProgramRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProgramRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProgramRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *CreateProgramRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *CreateProgramRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateProgramRequest) GetCreatedBy() uint32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *CreateProgramRequest) GetParticipants() []string {
	if x != nil {
		return x.Participants
	}
	return nil
}

func (x *CreateProgramRequest) GetCoaches() []string {
	if x != nil {
		return x.Coaches
	}
	return nil
}

func (x *CreateProgramRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateProgramRequest) GetPricings() []*Pricing {
	if x != nil {
		return x.Pricings
	}
	return nil
}

type UpdateProgramRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StartDate    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Status       string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	CreatedBy    uint32                 `protobuf:"varint,7,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	Participants []string               `protobuf:"bytes,8,rep,name=participants,proto3" json:"participants,omitempty"`
	Coaches      []string               `protobuf:"bytes,9,rep,name=coaches,proto3" json:"coaches,omitempty"`
	Tags         []string               `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	Pricings     []*Pricing             `protobuf:"bytes,11,rep,name=pricings,proto3" json:"pricings,omitempty"`
}

func (x *UpdateProgramRequest) Reset() {
	*x = UpdateProgramRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_program_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProgramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProgramRequest) ProtoMessage() {}

func (x *UpdateProgramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProgramRequest.ProtoReflect.Descriptor instead.
func (*UpdateProgramRequest) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProgramRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateProgramRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProgramRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateProgramRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *UpdateProgramRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *UpdateProgramRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateProgramRequest) GetCreatedBy() uint32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *UpdateProgramRequest) GetParticipants() []string {
	if x != nil {
		return x.Participants
	}
	return nil
}

func (x *UpdateProgramRequest) GetCoaches() []string {
	if x != nil {
		return x.Coaches
	}
	return nil
}

func (x *UpdateProgramRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdateProgramRequest) GetPricings() []*Pricing {
	if x != nil {
		return x.Pricings
	}
	return nil
}

type ListProgramRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor uint32 `protobuf:"varint,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListProgramRequest) Reset() {
	*x = ListProgramRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_program_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProgramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProgramRequest) ProtoMessage() {}

func (x *ListProgramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProgramRequest.ProtoReflect.Descriptor instead.
func (*ListProgramRequest) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{5}
}

func (x *ListProgramRequest) GetCursor() uint32 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *ListProgramRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListProgramResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Programs []*Program `protobuf:"bytes,1,rep,name=programs,proto3" json:"programs,omitempty"`
	Total    uint32     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Cursor   uint32     `protobuf:"varint,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListProgramResponse) Reset() {
	*x = ListProgramResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_program_program_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProgramResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProgramResponse) ProtoMessage() {}

func (x *ListProgramResponse) ProtoReflect() protoreflect.Message {
	mi := &file_program_program_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProgramResponse.ProtoReflect.Descriptor instead.
func (*ListProgramResponse) Descriptor() ([]byte, []int) {
	return file_program_program_proto_rawDescGZIP(), []int{6}
}

func (x *ListProgramResponse) GetPrograms() []*Program {
	if x != nil {
		return x.Programs
	}
	return nil
}

func (x *ListProgramResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListProgramResponse) GetCursor() uint32 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

var File_program_program_proto protoreflect.FileDescriptor

var file_program_program_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x33, 0x0a, 0x05, 0x43, 0x6f, 0x61,
	0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0xe9,
	0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x61,
	0x63, 0x68, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x70, 0x72,
	0x69, 0x63, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf2, 0x02, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x50, 0x72,
	0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x73, 0x22,
	0x82, 0x03, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x61, 0x63,
	0x68, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x61, 0x63, 0x68,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63,
	0x69, 0x6e, 0x67, 0x73, 0x22, 0x42, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x71, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x32, 0xdb, 0x02, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x7f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x74, 0x2d, 0x63, 0x6f, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0xca, 0x02, 0x07, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0xe2, 0x02, 0x13, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x07, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_program_program_proto_rawDescOnce sync.Once
	file_program_program_proto_rawDescData = file_program_program_proto_rawDesc
)

func file_program_program_proto_rawDescGZIP() []byte {
	file_program_program_proto_rawDescOnce.Do(func() {
		file_program_program_proto_rawDescData = protoimpl.X.CompressGZIP(file_program_program_proto_rawDescData)
	})
	return file_program_program_proto_rawDescData
}

var file_program_program_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_program_program_proto_goTypes = []any{
	(*Pricing)(nil),               // 0: program.Pricing
	(*Coach)(nil),                 // 1: program.Coach
	(*Program)(nil),               // 2: program.Program
	(*CreateProgramRequest)(nil),  // 3: program.CreateProgramRequest
	(*UpdateProgramRequest)(nil),  // 4: program.UpdateProgramRequest
	(*ListProgramRequest)(nil),    // 5: program.ListProgramRequest
	(*ListProgramResponse)(nil),   // 6: program.ListProgramResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*common.GetByIdRequest)(nil), // 8: common.GetByIdRequest
}
var file_program_program_proto_depIdxs = []int32{
	7,  // 0: program.Program.startDate:type_name -> google.protobuf.Timestamp
	7,  // 1: program.Program.endDate:type_name -> google.protobuf.Timestamp
	0,  // 2: program.Program.pricings:type_name -> program.Pricing
	7,  // 3: program.Program.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: program.Program.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: program.CreateProgramRequest.startDate:type_name -> google.protobuf.Timestamp
	7,  // 6: program.CreateProgramRequest.endDate:type_name -> google.protobuf.Timestamp
	0,  // 7: program.CreateProgramRequest.pricings:type_name -> program.Pricing
	7,  // 8: program.UpdateProgramRequest.startDate:type_name -> google.protobuf.Timestamp
	7,  // 9: program.UpdateProgramRequest.endDate:type_name -> google.protobuf.Timestamp
	0,  // 10: program.UpdateProgramRequest.pricings:type_name -> program.Pricing
	2,  // 11: program.ListProgramResponse.programs:type_name -> program.Program
	3,  // 12: program.ProgramService.CreateProgram:input_type -> program.CreateProgramRequest
	8,  // 13: program.ProgramService.GetProgram:input_type -> common.GetByIdRequest
	4,  // 14: program.ProgramService.UpdateProgram:input_type -> program.UpdateProgramRequest
	8,  // 15: program.ProgramService.DeleteProgram:input_type -> common.GetByIdRequest
	5,  // 16: program.ProgramService.ListProgram:input_type -> program.ListProgramRequest
	2,  // 17: program.ProgramService.CreateProgram:output_type -> program.Program
	2,  // 18: program.ProgramService.GetProgram:output_type -> program.Program
	2,  // 19: program.ProgramService.UpdateProgram:output_type -> program.Program
	2,  // 20: program.ProgramService.DeleteProgram:output_type -> program.Program
	6,  // 21: program.ProgramService.ListProgram:output_type -> program.ListProgramResponse
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_program_program_proto_init() }
func file_program_program_proto_init() {
	if File_program_program_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_program_program_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Pricing); i {
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
		file_program_program_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Coach); i {
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
		file_program_program_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Program); i {
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
		file_program_program_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateProgramRequest); i {
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
		file_program_program_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProgramRequest); i {
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
		file_program_program_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListProgramRequest); i {
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
		file_program_program_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListProgramResponse); i {
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
			RawDescriptor: file_program_program_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_program_program_proto_goTypes,
		DependencyIndexes: file_program_program_proto_depIdxs,
		MessageInfos:      file_program_program_proto_msgTypes,
	}.Build()
	File_program_program_proto = out.File
	file_program_program_proto_rawDesc = nil
	file_program_program_proto_goTypes = nil
	file_program_program_proto_depIdxs = nil
}
