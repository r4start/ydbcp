// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.27.1
// source: backup.proto

package ydbcp

import (
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

type Backup_Status int32

const (
	Backup_STATUS_UNSPECIFIED Backup_Status = 0
	Backup_PENDING            Backup_Status = 1
	Backup_AVAILABLE          Backup_Status = 2
	Backup_ERROR              Backup_Status = 3
	Backup_CANCELED           Backup_Status = 4
	Backup_DELETED            Backup_Status = 5
)

// Enum value maps for Backup_Status.
var (
	Backup_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "PENDING",
		2: "AVAILABLE",
		3: "ERROR",
		4: "CANCELED",
		5: "DELETED",
	}
	Backup_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"PENDING":            1,
		"AVAILABLE":          2,
		"ERROR":              3,
		"CANCELED":           4,
		"DELETED":            5,
	}
)

func (x Backup_Status) Enum() *Backup_Status {
	p := new(Backup_Status)
	*p = x
	return p
}

func (x Backup_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Backup_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_backup_proto_enumTypes[0].Descriptor()
}

func (Backup_Status) Type() protoreflect.EnumType {
	return &file_backup_proto_enumTypes[0]
}

func (x Backup_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Backup_Status.Descriptor instead.
func (Backup_Status) EnumDescriptor() ([]byte, []int) {
	return file_backup_proto_rawDescGZIP(), []int{0, 0}
}

type Backup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContainerId  string                 `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	DatabaseName string                 `protobuf:"bytes,3,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	Location     *S3Location            `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Audit        *AuditInfo             `protobuf:"bytes,5,opt,name=audit,proto3" json:"audit,omitempty"`
	Size         int64                  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Status       Backup_Status          `protobuf:"varint,7,opt,name=status,proto3,enum=ydbcp.Backup_Status" json:"status,omitempty"`
	Message      string                 `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
	ExpireAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
}

func (x *Backup) Reset() {
	*x = Backup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Backup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Backup) ProtoMessage() {}

func (x *Backup) ProtoReflect() protoreflect.Message {
	mi := &file_backup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Backup.ProtoReflect.Descriptor instead.
func (*Backup) Descriptor() ([]byte, []int) {
	return file_backup_proto_rawDescGZIP(), []int{0}
}

func (x *Backup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Backup) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *Backup) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *Backup) GetLocation() *S3Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Backup) GetAudit() *AuditInfo {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *Backup) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Backup) GetStatus() Backup_Status {
	if x != nil {
		return x.Status
	}
	return Backup_STATUS_UNSPECIFIED
}

func (x *Backup) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Backup) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

type S3Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint   string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Bucket     string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Region     string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	PathPrefix string `protobuf:"bytes,4,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
}

func (x *S3Location) Reset() {
	*x = S3Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Location) ProtoMessage() {}

func (x *S3Location) ProtoReflect() protoreflect.Message {
	mi := &file_backup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Location.ProtoReflect.Descriptor instead.
func (*S3Location) Descriptor() ([]byte, []int) {
	return file_backup_proto_rawDescGZIP(), []int{1}
}

func (x *S3Location) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *S3Location) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *S3Location) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *S3Location) GetPathPrefix() string {
	if x != nil {
		return x.PathPrefix
	}
	return ""
}

type AuditInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Creator     string                 `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CompletedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
}

func (x *AuditInfo) Reset() {
	*x = AuditInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditInfo) ProtoMessage() {}

func (x *AuditInfo) ProtoReflect() protoreflect.Message {
	mi := &file_backup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditInfo.ProtoReflect.Descriptor instead.
func (*AuditInfo) Descriptor() ([]byte, []int) {
	return file_backup_proto_rawDescGZIP(), []int{2}
}

func (x *AuditInfo) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *AuditInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AuditInfo) GetCompletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

var File_backup_proto protoreflect.FileDescriptor

var file_backup_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x79, 0x64, 0x62, 0x63, 0x70, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x03, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x79, 0x64,
	0x62, 0x63, 0x70, 0x2e, 0x53, 0x33, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x79, 0x64, 0x62, 0x63, 0x70, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x79, 0x64, 0x62, 0x63, 0x70, 0x2e, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x41, 0x74, 0x22, 0x62, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x05, 0x22, 0x79, 0x0a, 0x0a, 0x53, 0x33, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x22, 0x9f, 0x01, 0x0a, 0x09, 0x41, 0x75, 0x64, 0x69, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x79, 0x64, 0x62, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x79, 0x64, 0x62, 0x63, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backup_proto_rawDescOnce sync.Once
	file_backup_proto_rawDescData = file_backup_proto_rawDesc
)

func file_backup_proto_rawDescGZIP() []byte {
	file_backup_proto_rawDescOnce.Do(func() {
		file_backup_proto_rawDescData = protoimpl.X.CompressGZIP(file_backup_proto_rawDescData)
	})
	return file_backup_proto_rawDescData
}

var file_backup_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_backup_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_backup_proto_goTypes = []interface{}{
	(Backup_Status)(0),            // 0: ydbcp.Backup.Status
	(*Backup)(nil),                // 1: ydbcp.Backup
	(*S3Location)(nil),            // 2: ydbcp.S3Location
	(*AuditInfo)(nil),             // 3: ydbcp.AuditInfo
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_backup_proto_depIdxs = []int32{
	2, // 0: ydbcp.Backup.location:type_name -> ydbcp.S3Location
	3, // 1: ydbcp.Backup.audit:type_name -> ydbcp.AuditInfo
	0, // 2: ydbcp.Backup.status:type_name -> ydbcp.Backup.Status
	4, // 3: ydbcp.Backup.expire_at:type_name -> google.protobuf.Timestamp
	4, // 4: ydbcp.AuditInfo.created_at:type_name -> google.protobuf.Timestamp
	4, // 5: ydbcp.AuditInfo.completed_at:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_backup_proto_init() }
func file_backup_proto_init() {
	if File_backup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Backup); i {
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
		file_backup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Location); i {
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
		file_backup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditInfo); i {
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
			RawDescriptor: file_backup_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backup_proto_goTypes,
		DependencyIndexes: file_backup_proto_depIdxs,
		EnumInfos:         file_backup_proto_enumTypes,
		MessageInfos:      file_backup_proto_msgTypes,
	}.Build()
	File_backup_proto = out.File
	file_backup_proto_rawDesc = nil
	file_backup_proto_goTypes = nil
	file_backup_proto_depIdxs = nil
}
