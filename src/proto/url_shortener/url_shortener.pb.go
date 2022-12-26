// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: url_shortener/url_shortener.proto

package url_shortener

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

type ShortUrlDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationURL string `protobuf:"bytes,1,opt,name=destinationURL,proto3" json:"destinationURL,omitempty"`
	ShortURL       string `protobuf:"bytes,2,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
	CreatedAt      string `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *ShortUrlDTO) Reset() {
	*x = ShortUrlDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_url_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortUrlDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortUrlDTO) ProtoMessage() {}

func (x *ShortUrlDTO) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_url_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortUrlDTO.ProtoReflect.Descriptor instead.
func (*ShortUrlDTO) Descriptor() ([]byte, []int) {
	return file_url_shortener_url_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *ShortUrlDTO) GetDestinationURL() string {
	if x != nil {
		return x.DestinationURL
	}
	return ""
}

func (x *ShortUrlDTO) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

func (x *ShortUrlDTO) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationURL string `protobuf:"bytes,1,opt,name=destinationURL,proto3" json:"destinationURL,omitempty"`
}

func (x *CreateShortURLRequest) Reset() {
	*x = CreateShortURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_url_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLRequest) ProtoMessage() {}

func (x *CreateShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_url_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLRequest.ProtoReflect.Descriptor instead.
func (*CreateShortURLRequest) Descriptor() ([]byte, []int) {
	return file_url_shortener_url_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShortURLRequest) GetDestinationURL() string {
	if x != nil {
		return x.DestinationURL
	}
	return ""
}

type CreateShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL *ShortUrlDTO `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *CreateShortURLResponse) Reset() {
	*x = CreateShortURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_url_shortener_url_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLResponse) ProtoMessage() {}

func (x *CreateShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_url_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLResponse.ProtoReflect.Descriptor instead.
func (*CreateShortURLResponse) Descriptor() ([]byte, []int) {
	return file_url_shortener_url_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *CreateShortURLResponse) GetURL() *ShortUrlDTO {
	if x != nil {
		return x.URL
	}
	return nil
}

var File_url_shortener_url_shortener_proto protoreflect.FileDescriptor

var file_url_shortener_url_shortener_proto_rawDesc = []byte{
	0x0a, 0x21, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f,
	0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x22, 0x6f, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x44, 0x54,
	0x4f, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x3f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x52, 0x4c, 0x22, 0x46, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x72,
	0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x44, 0x54, 0x4f, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x32, 0x76, 0x0a, 0x13,
	0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x24, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x75, 0x72,
	0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xb6, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x72, 0x6c,
	0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x42, 0x11, 0x55, 0x72, 0x6c, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x75, 0x72, 0x62,
	0x61, 0x6a, 0x35, 0x31, 0x32, 0x34, 0x38, 0x37, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0xa2,
	0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0xca, 0x02, 0x0c, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0xe2, 0x02, 0x18, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0c, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_url_shortener_url_shortener_proto_rawDescOnce sync.Once
	file_url_shortener_url_shortener_proto_rawDescData = file_url_shortener_url_shortener_proto_rawDesc
)

func file_url_shortener_url_shortener_proto_rawDescGZIP() []byte {
	file_url_shortener_url_shortener_proto_rawDescOnce.Do(func() {
		file_url_shortener_url_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_shortener_url_shortener_proto_rawDescData)
	})
	return file_url_shortener_url_shortener_proto_rawDescData
}

var file_url_shortener_url_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_url_shortener_url_shortener_proto_goTypes = []interface{}{
	(*ShortUrlDTO)(nil),            // 0: url_shortener.ShortUrlDTO
	(*CreateShortURLRequest)(nil),  // 1: url_shortener.CreateShortURLRequest
	(*CreateShortURLResponse)(nil), // 2: url_shortener.CreateShortURLResponse
}
var file_url_shortener_url_shortener_proto_depIdxs = []int32{
	0, // 0: url_shortener.CreateShortURLResponse.URL:type_name -> url_shortener.ShortUrlDTO
	1, // 1: url_shortener.UrlShortenerService.CreateShortURL:input_type -> url_shortener.CreateShortURLRequest
	2, // 2: url_shortener.UrlShortenerService.CreateShortURL:output_type -> url_shortener.CreateShortURLResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_url_shortener_url_shortener_proto_init() }
func file_url_shortener_url_shortener_proto_init() {
	if File_url_shortener_url_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_url_shortener_url_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortUrlDTO); i {
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
		file_url_shortener_url_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortURLRequest); i {
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
		file_url_shortener_url_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortURLResponse); i {
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
			RawDescriptor: file_url_shortener_url_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_url_shortener_url_shortener_proto_goTypes,
		DependencyIndexes: file_url_shortener_url_shortener_proto_depIdxs,
		MessageInfos:      file_url_shortener_url_shortener_proto_msgTypes,
	}.Build()
	File_url_shortener_url_shortener_proto = out.File
	file_url_shortener_url_shortener_proto_rawDesc = nil
	file_url_shortener_url_shortener_proto_goTypes = nil
	file_url_shortener_url_shortener_proto_depIdxs = nil
}
