// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.8.0
// source: plant_service.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
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

var File_plant_service_proto protoreflect.FileDescriptor

var file_plant_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe1, 0x01, 0x0a, 0x0c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x22,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x1a, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x28, 0x01, 0x12, 0x27, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x28, 0x01, 0x12, 0x29, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x28, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_plant_service_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil), // 0: google.protobuf.Empty
	(*PlantID)(nil),       // 1: pb.PlantID
	(*Plant)(nil),         // 2: pb.Plant
}
var file_plant_service_proto_depIdxs = []int32{
	0, // 0: pb.PlantService.GetPlants:input_type -> google.protobuf.Empty
	1, // 1: pb.PlantService.GetPlant:input_type -> pb.PlantID
	2, // 2: pb.PlantService.CreatePlant:input_type -> pb.Plant
	2, // 3: pb.PlantService.UpdatePlant:input_type -> pb.Plant
	1, // 4: pb.PlantService.DeletePlant:input_type -> pb.PlantID
	2, // 5: pb.PlantService.GetPlants:output_type -> pb.Plant
	2, // 6: pb.PlantService.GetPlant:output_type -> pb.Plant
	1, // 7: pb.PlantService.CreatePlant:output_type -> pb.PlantID
	1, // 8: pb.PlantService.UpdatePlant:output_type -> pb.PlantID
	1, // 9: pb.PlantService.DeletePlant:output_type -> pb.PlantID
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plant_service_proto_init() }
func file_plant_service_proto_init() {
	if File_plant_service_proto != nil {
		return
	}
	file_plant_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plant_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plant_service_proto_goTypes,
		DependencyIndexes: file_plant_service_proto_depIdxs,
	}.Build()
	File_plant_service_proto = out.File
	file_plant_service_proto_rawDesc = nil
	file_plant_service_proto_goTypes = nil
	file_plant_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlantServiceClient is the client API for PlantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlantServiceClient interface {
	GetPlants(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PlantService_GetPlantsClient, error)
	GetPlant(ctx context.Context, in *PlantID, opts ...grpc.CallOption) (*Plant, error)
	CreatePlant(ctx context.Context, opts ...grpc.CallOption) (PlantService_CreatePlantClient, error)
	UpdatePlant(ctx context.Context, opts ...grpc.CallOption) (PlantService_UpdatePlantClient, error)
	DeletePlant(ctx context.Context, opts ...grpc.CallOption) (PlantService_DeletePlantClient, error)
}

type plantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlantServiceClient(cc grpc.ClientConnInterface) PlantServiceClient {
	return &plantServiceClient{cc}
}

func (c *plantServiceClient) GetPlants(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PlantService_GetPlantsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PlantService_serviceDesc.Streams[0], "/pb.PlantService/GetPlants", opts...)
	if err != nil {
		return nil, err
	}
	x := &plantServiceGetPlantsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PlantService_GetPlantsClient interface {
	Recv() (*Plant, error)
	grpc.ClientStream
}

type plantServiceGetPlantsClient struct {
	grpc.ClientStream
}

func (x *plantServiceGetPlantsClient) Recv() (*Plant, error) {
	m := new(Plant)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *plantServiceClient) GetPlant(ctx context.Context, in *PlantID, opts ...grpc.CallOption) (*Plant, error) {
	out := new(Plant)
	err := c.cc.Invoke(ctx, "/pb.PlantService/GetPlant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plantServiceClient) CreatePlant(ctx context.Context, opts ...grpc.CallOption) (PlantService_CreatePlantClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PlantService_serviceDesc.Streams[1], "/pb.PlantService/CreatePlant", opts...)
	if err != nil {
		return nil, err
	}
	x := &plantServiceCreatePlantClient{stream}
	return x, nil
}

type PlantService_CreatePlantClient interface {
	Send(*Plant) error
	CloseAndRecv() (*PlantID, error)
	grpc.ClientStream
}

type plantServiceCreatePlantClient struct {
	grpc.ClientStream
}

func (x *plantServiceCreatePlantClient) Send(m *Plant) error {
	return x.ClientStream.SendMsg(m)
}

func (x *plantServiceCreatePlantClient) CloseAndRecv() (*PlantID, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PlantID)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *plantServiceClient) UpdatePlant(ctx context.Context, opts ...grpc.CallOption) (PlantService_UpdatePlantClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PlantService_serviceDesc.Streams[2], "/pb.PlantService/UpdatePlant", opts...)
	if err != nil {
		return nil, err
	}
	x := &plantServiceUpdatePlantClient{stream}
	return x, nil
}

type PlantService_UpdatePlantClient interface {
	Send(*Plant) error
	CloseAndRecv() (*PlantID, error)
	grpc.ClientStream
}

type plantServiceUpdatePlantClient struct {
	grpc.ClientStream
}

func (x *plantServiceUpdatePlantClient) Send(m *Plant) error {
	return x.ClientStream.SendMsg(m)
}

