// Code generated by protoc-gen-go. DO NOT EDIT.
// source: survey-response.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SurveyItemResponse struct {
	Key  string        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Meta *ResponseMeta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	// for item groups:
	Items []*SurveyItemResponse `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	// for single items:
	Response             *ResponseItem `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SurveyItemResponse) Reset()         { *m = SurveyItemResponse{} }
func (m *SurveyItemResponse) String() string { return proto.CompactTextString(m) }
func (*SurveyItemResponse) ProtoMessage()    {}
func (*SurveyItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f7d4a107e829578, []int{0}
}

func (m *SurveyItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyItemResponse.Unmarshal(m, b)
}
func (m *SurveyItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyItemResponse.Marshal(b, m, deterministic)
}
func (m *SurveyItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyItemResponse.Merge(m, src)
}
func (m *SurveyItemResponse) XXX_Size() int {
	return xxx_messageInfo_SurveyItemResponse.Size(m)
}
func (m *SurveyItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyItemResponse proto.InternalMessageInfo

func (m *SurveyItemResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SurveyItemResponse) GetMeta() *ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *SurveyItemResponse) GetItems() []*SurveyItemResponse {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SurveyItemResponse) GetResponse() *ResponseItem {
	if m != nil {
		return m.Response
	}
	return nil
}

