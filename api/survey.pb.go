// Code generated by protoc-gen-go. DO NOT EDIT.
// source: survey.proto

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

type SurveyContext struct {
	Mode                 string            `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	PreviousResponses    []*SurveyResponse `protobuf:"bytes,2,rep,name=previous_responses,json=previousResponses,proto3" json:"previous_responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SurveyContext) Reset()         { *m = SurveyContext{} }
func (m *SurveyContext) String() string { return proto.CompactTextString(m) }
func (*SurveyContext) ProtoMessage()    {}
func (*SurveyContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{0}
}

func (m *SurveyContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyContext.Unmarshal(m, b)
}
func (m *SurveyContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyContext.Marshal(b, m, deterministic)
}
func (m *SurveyContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyContext.Merge(m, src)
}
func (m *SurveyContext) XXX_Size() int {
	return xxx_messageInfo_SurveyContext.Size(m)
}
func (m *SurveyContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyContext.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyContext proto.InternalMessageInfo

func (m *SurveyContext) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *SurveyContext) GetPreviousResponses() []*SurveyResponse {
	if m != nil {
		return m.PreviousResponses
	}
	return nil
}

// expression that define how the survey context should be generated by the
// server
type SurveyContextDef struct {
	Mode                 *ExpressionArg `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
	PreviousResponses    []*Expression  `protobuf:"bytes,4,rep,name=previous_responses,json=previousResponses,proto3" json:"previous_responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SurveyContextDef) Reset()         { *m = SurveyContextDef{} }
func (m *SurveyContextDef) String() string { return proto.CompactTextString(m) }
func (*SurveyContextDef) ProtoMessage()    {}
func (*SurveyContextDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{1}
}

func (m *SurveyContextDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyContextDef.Unmarshal(m, b)
}
func (m *SurveyContextDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyContextDef.Marshal(b, m, deterministic)
}
func (m *SurveyContextDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyContextDef.Merge(m, src)
}
func (m *SurveyContextDef) XXX_Size() int {
	return xxx_messageInfo_SurveyContextDef.Size(m)
}
func (m *SurveyContextDef) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyContextDef.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyContextDef proto.InternalMessageInfo

func (m *SurveyContextDef) GetMode() *ExpressionArg {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *SurveyContextDef) GetPreviousResponses() []*Expression {
	if m != nil {
		return m.PreviousResponses
	}
	return nil
}

type Survey struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 []*LocalisedObject `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	Current              *SurveyVersion     `protobuf:"bytes,3,opt,name=current,proto3" json:"current,omitempty"`
	History              []*SurveyVersion   `protobuf:"bytes,4,rep,name=history,proto3" json:"history,omitempty"`
	Description          []*LocalisedObject `protobuf:"bytes,5,rep,name=description,proto3" json:"description,omitempty"`
	PrefillRules         []*Expression      `protobuf:"bytes,6,rep,name=prefill_rules,json=prefillRules,proto3" json:"prefill_rules,omitempty"`
	ContextRules         *SurveyContextDef  `protobuf:"bytes,7,opt,name=context_rules,json=contextRules,proto3" json:"context_rules,omitempty"`
	MaxItemsPerPage      *MaxItemsPerPage   `protobuf:"bytes,8,opt,name=max_items_per_page,json=maxItemsPerPage,proto3" json:"max_items_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Survey) Reset()         { *m = Survey{} }
func (m *Survey) String() string { return proto.CompactTextString(m) }
func (*Survey) ProtoMessage()    {}
func (*Survey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{2}
}

func (m *Survey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Survey.Unmarshal(m, b)
}
func (m *Survey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Survey.Marshal(b, m, deterministic)
}
func (m *Survey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Survey.Merge(m, src)
}
func (m *Survey) XXX_Size() int {
	return xxx_messageInfo_Survey.Size(m)
}
func (m *Survey) XXX_DiscardUnknown() {
	xxx_messageInfo_Survey.DiscardUnknown(m)
}

var xxx_messageInfo_Survey proto.InternalMessageInfo

func (m *Survey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Survey) GetName() []*LocalisedObject {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Survey) GetCurrent() *SurveyVersion {
	if m != nil {
		return m.Current
	}
	return nil
}

func (m *Survey) GetHistory() []*SurveyVersion {
	if m != nil {
		return m.History
	}
	return nil
}

func (m *Survey) GetDescription() []*LocalisedObject {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Survey) GetPrefillRules() []*Expression {
	if m != nil {
		return m.PrefillRules
	}
	return nil
}

func (m *Survey) GetContextRules() *SurveyContextDef {
	if m != nil {
		return m.ContextRules
	}
	return nil
}

func (m *Survey) GetMaxItemsPerPage() *MaxItemsPerPage {
	if m != nil {
		return m.MaxItemsPerPage
	}
	return nil
}

type MaxItemsPerPage struct {
	Large                int32    `protobuf:"varint,1,opt,name=large,proto3" json:"large,omitempty"`
	Small                int32    `protobuf:"varint,2,opt,name=small,proto3" json:"small,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxItemsPerPage) Reset()         { *m = MaxItemsPerPage{} }
func (m *MaxItemsPerPage) String() string { return proto.CompactTextString(m) }
func (*MaxItemsPerPage) ProtoMessage()    {}
func (*MaxItemsPerPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{3}
}

func (m *MaxItemsPerPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxItemsPerPage.Unmarshal(m, b)
}
func (m *MaxItemsPerPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxItemsPerPage.Marshal(b, m, deterministic)
}
func (m *MaxItemsPerPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxItemsPerPage.Merge(m, src)
}
func (m *MaxItemsPerPage) XXX_Size() int {
	return xxx_messageInfo_MaxItemsPerPage.Size(m)
}
func (m *MaxItemsPerPage) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxItemsPerPage.DiscardUnknown(m)
}

var xxx_messageInfo_MaxItemsPerPage proto.InternalMessageInfo

func (m *MaxItemsPerPage) GetLarge() int32 {
	if m != nil {
		return m.Large
	}
	return 0
}

func (m *MaxItemsPerPage) GetSmall() int32 {
	if m != nil {
		return m.Small
	}
	return 0
}

type SurveyVersion struct {
	Published            int64       `protobuf:"varint,1,opt,name=published,proto3" json:"published,omitempty"`
	Unpublished          int64       `protobuf:"varint,2,opt,name=unpublished,proto3" json:"unpublished,omitempty"`
	SurveyDefinition     *SurveyItem `protobuf:"bytes,3,opt,name=survey_definition,json=surveyDefinition,proto3" json:"survey_definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SurveyVersion) Reset()         { *m = SurveyVersion{} }
func (m *SurveyVersion) String() string { return proto.CompactTextString(m) }
func (*SurveyVersion) ProtoMessage()    {}
func (*SurveyVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{4}
}

func (m *SurveyVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyVersion.Unmarshal(m, b)
}
func (m *SurveyVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyVersion.Marshal(b, m, deterministic)
}
func (m *SurveyVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyVersion.Merge(m, src)
}
func (m *SurveyVersion) XXX_Size() int {
	return xxx_messageInfo_SurveyVersion.Size(m)
}
func (m *SurveyVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyVersion.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyVersion proto.InternalMessageInfo

func (m *SurveyVersion) GetPublished() int64 {
	if m != nil {
		return m.Published
	}
	return 0
}

func (m *SurveyVersion) GetUnpublished() int64 {
	if m != nil {
		return m.Unpublished
	}
	return 0
}

func (m *SurveyVersion) GetSurveyDefinition() *SurveyItem {
	if m != nil {
		return m.SurveyDefinition
	}
	return nil
}

type SurveyItem struct {
	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key         string      `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Follows     []string    `protobuf:"bytes,3,rep,name=follows,proto3" json:"follows,omitempty"`
	Condition   *Expression `protobuf:"bytes,4,opt,name=condition,proto3" json:"condition,omitempty"`
	Priority    float32     `protobuf:"fixed32,5,opt,name=priority,proto3" json:"priority,omitempty"`
	Version     int32       `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	VersionTags []string    `protobuf:"bytes,7,rep,name=version_tags,json=versionTags,proto3" json:"version_tags,omitempty"`
	// Question group attributes ->
	Items           []*SurveyItem `protobuf:"bytes,8,rep,name=items,proto3" json:"items,omitempty"`
	SelectionMethod *Expression   `protobuf:"bytes,9,opt,name=selection_method,json=selectionMethod,proto3" json:"selection_method,omitempty"`
	// Question attributes ->
	Type                 string         `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Components           *ItemComponent `protobuf:"bytes,11,opt,name=components,proto3" json:"components,omitempty"`
	Validations          []*Validation  `protobuf:"bytes,12,rep,name=validations,proto3" json:"validations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SurveyItem) Reset()         { *m = SurveyItem{} }
func (m *SurveyItem) String() string { return proto.CompactTextString(m) }
func (*SurveyItem) ProtoMessage()    {}
func (*SurveyItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{5}
}

func (m *SurveyItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyItem.Unmarshal(m, b)
}
func (m *SurveyItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyItem.Marshal(b, m, deterministic)
}
func (m *SurveyItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyItem.Merge(m, src)
}
func (m *SurveyItem) XXX_Size() int {
	return xxx_messageInfo_SurveyItem.Size(m)
}
func (m *SurveyItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyItem.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyItem proto.InternalMessageInfo

func (m *SurveyItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SurveyItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SurveyItem) GetFollows() []string {
	if m != nil {
		return m.Follows
	}
	return nil
}

func (m *SurveyItem) GetCondition() *Expression {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *SurveyItem) GetPriority() float32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *SurveyItem) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *SurveyItem) GetVersionTags() []string {
	if m != nil {
		return m.VersionTags
	}
	return nil
}

func (m *SurveyItem) GetItems() []*SurveyItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SurveyItem) GetSelectionMethod() *Expression {
	if m != nil {
		return m.SelectionMethod
	}
	return nil
}

func (m *SurveyItem) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SurveyItem) GetComponents() *ItemComponent {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *SurveyItem) GetValidations() []*Validation {
	if m != nil {
		return m.Validations
	}
	return nil
}

type Validation struct {
	Key                  string      `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Type                 string      `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Rule                 *Expression `protobuf:"bytes,3,opt,name=rule,proto3" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Validation) Reset()         { *m = Validation{} }
func (m *Validation) String() string { return proto.CompactTextString(m) }
func (*Validation) ProtoMessage()    {}
func (*Validation) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{6}
}

func (m *Validation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validation.Unmarshal(m, b)
}
func (m *Validation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validation.Marshal(b, m, deterministic)
}
func (m *Validation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validation.Merge(m, src)
}
func (m *Validation) XXX_Size() int {
	return xxx_messageInfo_Validation.Size(m)
}
func (m *Validation) XXX_DiscardUnknown() {
	xxx_messageInfo_Validation.DiscardUnknown(m)
}

var xxx_messageInfo_Validation proto.InternalMessageInfo

func (m *Validation) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Validation) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Validation) GetRule() *Expression {
	if m != nil {
		return m.Rule
	}
	return nil
}

type ItemComponent struct {
	Role             string             `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Key              string             `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Content          []*LocalisedObject `protobuf:"bytes,3,rep,name=content,proto3" json:"content,omitempty"`
	DisplayCondition *Expression        `protobuf:"bytes,4,opt,name=display_condition,json=displayCondition,proto3" json:"display_condition,omitempty"`
	Disabled         *Expression        `protobuf:"bytes,5,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// group component
	Items []*ItemComponent `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
	Order *Expression      `protobuf:"bytes,7,opt,name=order,proto3" json:"order,omitempty"`
	// response compontent
	Dtype                string                    `protobuf:"bytes,8,opt,name=dtype,proto3" json:"dtype,omitempty"`
	Style                []*ItemComponent_Style    `protobuf:"bytes,9,rep,name=style,proto3" json:"style,omitempty"`
	Description          []*LocalisedObject        `protobuf:"bytes,10,rep,name=description,proto3" json:"description,omitempty"`
	Properties           *ItemComponent_Properties `protobuf:"bytes,11,opt,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ItemComponent) Reset()         { *m = ItemComponent{} }
func (m *ItemComponent) String() string { return proto.CompactTextString(m) }
func (*ItemComponent) ProtoMessage()    {}
func (*ItemComponent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{7}
}

func (m *ItemComponent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemComponent.Unmarshal(m, b)
}
func (m *ItemComponent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemComponent.Marshal(b, m, deterministic)
}
func (m *ItemComponent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemComponent.Merge(m, src)
}
func (m *ItemComponent) XXX_Size() int {
	return xxx_messageInfo_ItemComponent.Size(m)
}
func (m *ItemComponent) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemComponent.DiscardUnknown(m)
}

var xxx_messageInfo_ItemComponent proto.InternalMessageInfo

func (m *ItemComponent) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *ItemComponent) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ItemComponent) GetContent() []*LocalisedObject {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ItemComponent) GetDisplayCondition() *Expression {
	if m != nil {
		return m.DisplayCondition
	}
	return nil
}