func (x *plantServiceUpdatePlantClient) CloseAndRecv() (*PlantID, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PlantID)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *plantServiceClient) DeletePlant(ctx context.Context, opts ...grpc.CallOption) (PlantService_DeletePlantClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PlantService_serviceDesc.Streams[3], "/pb.PlantService/DeletePlant", opts...)
	if err != nil {
		return nil, err
	}
	x := &plantServiceDeletePlantClient{stream}
	return x, nil
}

type PlantService_DeletePlantClient interface {
	Send(*PlantID) error
	CloseAndRecv() (*PlantID, error)
	grpc.ClientStream
}

type plantServiceDeletePlantClient struct {
	grpc.ClientStream
}

func (x *plantServiceDeletePlantClient) Send(m *PlantID) error {
	return x.ClientStream.SendMsg(m)
}

func (x *plantServiceDeletePlantClient) CloseAndRecv() (*PlantID, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PlantID)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PlantServiceServer is the server API for PlantService service.
type PlantServiceServer interface {
	GetPlants(*emptypb.Empty, PlantService_GetPlantsServer) error
	GetPlant(context.Context, *PlantID) (*Plant, error)
	CreatePlant(PlantService_CreatePlantServer) error
	UpdatePlant(PlantService_UpdatePlantServer) error
	DeletePlant(PlantService_DeletePlantServer) error
}

// UnimplementedPlantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPlantServiceServer struct {
}

func (*UnimplementedPlantServiceServer) GetPlants(*emptypb.Empty, PlantService_GetPlantsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPlants not implemented")
}
func (*UnimplementedPlantServiceServer) GetPlant(context.Context, *PlantID) (*Plant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlant not implemented")
}
func (*UnimplementedPlantServiceServer) CreatePlant(PlantService_CreatePlantServer) error {
	return status.Errorf(codes.Unimplemented, "method CreatePlant not implemented")
}
func (*UnimplementedPlantServiceServer) UpdatePlant(PlantService_UpdatePlantServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdatePlant not implemented")
}
func (*UnimplementedPlantServiceServer) DeletePlant(PlantService_DeletePlantServer) error {
	return status.Errorf(codes.Unimplemented, "method DeletePlant not implemented")
}

func RegisterPlantServiceServer(s *grpc.Server, srv PlantServiceServer) {
	s.RegisterService(&_PlantService_serviceDesc, srv)
}

func _PlantService_GetPlants_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlantServiceServer).GetPlants(m, &plantServiceGetPlantsServer{stream})
}

type PlantService_GetPlantsServer interface {
	Send(*Plant) error
	grpc.ServerStream
}

type plantServiceGetPlantsServer struct {
	grpc.ServerStream
}

func (x *plantServiceGetPlantsServer) Send(m *Plant) error {
	return x.ServerStream.SendMsg(m)
}

func _PlantService_GetPlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlantID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlantServiceServer).GetPlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PlantService/GetPlant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlantServiceServer).GetPlant(ctx, req.(*PlantID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlantService_CreatePlant_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PlantServiceServer).CreatePlant(&plantServiceCreatePlantServer{stream})
}

type PlantService_CreatePlantServer interface {
	SendAndClose(*PlantID) error
	Recv() (*Plant, error)
	grpc.ServerStream
}

type plantServiceCreatePlantServer struct {
	grpc.ServerStream
}

func (x *plantServiceCreatePlantServer) SendAndClose(m *PlantID) error {
	return x.ServerStream.SendMsg(m)
}

func (x *plantServiceCreatePlantServer) Recv() (*Plant, error) {
	m := new(Plant)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PlantService_UpdatePlant_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PlantServiceServer).UpdatePlant(&plantServiceUpdatePlantServer{stream})
}

type PlantService_UpdatePlantServer interface {
	SendAndClose(*PlantID) error
	Recv() (*Plant, error)
	grpc.ServerStream
}

type plantServiceUpdatePlantServer struct {
	grpc.ServerStream
}

func (x *plantServiceUpdatePlantServer) SendAndClose(m *PlantID) error {
	return x.ServerStream.SendMsg(m)
}

func (x *plantServiceUpdatePlantServer) Recv() (*Plant, error) {
	m := new(Plant)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PlantService_DeletePlant_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PlantServiceServer).DeletePlant(&plantServiceDeletePlantServer{stream})
}

type PlantService_DeletePlantServer interface {
	SendAndClose(*PlantID) error
	Recv() (*PlantID, error)
	grpc.ServerStream
}

type plantServiceDeletePlantServer struct {
	grpc.ServerStream
}

func (x *plantServiceDeletePlantServer) SendAndClose(m *PlantID) error {
	return x.ServerStream.SendMsg(m)
}

func (x *plantServiceDeletePlantServer) Recv() (*PlantID, error) {
	m := new(PlantID)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PlantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PlantService",
	HandlerType: (*PlantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlant",
			Handler:    _PlantService_GetPlant_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPlants",
			Handler:       _PlantService_GetPlants_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreatePlant",
			Handler:       _PlantService_CreatePlant_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdatePlant",
			Handler:       _PlantService_UpdatePlant_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DeletePlant",
			Handler:       _PlantService_DeletePlant_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "plant_service.proto",
}
