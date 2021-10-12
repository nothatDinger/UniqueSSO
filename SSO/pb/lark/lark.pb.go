// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: lark.proto

package lark

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

type LarkUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"name,omitempty" gorm:"column:name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name,omitempty" gorm:"column:name"`
	// @gotags: json:"en_name,omitempty" gorm:"column:en_name"
	ENName string `protobuf:"bytes,2,opt,name=ENName,proto3" json:"en_name,omitempty" gorm:"column:en_name"`
	// @gotags: json:"avatar_url,omitempty" gorm:"column:avatar_url"
	AvatarURL string `protobuf:"bytes,3,opt,name=AvatarURL,proto3" json:"avatar_url,omitempty" gorm:"column:avatar_url"`
	// @gotags: json:"avatar_thumb,omitempty" gorm:"column:avatar_thumb"
	AvatarThumb string `protobuf:"bytes,4,opt,name=AvatarThumb,proto3" json:"avatar_thumb,omitempty" gorm:"column:avatar_thumb"`
	// @gotags: json:"avatar_middle,omitempty" gorm:"column:avatar_middle"
	AvatarMiddle string `protobuf:"bytes,5,opt,name=AvatarMiddle,proto3" json:"avatar_middle,omitempty" gorm:"column:avatar_middle"`
	// @gotags: json:"avatar_big,omitempty" gorm:"column:avatar_big"
	AvatarBig string `protobuf:"bytes,6,opt,name=AvatarBig,proto3" json:"avatar_big,omitempty" gorm:"column:avatar_big"`
	// @gotags: json:"open_id,omitempty" gorm:"column:open_id;index"
	OpenID string `protobuf:"bytes,7,opt,name=OpenID,proto3" json:"open_id,omitempty" gorm:"column:open_id;index"`
	// @gotags: json:"union_id,omitempty" gorm:"column:union_id;primaryKey"
	UnionID string `protobuf:"bytes,8,opt,name=UnionID,proto3" json:"union_id,omitempty" gorm:"column:union_id;primaryKey"`
	// @gotags: json:"email" gorm:"column:email"
	Email string `protobuf:"bytes,9,opt,name=Email,proto3" json:"email" gorm:"column:email"`
	// @gotags: json:"user_id,omitempty" gorm:"column:user_id"
	UserID string `protobuf:"bytes,10,opt,name=UserID,proto3" json:"user_id,omitempty" gorm:"column:user_id"`
	// @gotags: json:"mobile,omitempty" gorm:"column:mobile"
	Mobile string `protobuf:"bytes,11,opt,name=mobile,proto3" json:"mobile,omitempty" gorm:"column:mobile"`
	// @gotags: json:"tenant_key,omitempty" gorm:"column:tenant_key"
	TenantKey string `protobuf:"bytes,12,opt,name=TenantKey,proto3" json:"tenant_key,omitempty" gorm:"column:tenant_key"`
	// @gotags: json:"gender,omitempty" gorm:"gender"
	Gender int32 `protobuf:"varint,13,opt,name=Gender,proto3" json:"gender,omitempty" gorm:"gender"`
	// @gotags: json:"department_ids,omitempty" gorm:"-"
	DepartmentIds []string `protobuf:"bytes,14,rep,name=DepartmentIds,proto3" json:"department_ids,omitempty" gorm:"-"`
	// @gotags: json:"job_title,omitempty" gorm:"column:job_title"
	JobTitle string `protobuf:"bytes,15,opt,name=JobTitle,proto3" json:"job_title,omitempty" gorm:"column:job_title"`
	// @gotags: json:"employee_type" gorm:"column:employee_type"
	EmployeeType int32 `protobuf:"varint,16,opt,name=EmployeeType,proto3" json:"employee_type" gorm:"column:employee_type"`
}

func (x *LarkUserInfo) Reset() {
	*x = LarkUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LarkUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LarkUserInfo) ProtoMessage() {}

func (x *LarkUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_lark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LarkUserInfo.ProtoReflect.Descriptor instead.
func (*LarkUserInfo) Descriptor() ([]byte, []int) {
	return file_lark_proto_rawDescGZIP(), []int{0}
}

func (x *LarkUserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LarkUserInfo) GetENName() string {
	if x != nil {
		return x.ENName
	}
	return ""
}

func (x *LarkUserInfo) GetAvatarURL() string {
	if x != nil {
		return x.AvatarURL
	}
	return ""
}

func (x *LarkUserInfo) GetAvatarThumb() string {
	if x != nil {
		return x.AvatarThumb
	}
	return ""
}

func (x *LarkUserInfo) GetAvatarMiddle() string {
	if x != nil {
		return x.AvatarMiddle
	}
	return ""
}

func (x *LarkUserInfo) GetAvatarBig() string {
	if x != nil {
		return x.AvatarBig
	}
	return ""
}

func (x *LarkUserInfo) GetOpenID() string {
	if x != nil {
		return x.OpenID
	}
	return ""
}

func (x *LarkUserInfo) GetUnionID() string {
	if x != nil {
		return x.UnionID
	}
	return ""
}

func (x *LarkUserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LarkUserInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *LarkUserInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LarkUserInfo) GetTenantKey() string {
	if x != nil {
		return x.TenantKey
	}
	return ""
}

func (x *LarkUserInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *LarkUserInfo) GetDepartmentIds() []string {
	if x != nil {
		return x.DepartmentIds
	}
	return nil
}

func (x *LarkUserInfo) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

func (x *LarkUserInfo) GetEmployeeType() int32 {
	if x != nil {
		return x.EmployeeType
	}
	return 0
}

type Department struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"department"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"department"`
	// @gotags: json:"parent_department_id"
	ParentDepartmentID string `protobuf:"bytes,2,opt,name=ParentDepartmentID,proto3" json:"parent_department_id"`
	// @gotags: json:"department_id"
	DepartmentID string `protobuf:"bytes,3,opt,name=DepartmentID,proto3" json:"department_id"`
	// @gotags: json:"open_department_id"
	OpenDepartmentID string `protobuf:"bytes,4,opt,name=OpenDepartmentID,proto3" json:"open_department_id"`
	// @gotags: json:"leader_user_id"
	LeaderUserID string `protobuf:"bytes,5,opt,name=LeaderUserID,proto3" json:"leader_user_id"`
}