func (m *ItemComponent) GetDisabled() *Expression {
	if m != nil {
		return m.Disabled
	}
	return nil
}

func (m *ItemComponent) GetItems() []*ItemComponent {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ItemComponent) GetOrder() *Expression {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *ItemComponent) GetDtype() string {
	if m != nil {
		return m.Dtype
	}
	return ""
}

func (m *ItemComponent) GetStyle() []*ItemComponent_Style {
	if m != nil {
		return m.Style
	}
	return nil
}

func (m *ItemComponent) GetDescription() []*LocalisedObject {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *ItemComponent) GetProperties() *ItemComponent_Properties {
	if m != nil {
		return m.Properties
	}
	return nil
}

type ItemComponent_Style struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemComponent_Style) Reset()         { *m = ItemComponent_Style{} }
func (m *ItemComponent_Style) String() string { return proto.CompactTextString(m) }
func (*ItemComponent_Style) ProtoMessage()    {}
func (*ItemComponent_Style) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{7, 0}
}

func (m *ItemComponent_Style) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemComponent_Style.Unmarshal(m, b)
}
func (m *ItemComponent_Style) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemComponent_Style.Marshal(b, m, deterministic)
}
func (m *ItemComponent_Style) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemComponent_Style.Merge(m, src)
}
func (m *ItemComponent_Style) XXX_Size() int {
	return xxx_messageInfo_ItemComponent_Style.Size(m)
}
func (m *ItemComponent_Style) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemComponent_Style.DiscardUnknown(m)
}

