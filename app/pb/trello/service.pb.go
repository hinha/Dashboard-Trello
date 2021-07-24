// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.15.6
// source: internal/trello/proto/service.proto

package trello

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_trello_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_internal_trello_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_internal_trello_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *Board) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Board) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Board  *Board `protobuf:"bytes,3,opt,name=board,proto3" json:"board,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_trello_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_internal_trello_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_internal_trello_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *Request) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Request) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_trello_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_internal_trello_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_internal_trello_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *Member) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Member) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardId               string `protobuf:"bytes,1,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	CardCategory         string `protobuf:"bytes,2,opt,name=card_category,json=cardCategory,proto3" json:"card_category,omitempty"`
	CardName             string `protobuf:"bytes,3,opt,name=card_name,json=cardName,proto3" json:"card_name,omitempty"`
	CardVotes            int64  `protobuf:"varint,4,opt,name=card_votes,json=cardVotes,proto3" json:"card_votes,omitempty"`
	CountCheckItems      int64  `protobuf:"varint,5,opt,name=count_check_items,json=countCheckItems,proto3" json:"count_check_items,omitempty"`
	CountCheckLists      int64  `protobuf:"varint,6,opt,name=count_check_lists,json=countCheckLists,proto3" json:"count_check_lists,omitempty"`
	CheckItemsComplete   int64  `protobuf:"varint,7,opt,name=check_items_complete,json=checkItemsComplete,proto3" json:"check_items_complete,omitempty"`
	CheckItemsIncomplete int64  `protobuf:"varint,8,opt,name=check_items_incomplete,json=checkItemsIncomplete,proto3" json:"check_items_incomplete,omitempty"`
	CommentCount         int64  `protobuf:"varint,9,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	AttachmentsCount     int64  `protobuf:"varint,10,opt,name=attachments_count,json=attachmentsCount,proto3" json:"attachments_count,omitempty"`
	Url                  string `protobuf:"bytes,11,opt,name=url,proto3" json:"url,omitempty"`
	MemberId             string `protobuf:"bytes,12,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	MemberName           string `protobuf:"bytes,13,opt,name=member_name,json=memberName,proto3" json:"member_name,omitempty"`
	MemberUsername       string `protobuf:"bytes,14,opt,name=member_username,json=memberUsername,proto3" json:"member_username,omitempty"`
	CreatedAt            int64  `protobuf:"varint,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_trello_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_internal_trello_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_internal_trello_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *Card) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *Card) GetCardCategory() string {
	if x != nil {
		return x.CardCategory
	}
	return ""
}

func (x *Card) GetCardName() string {
	if x != nil {
		return x.CardName
	}
	return ""
}

func (x *Card) GetCardVotes() int64 {
	if x != nil {
		return x.CardVotes
	}
	return 0
}

func (x *Card) GetCountCheckItems() int64 {
	if x != nil {
		return x.CountCheckItems
	}
	return 0
}

func (x *Card) GetCountCheckLists() int64 {
	if x != nil {
		return x.CountCheckLists
	}
	return 0
}

func (x *Card) GetCheckItemsComplete() int64 {
	if x != nil {
		return x.CheckItemsComplete
	}
	return 0
}

func (x *Card) GetCheckItemsIncomplete() int64 {
	if x != nil {
		return x.CheckItemsIncomplete
	}
	return 0
}

func (x *Card) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Card) GetAttachmentsCount() int64 {
	if x != nil {
		return x.AttachmentsCount
	}
	return 0
}

func (x *Card) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Card) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Card) GetMemberName() string {
	if x != nil {
		return x.MemberName
	}
	return ""
}

func (x *Card) GetMemberUsername() string {
	if x != nil {
		return x.MemberUsername
	}
	return ""
}

func (x *Card) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*Card `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	LastUpdate string  `protobuf:"bytes,2,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	Error      string  `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_trello_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_trello_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_internal_trello_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetData() []*Card {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetLastUpdate() string {
	if x != nil {
		return x.LastUpdate
	}
	return ""
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_internal_trello_proto_service_proto protoreflect.FileDescriptor

var file_internal_trello_proto_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x72, 0x65, 0x6c, 0x6c,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x2b, 0x0a, 0x05, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x20, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x05, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x22, 0x48, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xaa, 0x04,
	0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x61, 0x72, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2a, 0x0a, 0x11,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x60, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x32, 0x0a, 0x06,
	0x54, 0x72, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x28, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x16, 0x5a, 0x14, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x65, 0x6c, 0x6c,
	0x6f, 0x3b, 0x74, 0x72, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_trello_proto_service_proto_rawDescOnce sync.Once
	file_internal_trello_proto_service_proto_rawDescData = file_internal_trello_proto_service_proto_rawDesc
)

func file_internal_trello_proto_service_proto_rawDescGZIP() []byte {
	file_internal_trello_proto_service_proto_rawDescOnce.Do(func() {
		file_internal_trello_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_trello_proto_service_proto_rawDescData)
	})
	return file_internal_trello_proto_service_proto_rawDescData
}

var file_internal_trello_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_trello_proto_service_proto_goTypes = []interface{}{
	(*Board)(nil),    // 0: api.Board
	(*Request)(nil),  // 1: api.Request
	(*Member)(nil),   // 2: api.Member
	(*Card)(nil),     // 3: api.Card
	(*Response)(nil), // 4: api.Response
}
var file_internal_trello_proto_service_proto_depIdxs = []int32{
	0, // 0: api.Request.board:type_name -> api.Board
	3, // 1: api.Response.data:type_name -> api.Card
	1, // 2: api.Trello.getCard:input_type -> api.Request
	4, // 3: api.Trello.getCard:output_type -> api.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_trello_proto_service_proto_init() }
func file_internal_trello_proto_service_proto_init() {
	if File_internal_trello_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_trello_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
		file_internal_trello_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_internal_trello_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_internal_trello_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_internal_trello_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_internal_trello_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_trello_proto_service_proto_goTypes,
		DependencyIndexes: file_internal_trello_proto_service_proto_depIdxs,
		MessageInfos:      file_internal_trello_proto_service_proto_msgTypes,
	}.Build()
	File_internal_trello_proto_service_proto = out.File
	file_internal_trello_proto_service_proto_rawDesc = nil
	file_internal_trello_proto_service_proto_goTypes = nil
	file_internal_trello_proto_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TrelloClient is the client API for Trello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrelloClient interface {
	GetCard(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type trelloClient struct {
	cc grpc.ClientConnInterface
}

func NewTrelloClient(cc grpc.ClientConnInterface) TrelloClient {
	return &trelloClient{cc}
}

func (c *trelloClient) GetCard(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.Trello/getCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrelloServer is the server API for Trello service.
type TrelloServer interface {
	GetCard(context.Context, *Request) (*Response, error)
}

// UnimplementedTrelloServer can be embedded to have forward compatible implementations.
type UnimplementedTrelloServer struct {
}

func (*UnimplementedTrelloServer) GetCard(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}

func RegisterTrelloServer(s *grpc.Server, srv TrelloServer) {
	s.RegisterService(&_Trello_serviceDesc, srv)
}

func _Trello_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrelloServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Trello/GetCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrelloServer).GetCard(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Trello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Trello",
	HandlerType: (*TrelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCard",
			Handler:    _Trello_GetCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/trello/proto/service.proto",
}
