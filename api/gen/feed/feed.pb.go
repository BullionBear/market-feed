// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: api/proto/feed.proto

package feed

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Status int32

const (
	Status_CREATED      Status = 0
	Status_OK           Status = 1
	Status_INITIALIZING Status = 2
	Status_ERROR        Status = 3
	Status_UNKNOWN      Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "CREATED",
		1: "OK",
		2: "INITIALIZING",
		3: "ERROR",
		4: "UNKNOWN",
	}
	Status_value = map[string]int32{
		"CREATED":      0,
		"OK":           1,
		"INITIALIZING": 2,
		"ERROR":        3,
		"UNKNOWN":      4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_feed_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_proto_feed_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_feed_proto_rawDescGZIP(), []int{0}
}

type Kline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenTime                 int64   `protobuf:"varint,1,opt,name=openTime,proto3" json:"openTime,omitempty"`
	Open                     float64 `protobuf:"fixed64,2,opt,name=open,proto3" json:"open,omitempty"`
	High                     float64 `protobuf:"fixed64,3,opt,name=high,proto3" json:"high,omitempty"`
	Low                      float64 `protobuf:"fixed64,4,opt,name=low,proto3" json:"low,omitempty"`
	Close                    float64 `protobuf:"fixed64,5,opt,name=close,proto3" json:"close,omitempty"`
	Volume                   float64 `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume,omitempty"`
	CloseTime                int64   `protobuf:"varint,7,opt,name=closeTime,proto3" json:"closeTime,omitempty"`
	QuoteAssetVolume         float64 `protobuf:"fixed64,8,opt,name=quoteAssetVolume,proto3" json:"quoteAssetVolume,omitempty"`
	TradeNum                 int64   `protobuf:"varint,9,opt,name=tradeNum,proto3" json:"tradeNum,omitempty"`
	TakerBuyBaseAssetVolume  float64 `protobuf:"fixed64,10,opt,name=takerBuyBaseAssetVolume,proto3" json:"takerBuyBaseAssetVolume,omitempty"`
	TakerBuyQuoteAssetVolume float64 `protobuf:"fixed64,11,opt,name=takerBuyQuoteAssetVolume,proto3" json:"takerBuyQuoteAssetVolume,omitempty"`
}

func (x *Kline) Reset() {
	*x = Kline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kline) ProtoMessage() {}

func (x *Kline) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kline.ProtoReflect.Descriptor instead.
func (*Kline) Descriptor() ([]byte, []int) {
	return file_api_proto_feed_proto_rawDescGZIP(), []int{0}
}

func (x *Kline) GetOpenTime() int64 {
	if x != nil {
		return x.OpenTime
	}
	return 0
}

func (x *Kline) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Kline) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Kline) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Kline) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *Kline) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Kline) GetCloseTime() int64 {
	if x != nil {
		return x.CloseTime
	}
	return 0
}

func (x *Kline) GetQuoteAssetVolume() float64 {
	if x != nil {
		return x.QuoteAssetVolume
	}
	return 0
}

func (x *Kline) GetTradeNum() int64 {
	if x != nil {
		return x.TradeNum
	}
	return 0
}

func (x *Kline) GetTakerBuyBaseAssetVolume() float64 {
	if x != nil {
		return x.TakerBuyBaseAssetVolume
	}
	return 0
}

func (x *Kline) GetTakerBuyQuoteAssetVolume() float64 {
	if x != nil {
		return x.TakerBuyQuoteAssetVolume
	}
	return 0
}

type ReadKlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ReadKlineRequest) Reset() {
	*x = ReadKlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadKlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadKlineRequest) ProtoMessage() {}

