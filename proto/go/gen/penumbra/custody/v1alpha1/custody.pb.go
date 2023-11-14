// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: penumbra/custody/v1alpha1/custody.proto

package custodyv1alpha1

import (
	v1alpha11 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/keys/v1alpha1"
	v1alpha1 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/transaction/v1alpha1"
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

type AuthorizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The transaction plan to authorize.
	Plan *v1alpha1.TransactionPlan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	// Optionally, pre-authorization data, if required by the custodian.
	//
	// Pre-authorization data is backend-specific, and backends are free to ignore it.
	//
	// Multiple `PreAuthorization` packets can be included in a single request,
	// to support multi-party pre-authorizations.
	PreAuthorizations []*PreAuthorization `protobuf:"bytes,3,rep,name=pre_authorizations,json=preAuthorizations,proto3" json:"pre_authorizations,omitempty"`
}

func (x *AuthorizeRequest) Reset() {
	*x = AuthorizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeRequest) ProtoMessage() {}

func (x *AuthorizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizeRequest) GetPlan() *v1alpha1.TransactionPlan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *AuthorizeRequest) GetPreAuthorizations() []*PreAuthorization {
	if x != nil {
		return x.PreAuthorizations
	}
	return nil
}

type AuthorizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *v1alpha1.AuthorizationData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AuthorizeResponse) Reset() {
	*x = AuthorizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeResponse) ProtoMessage() {}

func (x *AuthorizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeResponse) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizeResponse) GetData() *v1alpha1.AuthorizationData {
	if x != nil {
		return x.Data
	}
	return nil
}

// A pre-authorization packet.  This allows a custodian to delegate (partial)
// signing authority to other authorization mechanisms.  Details of how a
// custodian manages those keys are out-of-scope for the custody protocol and
// are custodian-specific.
type PreAuthorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PreAuthorization:
	//
	//	*PreAuthorization_Ed25519_
	PreAuthorization isPreAuthorization_PreAuthorization `protobuf_oneof:"pre_authorization"`
}

func (x *PreAuthorization) Reset() {
	*x = PreAuthorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreAuthorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreAuthorization) ProtoMessage() {}

func (x *PreAuthorization) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreAuthorization.ProtoReflect.Descriptor instead.
func (*PreAuthorization) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{2}
}

func (m *PreAuthorization) GetPreAuthorization() isPreAuthorization_PreAuthorization {
	if m != nil {
		return m.PreAuthorization
	}
	return nil
}

func (x *PreAuthorization) GetEd25519() *PreAuthorization_Ed25519 {
	if x, ok := x.GetPreAuthorization().(*PreAuthorization_Ed25519_); ok {
		return x.Ed25519
	}
	return nil
}

type isPreAuthorization_PreAuthorization interface {
	isPreAuthorization_PreAuthorization()
}