var xxx_messageInfo_ItemComponent_Style proto.InternalMessageInfo

func (m *ItemComponent_Style) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ItemComponent_Style) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ItemComponent_Properties struct {
	Min                  *ExpressionArg `protobuf:"bytes,1,opt,name=min,proto3" json:"min,omitempty"`
	Max                  *ExpressionArg `protobuf:"bytes,2,opt,name=max,proto3" json:"max,omitempty"`
	StepSize             *ExpressionArg `protobuf:"bytes,3,opt,name=stepSize,proto3" json:"stepSize,omitempty"`
	DateInputMode        *ExpressionArg `protobuf:"bytes,4,opt,name=dateInputMode,proto3" json:"dateInputMode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ItemComponent_Properties) Reset()         { *m = ItemComponent_Properties{} }
func (m *ItemComponent_Properties) String() string { return proto.CompactTextString(m) }
func (*ItemComponent_Properties) ProtoMessage()    {}
func (*ItemComponent_Properties) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{7, 1}
}

func (m *ItemComponent_Properties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemComponent_Properties.Unmarshal(m, b)
}
func (m *ItemComponent_Properties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemComponent_Properties.Marshal(b, m, deterministic)
}
func (m *ItemComponent_Properties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemComponent_Properties.Merge(m, src)
}
func (m *ItemComponent_Properties) XXX_Size() int {
	return xxx_messageInfo_ItemComponent_Properties.Size(m)
}
func (m *ItemComponent_Properties) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemComponent_Properties.DiscardUnknown(m)
}

var xxx_messageInfo_ItemComponent_Properties proto.InternalMessageInfo

func (m *ItemComponent_Properties) GetMin() *ExpressionArg {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *ItemComponent_Properties) GetMax() *ExpressionArg {
	if m != nil {
		return m.Max
	}
	return nil
}

func (m *ItemComponent_Properties) GetStepSize() *ExpressionArg {
	if m != nil {
		return m.StepSize
	}
	return nil
}

func (m *ItemComponent_Properties) GetDateInputMode() *ExpressionArg {
	if m != nil {
		return m.DateInputMode
	}
	return nil
}

type LocalisedObject struct {
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// for texts:
	Parts                []*ExpressionArg `protobuf:"bytes,2,rep,name=parts,proto3" json:"parts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LocalisedObject) Reset()         { *m = LocalisedObject{} }
func (m *LocalisedObject) String() string { return proto.CompactTextString(m) }
func (*LocalisedObject) ProtoMessage()    {}
func (*LocalisedObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_a40f94eaa8e6ca46, []int{8}
}

func (m *LocalisedObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalisedObject.Unmarshal(m, b)
}
func (m *LocalisedObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalisedObject.Marshal(b, m, deterministic)
}
func (m *LocalisedObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalisedObject.Merge(m, src)
}
func (m *LocalisedObject) XXX_Size() int {
	return xxx_messageInfo_LocalisedObject.Size(m)
}
func (m *LocalisedObject) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalisedObject.DiscardUnknown(m)
}

