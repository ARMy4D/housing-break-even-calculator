// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: models/pb/calculator.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type PaymentType int32

const (
	PaymentType_RENT     PaymentType = 0
	PaymentType_MORTGAGE PaymentType = 1
	PaymentType_TOSS_UP  PaymentType = 2
)

// Enum value maps for PaymentType.
var (
	PaymentType_name = map[int32]string{
		0: "RENT",
		1: "MORTGAGE",
		2: "TOSS_UP",
	}
	PaymentType_value = map[string]int32{
		"RENT":     0,
		"MORTGAGE": 1,
		"TOSS_UP":  2,
	}
)

func (x PaymentType) Enum() *PaymentType {
	p := new(PaymentType)
	*p = x
	return p
}

func (x PaymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_models_pb_calculator_proto_enumTypes[0].Descriptor()
}

func (PaymentType) Type() protoreflect.EnumType {
	return &file_models_pb_calculator_proto_enumTypes[0]
}

func (x PaymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentType.Descriptor instead.
func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return file_models_pb_calculator_proto_rawDescGZIP(), []int{0}
}

type CalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	House    *CalculatorRequest_HouseSetting    `protobuf:"bytes,1,opt,name=house,proto3" json:"house,omitempty"`
	Rent     *CalculatorRequest_RentSetting     `protobuf:"bytes,2,opt,name=rent,proto3" json:"rent,omitempty"`
	Mortgage *CalculatorRequest_MortgageSetting `protobuf:"bytes,3,opt,name=mortgage,proto3" json:"mortgage,omitempty"`
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_pb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_pb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_models_pb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *CalculatorRequest) GetHouse() *CalculatorRequest_HouseSetting {
	if x != nil {
		return x.House
	}
	return nil
}

func (x *CalculatorRequest) GetRent() *CalculatorRequest_RentSetting {
	if x != nil {
		return x.Rent
	}
	return nil
}

func (x *CalculatorRequest) GetMortgage() *CalculatorRequest_MortgageSetting {
	if x != nil {
		return x.Mortgage
	}
	return nil
}

type CalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BreakEvenYearIntrest                   int32       `protobuf:"varint,1,opt,name=break_even_year_intrest,json=breakEvenYearIntrest,proto3" json:"break_even_year_intrest,omitempty"`
	BreakEvenYearOverall                   int32       `protobuf:"varint,2,opt,name=break_even_year_overall,json=breakEvenYearOverall,proto3" json:"break_even_year_overall,omitempty"`
	MortgageIntrestCostOverResidancePeriod float32     `protobuf:"fixed32,3,opt,name=mortgage_intrest_cost_over_residance_period,json=mortgageIntrestCostOverResidancePeriod,proto3" json:"mortgage_intrest_cost_over_residance_period,omitempty"`
	MortgageCostOverResidancePeriod        float32     `protobuf:"fixed32,4,opt,name=mortgage_cost_over_residance_period,json=mortgageCostOverResidancePeriod,proto3" json:"mortgage_cost_over_residance_period,omitempty"`
	RentCostOverResidancePeriod            float32     `protobuf:"fixed32,5,opt,name=rent_cost_over_residance_period,json=rentCostOverResidancePeriod,proto3" json:"rent_cost_over_residance_period,omitempty"`
	BestPaymentTypeIntrestOnly             PaymentType `protobuf:"varint,6,opt,name=best_payment_type_intrest_only,json=bestPaymentTypeIntrestOnly,proto3,enum=PaymentType" json:"best_payment_type_intrest_only,omitempty"`
	BestPaymentTypeOverall                 PaymentType `protobuf:"varint,7,opt,name=best_payment_type_overall,json=bestPaymentTypeOverall,proto3,enum=PaymentType" json:"best_payment_type_overall,omitempty"`
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_pb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_pb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_models_pb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatorResponse) GetBreakEvenYearIntrest() int32 {
	if x != nil {
		return x.BreakEvenYearIntrest
	}
	return 0
}

func (x *CalculatorResponse) GetBreakEvenYearOverall() int32 {
	if x != nil {
		return x.BreakEvenYearOverall
	}
	return 0
}

func (x *CalculatorResponse) GetMortgageIntrestCostOverResidancePeriod() float32 {
	if x != nil {
		return x.MortgageIntrestCostOverResidancePeriod
	}
	return 0
}

func (x *CalculatorResponse) GetMortgageCostOverResidancePeriod() float32 {
	if x != nil {
		return x.MortgageCostOverResidancePeriod
	}
	return 0
}

func (x *CalculatorResponse) GetRentCostOverResidancePeriod() float32 {
	if x != nil {
		return x.RentCostOverResidancePeriod
	}
	return 0
}

func (x *CalculatorResponse) GetBestPaymentTypeIntrestOnly() PaymentType {
	if x != nil {
		return x.BestPaymentTypeIntrestOnly
	}
	return PaymentType_RENT
}

func (x *CalculatorResponse) GetBestPaymentTypeOverall() PaymentType {
	if x != nil {
		return x.BestPaymentTypeOverall
	}
	return PaymentType_RENT
}

