// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/certificates/certificate_request.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//CertificateRequests are generated by the CertificateRequesting Agent installed on managed clusters.
//They are used to request a signed certificate from Service Mesh Hub based on a private key
//generated by the Agent (which never leaves the managed cluster).
//
//When Service Mesh Hub creates an IssuedCertificate on a managed cluster, the local CertificateRequesting Agent
//will generate a CertificateRequest corresponding to it.
//
//Service Mesh Hub will then process the Certificate Signing Request contained in the
//CertificateRequestSpec, and write the signed SSL certificate back as a secret in the managed cluster,
//and the CertificateRequest Status to point to that secret.
type CertificateRequestSpec struct {
	// Base64-encoded data for the PKCS#10 Certificate Signing Request issued
	// by the CertificateRequesting Agent living in the managed cluster, corresponding
	// to the IssuedRequest received by the CertificateRequesting Agent.
	CertificateSigningRequest []byte   `protobuf:"bytes,1,opt,name=certificate_signing_request,json=certificateSigningRequest,proto3" json:"certificate_signing_request,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *CertificateRequestSpec) Reset()         { *m = CertificateRequestSpec{} }
func (m *CertificateRequestSpec) String() string { return proto.CompactTextString(m) }
func (*CertificateRequestSpec) ProtoMessage()    {}
func (*CertificateRequestSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbb195e6f26dce40, []int{0}
}
func (m *CertificateRequestSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequestSpec.Unmarshal(m, b)
}
func (m *CertificateRequestSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequestSpec.Marshal(b, m, deterministic)
}
func (m *CertificateRequestSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequestSpec.Merge(m, src)
}
func (m *CertificateRequestSpec) XXX_Size() int {
	return xxx_messageInfo_CertificateRequestSpec.Size(m)
}
func (m *CertificateRequestSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequestSpec.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequestSpec proto.InternalMessageInfo

func (m *CertificateRequestSpec) GetCertificateSigningRequest() []byte {
	if m != nil {
		return m.CertificateSigningRequest
	}
	return nil
}

type CertificateRequestStatus struct {
	// The most recent generation observed in the the CertificateRequest metadata.
	// if the observedGeneration does not match generation, the CA has not processed the most
	// recent version of this request.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// any errors observed which prevented the CertificateRequest from being processed.
	// if errors is empty, the request has been processed successfully
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	// the name of the secret containing the issued SSL certificate, once it has been
	// populated by the issuing CA (Service Mesh Hub).
	SecretName           string   `protobuf:"bytes,3,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateRequestStatus) Reset()         { *m = CertificateRequestStatus{} }
func (m *CertificateRequestStatus) String() string { return proto.CompactTextString(m) }
func (*CertificateRequestStatus) ProtoMessage()    {}
func (*CertificateRequestStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbb195e6f26dce40, []int{1}
}
func (m *CertificateRequestStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequestStatus.Unmarshal(m, b)
}
func (m *CertificateRequestStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequestStatus.Marshal(b, m, deterministic)
}
func (m *CertificateRequestStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequestStatus.Merge(m, src)
}
func (m *CertificateRequestStatus) XXX_Size() int {
	return xxx_messageInfo_CertificateRequestStatus.Size(m)
}
func (m *CertificateRequestStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequestStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequestStatus proto.InternalMessageInfo

func (m *CertificateRequestStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *CertificateRequestStatus) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *CertificateRequestStatus) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func init() {
	proto.RegisterType((*CertificateRequestSpec)(nil), "certificates.smh.solo.io.CertificateRequestSpec")
	proto.RegisterType((*CertificateRequestStatus)(nil), "certificates.smh.solo.io.CertificateRequestStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/certificates/certificate_request.proto", fileDescriptor_bbb195e6f26dce40)
}

var fileDescriptor_bbb195e6f26dce40 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0x95, 0xaf, 0x52, 0xa5, 0xfa, 0x63, 0x0a, 0xa8, 0x0a, 0x45, 0x6a, 0xab, 0x4e, 0x5d,
	0x1a, 0xab, 0x65, 0x67, 0x80, 0x81, 0x89, 0x0e, 0xe9, 0x82, 0x58, 0x2a, 0xc7, 0x3d, 0x75, 0xac,
	0xd6, 0x39, 0xc6, 0x3f, 0xb9, 0x02, 0x2e, 0x86, 0xeb, 0xe2, 0x4a, 0x50, 0xe2, 0x14, 0x22, 0xe8,
	0xc0, 0xe4, 0xf3, 0xf3, 0xea, 0xf1, 0x39, 0xe7, 0x25, 0x6b, 0x21, 0x5d, 0xe1, 0xf3, 0x94, 0xa3,
	0xa2, 0x16, 0x8f, 0xb8, 0x90, 0x48, 0x2d, 0x98, 0x4a, 0x72, 0x58, 0x28, 0xb0, 0xc5, 0xa2, 0xf0,
	0x39, 0x65, 0x5a, 0x52, 0x0e, 0xc6, 0xc9, 0xbd, 0xe4, 0xcc, 0x81, 0xed, 0x26, 0x5b, 0x03, 0xaf,
	0x1e, 0xac, 0x4b, 0xb5, 0x41, 0x87, 0x71, 0xd2, 0xd5, 0xa5, 0x56, 0x15, 0x69, 0x4d, 0x4d, 0x25,
	0x8e, 0xd2, 0x73, 0x3f, 0x1d, 0xaa, 0x55, 0xa0, 0xa3, 0x01, 0x5a, 0x2d, 0x9b, 0x37, 0x90, 0x46,
	0x13, 0x81, 0x28, 0x8e, 0x40, 0x9b, 0x2c, 0xf7, 0x7b, 0xea, 0xa4, 0x02, 0xeb, 0x98, 0xd2, 0xad,
	0x60, 0xfc, 0x53, 0xb0, 0xf3, 0x86, 0x39, 0x89, 0x65, 0xdb, 0xbf, 0x12, 0x28, 0xb0, 0x09, 0x69,
	0x1d, 0x85, 0xea, 0xec, 0x99, 0x0c, 0x1f, 0xbe, 0x47, 0xcc, 0xc2, 0xf0, 0x1b, 0x0d, 0x3c, 0xbe,
	0x23, 0x37, 0xdd, 0xbd, 0xac, 0x14, 0xa5, 0x2c, 0xc5, 0x69, 0xbf, 0x24, 0x9a, 0x46, 0xf3, 0x8b,
	0xec, 0xba, 0x23, 0xd9, 0x04, 0x45, 0xcb, 0x98, 0xbd, 0x45, 0x24, 0x39, 0x83, 0x76, 0xcc, 0x79,
	0x1b, 0x53, 0x72, 0x89, 0x79, 0x7d, 0x58, 0xd8, 0x6d, 0x05, 0x94, 0x10, 0x26, 0x6d, 0xa0, 0xbd,
	0x2c, 0x3e, 0xb5, 0x1e, 0xbf, 0x3a, 0xf1, 0x90, 0xf4, 0xc1, 0x18, 0x34, 0x36, 0xf9, 0x37, 0xed,
	0xcd, 0x07, 0x59, 0x9b, 0xc5, 0x13, 0xf2, 0xdf, 0x02, 0x37, 0xe0, 0xb6, 0x25, 0x53, 0x90, 0xf4,
	0xa6, 0xd1, 0x7c, 0x90, 0x91, 0x50, 0x5a, 0x33, 0x05, 0xf7, 0x9b, 0xf7, 0x8f, 0x71, 0xf4, 0xf2,
	0xf4, 0x17, 0x5f, 0xf5, 0x41, 0xfc, 0xf2, 0xb6, 0xeb, 0x19, 0xad, 0x96, 0xec, 0xa8, 0x0b, 0xb6,
	0xca, 0xfb, 0xcd, 0xf1, 0x6e, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x3b, 0xcd, 0xfb, 0x2f,
	0x02, 0x00, 0x00,
}

func (this *CertificateRequestSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CertificateRequestSpec)
	if !ok {
		that2, ok := that.(CertificateRequestSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.CertificateSigningRequest, that1.CertificateSigningRequest) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CertificateRequestStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CertificateRequestStatus)
	if !ok {
		that2, ok := that.(CertificateRequestStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if len(this.Errors) != len(that1.Errors) {
		return false
	}
	for i := range this.Errors {
		if this.Errors[i] != that1.Errors[i] {
			return false
		}
	}
	if this.SecretName != that1.SecretName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}