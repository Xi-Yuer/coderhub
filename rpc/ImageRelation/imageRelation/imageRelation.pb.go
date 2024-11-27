// goctl rpc protoc imageRelation.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: imageRelation.proto

package imageRelation

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

// 图片关系实体
type ImageRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                  // 关系ID
	ImageId    int64  `protobuf:"varint,2,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`         // 图片ID
	EntityId   int64  `protobuf:"varint,3,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`      // 关联实体ID
	EntityType string `protobuf:"bytes,4,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"` // 关联实体类型(comment/article等)
	Sort       int32  `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"`                              // 排序号
	CreatedAt  string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`    // 创建时间
}

func (x *ImageRelation) Reset() {
	*x = ImageRelation{}
	mi := &file_imageRelation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageRelation) ProtoMessage() {}

func (x *ImageRelation) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageRelation.ProtoReflect.Descriptor instead.
func (*ImageRelation) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{0}
}

func (x *ImageRelation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ImageRelation) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *ImageRelation) GetEntityId() int64 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *ImageRelation) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *ImageRelation) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *ImageRelation) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// 创建关系请求
type CreateRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId    int64  `protobuf:"varint,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`         // 图片ID
	EntityId   int64  `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`      // 关联实体ID
	EntityType string `protobuf:"bytes,3,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"` // 关联实体类型
	Sort       int32  `protobuf:"varint,4,opt,name=sort,proto3" json:"sort,omitempty"`                              // 排序号
}

func (x *CreateRelationRequest) Reset() {
	*x = CreateRelationRequest{}
	mi := &file_imageRelation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRelationRequest) ProtoMessage() {}

func (x *CreateRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRelationRequest.ProtoReflect.Descriptor instead.
func (*CreateRelationRequest) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRelationRequest) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *CreateRelationRequest) GetEntityId() int64 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *CreateRelationRequest) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *CreateRelationRequest) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type CreateRelationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relation *ImageRelation `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"` // 创建的关系
}

func (x *CreateRelationResponse) Reset() {
	*x = CreateRelationResponse{}
	mi := &file_imageRelation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRelationResponse) ProtoMessage() {}

func (x *CreateRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRelationResponse.ProtoReflect.Descriptor instead.
func (*CreateRelationResponse) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRelationResponse) GetRelation() *ImageRelation {
	if x != nil {
		return x.Relation
	}
	return nil
}

// 批量创建关系请求
type BatchCreateRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relations []*CreateRelationRequest `protobuf:"bytes,1,rep,name=relations,proto3" json:"relations,omitempty"` // 关系列表
}

func (x *BatchCreateRelationRequest) Reset() {
	*x = BatchCreateRelationRequest{}
	mi := &file_imageRelation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchCreateRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateRelationRequest) ProtoMessage() {}

func (x *BatchCreateRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateRelationRequest.ProtoReflect.Descriptor instead.
func (*BatchCreateRelationRequest) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{3}
}

func (x *BatchCreateRelationRequest) GetRelations() []*CreateRelationRequest {
	if x != nil {
		return x.Relations
	}
	return nil
}

type BatchCreateRelationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relations []*ImageRelation `protobuf:"bytes,1,rep,name=relations,proto3" json:"relations,omitempty"` // 创建的关系列表
}

func (x *BatchCreateRelationResponse) Reset() {
	*x = BatchCreateRelationResponse{}
	mi := &file_imageRelation_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchCreateRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateRelationResponse) ProtoMessage() {}

func (x *BatchCreateRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateRelationResponse.ProtoReflect.Descriptor instead.
func (*BatchCreateRelationResponse) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{4}
}

func (x *BatchCreateRelationResponse) GetRelations() []*ImageRelation {
	if x != nil {
		return x.Relations
	}
	return nil
}

// 获取实体关联的图片列表请求
type GetImagesByEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId   int64  `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`      // 关联实体ID
	EntityType string `protobuf:"bytes,2,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"` // 关联实体类型
}

