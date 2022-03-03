// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: api/data.proto

package api

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

type TradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sym string `protobuf:"bytes,1,opt,name=sym,proto3" json:"sym,omitempty"`
}

func (x *TradeRequest) Reset() {
	*x = TradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeRequest) ProtoMessage() {}

func (x *TradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeRequest.ProtoReflect.Descriptor instead.
func (*TradeRequest) Descriptor() ([]byte, []int) {
	return file_api_data_proto_rawDescGZIP(), []int{0}
}

func (x *TradeRequest) GetSym() string {
	if x != nil {
		return x.Sym
	}
	return ""
}

type TradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trades []*Trade `protobuf:"bytes,1,rep,name=trades,proto3" json:"trades,omitempty"`
}

func (x *TradeResponse) Reset() {
	*x = TradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeResponse) ProtoMessage() {}

func (x *TradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeResponse.ProtoReflect.Descriptor instead.
func (*TradeResponse) Descriptor() ([]byte, []int) {
	return file_api_data_proto_rawDescGZIP(), []int{1}
}

func (x *TradeResponse) GetTrades() []*Trade {
	if x != nil {
		return x.Trades
	}
	return nil
}

type Trade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Sym   string                 `protobuf:"bytes,2,opt,name=sym,proto3" json:"sym,omitempty"`
	Price float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Qty   int64                  `protobuf:"varint,4,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *Trade) Reset() {
	*x = Trade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trade) ProtoMessage() {}

func (x *Trade) ProtoReflect() protoreflect.Message {
	mi := &file_api_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trade.ProtoReflect.Descriptor instead.
func (*Trade) Descriptor() ([]byte, []int) {
	return file_api_data_proto_rawDescGZIP(), []int{2}
}

func (x *Trade) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Trade) GetSym() string {
	if x != nil {
		return x.Sym
	}
	return ""
}

func (x *Trade) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Trade) GetQty() int64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

var File_api_data_proto protoreflect.FileDescriptor

var file_api_data_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x79, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x79, 0x6d, 0x22, 0x33, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x06, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x22, 0x71, 0x0a,
	0x05, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x79, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x79, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x71, 0x74, 0x79,
	0x32, 0x42, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x33, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_data_proto_rawDescOnce sync.Once
	file_api_data_proto_rawDescData = file_api_data_proto_rawDesc
)

func file_api_data_proto_rawDescGZIP() []byte {
	file_api_data_proto_rawDescOnce.Do(func() {
		file_api_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_data_proto_rawDescData)
	})
	return file_api_data_proto_rawDescData
}

var file_api_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_data_proto_goTypes = []interface{}{
	(*TradeRequest)(nil),          // 0: api.TradeRequest
	(*TradeResponse)(nil),         // 1: api.TradeResponse
	(*Trade)(nil),                 // 2: api.Trade
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_api_data_proto_depIdxs = []int32{
	2, // 0: api.TradeResponse.trades:type_name -> api.Trade
	3, // 1: api.Trade.time:type_name -> google.protobuf.Timestamp
	0, // 2: api.DataService.GetTrade:input_type -> api.TradeRequest
	1, // 3: api.DataService.GetTrade:output_type -> api.TradeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_data_proto_init() }
func file_api_data_proto_init() {
	if File_api_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeRequest); i {
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
		file_api_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeResponse); i {
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
		file_api_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trade); i {
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
			RawDescriptor: file_api_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_data_proto_goTypes,
		DependencyIndexes: file_api_data_proto_depIdxs,
		MessageInfos:      file_api_data_proto_msgTypes,
	}.Build()
	File_api_data_proto = out.File
	file_api_data_proto_rawDesc = nil
	file_api_data_proto_goTypes = nil
	file_api_data_proto_depIdxs = nil
}
