// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.0
// source: protos/asset.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Aggregate1 string `protobuf:"bytes,2,opt,name=aggregate1,proto3" json:"aggregate1,omitempty"`
	Aggregate2 string `protobuf:"bytes,3,opt,name=aggregate2,proto3" json:"aggregate2,omitempty"`
}

func (x *AssetRequest) Reset() {
	*x = AssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetRequest) ProtoMessage() {}

func (x *AssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetRequest.ProtoReflect.Descriptor instead.
func (*AssetRequest) Descriptor() ([]byte, []int) {
	return file_protos_asset_proto_rawDescGZIP(), []int{0}
}

func (x *AssetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AssetRequest) GetAggregate1() string {
	if x != nil {
		return x.Aggregate1
	}
	return ""
}

func (x *AssetRequest) GetAggregate2() string {
	if x != nil {
		return x.Aggregate2
	}
	return ""
}

type AssetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfAssets int32 `protobuf:"varint,1,opt,name=numberOfAssets,proto3" json:"numberOfAssets,omitempty"`
}

func (x *AssetResponse) Reset() {
	*x = AssetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_asset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetResponse) ProtoMessage() {}

func (x *AssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_asset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetResponse.ProtoReflect.Descriptor instead.
func (*AssetResponse) Descriptor() ([]byte, []int) {
	return file_protos_asset_proto_rawDescGZIP(), []int{1}
}

func (x *AssetResponse) GetNumberOfAssets() int32 {
	if x != nil {
		return x.NumberOfAssets
	}
	return 0
}

var File_protos_asset_proto protoreflect.FileDescriptor

var file_protos_asset_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x62, 0x0a, 0x0c,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x31, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x31,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x32, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x32,
	0x22, 0x37, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x66, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x32, 0x4b, 0x0a, 0x05, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_asset_proto_rawDescOnce sync.Once
	file_protos_asset_proto_rawDescData = file_protos_asset_proto_rawDesc
)

func file_protos_asset_proto_rawDescGZIP() []byte {
	file_protos_asset_proto_rawDescOnce.Do(func() {
		file_protos_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_asset_proto_rawDescData)
	})
	return file_protos_asset_proto_rawDescData
}

var file_protos_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_asset_proto_goTypes = []interface{}{
	(*AssetRequest)(nil),  // 0: protos.AssetRequest
	(*AssetResponse)(nil), // 1: protos.AssetResponse
}
var file_protos_asset_proto_depIdxs = []int32{
	0, // 0: protos.Asset.GetAssetChanges:input_type -> protos.AssetRequest
	1, // 1: protos.Asset.GetAssetChanges:output_type -> protos.AssetResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_asset_proto_init() }
func file_protos_asset_proto_init() {
	if File_protos_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetRequest); i {
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
		file_protos_asset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetResponse); i {
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
			RawDescriptor: file_protos_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_asset_proto_goTypes,
		DependencyIndexes: file_protos_asset_proto_depIdxs,
		MessageInfos:      file_protos_asset_proto_msgTypes,
	}.Build()
	File_protos_asset_proto = out.File
	file_protos_asset_proto_rawDesc = nil
	file_protos_asset_proto_goTypes = nil
	file_protos_asset_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AssetClient is the client API for Asset service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssetClient interface {
	GetAssetChanges(ctx context.Context, in *AssetRequest, opts ...grpc.CallOption) (Asset_GetAssetChangesClient, error)
}

type assetClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetClient(cc grpc.ClientConnInterface) AssetClient {
	return &assetClient{cc}
}

func (c *assetClient) GetAssetChanges(ctx context.Context, in *AssetRequest, opts ...grpc.CallOption) (Asset_GetAssetChangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Asset_serviceDesc.Streams[0], "/protos.Asset/GetAssetChanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &assetGetAssetChangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Asset_GetAssetChangesClient interface {
	Recv() (*AssetResponse, error)
	grpc.ClientStream
}

type assetGetAssetChangesClient struct {
	grpc.ClientStream
}

func (x *assetGetAssetChangesClient) Recv() (*AssetResponse, error) {
	m := new(AssetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AssetServer is the server API for Asset service.
type AssetServer interface {
	GetAssetChanges(*AssetRequest, Asset_GetAssetChangesServer) error
}

// UnimplementedAssetServer can be embedded to have forward compatible implementations.
type UnimplementedAssetServer struct {
}

func (*UnimplementedAssetServer) GetAssetChanges(*AssetRequest, Asset_GetAssetChangesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAssetChanges not implemented")
}

func RegisterAssetServer(s *grpc.Server, srv AssetServer) {
	s.RegisterService(&_Asset_serviceDesc, srv)
}

func _Asset_GetAssetChanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AssetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AssetServer).GetAssetChanges(m, &assetGetAssetChangesServer{stream})
}

type Asset_GetAssetChangesServer interface {
	Send(*AssetResponse) error
	grpc.ServerStream
}

type assetGetAssetChangesServer struct {
	grpc.ServerStream
}

func (x *assetGetAssetChangesServer) Send(m *AssetResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Asset_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Asset",
	HandlerType: (*AssetServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAssetChanges",
			Handler:       _Asset_GetAssetChanges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/asset.proto",
}