func (x *Department) Reset() {
	*x = Department{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lark_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Department) ProtoMessage() {}

func (x *Department) ProtoReflect() protoreflect.Message {
	mi := &file_lark_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Department.ProtoReflect.Descriptor instead.
func (*Department) Descriptor() ([]byte, []int) {
	return file_lark_proto_rawDescGZIP(), []int{1}
}

func (x *Department) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Department) GetParentDepartmentID() string {
	if x != nil {
		return x.ParentDepartmentID
	}
	return ""
}

func (x *Department) GetDepartmentID() string {
	if x != nil {
		return x.DepartmentID
	}
	return ""
}

func (x *Department) GetOpenDepartmentID() string {
	if x != nil {
		return x.OpenDepartmentID
	}
	return ""
}

func (x *Department) GetLeaderUserID() string {
	if x != nil {
		return x.LeaderUserID
	}
	return ""
}

var File_lark_proto protoreflect.FileDescriptor

var file_lark_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x61,
	0x72, 0x6b, 0x22, 0xd0, 0x03, 0x0a, 0x0c, 0x4c, 0x61, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x4e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x4e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x20, 0x0a,
	0x0b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x12,
	0x22, 0x0a, 0x0c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x69, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x69,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x6e, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x55, 0x6e, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x24, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4a, 0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10,
	0x4f, 0x70, 0x65, 0x6e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x53, 0x53,
	0x4f, 0x2f, 0x70, 0x62, 0x2f, 0x6c, 0x61, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lark_proto_rawDescOnce sync.Once
	file_lark_proto_rawDescData = file_lark_proto_rawDesc
)

func file_lark_proto_rawDescGZIP() []byte {
	file_lark_proto_rawDescOnce.Do(func() {
		file_lark_proto_rawDescData = protoimpl.X.CompressGZIP(file_lark_proto_rawDescData)
	})
	return file_lark_proto_rawDescData
}

var file_lark_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lark_proto_goTypes = []interface{}{
	(*LarkUserInfo)(nil), // 0: lark.LarkUserInfo
	(*Department)(nil),   // 1: lark.Department
}
var file_lark_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lark_proto_init() }
func file_lark_proto_init() {
	if File_lark_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lark_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LarkUserInfo); i {
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
		file_lark_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Department); i {
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
			RawDescriptor: file_lark_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lark_proto_goTypes,
		DependencyIndexes: file_lark_proto_depIdxs,
		MessageInfos:      file_lark_proto_msgTypes,
	}.Build()
	File_lark_proto = out.File
	file_lark_proto_rawDesc = nil
	file_lark_proto_goTypes = nil
	file_lark_proto_depIdxs = nil
}
