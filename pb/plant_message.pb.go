// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.8.0
// source: plant_message.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Plant_Category int32

const (
	Plant_TREE   Plant_Category = 0
	Plant_FLOWER Plant_Category = 1
	Plant_FRUIT  Plant_Category = 2
)

// Enum value maps for Plant_Category.
var (
	Plant_Category_name = map[int32]string{
		0: "TREE",
		1: "FLOWER",
		2: "FRUIT",
	}
	Plant_Category_value = map[string]int32{
		"TREE":   0,
		"FLOWER": 1,
		"FRUIT":  2,
	}
)

func (x Plant_Category) Enum() *Plant_Category {
	p := new(Plant_Category)
	*p = x
	return p
}

func (x Plant_Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Plant_Category) Descriptor() protoreflect.EnumDescriptor {
	return file_plant_message_proto_enumTypes[0].Descriptor()
}

func (Plant_Category) Type() protoreflect.EnumType {
	return &file_plant_message_proto_enumTypes[0]
}

func (x Plant_Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Plant_Category.Descriptor instead.
func (Plant_Category) EnumDescriptor() ([]byte, []int) {
	return file_plant_message_proto_rawDescGZIP(), []int{0, 0}
}

type Plant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category    Plant_Category         `protobuf:"varint,3,opt,name=category,proto3,enum=pb.Plant_Category" json:"category,omitempty"`
	Price       float32                `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Description string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	User        *User                  `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Plant) Reset() {
	*x = Plant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plant_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plant) ProtoMessage() {}

func (x *Plant) ProtoReflect() protoreflect.Message {
	mi := &file_plant_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plant.ProtoReflect.Descriptor instead.
func (*Plant) Descriptor() ([]byte, []int) {
	return file_plant_message_proto_rawDescGZIP(), []int{0}
}

func (x *Plant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plant) GetCategory() Plant_Category {
	if x != nil {
		return x.Category
	}
	return Plant_TREE
}

func (x *Plant) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Plant) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Plant) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Plant) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type PlantID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PlantID) Reset() {
	*x = PlantID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plant_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlantID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlantID) ProtoMessage() {}

func (x *PlantID) ProtoReflect() protoreflect.Message {
	mi := &file_plant_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlantID.ProtoReflect.Descriptor instead.
func (*PlantID) Descriptor() ([]byte, []int) {
	return file_plant_message_proto_rawDescGZIP(), []int{1}
}

func (x *PlantID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_plant_message_proto protoreflect.FileDescriptor

var file_plant_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d,
	0x02, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x2b, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x08, 0x0a,
	0x04, 0x54, 0x52, 0x45, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x4c, 0x4f, 0x57, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x52, 0x55, 0x49, 0x54, 0x10, 0x02, 0x22, 0x1f,
	0x0a, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plant_message_proto_rawDescOnce sync.Once
	file_plant_message_proto_rawDescData = file_plant_message_proto_rawDesc
)

func file_plant_message_proto_rawDescGZIP() []byte {
	file_plant_message_proto_rawDescOnce.Do(func() {
		file_plant_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_plant_message_proto_rawDescData)
	})
	return file_plant_message_proto_rawDescData
}

var file_plant_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_plant_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_plant_message_proto_goTypes = []interface{}{
	(Plant_Category)(0),           // 0: pb.Plant.Category
	(*Plant)(nil),                 // 1: pb.Plant
	(*PlantID)(nil),               // 2: pb.PlantID
	(*User)(nil),                  // 3: pb.User
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_plant_message_proto_depIdxs = []int32{
	0, // 0: pb.Plant.category:type_name -> pb.Plant.Category
	3, // 1: pb.Plant.user:type_name -> pb.User
	4, // 2: pb.Plant.last_updated:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_plant_message_proto_init() }
func file_plant_message_proto_init() {
	if File_plant_message_proto != nil {
		return
	}
	file_user_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_plant_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plant); i {
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
		file_plant_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlantID); i {
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
			RawDescriptor: file_plant_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plant_message_proto_goTypes,
		DependencyIndexes: file_plant_message_proto_depIdxs,
		EnumInfos:         file_plant_message_proto_enumTypes,
		MessageInfos:      file_plant_message_proto_msgTypes,
	}.Build()
	File_plant_message_proto = out.File
	file_plant_message_proto_rawDesc = nil
	file_plant_message_proto_goTypes = nil
	file_plant_message_proto_depIdxs = nil
}