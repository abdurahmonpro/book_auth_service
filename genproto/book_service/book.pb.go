// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: book.proto

package book_service

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Isbn      string `protobuf:"bytes,2,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Cover     string `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Author    string `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	Published string `protobuf:"bytes,6,opt,name=published,proto3" json:"published,omitempty"`
	Pages     int32  `protobuf:"varint,7,opt,name=pages,proto3" json:"pages,omitempty"`
	Status    int32  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"` // 0-new, 1-reading, 2-finished,
	CreatedAt string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetIsbn() string {
	if x != nil {
		return x.Isbn
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *Book) GetPages() int32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *Book) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Book) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Book) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isbn      string `protobuf:"bytes,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Cover     string `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Author    string `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	Published string `protobuf:"bytes,5,opt,name=published,proto3" json:"published,omitempty"`
	Pages     int32  `protobuf:"varint,7,opt,name=pages,proto3" json:"pages,omitempty"`
	Status    int32  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"` // 0-new, 1-reading, 2-finished,
}

func (x *CreateBook) Reset() {
	*x = CreateBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBook) ProtoMessage() {}

func (x *CreateBook) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBook.ProtoReflect.Descriptor instead.
func (*CreateBook) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBook) GetIsbn() string {
	if x != nil {
		return x.Isbn
	}
	return ""
}

func (x *CreateBook) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBook) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *CreateBook) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *CreateBook) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *CreateBook) GetPages() int32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *CreateBook) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UpdateBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Isbn      string `protobuf:"bytes,2,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Cover     string `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Author    string `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	Published string `protobuf:"bytes,6,opt,name=published,proto3" json:"published,omitempty"`
	Pages     int32  `protobuf:"varint,7,opt,name=pages,proto3" json:"pages,omitempty"`
	Status    int32  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"` // 0-new, 1-reading, 2-finished,
}

func (x *UpdateBook) Reset() {
	*x = UpdateBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBook) ProtoMessage() {}

func (x *UpdateBook) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBook.ProtoReflect.Descriptor instead.
func (*UpdateBook) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBook) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBook) GetIsbn() string {
	if x != nil {
		return x.Isbn
	}
	return ""
}

func (x *UpdateBook) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateBook) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *UpdateBook) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *UpdateBook) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *UpdateBook) GetPages() int32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *UpdateBook) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UpdatePatchBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields *_struct.Struct `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *UpdatePatchBook) Reset() {
	*x = UpdatePatchBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePatchBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePatchBook) ProtoMessage() {}

func (x *UpdatePatchBook) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePatchBook.ProtoReflect.Descriptor instead.
func (*UpdatePatchBook) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePatchBook) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePatchBook) GetFields() *_struct.Struct {
	if x != nil {
		return x.Fields
	}
	return nil
}

type BookPK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BookPK) Reset() {
	*x = BookPK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookPK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookPK) ProtoMessage() {}

func (x *BookPK) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookPK.ProtoReflect.Descriptor instead.
func (*BookPK) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{4}
}

func (x *BookPK) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BookByTitle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *BookByTitle) Reset() {
	*x = BookByTitle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookByTitle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookByTitle) ProtoMessage() {}

func (x *BookByTitle) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookByTitle.ProtoReflect.Descriptor instead.
func (*BookByTitle) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{5}
}

func (x *BookByTitle) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type BookListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *BookListRequest) Reset() {
	*x = BookListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookListRequest) ProtoMessage() {}

func (x *BookListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookListRequest.ProtoReflect.Descriptor instead.
func (*BookListRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{6}
}

func (x *BookListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *BookListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *BookListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type BookListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Books []*Book `protobuf:"bytes,2,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BookListResponse) Reset() {
	*x = BookListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookListResponse) ProtoMessage() {}

func (x *BookListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookListResponse.ProtoReflect.Descriptor instead.
func (*BookListResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{7}
}

func (x *BookListResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *BookListResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_book_proto protoreflect.FileDescriptor

var file_book_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x73, 0x62, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x52, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x18, 0x0a,
	0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x4b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x42,
	0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x57, 0x0a, 0x0f,
	0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x52, 0x0a, 0x10, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x28, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_proto_rawDescOnce sync.Once
	file_book_proto_rawDescData = file_book_proto_rawDesc
)

func file_book_proto_rawDescGZIP() []byte {
	file_book_proto_rawDescOnce.Do(func() {
		file_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_proto_rawDescData)
	})
	return file_book_proto_rawDescData
}

var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_book_proto_goTypes = []interface{}{
	(*Book)(nil),             // 0: book_service.Book
	(*CreateBook)(nil),       // 1: book_service.CreateBook
	(*UpdateBook)(nil),       // 2: book_service.UpdateBook
	(*UpdatePatchBook)(nil),  // 3: book_service.UpdatePatchBook
	(*BookPK)(nil),           // 4: book_service.BookPK
	(*BookByTitle)(nil),      // 5: book_service.BookByTitle
	(*BookListRequest)(nil),  // 6: book_service.BookListRequest
	(*BookListResponse)(nil), // 7: book_service.BookListResponse
	(*_struct.Struct)(nil),   // 8: google.protobuf.Struct
}
var file_book_proto_depIdxs = []int32{
	8, // 0: book_service.UpdatePatchBook.fields:type_name -> google.protobuf.Struct
	0, // 1: book_service.BookListResponse.books:type_name -> book_service.Book
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_book_proto_init() }
func file_book_proto_init() {
	if File_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBook); i {
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
		file_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBook); i {
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
		file_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePatchBook); i {
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
		file_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookPK); i {
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
		file_book_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookByTitle); i {
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
		file_book_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookListRequest); i {
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
		file_book_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookListResponse); i {
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
			RawDescriptor: file_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_book_proto_goTypes,
		DependencyIndexes: file_book_proto_depIdxs,
		MessageInfos:      file_book_proto_msgTypes,
	}.Build()
	File_book_proto = out.File
	file_book_proto_rawDesc = nil
	file_book_proto_goTypes = nil
	file_book_proto_depIdxs = nil
}