func (x *ReadKlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadKlineRequest.ProtoReflect.Descriptor instead.
func (*ReadKlineRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_feed_proto_rawDescGZIP(), []int{1}
}

func (x *ReadKlineRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ReadKlineRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    Status `protobuf:"varint,1,opt,name=status,proto3,enum=feed.Status" json:"status,omitempty"`
	Start     int64  `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End       int64  `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Size      int64  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_feed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_feed_proto_rawDescGZIP(), []int{2}
}

func (x *StatusResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_CREATED
}

func (x *StatusResponse) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *StatusResponse) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *StatusResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *StatusResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Length int64  `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_feed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_feed_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ConfigResponse) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type SubscriberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscribers []int64 `protobuf:"varint,1,rep,packed,name=subscribers,proto3" json:"subscribers,omitempty"`
}

func (x *SubscriberResponse) Reset() {
	*x = SubscriberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_feed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriberResponse) ProtoMessage() {}

func (x *SubscriberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_feed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriberResponse.ProtoReflect.Descriptor instead.
func (*SubscriberResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_feed_proto_rawDescGZIP(), []int{4}
}

func (x *SubscriberResponse) GetSubscribers() []int64 {
	if x != nil {
		return x.Subscribers
	}
	return nil
}

type KlineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kline *Kline `protobuf:"bytes,1,opt,name=kline,proto3" json:"kline,omitempty"`
}

func (x *KlineResponse) Reset() {
	*x = KlineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_feed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KlineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KlineResponse) ProtoMessage() {}

func (x *KlineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_feed_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KlineResponse.ProtoReflect.Descriptor instead.
func (*KlineResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_feed_proto_rawDescGZIP(), []int{5}
}

func (x *KlineResponse) GetKline() *Kline {
	if x != nil {
		return x.Kline
	}
	return nil
}

var File_api_proto_feed_proto protoreflect.FileDescriptor

var file_api_proto_feed_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x65, 0x65, 0x64, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02, 0x0a, 0x05, 0x4b, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f,
	0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x10, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x38, 0x0a,
	0x17, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x42, 0x75, 0x79, 0x42, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x17,
	0x74, 0x61, 0x6b, 0x65, 0x72, 0x42, 0x75, 0x79, 0x42, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x74, 0x61, 0x6b, 0x65, 0x72,
	0x42, 0x75, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x18, 0x74, 0x61, 0x6b, 0x65, 0x72,
	0x42, 0x75, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x4b, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22,
	0x90, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x22, 0x40, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x22, 0x36, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x0b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x22, 0x32, 0x0a, 0x0d,
	0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x05, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66,
	0x65, 0x65, 0x64, 0x2e, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x05, 0x6b, 0x6c, 0x69, 0x6e, 0x65,
	0x2a, 0x47, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x04, 0x32, 0xc6, 0x02, 0x0a, 0x04, 0x46, 0x65,
	0x65, 0x64, 0x12, 0x39, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x14, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x18, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x4b, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x13,
	0x52, 0x65, 0x61, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x4b, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4b,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x2e, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x66, 0x65, 0x65, 0x64, 0x3b, 0x66, 0x65, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_feed_proto_rawDescOnce sync.Once
	file_api_proto_feed_proto_rawDescData = file_api_proto_feed_proto_rawDesc
)

func file_api_proto_feed_proto_rawDescGZIP() []byte {
	file_api_proto_feed_proto_rawDescOnce.Do(func() {
		file_api_proto_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_feed_proto_rawDescData)
	})
	return file_api_proto_feed_proto_rawDescData
}

var file_api_proto_feed_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_feed_proto_goTypes = []interface{}{
	(Status)(0),                // 0: feed.Status
	(*Kline)(nil),              // 1: feed.Kline
	(*ReadKlineRequest)(nil),   // 2: feed.ReadKlineRequest
	(*StatusResponse)(nil),     // 3: feed.StatusResponse
	(*ConfigResponse)(nil),     // 4: feed.ConfigResponse
	(*SubscriberResponse)(nil), // 5: feed.SubscriberResponse
	(*KlineResponse)(nil),      // 6: feed.KlineResponse
	(*empty.Empty)(nil),        // 7: google.protobuf.Empty
}
var file_api_proto_feed_proto_depIdxs = []int32{
	0, // 0: feed.StatusResponse.status:type_name -> feed.Status
	1, // 1: feed.KlineResponse.kline:type_name -> feed.Kline
	7, // 2: feed.Feed.GetConfig:input_type -> google.protobuf.Empty
	7, // 3: feed.Feed.GetStatus:input_type -> google.protobuf.Empty
	7, // 4: feed.Feed.GetSubscriber:input_type -> google.protobuf.Empty
	7, // 5: feed.Feed.SubscribeKline:input_type -> google.protobuf.Empty
	2, // 6: feed.Feed.ReadHistoricalKline:input_type -> feed.ReadKlineRequest
	4, // 7: feed.Feed.GetConfig:output_type -> feed.ConfigResponse
	3, // 8: feed.Feed.GetStatus:output_type -> feed.StatusResponse
	5, // 9: feed.Feed.GetSubscriber:output_type -> feed.SubscriberResponse
	6, // 10: feed.Feed.SubscribeKline:output_type -> feed.KlineResponse
	6, // 11: feed.Feed.ReadHistoricalKline:output_type -> feed.KlineResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_feed_proto_init() }
func file_api_proto_feed_proto_init() {
	if File_api_proto_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kline); i {
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
		file_api_proto_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadKlineRequest); i {
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
		file_api_proto_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_api_proto_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
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
		file_api_proto_feed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriberResponse); i {
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
		file_api_proto_feed_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KlineResponse); i {
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
			RawDescriptor: file_api_proto_feed_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_feed_proto_goTypes,
		DependencyIndexes: file_api_proto_feed_proto_depIdxs,
		EnumInfos:         file_api_proto_feed_proto_enumTypes,
		MessageInfos:      file_api_proto_feed_proto_msgTypes,
	}.Build()
	File_api_proto_feed_proto = out.File
	file_api_proto_feed_proto_rawDesc = nil
	file_api_proto_feed_proto_goTypes = nil
	file_api_proto_feed_proto_depIdxs = nil
}