func (x *GetImagesByEntityRequest) Reset() {
	*x = GetImagesByEntityRequest{}
	mi := &file_imageRelation_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetImagesByEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesByEntityRequest) ProtoMessage() {}

func (x *GetImagesByEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesByEntityRequest.ProtoReflect.Descriptor instead.
func (*GetImagesByEntityRequest) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{5}
}

func (x *GetImagesByEntityRequest) GetEntityId() int64 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *GetImagesByEntityRequest) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

type GetImagesByEntityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*ImageInfo `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"` // 图片列表
}

func (x *GetImagesByEntityResponse) Reset() {
	*x = GetImagesByEntityResponse{}
	mi := &file_imageRelation_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetImagesByEntityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesByEntityResponse) ProtoMessage() {}

func (x *GetImagesByEntityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesByEntityResponse.ProtoReflect.Descriptor instead.
func (*GetImagesByEntityResponse) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{6}
}

func (x *GetImagesByEntityResponse) GetImages() []*ImageInfo {
	if x != nil {
		return x.Images
	}
	return nil
}

// 获取图片关联的实体列表请求
type GetEntitiesByImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId    int64  `protobuf:"varint,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`         // 图片ID
	EntityType string `protobuf:"bytes,2,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"` // 关联实体类型（可选）
}

func (x *GetEntitiesByImageRequest) Reset() {
	*x = GetEntitiesByImageRequest{}
	mi := &file_imageRelation_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEntitiesByImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntitiesByImageRequest) ProtoMessage() {}

func (x *GetEntitiesByImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntitiesByImageRequest.ProtoReflect.Descriptor instead.
func (*GetEntitiesByImageRequest) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{7}
}

func (x *GetEntitiesByImageRequest) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *GetEntitiesByImageRequest) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

type GetEntitiesByImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities []*EntityInfo `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"` // 实体列表
}

func (x *GetEntitiesByImageResponse) Reset() {
	*x = GetEntitiesByImageResponse{}
	mi := &file_imageRelation_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEntitiesByImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntitiesByImageResponse) ProtoMessage() {}

func (x *GetEntitiesByImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntitiesByImageResponse.ProtoReflect.Descriptor instead.
func (*GetEntitiesByImageResponse) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{8}
}

func (x *GetEntitiesByImageResponse) GetEntities() []*EntityInfo {
	if x != nil {
		return x.Entities
	}
	return nil
}

// 图片信息
type ImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId      int64  `protobuf:"varint,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Url          string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	ThumbnailUrl string `protobuf:"bytes,3,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	Width        int32  `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	Height       int32  `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Sort         int32  `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"` // 排序号
}

func (x *ImageInfo) Reset() {
	*x = ImageInfo{}
	mi := &file_imageRelation_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfo) ProtoMessage() {}

func (x *ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{9}
}

func (x *ImageInfo) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *ImageInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ImageInfo) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *ImageInfo) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageInfo) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ImageInfo) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

// 实体信息
type EntityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityType string `protobuf:"bytes,1,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"` // 实体类型
	EntityId   int64  `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`      // 实体ID
	CreatedAt  int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`   // 关联创建时间
}

func (x *EntityInfo) Reset() {
	*x = EntityInfo{}
	mi := &file_imageRelation_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityInfo) ProtoMessage() {}

func (x *EntityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_imageRelation_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityInfo.ProtoReflect.Descriptor instead.
func (*EntityInfo) Descriptor() ([]byte, []int) {
	return file_imageRelation_proto_rawDescGZIP(), []int{10}
}

func (x *EntityInfo) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *EntityInfo) GetEntityId() int64 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EntityInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_imageRelation_proto protoreflect.FileDescriptor

var file_imageRelation_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x52, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x60, 0x0a,
	0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x09, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x59, 0x0a, 0x1b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x58, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x4d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x42, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x53, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x22, 0x9f, 0x01, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x22, 0x69, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xbe,
	0x03, 0x0a, 0x14, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x13, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x29, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x27, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x42, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x28, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imageRelation_proto_rawDescOnce sync.Once
	file_imageRelation_proto_rawDescData = file_imageRelation_proto_rawDesc
)