type PreAuthorization_Ed25519_ struct {
	Ed25519 *PreAuthorization_Ed25519 `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}

func (*PreAuthorization_Ed25519_) isPreAuthorization_PreAuthorization() {}

type ExportFullViewingKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExportFullViewingKeyRequest) Reset() {
	*x = ExportFullViewingKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportFullViewingKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportFullViewingKeyRequest) ProtoMessage() {}

func (x *ExportFullViewingKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportFullViewingKeyRequest.ProtoReflect.Descriptor instead.
func (*ExportFullViewingKeyRequest) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{3}
}

type ExportFullViewingKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The full viewing key.
	FullViewingKey *v1alpha11.FullViewingKey `protobuf:"bytes,1,opt,name=full_viewing_key,json=fullViewingKey,proto3" json:"full_viewing_key,omitempty"`
}

func (x *ExportFullViewingKeyResponse) Reset() {
	*x = ExportFullViewingKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportFullViewingKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportFullViewingKeyResponse) ProtoMessage() {}

func (x *ExportFullViewingKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportFullViewingKeyResponse.ProtoReflect.Descriptor instead.
func (*ExportFullViewingKeyResponse) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{4}
}

func (x *ExportFullViewingKeyResponse) GetFullViewingKey() *v1alpha11.FullViewingKey {
	if x != nil {
		return x.FullViewingKey
	}
	return nil
}

type ConfirmAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressIndex *v1alpha11.AddressIndex `protobuf:"bytes,1,opt,name=address_index,json=addressIndex,proto3" json:"address_index,omitempty"`
}

func (x *ConfirmAddressRequest) Reset() {
	*x = ConfirmAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmAddressRequest) ProtoMessage() {}

func (x *ConfirmAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmAddressRequest.ProtoReflect.Descriptor instead.
func (*ConfirmAddressRequest) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{5}
}

func (x *ConfirmAddressRequest) GetAddressIndex() *v1alpha11.AddressIndex {
	if x != nil {
		return x.AddressIndex
	}
	return nil
}

type ConfirmAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *v1alpha11.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ConfirmAddressResponse) Reset() {
	*x = ConfirmAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmAddressResponse) ProtoMessage() {}

func (x *ConfirmAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmAddressResponse.ProtoReflect.Descriptor instead.
func (*ConfirmAddressResponse) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{6}
}

func (x *ConfirmAddressResponse) GetAddress() *v1alpha11.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

// An Ed25519-based preauthorization, containing an Ed25519 signature over the
// `TransactionPlan`.
type PreAuthorization_Ed25519 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Ed25519 verification key used to verify the signature.
	Vk []byte `protobuf:"bytes,1,opt,name=vk,proto3" json:"vk,omitempty"`
	// The Ed25519 signature over the `TransactionPlan`.
	Sig []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *PreAuthorization_Ed25519) Reset() {
	*x = PreAuthorization_Ed25519{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreAuthorization_Ed25519) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreAuthorization_Ed25519) ProtoMessage() {}

func (x *PreAuthorization_Ed25519) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreAuthorization_Ed25519.ProtoReflect.Descriptor instead.
func (*PreAuthorization_Ed25519) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PreAuthorization_Ed25519) GetVk() []byte {
	if x != nil {
		return x.Vk
	}
	return nil
}

func (x *PreAuthorization_Ed25519) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

var File_penumbra_custody_v1alpha1_custody_proto protoreflect.FileDescriptor

var file_penumbra_custody_v1alpha1_custody_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x64, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x70, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x26, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x70, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e,
	0x12, 0x5a, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70,
	0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x70, 0x72, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x5e, 0x0a, 0x11,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa5, 0x01, 0x0a,
	0x10, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4f, 0x0a, 0x07, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x48, 0x00, 0x52, 0x07, 0x65, 0x64, 0x32, 0x35, 0x35,
	0x31, 0x39, 0x1a, 0x2b, 0x0a, 0x07, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x12, 0x0e, 0x0a,
	0x02, 0x76, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x76, 0x6b, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x42,
	0x13, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x1b, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x75,
	0x6c, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x75, 0x0a, 0x1c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x75, 0x6c,
	0x6c, 0x56, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x69, 0x65, 0x77,
	0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65,
	0x79, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x75, 0x6c, 0x6c,
	0x56, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x0e, 0x66, 0x75, 0x6c, 0x6c,
	0x56, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x22, 0x67, 0x0a, 0x15, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x65, 0x6e,
	0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x22, 0x58, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6b,
	0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x81, 0x03,
	0x0a, 0x16, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x87, 0x01, 0x0a, 0x14, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x56,
	0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x2e, 0x70, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x75, 0x6c, 0x6c,
	0x56, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x37, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x56, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x0e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x2e, 0x70,
	0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64,
	0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x8d, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62,
	0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x42, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2d, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x64, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x64, 0x79, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x50, 0x43, 0x58, 0xaa, 0x02, 0x19, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca,
	0x02, 0x19, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x64, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x50, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x3a, 0x3a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_penumbra_custody_v1alpha1_custody_proto_rawDescOnce sync.Once
	file_penumbra_custody_v1alpha1_custody_proto_rawDescData = file_penumbra_custody_v1alpha1_custody_proto_rawDesc
)

func file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP() []byte {
	file_penumbra_custody_v1alpha1_custody_proto_rawDescOnce.Do(func() {
		file_penumbra_custody_v1alpha1_custody_proto_rawDescData = protoimpl.X.CompressGZIP(file_penumbra_custody_v1alpha1_custody_proto_rawDescData)
	})
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescData
}

var file_penumbra_custody_v1alpha1_custody_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_penumbra_custody_v1alpha1_custody_proto_goTypes = []interface{}{
	(*AuthorizeRequest)(nil),             // 0: penumbra.custody.v1alpha1.AuthorizeRequest
	(*AuthorizeResponse)(nil),            // 1: penumbra.custody.v1alpha1.AuthorizeResponse
	(*PreAuthorization)(nil),             // 2: penumbra.custody.v1alpha1.PreAuthorization
	(*ExportFullViewingKeyRequest)(nil),  // 3: penumbra.custody.v1alpha1.ExportFullViewingKeyRequest
	(*ExportFullViewingKeyResponse)(nil), // 4: penumbra.custody.v1alpha1.ExportFullViewingKeyResponse
	(*ConfirmAddressRequest)(nil),        // 5: penumbra.custody.v1alpha1.ConfirmAddressRequest
	(*ConfirmAddressResponse)(nil),       // 6: penumbra.custody.v1alpha1.ConfirmAddressResponse
	(*PreAuthorization_Ed25519)(nil),     // 7: penumbra.custody.v1alpha1.PreAuthorization.Ed25519
	(*v1alpha1.TransactionPlan)(nil),     // 8: penumbra.core.transaction.v1alpha1.TransactionPlan
	(*v1alpha1.AuthorizationData)(nil),   // 9: penumbra.core.transaction.v1alpha1.AuthorizationData
	(*v1alpha11.FullViewingKey)(nil),     // 10: penumbra.core.keys.v1alpha1.FullViewingKey
	(*v1alpha11.AddressIndex)(nil),       // 11: penumbra.core.keys.v1alpha1.AddressIndex
	(*v1alpha11.Address)(nil),            // 12: penumbra.core.keys.v1alpha1.Address
}
var file_penumbra_custody_v1alpha1_custody_proto_depIdxs = []int32{
	8,  // 0: penumbra.custody.v1alpha1.AuthorizeRequest.plan:type_name -> penumbra.core.transaction.v1alpha1.TransactionPlan
	2,  // 1: penumbra.custody.v1alpha1.AuthorizeRequest.pre_authorizations:type_name -> penumbra.custody.v1alpha1.PreAuthorization
	9,  // 2: penumbra.custody.v1alpha1.AuthorizeResponse.data:type_name -> penumbra.core.transaction.v1alpha1.AuthorizationData
	7,  // 3: penumbra.custody.v1alpha1.PreAuthorization.ed25519:type_name -> penumbra.custody.v1alpha1.PreAuthorization.Ed25519
	10, // 4: penumbra.custody.v1alpha1.ExportFullViewingKeyResponse.full_viewing_key:type_name -> penumbra.core.keys.v1alpha1.FullViewingKey
	11, // 5: penumbra.custody.v1alpha1.ConfirmAddressRequest.address_index:type_name -> penumbra.core.keys.v1alpha1.AddressIndex
	12, // 6: penumbra.custody.v1alpha1.ConfirmAddressResponse.address:type_name -> penumbra.core.keys.v1alpha1.Address
	0,  // 7: penumbra.custody.v1alpha1.CustodyProtocolService.Authorize:input_type -> penumbra.custody.v1alpha1.AuthorizeRequest
	3,  // 8: penumbra.custody.v1alpha1.CustodyProtocolService.ExportFullViewingKey:input_type -> penumbra.custody.v1alpha1.ExportFullViewingKeyRequest
	5,  // 9: penumbra.custody.v1alpha1.CustodyProtocolService.ConfirmAddress:input_type -> penumbra.custody.v1alpha1.ConfirmAddressRequest
	1,  // 10: penumbra.custody.v1alpha1.CustodyProtocolService.Authorize:output_type -> penumbra.custody.v1alpha1.AuthorizeResponse
	4,  // 11: penumbra.custody.v1alpha1.CustodyProtocolService.ExportFullViewingKey:output_type -> penumbra.custody.v1alpha1.ExportFullViewingKeyResponse
	6,  // 12: penumbra.custody.v1alpha1.CustodyProtocolService.ConfirmAddress:output_type -> penumbra.custody.v1alpha1.ConfirmAddressResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_penumbra_custody_v1alpha1_custody_proto_init() }
func file_penumbra_custody_v1alpha1_custody_proto_init() {
	if File_penumbra_custody_v1alpha1_custody_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeRequest); i {
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
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeResponse); i {
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
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreAuthorization); i {
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
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportFullViewingKeyRequest); i {
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
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportFullViewingKeyResponse); i {
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
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmAddressRequest); i {
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
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmAddressResponse); i {
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
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreAuthorization_Ed25519); i {
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
	file_penumbra_custody_v1alpha1_custody_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PreAuthorization_Ed25519_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_penumbra_custody_v1alpha1_custody_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_penumbra_custody_v1alpha1_custody_proto_goTypes,
		DependencyIndexes: file_penumbra_custody_v1alpha1_custody_proto_depIdxs,
		MessageInfos:      file_penumbra_custody_v1alpha1_custody_proto_msgTypes,
	}.Build()
	File_penumbra_custody_v1alpha1_custody_proto = out.File
	file_penumbra_custody_v1alpha1_custody_proto_rawDesc = nil
	file_penumbra_custody_v1alpha1_custody_proto_goTypes = nil
	file_penumbra_custody_v1alpha1_custody_proto_depIdxs = nil
}