var xxx_messageInfo_LocalisedObject proto.InternalMessageInfo

func (m *LocalisedObject) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *LocalisedObject) GetParts() []*ExpressionArg {
	if m != nil {
		return m.Parts
	}
	return nil
}

func init() {
	proto.RegisterType((*SurveyContext)(nil), "inf.survey.SurveyContext")
	proto.RegisterType((*SurveyContextDef)(nil), "inf.survey.SurveyContextDef")
	proto.RegisterType((*Survey)(nil), "inf.survey.Survey")
	proto.RegisterType((*MaxItemsPerPage)(nil), "inf.survey.MaxItemsPerPage")
	proto.RegisterType((*SurveyVersion)(nil), "inf.survey.SurveyVersion")
	proto.RegisterType((*SurveyItem)(nil), "inf.survey.SurveyItem")
	proto.RegisterType((*Validation)(nil), "inf.survey.Validation")
	proto.RegisterType((*ItemComponent)(nil), "inf.survey.ItemComponent")
	proto.RegisterType((*ItemComponent_Style)(nil), "inf.survey.ItemComponent.Style")
	proto.RegisterType((*ItemComponent_Properties)(nil), "inf.survey.ItemComponent.Properties")
	proto.RegisterType((*LocalisedObject)(nil), "inf.survey.LocalisedObject")
}