func file_imageRelation_proto_rawDescGZIP() []byte {
	file_imageRelation_proto_rawDescOnce.Do(func() {
		file_imageRelation_proto_rawDescData = protoimpl.X.CompressGZIP(file_imageRelation_proto_rawDescData)
	})
	return file_imageRelation_proto_rawDescData
}

var file_imageRelation_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_imageRelation_proto_goTypes = []any{
	(*ImageRelation)(nil),               // 0: imageRelation.ImageRelation
	(*CreateRelationRequest)(nil),       // 1: imageRelation.CreateRelationRequest
	(*CreateRelationResponse)(nil),      // 2: imageRelation.CreateRelationResponse
	(*BatchCreateRelationRequest)(nil),  // 3: imageRelation.BatchCreateRelationRequest
	(*BatchCreateRelationResponse)(nil), // 4: imageRelation.BatchCreateRelationResponse
	(*GetImagesByEntityRequest)(nil),    // 5: imageRelation.GetImagesByEntityRequest
	(*GetImagesByEntityResponse)(nil),   // 6: imageRelation.GetImagesByEntityResponse
	(*GetEntitiesByImageRequest)(nil),   // 7: imageRelation.GetEntitiesByImageRequest
	(*GetEntitiesByImageResponse)(nil),  // 8: imageRelation.GetEntitiesByImageResponse
	(*ImageInfo)(nil),                   // 9: imageRelation.ImageInfo
	(*EntityInfo)(nil),                  // 10: imageRelation.EntityInfo
}
var file_imageRelation_proto_depIdxs = []int32{
	0,  // 0: imageRelation.CreateRelationResponse.relation:type_name -> imageRelation.ImageRelation
	1,  // 1: imageRelation.BatchCreateRelationRequest.relations:type_name -> imageRelation.CreateRelationRequest
	0,  // 2: imageRelation.BatchCreateRelationResponse.relations:type_name -> imageRelation.ImageRelation
	9,  // 3: imageRelation.GetImagesByEntityResponse.images:type_name -> imageRelation.ImageInfo
	10, // 4: imageRelation.GetEntitiesByImageResponse.entities:type_name -> imageRelation.EntityInfo
	1,  // 5: imageRelation.ImageRelationService.CreateRelation:input_type -> imageRelation.CreateRelationRequest
	3,  // 6: imageRelation.ImageRelationService.BatchCreateRelation:input_type -> imageRelation.BatchCreateRelationRequest
	5,  // 7: imageRelation.ImageRelationService.GetImagesByEntity:input_type -> imageRelation.GetImagesByEntityRequest
	7,  // 8: imageRelation.ImageRelationService.GetEntitiesByImage:input_type -> imageRelation.GetEntitiesByImageRequest
	2,  // 9: imageRelation.ImageRelationService.CreateRelation:output_type -> imageRelation.CreateRelationResponse
	4,  // 10: imageRelation.ImageRelationService.BatchCreateRelation:output_type -> imageRelation.BatchCreateRelationResponse
	6,  // 11: imageRelation.ImageRelationService.GetImagesByEntity:output_type -> imageRelation.GetImagesByEntityResponse
	8,  // 12: imageRelation.ImageRelationService.GetEntitiesByImage:output_type -> imageRelation.GetEntitiesByImageResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_imageRelation_proto_init() }
func file_imageRelation_proto_init() {
	if File_imageRelation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_imageRelation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imageRelation_proto_goTypes,
		DependencyIndexes: file_imageRelation_proto_depIdxs,
		MessageInfos:      file_imageRelation_proto_msgTypes,
	}.Build()
	File_imageRelation_proto = out.File
	file_imageRelation_proto_rawDesc = nil
	file_imageRelation_proto_goTypes = nil
	file_imageRelation_proto_depIdxs = nil
}
