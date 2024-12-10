// goctl rpc protoc academic_navigator.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.3
// source: academic_navigator.proto

package academic_navigator

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

// 新增学术导航
type AddAcademicNavigatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`   // 用户 ID
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                // 内容
	Education string `protobuf:"bytes,3,opt,name=education,proto3" json:"education,omitempty"`            // 学历
	Major     string `protobuf:"bytes,4,opt,name=major,proto3" json:"major,omitempty"`                    // 专业
	School    string `protobuf:"bytes,5,opt,name=school,proto3" json:"school,omitempty"`                  // 学校
	WorkExp   string `protobuf:"bytes,6,opt,name=work_exp,json=workExp,proto3" json:"work_exp,omitempty"` // 工作经验
}

func (x *AddAcademicNavigatorRequest) Reset() {
	*x = AddAcademicNavigatorRequest{}
	mi := &file_academic_navigator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddAcademicNavigatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAcademicNavigatorRequest) ProtoMessage() {}

func (x *AddAcademicNavigatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_academic_navigator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAcademicNavigatorRequest.ProtoReflect.Descriptor instead.
func (*AddAcademicNavigatorRequest) Descriptor() ([]byte, []int) {
	return file_academic_navigator_proto_rawDescGZIP(), []int{0}
}

func (x *AddAcademicNavigatorRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddAcademicNavigatorRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AddAcademicNavigatorRequest) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *AddAcademicNavigatorRequest) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *AddAcademicNavigatorRequest) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *AddAcademicNavigatorRequest) GetWorkExp() string {
	if x != nil {
		return x.WorkExp
	}
	return ""
}

// 获取学术导航
type GetAcademicNavigatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`       // 用户 ID
	Education string `protobuf:"bytes,2,opt,name=education,proto3" json:"education,omitempty"`                // 学历
	Major     string `protobuf:"bytes,3,opt,name=major,proto3" json:"major,omitempty"`                        // 专业
	School    string `protobuf:"bytes,4,opt,name=school,proto3" json:"school,omitempty"`                      // 学校
	WorkExp   string `protobuf:"bytes,5,opt,name=work_exp,json=workExp,proto3" json:"work_exp,omitempty"`     // 工作经验
	Page      int64  `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`                         // 页码
	PageSize  int64  `protobuf:"varint,7,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"` // 每页大小
}

func (x *GetAcademicNavigatorRequest) Reset() {
	*x = GetAcademicNavigatorRequest{}
	mi := &file_academic_navigator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAcademicNavigatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAcademicNavigatorRequest) ProtoMessage() {}

func (x *GetAcademicNavigatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_academic_navigator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAcademicNavigatorRequest.ProtoReflect.Descriptor instead.
func (*GetAcademicNavigatorRequest) Descriptor() ([]byte, []int) {
	return file_academic_navigator_proto_rawDescGZIP(), []int{1}
}

func (x *GetAcademicNavigatorRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetAcademicNavigatorRequest) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *GetAcademicNavigatorRequest) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *GetAcademicNavigatorRequest) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *GetAcademicNavigatorRequest) GetWorkExp() string {
	if x != nil {
		return x.WorkExp
	}
	return ""
}

func (x *GetAcademicNavigatorRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAcademicNavigatorRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 删除学术导航
type DeleteAcademicNavigatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                       // 学术导航 ID
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户 ID
}

func (x *DeleteAcademicNavigatorRequest) Reset() {
	*x = DeleteAcademicNavigatorRequest{}
	mi := &file_academic_navigator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAcademicNavigatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAcademicNavigatorRequest) ProtoMessage() {}

func (x *DeleteAcademicNavigatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_academic_navigator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAcademicNavigatorRequest.ProtoReflect.Descriptor instead.
func (*DeleteAcademicNavigatorRequest) Descriptor() ([]byte, []int) {
	return file_academic_navigator_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteAcademicNavigatorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteAcademicNavigatorRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 是否成功
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_academic_navigator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_academic_navigator_proto_msgTypes[3]
	if x != nil {
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
	return file_academic_navigator_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_academic_navigator_proto protoreflect.FileDescriptor

var file_academic_navigator_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x63, 0x61, 0x64,
	0x65, 0x6d, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x22, 0xb7,
	0x01, 0x0a, 0x1b, 0x41, 0x64, 0x64, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x4e, 0x61,
	0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x19, 0x0a,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x78, 0x70, 0x22, 0xce, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x19, 0x0a,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x78, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x49, 0x0a, 0x1e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x4e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xdb, 0x02, 0x0a, 0x18, 0x41,
	0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x41, 0x63,
	0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x2f, 0x2e, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63,
	0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x76, 0x69,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x67, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x4e,
	0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2f, 0x2e, 0x61, 0x63, 0x61, 0x64, 0x65,
	0x6d, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x63, 0x61, 0x64,
	0x65, 0x6d, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x17, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x4e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x32, 0x2e, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x5f,
	0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x63, 0x61, 0x64, 0x65,
	0x6d, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x61, 0x63,
	0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_academic_navigator_proto_rawDescOnce sync.Once
	file_academic_navigator_proto_rawDescData = file_academic_navigator_proto_rawDesc
)

func file_academic_navigator_proto_rawDescGZIP() []byte {
	file_academic_navigator_proto_rawDescOnce.Do(func() {
		file_academic_navigator_proto_rawDescData = protoimpl.X.CompressGZIP(file_academic_navigator_proto_rawDescData)
	})
	return file_academic_navigator_proto_rawDescData
}

var file_academic_navigator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_academic_navigator_proto_goTypes = []any{
	(*AddAcademicNavigatorRequest)(nil),    // 0: academic_navigator.AddAcademicNavigatorRequest
	(*GetAcademicNavigatorRequest)(nil),    // 1: academic_navigator.GetAcademicNavigatorRequest
	(*DeleteAcademicNavigatorRequest)(nil), // 2: academic_navigator.DeleteAcademicNavigatorRequest
	(*Response)(nil),                       // 3: academic_navigator.Response
}
var file_academic_navigator_proto_depIdxs = []int32{
	0, // 0: academic_navigator.AcademicNavigatorService.AddAcademicNavigator:input_type -> academic_navigator.AddAcademicNavigatorRequest
	1, // 1: academic_navigator.AcademicNavigatorService.GetAcademicNavigator:input_type -> academic_navigator.GetAcademicNavigatorRequest
	2, // 2: academic_navigator.AcademicNavigatorService.DeleteAcademicNavigator:input_type -> academic_navigator.DeleteAcademicNavigatorRequest
	3, // 3: academic_navigator.AcademicNavigatorService.AddAcademicNavigator:output_type -> academic_navigator.Response
	3, // 4: academic_navigator.AcademicNavigatorService.GetAcademicNavigator:output_type -> academic_navigator.Response
	3, // 5: academic_navigator.AcademicNavigatorService.DeleteAcademicNavigator:output_type -> academic_navigator.Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_academic_navigator_proto_init() }
func file_academic_navigator_proto_init() {
	if File_academic_navigator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_academic_navigator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_academic_navigator_proto_goTypes,
		DependencyIndexes: file_academic_navigator_proto_depIdxs,
		MessageInfos:      file_academic_navigator_proto_msgTypes,
	}.Build()
	File_academic_navigator_proto = out.File
	file_academic_navigator_proto_rawDesc = nil
	file_academic_navigator_proto_goTypes = nil
	file_academic_navigator_proto_depIdxs = nil
}