func init() {
	proto.RegisterFile("survey.proto", fileDescriptor_a40f94eaa8e6ca46)
}

var fileDescriptor_a40f94eaa8e6ca46 = []byte{
	// 942 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x96, 0x3d, 0x1e, 0xc7, 0x2e, 0xc7, 0xc4, 0x69, 0x2d, 0x68, 0xf0, 0xae, 0xc0, 0x18, 0x0e,
	0x3e, 0x80, 0x03, 0x59, 0x56, 0xb0, 0x48, 0x2b, 0xb1, 0x24, 0xa0, 0x5d, 0x44, 0x44, 0x34, 0x41,
	0x7b, 0x80, 0x83, 0xd5, 0x99, 0x29, 0x3b, 0x0d, 0xe3, 0xe9, 0x51, 0x77, 0xdb, 0x6b, 0xf3, 0x18,
	0x88, 0x17, 0xe2, 0xcc, 0x23, 0xf0, 0x16, 0x3c, 0xc1, 0xaa, 0x7f, 0xe6, 0xc7, 0x71, 0x12, 0x67,
	0x6f, 0xdd, 0x55, 0x5f, 0x55, 0x7d, 0xf5, 0xd3, 0x35, 0x03, 0xfb, 0x72, 0x21, 0x96, 0xb8, 0x1e,
	0x67, 0x82, 0x2b, 0x4e, 0x80, 0xa5, 0xd3, 0xb1, 0x95, 0xf4, 0x7b, 0xb8, 0xca, 0x04, 0x4a, 0xc9,
	0x78, 0x6a, 0xb5, 0xfd, 0x77, 0xad, 0xe6, 0x33, 0x81, 0x32, 0xe3, 0xa9, 0x44, 0x2b, 0x1e, 0xbe,
	0x86, 0xee, 0x85, 0x51, 0x9c, 0xf0, 0x54, 0xe1, 0x4a, 0x11, 0x02, 0x8d, 0x39, 0x8f, 0x31, 0xa8,
	0x0d, 0x6a, 0xa3, 0x76, 0x68, 0xce, 0x24, 0x04, 0x92, 0x09, 0x5c, 0x32, 0xbe, 0x90, 0x93, 0xdc,
	0x5e, 0x06, 0xf5, 0x81, 0x37, 0xea, 0x1c, 0x7f, 0x3c, 0x2e, 0xc3, 0x16, 0xca, 0xb1, 0xf5, 0x19,
	0xba, 0x6b, 0x78, 0x98, 0x9b, 0xe7, 0x12, 0x39, 0xfc, 0xab, 0x06, 0xbd, 0x8d, 0xc8, 0xa7, 0x38,
	0x25, 0xc7, 0x2e, 0xb8, 0x37, 0xa8, 0x8d, 0x3a, 0xc7, 0x1f, 0x18, 0xd7, 0x65, 0x26, 0x72, 0xfc,
	0x7d, 0x71, 0x7e, 0x2e, 0x66, 0x8e, 0xdc, 0x8f, 0x37, 0x92, 0x6b, 0x18, 0x72, 0x0f, 0xef, 0xf0,
	0x70, 0x13, 0xa9, 0x7f, 0x3c, 0x68, 0x5a, 0x52, 0xe4, 0x1d, 0xa8, 0xb3, 0xd8, 0x55, 0xa1, 0xce,
	0x62, 0x72, 0x04, 0x8d, 0x94, 0xce, 0xd1, 0x65, 0xfd, 0xb0, 0x92, 0xf5, 0xf8, 0x27, 0x1e, 0xd1,
	0x84, 0x49, 0x8c, 0x7f, 0xbe, 0xfc, 0x1d, 0x23, 0x15, 0x1a, 0x20, 0x79, 0x0c, 0x7b, 0xd1, 0x42,
	0x08, 0x4c, 0x95, 0x4b, 0xe7, 0xfd, 0xaa, 0x8d, 0x8d, 0xf2, 0x0a, 0x85, 0xa1, 0x92, 0x23, 0xb5,
	0xd1, 0x15, 0x93, 0x8a, 0x8b, 0xb5, 0xcb, 0xe0, 0x2e, 0x23, 0x87, 0x24, 0xcf, 0xa0, 0x13, 0xa3,
	0x8c, 0x04, 0xcb, 0x14, 0xe3, 0x69, 0xe0, 0xef, 0x66, 0x58, 0xc5, 0x93, 0x6f, 0xa1, 0x9b, 0x09,
	0x9c, 0xb2, 0x24, 0x99, 0x88, 0x45, 0x82, 0x32, 0x68, 0xee, 0xae, 0xdd, 0xbe, 0xb3, 0x08, 0xb5,
	0x01, 0x79, 0x0e, 0xdd, 0xc8, 0x36, 0xd1, 0x79, 0xd8, 0x33, 0x09, 0x3f, 0xda, 0xe6, 0x5e, 0xf6,
	0x3a, 0xdc, 0x77, 0x26, 0xd6, 0xc5, 0x0b, 0x20, 0x73, 0xba, 0x9a, 0x30, 0x85, 0x73, 0x39, 0xc9,
	0x50, 0x4c, 0x32, 0x3a, 0xc3, 0xa0, 0x65, 0xfc, 0x6c, 0xa4, 0x72, 0x46, 0x57, 0x2f, 0x35, 0xe8,
	0x1c, 0xc5, 0x39, 0x9d, 0x61, 0x78, 0x30, 0xdf, 0x14, 0x0c, 0x9f, 0xc1, 0xc1, 0x35, 0x0c, 0x79,
	0x00, 0x7e, 0x42, 0xc5, 0xcc, 0x0e, 0xb5, 0x1f, 0xda, 0x8b, 0x96, 0xca, 0x39, 0x4d, 0x92, 0xa0,
	0x6e, 0xa5, 0xe6, 0x32, 0xfc, 0xbb, 0x96, 0xbf, 0x08, 0x57, 0x67, 0xf2, 0x08, 0xda, 0xd9, 0xe2,
	0x32, 0x61, 0xf2, 0x0a, 0xed, 0x40, 0x78, 0x61, 0x29, 0x20, 0x03, 0xe8, 0x2c, 0xd2, 0x52, 0x5f,
	0x37, 0xfa, 0xaa, 0x88, 0x9c, 0xc0, 0xa1, 0x7b, 0x1e, 0x31, 0x4e, 0x59, 0xca, 0x4c, 0x93, 0xec,
	0x48, 0xbc, 0xb7, 0x5d, 0x21, 0x4d, 0x3c, 0xec, 0x59, 0xd1, 0x69, 0x81, 0x1f, 0xfe, 0xe7, 0x01,
	0x94, 0x80, 0xad, 0xe9, 0xec, 0x81, 0xf7, 0x07, 0xae, 0x4d, 0xf4, 0x76, 0xa8, 0x8f, 0x24, 0x80,
	0xbd, 0x29, 0x4f, 0x12, 0xfe, 0x5a, 0x06, 0xde, 0xc0, 0x1b, 0xb5, 0xc3, 0xfc, 0x4a, 0x9e, 0x42,
	0x3b, 0xe2, 0x69, 0x6c, 0x79, 0x34, 0x2a, 0x15, 0xbe, 0xa5, 0xd7, 0x25, 0x9a, 0xf4, 0xa1, 0x95,
	0x09, 0xc6, 0x05, 0x53, 0xeb, 0xc0, 0x1f, 0xd4, 0x46, 0xf5, 0xb0, 0xb8, 0xeb, 0x80, 0x4b, 0x5b,
	0xb1, 0xa0, 0x69, 0x0a, 0x9a, 0x5f, 0xc9, 0x47, 0xb0, 0xef, 0x8e, 0x13, 0x45, 0x67, 0x7a, 0x3a,
	0x34, 0x9f, 0x8e, 0x93, 0xfd, 0x42, 0x67, 0x92, 0x7c, 0x0a, 0xbe, 0x69, 0x7d, 0xd0, 0x32, 0xb3,
	0x77, 0x5b, 0x5d, 0x2c, 0x88, 0xfc, 0x00, 0x3d, 0x89, 0x09, 0x46, 0x9a, 0xd3, 0x64, 0x8e, 0xea,
	0x8a, 0xc7, 0x41, 0x7b, 0x77, 0x22, 0x07, 0x85, 0xd1, 0x99, 0xb1, 0xd1, 0xbb, 0x4e, 0xad, 0x33,
	0x0c, 0xc0, 0xee, 0x3a, 0x7d, 0x26, 0x4f, 0x01, 0x22, 0x3e, 0xcf, 0x78, 0x8a, 0xa9, 0x92, 0x41,
	0x67, 0xfb, 0xe5, 0x6a, 0x22, 0x27, 0x39, 0x22, 0xac, 0x80, 0xc9, 0xd7, 0xd0, 0x59, 0xd2, 0x84,
	0xc5, 0x54, 0x87, 0x90, 0xc1, 0xfe, 0x76, 0x2a, 0xaf, 0x0a, 0x75, 0x58, 0x85, 0x0e, 0x23, 0x80,
	0x52, 0x95, 0x37, 0xb3, 0x56, 0x36, 0x33, 0x27, 0x5a, 0xaf, 0x10, 0x3d, 0x82, 0x86, 0x7e, 0x6c,
	0x6e, 0x92, 0xee, 0x4c, 0xdc, 0x00, 0x87, 0xff, 0x36, 0xa1, 0xbb, 0x41, 0x5e, 0xbb, 0x15, 0x3c,
	0x29, 0x76, 0xbd, 0x3e, 0xdf, 0x30, 0x49, 0x4f, 0x60, 0xcf, 0x3c, 0x55, 0xb3, 0xc8, 0x76, 0xae,
	0x96, 0x1c, 0x4b, 0x5e, 0xc0, 0x61, 0xcc, 0x64, 0x96, 0xd0, 0xf5, 0xe4, 0xad, 0xc6, 0xad, 0xe7,
	0xac, 0x4e, 0x8a, 0xa9, 0xfb, 0x0a, 0x5a, 0x31, 0x93, 0xf4, 0x32, 0xc1, 0xd8, 0x4c, 0xdd, 0x0e,
	0x07, 0x05, 0x98, 0x1c, 0xe5, 0x53, 0xd5, 0xdc, 0xde, 0xa5, 0x9b, 0x6d, 0x74, 0x83, 0xf5, 0x05,
	0xf8, 0x5c, 0xc4, 0x28, 0xdc, 0x02, 0xbb, 0x33, 0x8c, 0x45, 0xea, 0x2d, 0x12, 0x9b, 0xde, 0xb4,
	0x4c, 0xc5, 0xec, 0x85, 0x3c, 0x01, 0x5f, 0xaa, 0x75, 0x82, 0x41, 0xdb, 0x44, 0xfe, 0xf0, 0xd6,
	0xc8, 0xe3, 0x0b, 0x0d, 0x0b, 0x2d, 0xfa, 0xfa, 0x26, 0x87, 0xb7, 0xdc, 0xe4, 0xa7, 0x00, 0x99,
	0xe0, 0x19, 0x0a, 0xc5, 0x30, 0x9f, 0xdd, 0x4f, 0x6e, 0x0f, 0x7d, 0x5e, 0x60, 0xc3, 0x8a, 0x5d,
	0xff, 0x08, 0x7c, 0x43, 0xea, 0x86, 0x39, 0x7c, 0x00, 0xfe, 0x92, 0x26, 0x8b, 0x7c, 0x10, 0xed,
	0xa5, 0xff, 0x7f, 0x0d, 0xa0, 0xf4, 0x45, 0x3e, 0x07, 0x6f, 0xce, 0x52, 0x63, 0xb6, 0xfb, 0x1b,
	0xae, 0xa1, 0xc6, 0x82, 0xae, 0x8c, 0xd3, 0xfb, 0x58, 0xd0, 0x15, 0xf9, 0x06, 0x5a, 0x52, 0x61,
	0x76, 0xc1, 0xfe, 0xbc, 0xef, 0xcf, 0x42, 0x81, 0x27, 0xa7, 0xd0, 0x8d, 0xa9, 0xc2, 0x97, 0x69,
	0xb6, 0x50, 0x67, 0xfa, 0x6f, 0xa3, 0x71, 0x2f, 0x07, 0x9b, 0x46, 0xc3, 0xdf, 0xe0, 0xe0, 0x5a,
	0x2f, 0xf4, 0x73, 0x8a, 0x2a, 0xbf, 0x4e, 0xfa, 0x4c, 0xbe, 0x04, 0x3f, 0xa3, 0x42, 0xe5, 0x7f,
	0x4b, 0xbb, 0x82, 0x58, 0xf0, 0x77, 0xfe, 0xaf, 0x1e, 0xcd, 0xd8, 0x65, 0xd3, 0xfc, 0xa3, 0x3d,
	0x7e, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x9f, 0xd6, 0xf4, 0xe8, 0x09, 0x00, 0x00,
}