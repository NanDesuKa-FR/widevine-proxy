// Code generated by protoc-gen-go. DO NOT EDIT.
// source: WidevineCencHeader.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	WidevineCencHeader.proto

It has these top-level messages:
	WidevineCencHeader
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type WidevineCencHeader_Algorithm int32

const (
	WidevineCencHeader_UNENCRYPTED WidevineCencHeader_Algorithm = 0
	WidevineCencHeader_AESCTR      WidevineCencHeader_Algorithm = 1
)

var WidevineCencHeader_Algorithm_name = map[int32]string{
	0: "UNENCRYPTED",
	1: "AESCTR",
}
var WidevineCencHeader_Algorithm_value = map[string]int32{
	"UNENCRYPTED": 0,
	"AESCTR":      1,
}

func (x WidevineCencHeader_Algorithm) Enum() *WidevineCencHeader_Algorithm {
	p := new(WidevineCencHeader_Algorithm)
	*p = x
	return p
}
func (x WidevineCencHeader_Algorithm) String() string {
	return proto1.EnumName(WidevineCencHeader_Algorithm_name, int32(x))
}
func (x *WidevineCencHeader_Algorithm) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(WidevineCencHeader_Algorithm_value, data, "WidevineCencHeader_Algorithm")
	if err != nil {
		return err
	}
	*x = WidevineCencHeader_Algorithm(value)
	return nil
}
func (WidevineCencHeader_Algorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type WidevineCencHeader struct {
	Algorithm *WidevineCencHeader_Algorithm `protobuf:"varint,1,opt,name=algorithm,enum=proto.WidevineCencHeader_Algorithm" json:"algorithm,omitempty"`
	KeyId     [][]byte                      `protobuf:"bytes,2,rep,name=key_id,json=keyId" json:"key_id,omitempty"`
	// Content provider name.
	Provider *string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	// A content identifier, specificed by content provider.
	ContentId []byte `protobuf:"bytes,4,opt,name=content_id,json=contentId" json:"content_id,omitempty"`
	// Track type. Acceptable values are SD, HD, and AUDIO. Used to
	// differentiate content keys used by an asset.
	TrackTypeDeprecated *string `protobuf:"bytes,5,opt,name=track_type_deprecated,json=trackTypeDeprecated" json:"track_type_deprecated,omitempty"`
	// The name of a registered policy to be used for this asset.
	Policy *string `protobuf:"bytes,6,opt,name=policy" json:"policy,omitempty"`
	// Crypto period index, for media using key rotation.
	CryptoPeriodIndex *uint32 `protobuf:"varint,7,opt,name=crypto_period_index,json=cryptoPeriodIndex" json:"crypto_period_index,omitempty"`
	// Optional protected context for group content. The grouped_license is a
	// serialized SignedMessage.
	GroupedLicense []byte `protobuf:"bytes,8,opt,name=grouped_license,json=groupedLicense" json:"grouped_license,omitempty"`
	// Protection scheme identifying the encryption algorithm.
	// Represented as one of the following 4CC values:
	// 'cenc' (AES??CTR), 'cbc1' (AES??CBC),
	// 'cens' (AES??CTR subsample), 'cbcs' (AES??CBC subsample).
	ProtectionScheme *uint32 `protobuf:"varint,9,opt,name=protection_scheme,json=protectionScheme" json:"protection_scheme,omitempty"`
	// Optional. For media using key rotation, this represents the duration
	// of each crypto period in seconds.
	CryptoPeriodSeconds *uint32 `protobuf:"varint,10,opt,name=crypto_period_seconds,json=cryptoPeriodSeconds" json:"crypto_period_seconds,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *WidevineCencHeader) Reset()                    { *m = WidevineCencHeader{} }
func (m *WidevineCencHeader) String() string            { return proto1.CompactTextString(m) }
func (*WidevineCencHeader) ProtoMessage()               {}
func (*WidevineCencHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WidevineCencHeader) GetAlgorithm() WidevineCencHeader_Algorithm {
	if m != nil && m.Algorithm != nil {
		return *m.Algorithm
	}
	return WidevineCencHeader_UNENCRYPTED
}

func (m *WidevineCencHeader) GetKeyId() [][]byte {
	if m != nil {
		return m.KeyId
	}
	return nil
}

func (m *WidevineCencHeader) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

func (m *WidevineCencHeader) GetContentId() []byte {
	if m != nil {
		return m.ContentId
	}
	return nil
}

func (m *WidevineCencHeader) GetTrackTypeDeprecated() string {
	if m != nil && m.TrackTypeDeprecated != nil {
		return *m.TrackTypeDeprecated
	}
	return ""
}

func (m *WidevineCencHeader) GetPolicy() string {
	if m != nil && m.Policy != nil {
		return *m.Policy
	}
	return ""
}

func (m *WidevineCencHeader) GetCryptoPeriodIndex() uint32 {
	if m != nil && m.CryptoPeriodIndex != nil {
		return *m.CryptoPeriodIndex
	}
	return 0
}

func (m *WidevineCencHeader) GetGroupedLicense() []byte {
	if m != nil {
		return m.GroupedLicense
	}
	return nil
}

func (m *WidevineCencHeader) GetProtectionScheme() uint32 {
	if m != nil && m.ProtectionScheme != nil {
		return *m.ProtectionScheme
	}
	return 0
}

func (m *WidevineCencHeader) GetCryptoPeriodSeconds() uint32 {
	if m != nil && m.CryptoPeriodSeconds != nil {
		return *m.CryptoPeriodSeconds
	}
	return 0
}

func init() {
	proto1.RegisterType((*WidevineCencHeader)(nil), "proto.WidevineCencHeader")
	proto1.RegisterEnum("proto.WidevineCencHeader_Algorithm", WidevineCencHeader_Algorithm_name, WidevineCencHeader_Algorithm_value)
}

func init() { proto1.RegisterFile("WidevineCencHeader.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x5f, 0x4b, 0x02, 0x41,
	0x14, 0xc5, 0xdb, 0xcc, 0xcd, 0xbd, 0x99, 0x7f, 0x46, 0x8c, 0x21, 0x08, 0x16, 0x7b, 0x68, 0x21,
	0xf0, 0xc1, 0x6f, 0x20, 0x2a, 0x24, 0x84, 0xc8, 0x6a, 0x44, 0x4f, 0xcb, 0x32, 0x73, 0xd1, 0x41,
	0x9d, 0x19, 0x66, 0x27, 0x69, 0x3f, 0x6a, 0xdf, 0x26, 0x76, 0xd6, 0xb4, 0xa8, 0xa7, 0xe1, 0x9e,
	0xdf, 0x3d, 0x87, 0x33, 0x17, 0xe8, 0xab, 0xe0, 0xb8, 0x17, 0x12, 0x47, 0x28, 0xd9, 0x13, 0xa6,
	0x1c, 0x4d, 0x5f, 0x1b, 0x65, 0x15, 0xa9, 0xba, 0xa7, 0xf7, 0x59, 0x01, 0xf2, 0x77, 0x87, 0x0c,
	0x21, 0x48, 0xb7, 0x2b, 0x65, 0x84, 0x5d, 0xef, 0xa8, 0x17, 0x7a, 0x51, 0x63, 0x70, 0x5f, 0x1a,
	0xfb, 0xff, 0x24, 0x0e, 0xbf, 0x57, 0xe3, 0x93, 0x8b, 0x74, 0xc1, 0xdf, 0x60, 0x9e, 0x08, 0x4e,
	0xcf, 0xc3, 0x4a, 0x54, 0x8f, 0xab, 0x1b, 0xcc, 0xa7, 0x9c, 0xdc, 0x42, 0x4d, 0x1b, 0xb5, 0x17,
	0x1c, 0x0d, 0xad, 0x84, 0x5e, 0x14, 0xc4, 0xc7, 0x99, 0xdc, 0x01, 0x30, 0x25, 0x2d, 0x4a, 0x5b,
	0xd8, 0x2e, 0x42, 0x2f, 0xaa, 0xc7, 0xc1, 0x41, 0x99, 0x72, 0x32, 0x80, 0xae, 0x35, 0x29, 0xdb,
	0x24, 0x36, 0xd7, 0x98, 0x70, 0xd4, 0x06, 0x59, 0x6a, 0x91, 0xd3, 0xaa, 0xcb, 0xe9, 0x38, 0xb8,
	0xcc, 0x35, 0x8e, 0x8f, 0x88, 0xdc, 0x80, 0xaf, 0xd5, 0x56, 0xb0, 0x9c, 0xfa, 0x6e, 0xe9, 0x30,
	0x91, 0x3e, 0x74, 0x98, 0xc9, 0xb5, 0x55, 0x89, 0x46, 0x23, 0x14, 0x4f, 0x84, 0xe4, 0xf8, 0x41,
	0x2f, 0x43, 0x2f, 0xba, 0x8e, 0xdb, 0x25, 0x9a, 0x3b, 0x32, 0x2d, 0x00, 0x79, 0x80, 0xe6, 0xca,
	0xa8, 0x77, 0x8d, 0x3c, 0xd9, 0x0a, 0x86, 0x32, 0x43, 0x5a, 0x73, 0xfd, 0x1a, 0x07, 0xf9, 0xb9,
	0x54, 0xc9, 0x23, 0xb4, 0x8b, 0x3b, 0x21, 0xb3, 0x42, 0xc9, 0x24, 0x63, 0x6b, 0xdc, 0x21, 0x0d,
	0x5c, 0x6c, 0xeb, 0x04, 0x16, 0x4e, 0x2f, 0x7e, 0xf4, 0xbb, 0x45, 0x86, 0x4c, 0x49, 0x9e, 0x51,
	0x70, 0x86, 0xce, 0xcf, 0x1e, 0x8b, 0x12, 0xf5, 0x22, 0x08, 0x8e, 0xf7, 0x26, 0x4d, 0xb8, 0x7a,
	0x99, 0x4d, 0x66, 0xa3, 0xf8, 0x6d, 0xbe, 0x9c, 0x8c, 0x5b, 0x67, 0x04, 0xc0, 0x1f, 0x4e, 0x16,
	0xa3, 0x65, 0xdc, 0xf2, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x20, 0x71, 0x0a, 0xe2, 0xfd, 0x01,
	0x00, 0x00,
}