type CalculatorRequest_RentSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rent                   float32 `protobuf:"fixed32,1,opt,name=rent,proto3" json:"rent,omitempty"`
	YearlyRentIncreaseRate float32 `protobuf:"fixed32,2,opt,name=yearly_rent_increase_rate,json=yearlyRentIncreaseRate,proto3" json:"yearly_rent_increase_rate,omitempty"`
}

func (x *CalculatorRequest_RentSetting) Reset() {
	*x = CalculatorRequest_RentSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_pb_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest_RentSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest_RentSetting) ProtoMessage() {}

func (x *CalculatorRequest_RentSetting) ProtoReflect() protoreflect.Message {
	mi := &file_models_pb_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest_RentSetting.ProtoReflect.Descriptor instead.
func (*CalculatorRequest_RentSetting) Descriptor() ([]byte, []int) {
	return file_models_pb_calculator_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CalculatorRequest_RentSetting) GetRent() float32 {
	if x != nil {
		return x.Rent
	}
	return 0
}

func (x *CalculatorRequest_RentSetting) GetYearlyRentIncreaseRate() float32 {
	if x != nil {
		return x.YearlyRentIncreaseRate
	}
	return 0
}

type CalculatorRequest_HouseSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price                   float32 `protobuf:"fixed32,1,opt,name=price,proto3" json:"price,omitempty"`
	PropertyTaxRate         float32 `protobuf:"fixed32,2,opt,name=property_tax_rate,json=propertyTaxRate,proto3" json:"property_tax_rate,omitempty"`
	PropertyTransferTaxRate float32 `protobuf:"fixed32,3,opt,name=property_transfer_tax_rate,json=propertyTransferTaxRate,proto3" json:"property_transfer_tax_rate,omitempty"`
	YearsExpectedToReside   int32   `protobuf:"varint,4,opt,name=years_expected_to_reside,json=yearsExpectedToReside,proto3" json:"years_expected_to_reside,omitempty"`
}

func (x *CalculatorRequest_HouseSetting) Reset() {
	*x = CalculatorRequest_HouseSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_pb_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest_HouseSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest_HouseSetting) ProtoMessage() {}

func (x *CalculatorRequest_HouseSetting) ProtoReflect() protoreflect.Message {
	mi := &file_models_pb_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest_HouseSetting.ProtoReflect.Descriptor instead.
func (*CalculatorRequest_HouseSetting) Descriptor() ([]byte, []int) {
	return file_models_pb_calculator_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CalculatorRequest_HouseSetting) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CalculatorRequest_HouseSetting) GetPropertyTaxRate() float32 {
	if x != nil {
		return x.PropertyTaxRate
	}
	return 0
}

func (x *CalculatorRequest_HouseSetting) GetPropertyTransferTaxRate() float32 {
	if x != nil {
		return x.PropertyTransferTaxRate
	}
	return 0
}

func (x *CalculatorRequest_HouseSetting) GetYearsExpectedToReside() int32 {
	if x != nil {
		return x.YearsExpectedToReside
	}
	return 0
}

