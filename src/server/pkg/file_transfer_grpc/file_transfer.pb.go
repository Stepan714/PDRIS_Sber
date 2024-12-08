// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: file_transfer.proto

package file_transfer_grpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_file_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *GetFileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkData []byte `protobuf:"bytes,1,opt,name=chunk_data,json=chunkData,proto3" json:"chunk_data,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_file_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_file_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *FileChunk) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

type GetFileListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFileListRequest) Reset() {
	*x = GetFileListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileListRequest) ProtoMessage() {}

func (x *GetFileListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileListRequest.ProtoReflect.Descriptor instead.
func (*GetFileListRequest) Descriptor() ([]byte, []int) {
	return file_file_transfer_proto_rawDescGZIP(), []int{2}
}

type GetFileListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *GetFileListResponse) Reset() {
	*x = GetFileListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileListResponse) ProtoMessage() {}

func (x *GetFileListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileListResponse.ProtoReflect.Descriptor instead.
func (*GetFileListResponse) Descriptor() ([]byte, []int) {
	return file_file_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *GetFileListResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type GetFileInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetFileInfoRequest) Reset() {
	*x = GetFileInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_transfer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileInfoRequest) ProtoMessage() {}

func (x *GetFileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_transfer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileInfoRequest.ProtoReflect.Descriptor instead.
func (*GetFileInfoRequest) Descriptor() ([]byte, []int) {
	return file_file_transfer_proto_rawDescGZIP(), []int{4}
}

func (x *GetFileInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetFileInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size             string `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
	Type             string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	CreationDate     string `protobuf:"bytes,4,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	LastModifiedDate string `protobuf:"bytes,5,opt,name=last_modified_date,json=lastModifiedDate,proto3" json:"last_modified_date,omitempty"`
	AccessRights     string `protobuf:"bytes,6,opt,name=access_rights,json=accessRights,proto3" json:"access_rights,omitempty"`
	Location         string `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *GetFileInfoResponse) Reset() {
	*x = GetFileInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_transfer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileInfoResponse) ProtoMessage() {}

func (x *GetFileInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_transfer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileInfoResponse.ProtoReflect.Descriptor instead.
func (*GetFileInfoResponse) Descriptor() ([]byte, []int) {
	return file_file_transfer_proto_rawDescGZIP(), []int{5}
}

func (x *GetFileInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetFileInfoResponse) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *GetFileInfoResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetFileInfoResponse) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
}

func (x *GetFileInfoResponse) GetLastModifiedDate() string {
	if x != nil {
		return x.LastModifiedDate
	}
	return ""
}

func (x *GetFileInfoResponse) GetAccessRights() string {
	if x != nil {
		return x.AccessRights
	}
	return ""
}

func (x *GetFileInfoResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_file_transfer_proto protoreflect.FileDescriptor

var file_file_transfer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a,
	0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x2b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x61, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xf0,
	0x02, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x22, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x6e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x7a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x69, 0x6e, 0x66,
	0x6f, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_file_transfer_proto_rawDescOnce sync.Once
	file_file_transfer_proto_rawDescData = file_file_transfer_proto_rawDesc
)

func file_file_transfer_proto_rawDescGZIP() []byte {
	file_file_transfer_proto_rawDescOnce.Do(func() {
		file_file_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_transfer_proto_rawDescData)
	})
	return file_file_transfer_proto_rawDescData
}

var file_file_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_file_transfer_proto_goTypes = []interface{}{
	(*GetFileRequest)(nil),      // 0: file_transfer_grpc.GetFileRequest
	(*FileChunk)(nil),           // 1: file_transfer_grpc.FileChunk
	(*GetFileListRequest)(nil),  // 2: file_transfer_grpc.GetFileListRequest
	(*GetFileListResponse)(nil), // 3: file_transfer_grpc.GetFileListResponse
	(*GetFileInfoRequest)(nil),  // 4: file_transfer_grpc.GetFileInfoRequest
	(*GetFileInfoResponse)(nil), // 5: file_transfer_grpc.GetFileInfoResponse
}
var file_file_transfer_proto_depIdxs = []int32{
	0, // 0: file_transfer_grpc.FileTransferService.GetFile:input_type -> file_transfer_grpc.GetFileRequest
	2, // 1: file_transfer_grpc.FileTransferService.GetFileList:input_type -> file_transfer_grpc.GetFileListRequest
	4, // 2: file_transfer_grpc.FileTransferService.GetFileInfo:input_type -> file_transfer_grpc.GetFileInfoRequest
	1, // 3: file_transfer_grpc.FileTransferService.GetFile:output_type -> file_transfer_grpc.FileChunk
	3, // 4: file_transfer_grpc.FileTransferService.GetFileList:output_type -> file_transfer_grpc.GetFileListResponse
	5, // 5: file_transfer_grpc.FileTransferService.GetFileInfo:output_type -> file_transfer_grpc.GetFileInfoResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_transfer_proto_init() }
func file_file_transfer_proto_init() {
	if File_file_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
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
		file_file_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChunk); i {
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
		file_file_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileListRequest); i {
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
		file_file_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileListResponse); i {
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
		file_file_transfer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileInfoRequest); i {
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
		file_file_transfer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileInfoResponse); i {
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
			RawDescriptor: file_file_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_transfer_proto_goTypes,
		DependencyIndexes: file_file_transfer_proto_depIdxs,
		MessageInfos:      file_file_transfer_proto_msgTypes,
	}.Build()
	File_file_transfer_proto = out.File
	file_file_transfer_proto_rawDesc = nil
	file_file_transfer_proto_goTypes = nil
	file_file_transfer_proto_depIdxs = nil
}