type ResponseItem struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Dtype string `protobuf:"bytes,3,opt,name=dtype,proto3" json:"dtype,omitempty"`
	// For response option groups:
	Items                []*ResponseItem `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ResponseItem) Reset()         { *m = ResponseItem{} }
func (m *ResponseItem) String() string { return proto.CompactTextString(m) }
func (*ResponseItem) ProtoMessage()    {}
func (*ResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f7d4a107e829578, []int{1}
}

func (m *ResponseItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseItem.Unmarshal(m, b)
}
func (m *ResponseItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseItem.Marshal(b, m, deterministic)
}
func (m *ResponseItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseItem.Merge(m, src)
}
func (m *ResponseItem) XXX_Size() int {
	return xxx_messageInfo_ResponseItem.Size(m)
}
func (m *ResponseItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseItem.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseItem proto.InternalMessageInfo

func (m *ResponseItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ResponseItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ResponseItem) GetDtype() string {
	if m != nil {
		return m.Dtype
	}
	return ""
}

func (m *ResponseItem) GetItems() []*ResponseItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type ResponseMeta struct {
	Position   int32  `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	LocaleCode string `protobuf:"bytes,2,opt,name=locale_code,json=localeCode,proto3" json:"locale_code,omitempty"`
	Version    int32  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// timestamps:
	Rendered             []int64  `protobuf:"varint,4,rep,packed,name=rendered,proto3" json:"rendered,omitempty"`
	Displayed            []int64  `protobuf:"varint,5,rep,packed,name=displayed,proto3" json:"displayed,omitempty"`
	Responded            []int64  `protobuf:"varint,6,rep,packed,name=responded,proto3" json:"responded,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMeta) Reset()         { *m = ResponseMeta{} }
func (m *ResponseMeta) String() string { return proto.CompactTextString(m) }
func (*ResponseMeta) ProtoMessage()    {}
func (*ResponseMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f7d4a107e829578, []int{2}
}

func (m *ResponseMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMeta.Unmarshal(m, b)
}
func (m *ResponseMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMeta.Marshal(b, m, deterministic)
}
func (m *ResponseMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMeta.Merge(m, src)
}
func (m *ResponseMeta) XXX_Size() int {
	return xxx_messageInfo_ResponseMeta.Size(m)
}
func (m *ResponseMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMeta proto.InternalMessageInfo

func (m *ResponseMeta) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *ResponseMeta) GetLocaleCode() string {
	if m != nil {
		return m.LocaleCode
	}
	return ""
}

func (m *ResponseMeta) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ResponseMeta) GetRendered() []int64 {
	if m != nil {
		return m.Rendered
	}
	return nil
}

func (m *ResponseMeta) GetDisplayed() []int64 {
	if m != nil {
		return m.Displayed
	}
	return nil
}

func (m *ResponseMeta) GetResponded() []int64 {
	if m != nil {
		return m.Responded
	}
	return nil
}

func init() {
	proto.RegisterType((*SurveyItemResponse)(nil), "inf.survey_response.SurveyItemResponse")
	proto.RegisterType((*ResponseItem)(nil), "inf.survey_response.ResponseItem")
	proto.RegisterType((*ResponseMeta)(nil), "inf.survey_response.ResponseMeta")
}

func init() {
	proto.RegisterFile("survey-response.proto", fileDescriptor_9f7d4a107e829578)
}

var fileDescriptor_9f7d4a107e829578 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0x4a, 0x43, 0x31,
	0x14, 0xe5, 0x99, 0xbe, 0xda, 0xde, 0x3a, 0x48, 0x54, 0x08, 0x22, 0x58, 0xbb, 0xd8, 0xc5, 0x37,
	0x54, 0xc4, 0xa9, 0x8b, 0x4e, 0x0e, 0x2e, 0x71, 0x73, 0x29, 0xb1, 0xb9, 0x42, 0xf0, 0xf5, 0x25,
	0x24, 0x69, 0xe1, 0x7d, 0x80, 0x9f, 0xe5, 0x97, 0xf8, 0x33, 0x92, 0xc4, 0x3c, 0x0b, 0x76, 0xe8,
	0x96, 0x73, 0xce, 0x3d, 0x37, 0xe7, 0x84, 0xc0, 0x99, 0x5b, 0xdb, 0x0d, 0xb6, 0x37, 0x16, 0x9d,
	0xd1, 0x8d, 0xc3, 0xca, 0x58, 0xed, 0x35, 0x3d, 0x51, 0xcd, 0x7b, 0x95, 0xa4, 0x45, 0x96, 0x26,
	0xdf, 0x05, 0xd0, 0x97, 0xc8, 0x3d, 0x79, 0x5c, 0xf1, 0x5f, 0x9a, 0x1e, 0x03, 0xf9, 0xc0, 0x96,
	0x15, 0xe3, 0x62, 0x3a, 0xe4, 0xe1, 0x48, 0xef, 0xa0, 0xb7, 0x42, 0x2f, 0xd8, 0xc1, 0xb8, 0x98,
	0x8e, 0x66, 0x57, 0xd5, 0x8e, 0x65, 0x55, 0xb6, 0x3f, 0xa3, 0x17, 0x3c, 0x8e, 0xd3, 0x39, 0x94,
	0xca, 0xe3, 0xca, 0x31, 0x32, 0x26, 0xd3, 0xd1, 0xec, 0x7a, 0xa7, 0xef, 0x7f, 0x00, 0x9e, 0x5c,
	0x74, 0x0e, 0x83, 0x3c, 0xc5, 0x7a, 0x7b, 0xdc, 0x1c, 0x77, 0x74, 0x96, 0xc9, 0x67, 0x01, 0x47,
	0xdb, 0xd2, 0x8e, 0x5e, 0xa7, 0x50, 0x6e, 0x44, 0xbd, 0xc6, 0x58, 0x6c, 0xc8, 0x13, 0x08, 0xac,
	0xf4, 0xad, 0x41, 0x46, 0x12, 0x1b, 0x01, 0xbd, 0xcf, 0x65, 0x7a, 0xb1, 0xcc, 0x1e, 0x51, 0xd2,
	0xfc, 0xe4, 0x6b, 0x2b, 0x47, 0x78, 0x1c, 0x7a, 0x0e, 0x03, 0xa3, 0x9d, 0xf2, 0x4a, 0x37, 0x31,
	0x4c, 0xc9, 0x3b, 0x4c, 0x2f, 0x61, 0x54, 0xeb, 0xa5, 0xa8, 0x71, 0xb1, 0xd4, 0x32, 0xe7, 0x82,
	0x44, 0x3d, 0x6a, 0x89, 0x94, 0xc1, 0xe1, 0x06, 0xad, 0x0b, 0x5e, 0x12, 0xbd, 0x19, 0x86, 0xb5,
	0x16, 0x1b, 0x89, 0x16, 0x65, 0xcc, 0x48, 0x78, 0x87, 0xe9, 0x05, 0x0c, 0xa5, 0x72, 0xa6, 0x16,
	0x2d, 0x4a, 0x56, 0x46, 0xf1, 0x8f, 0x08, 0x6a, 0x6a, 0x20, 0x51, 0xb2, 0x7e, 0x52, 0x3b, 0xe2,
	0xa1, 0x7c, 0x25, 0xc2, 0xa8, 0xb7, 0x7e, 0xfc, 0x48, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x19, 0x4d, 0x00, 0xad, 0x61, 0x02, 0x00, 0x00,
}