type CalculatorRequest_MortgageSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DownPayment float32 `protobuf:"fixed32,1,opt,name=down_payment,json=downPayment,proto3" json:"down_payment,omitempty"`
	IntrestRate float32 `protobuf:"fixed32,2,opt,name=intrest_rate,json=intrestRate,proto3" json:"intrest_rate,omitempty"`
	Term        int32   `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *CalculatorRequest_MortgageSetting) Reset() {
	*x = CalculatorRequest_MortgageSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_pb_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest_MortgageSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest_MortgageSetting) ProtoMessage() {}

func (x *CalculatorRequest_MortgageSetting) ProtoReflect() protoreflect.Message {
	mi := &file_models_pb_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest_MortgageSetting.ProtoReflect.Descriptor instead.
func (*CalculatorRequest_MortgageSetting) Descriptor() ([]byte, []int) {
	return file_models_pb_calculator_proto_rawDescGZIP(), []int{0, 2}
}

func (x *CalculatorRequest_MortgageSetting) GetDownPayment() float32 {
	if x != nil {
		return x.DownPayment
	}
	return 0
}

func (x *CalculatorRequest_MortgageSetting) GetIntrestRate() float32 {
	if x != nil {
		return x.IntrestRate
	}
	return 0
}

func (x *CalculatorRequest_MortgageSetting) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

var File_models_pb_calculator_proto protoreflect.FileDescriptor

var file_models_pb_calculator_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x04, 0x0a,
	0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a,
	0x08, 0x6d, 0x6f, 0x72, 0x74, 0x67, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x72, 0x74, 0x67, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x08, 0x6d, 0x6f, 0x72, 0x74, 0x67, 0x61, 0x67, 0x65, 0x1a, 0x5c, 0x0a,
	0x0b, 0x52, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x39, 0x0a, 0x19, 0x79, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x16, 0x79, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x52, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x52, 0x61, 0x74, 0x65, 0x1a, 0xc6, 0x01, 0x0a, 0x0c,
	0x48, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x74,
	0x61, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x12, 0x3b,
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x54, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x79,
	0x65, 0x61, 0x72, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f,
	0x5f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x79,
	0x65, 0x61, 0x72, 0x73, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x52, 0x65,
	0x73, 0x69, 0x64, 0x65, 0x1a, 0x6b, 0x0a, 0x0f, 0x4d, 0x6f, 0x72, 0x74, 0x67, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x5f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x64,
	0x6f, 0x77, 0x6e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e,
	0x74, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x69, 0x6e, 0x74, 0x72, 0x65, 0x73, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x22, 0x8e, 0x04, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x62, 0x72, 0x65, 0x61,
	0x6b, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x72,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x62, 0x72, 0x65, 0x61, 0x6b,
	0x45, 0x76, 0x65, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x49, 0x6e, 0x74, 0x72, 0x65, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x17, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x5f, 0x79, 0x65,
	0x61, 0x72, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x14, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x4f,
	0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x12, 0x5b, 0x0a, 0x2b, 0x6d, 0x6f, 0x72, 0x74, 0x67, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f,
	0x6f, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x26, 0x6d, 0x6f, 0x72,
	0x74, 0x67, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x74, 0x72, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x73, 0x74,
	0x4f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x4c, 0x0a, 0x23, 0x6d, 0x6f, 0x72, 0x74, 0x67, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x1f, 0x6d, 0x6f, 0x72, 0x74, 0x67, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x4f, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x44, 0x0a, 0x1f, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x1b, 0x72, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x73, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x69, 0x64, 0x61, 0x6e, 0x63,
	0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x50, 0x0a, 0x1e, 0x62, 0x65, 0x73, 0x74, 0x5f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x6e, 0x74,
	0x72, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x1a, 0x62,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e,
	0x74, 0x72, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x47, 0x0a, 0x19, 0x62, 0x65, 0x73,
	0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6f,
	0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x16, 0x62, 0x65, 0x73, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x61,
	0x6c, 0x6c, 0x2a, 0x32, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d,
	0x4f, 0x52, 0x54, 0x47, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x4f, 0x53,
	0x53, 0x5f, 0x55, 0x50, 0x10, 0x02, 0x32, 0x42, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x6d, 0x79, 0x34, 0x64, 0x2f,
	0x68, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x2d, 0x65, 0x76,
	0x65, 0x6e, 0x2d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_pb_calculator_proto_rawDescOnce sync.Once
	file_models_pb_calculator_proto_rawDescData = file_models_pb_calculator_proto_rawDesc
)

func file_models_pb_calculator_proto_rawDescGZIP() []byte {
	file_models_pb_calculator_proto_rawDescOnce.Do(func() {
		file_models_pb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_pb_calculator_proto_rawDescData)
	})
	return file_models_pb_calculator_proto_rawDescData
}

var file_models_pb_calculator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_pb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_models_pb_calculator_proto_goTypes = []interface{}{
	(PaymentType)(0),                          // 0: PaymentType
	(*CalculatorRequest)(nil),                 // 1: CalculatorRequest
	(*CalculatorResponse)(nil),                // 2: CalculatorResponse
	(*CalculatorRequest_RentSetting)(nil),     // 3: CalculatorRequest.RentSetting
	(*CalculatorRequest_HouseSetting)(nil),    // 4: CalculatorRequest.HouseSetting
	(*CalculatorRequest_MortgageSetting)(nil), // 5: CalculatorRequest.MortgageSetting
}
var file_models_pb_calculator_proto_depIdxs = []int32{
	4, // 0: CalculatorRequest.house:type_name -> CalculatorRequest.HouseSetting
	3, // 1: CalculatorRequest.rent:type_name -> CalculatorRequest.RentSetting
	5, // 2: CalculatorRequest.mortgage:type_name -> CalculatorRequest.MortgageSetting
	0, // 3: CalculatorResponse.best_payment_type_intrest_only:type_name -> PaymentType
	0, // 4: CalculatorResponse.best_payment_type_overall:type_name -> PaymentType
	1, // 5: Calculator.Calculate:input_type -> CalculatorRequest
	2, // 6: Calculator.Calculate:output_type -> CalculatorResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_models_pb_calculator_proto_init() }
func file_models_pb_calculator_proto_init() {
	if File_models_pb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_pb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest); i {
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
		file_models_pb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorResponse); i {
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
		file_models_pb_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest_RentSetting); i {
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
		file_models_pb_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest_HouseSetting); i {
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
		file_models_pb_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest_MortgageSetting); i {
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
			RawDescriptor: file_models_pb_calculator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_models_pb_calculator_proto_goTypes,
		DependencyIndexes: file_models_pb_calculator_proto_depIdxs,
		EnumInfos:         file_models_pb_calculator_proto_enumTypes,
		MessageInfos:      file_models_pb_calculator_proto_msgTypes,
	}.Build()
	File_models_pb_calculator_proto = out.File
	file_models_pb_calculator_proto_rawDesc = nil
	file_models_pb_calculator_proto_goTypes = nil
	file_models_pb_calculator_proto_depIdxs = nil
}
