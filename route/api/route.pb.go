// Code generated by protoc-gen-go. DO NOT EDIT.
// source: route/api/route.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	route/api/route.proto

It has these top-level messages:
	Route
	Path
	StaticPath
	BGPPath
	ASPathSegment
	LargeCommunity
	UnknownAttribute
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bio_net "github.com/bio-routing/bio-rd/net/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Path_Type int32

const (
	Path_Static Path_Type = 0
	Path_BGP    Path_Type = 1
)

var Path_Type_name = map[int32]string{
	0: "Static",
	1: "BGP",
}
var Path_Type_value = map[string]int32{
	"Static": 0,
	"BGP":    1,
}

func (x Path_Type) String() string {
	return proto.EnumName(Path_Type_name, int32(x))
}
func (Path_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Route struct {
	Pfx   *bio_net.Prefix `protobuf:"bytes,1,opt,name=pfx" json:"pfx,omitempty"`
	Paths []*Path         `protobuf:"bytes,2,rep,name=paths" json:"paths,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Route) GetPfx() *bio_net.Prefix {
	if m != nil {
		return m.Pfx
	}
	return nil
}

func (m *Route) GetPaths() []*Path {
	if m != nil {
		return m.Paths
	}
	return nil
}

type Path struct {
	Type       Path_Type   `protobuf:"varint,1,opt,name=type,enum=bio.route.Path_Type" json:"type,omitempty"`
	StaticPath *StaticPath `protobuf:"bytes,2,opt,name=static_path,json=staticPath" json:"static_path,omitempty"`
	BGPPath    *BGPPath    `protobuf:"bytes,3,opt,name=BGP_path,json=BGPPath" json:"BGP_path,omitempty"`
}

func (m *Path) Reset()                    { *m = Path{} }
func (m *Path) String() string            { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()               {}
func (*Path) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Path) GetType() Path_Type {
	if m != nil {
		return m.Type
	}
	return Path_Static
}

func (m *Path) GetStaticPath() *StaticPath {
	if m != nil {
		return m.StaticPath
	}
	return nil
}

func (m *Path) GetBGPPath() *BGPPath {
	if m != nil {
		return m.BGPPath
	}
	return nil
}

type StaticPath struct {
	NextHop *bio_net.IP `protobuf:"bytes,1,opt,name=next_hop,json=nextHop" json:"next_hop,omitempty"`
}

func (m *StaticPath) Reset()                    { *m = StaticPath{} }
func (m *StaticPath) String() string            { return proto.CompactTextString(m) }
func (*StaticPath) ProtoMessage()               {}
func (*StaticPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StaticPath) GetNextHop() *bio_net.IP {
	if m != nil {
		return m.NextHop
	}
	return nil
}

type BGPPath struct {
	PathIdentifier    uint32              `protobuf:"varint,1,opt,name=path_identifier,json=pathIdentifier" json:"path_identifier,omitempty"`
	NextHop           *bio_net.IP         `protobuf:"bytes,2,opt,name=next_hop,json=nextHop" json:"next_hop,omitempty"`
	LocalPref         uint32              `protobuf:"varint,3,opt,name=local_pref,json=localPref" json:"local_pref,omitempty"`
	ASPath            []*ASPathSegment    `protobuf:"bytes,4,rep,name=AS_path,json=ASPath" json:"AS_path,omitempty"`
	Origin            uint32              `protobuf:"varint,5,opt,name=origin" json:"origin,omitempty"`
	MED               uint32              `protobuf:"varint,6,opt,name=MED" json:"MED,omitempty"`
	EBGP              bool                `protobuf:"varint,7,opt,name=EBGP" json:"EBGP,omitempty"`
	BGPIdentifier     uint32              `protobuf:"varint,8,opt,name=BGP_identifier,json=BGPIdentifier" json:"BGP_identifier,omitempty"`
	Source            *bio_net.IP         `protobuf:"bytes,9,opt,name=source" json:"source,omitempty"`
	Communities       []uint32            `protobuf:"varint,10,rep,packed,name=communities" json:"communities,omitempty"`
	LargeCommunities  []*LargeCommunity   `protobuf:"bytes,11,rep,name=large_communities,json=largeCommunities" json:"large_communities,omitempty"`
	OriginatorId      uint32              `protobuf:"varint,12,opt,name=originator_id,json=originatorId" json:"originator_id,omitempty"`
	ClusterList       []uint32            `protobuf:"varint,13,rep,packed,name=cluster_list,json=clusterList" json:"cluster_list,omitempty"`
	UnknownAttributes []*UnknownAttribute `protobuf:"bytes,14,rep,name=unknown_attributes,json=unknownAttributes" json:"unknown_attributes,omitempty"`
}

func (m *BGPPath) Reset()                    { *m = BGPPath{} }
func (m *BGPPath) String() string            { return proto.CompactTextString(m) }
func (*BGPPath) ProtoMessage()               {}
func (*BGPPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BGPPath) GetPathIdentifier() uint32 {
	if m != nil {
		return m.PathIdentifier
	}
	return 0
}

func (m *BGPPath) GetNextHop() *bio_net.IP {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *BGPPath) GetLocalPref() uint32 {
	if m != nil {
		return m.LocalPref
	}
	return 0
}

func (m *BGPPath) GetASPath() []*ASPathSegment {
	if m != nil {
		return m.ASPath
	}
	return nil
}

func (m *BGPPath) GetOrigin() uint32 {
	if m != nil {
		return m.Origin
	}
	return 0
}

func (m *BGPPath) GetMED() uint32 {
	if m != nil {
		return m.MED
	}
	return 0
}

func (m *BGPPath) GetEBGP() bool {
	if m != nil {
		return m.EBGP
	}
	return false
}

func (m *BGPPath) GetBGPIdentifier() uint32 {
	if m != nil {
		return m.BGPIdentifier
	}
	return 0
}

func (m *BGPPath) GetSource() *bio_net.IP {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *BGPPath) GetCommunities() []uint32 {
	if m != nil {
		return m.Communities
	}
	return nil
}

func (m *BGPPath) GetLargeCommunities() []*LargeCommunity {
	if m != nil {
		return m.LargeCommunities
	}
	return nil
}

func (m *BGPPath) GetOriginatorId() uint32 {
	if m != nil {
		return m.OriginatorId
	}
	return 0
}

func (m *BGPPath) GetClusterList() []uint32 {
	if m != nil {
		return m.ClusterList
	}
	return nil
}

func (m *BGPPath) GetUnknownAttributes() []*UnknownAttribute {
	if m != nil {
		return m.UnknownAttributes
	}
	return nil
}

type ASPathSegment struct {
	ASSequence bool     `protobuf:"varint,1,opt,name=AS_sequence,json=ASSequence" json:"AS_sequence,omitempty"`
	ASNs       []uint32 `protobuf:"varint,2,rep,packed,name=ASNs" json:"ASNs,omitempty"`
}

func (m *ASPathSegment) Reset()                    { *m = ASPathSegment{} }
func (m *ASPathSegment) String() string            { return proto.CompactTextString(m) }
func (*ASPathSegment) ProtoMessage()               {}
func (*ASPathSegment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ASPathSegment) GetASSequence() bool {
	if m != nil {
		return m.ASSequence
	}
	return false
}

func (m *ASPathSegment) GetASNs() []uint32 {
	if m != nil {
		return m.ASNs
	}
	return nil
}

type LargeCommunity struct {
	GlobalAdministrator uint32 `protobuf:"varint,1,opt,name=global_administrator,json=globalAdministrator" json:"global_administrator,omitempty"`
	DataPart1           uint32 `protobuf:"varint,2,opt,name=data_part1,json=dataPart1" json:"data_part1,omitempty"`
	DataPart2           uint32 `protobuf:"varint,3,opt,name=data_part2,json=dataPart2" json:"data_part2,omitempty"`
}

func (m *LargeCommunity) Reset()                    { *m = LargeCommunity{} }
func (m *LargeCommunity) String() string            { return proto.CompactTextString(m) }
func (*LargeCommunity) ProtoMessage()               {}
func (*LargeCommunity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LargeCommunity) GetGlobalAdministrator() uint32 {
	if m != nil {
		return m.GlobalAdministrator
	}
	return 0
}

func (m *LargeCommunity) GetDataPart1() uint32 {
	if m != nil {
		return m.DataPart1
	}
	return 0
}

func (m *LargeCommunity) GetDataPart2() uint32 {
	if m != nil {
		return m.DataPart2
	}
	return 0
}

type UnknownAttribute struct {
	Optional   bool   `protobuf:"varint,1,opt,name=optional" json:"optional,omitempty"`
	Transitive bool   `protobuf:"varint,2,opt,name=transitive" json:"transitive,omitempty"`
	Partial    bool   `protobuf:"varint,3,opt,name=partial" json:"partial,omitempty"`
	TypeCode   uint32 `protobuf:"varint,4,opt,name=type_code,json=typeCode" json:"type_code,omitempty"`
	Value      []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *UnknownAttribute) Reset()                    { *m = UnknownAttribute{} }
func (m *UnknownAttribute) String() string            { return proto.CompactTextString(m) }
func (*UnknownAttribute) ProtoMessage()               {}
func (*UnknownAttribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UnknownAttribute) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *UnknownAttribute) GetTransitive() bool {
	if m != nil {
		return m.Transitive
	}
	return false
}

func (m *UnknownAttribute) GetPartial() bool {
	if m != nil {
		return m.Partial
	}
	return false
}

func (m *UnknownAttribute) GetTypeCode() uint32 {
	if m != nil {
		return m.TypeCode
	}
	return 0
}

func (m *UnknownAttribute) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Route)(nil), "bio.route.Route")
	proto.RegisterType((*Path)(nil), "bio.route.Path")
	proto.RegisterType((*StaticPath)(nil), "bio.route.StaticPath")
	proto.RegisterType((*BGPPath)(nil), "bio.route.BGPPath")
	proto.RegisterType((*ASPathSegment)(nil), "bio.route.ASPathSegment")
	proto.RegisterType((*LargeCommunity)(nil), "bio.route.LargeCommunity")
	proto.RegisterType((*UnknownAttribute)(nil), "bio.route.UnknownAttribute")
	proto.RegisterEnum("bio.route.Path_Type", Path_Type_name, Path_Type_value)
}

func init() { proto.RegisterFile("route/api/route.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5d, 0x53, 0xd3, 0x4c,
	0x14, 0x7e, 0x4b, 0xbf, 0xd2, 0xd3, 0xa6, 0x94, 0x7d, 0xc1, 0x89, 0x30, 0x6a, 0x09, 0x83, 0xd4,
	0x0b, 0xda, 0xa1, 0x3a, 0xde, 0xb7, 0x80, 0x15, 0x07, 0x9d, 0xb8, 0xd1, 0x1b, 0x6f, 0x32, 0xdb,
	0x64, 0xdb, 0xee, 0x98, 0x66, 0x63, 0xb2, 0x41, 0xb8, 0xf4, 0x77, 0xf8, 0x37, 0xf8, 0x81, 0xce,
	0xee, 0xa6, 0x25, 0xc5, 0xf1, 0xee, 0x9c, 0xe7, 0x39, 0x5f, 0x4f, 0xce, 0xc9, 0xc2, 0x5e, 0xc2,
	0x33, 0x41, 0x07, 0x24, 0x66, 0x03, 0x65, 0xf5, 0xe3, 0x84, 0x0b, 0x8e, 0x1a, 0x53, 0xc6, 0xfb,
	0x0a, 0xd8, 0xdf, 0x89, 0xa8, 0x50, 0x7c, 0x44, 0x85, 0x66, 0xed, 0xcf, 0x50, 0xc5, 0x92, 0x43,
	0x87, 0x50, 0x8e, 0x67, 0xb7, 0x56, 0xa9, 0x5b, 0xea, 0x35, 0x87, 0xdb, 0x7d, 0x99, 0x24, 0xa3,
	0x9c, 0x84, 0xce, 0xd8, 0x2d, 0x96, 0x1c, 0x3a, 0x86, 0x6a, 0x4c, 0xc4, 0x22, 0xb5, 0xb6, 0xba,
	0xe5, 0x75, 0x90, 0x6e, 0xe5, 0x10, 0xb1, 0xc0, 0x9a, 0xb5, 0xef, 0x4b, 0x50, 0x91, 0x3e, 0xea,
	0x41, 0x45, 0xdc, 0xc5, 0x54, 0xd5, 0x6c, 0x0f, 0x77, 0x1f, 0x85, 0xf7, 0xbf, 0xdc, 0xc5, 0x14,
	0xab, 0x08, 0xf4, 0x16, 0x9a, 0xa9, 0x20, 0x82, 0xf9, 0x9e, 0x2c, 0x61, 0x6d, 0xa9, 0x21, 0xf6,
	0x0a, 0x09, 0xae, 0x62, 0x55, 0x17, 0x48, 0xd7, 0x36, 0x3a, 0x05, 0x63, 0x3c, 0x71, 0x74, 0x52,
	0x59, 0x25, 0xa1, 0x42, 0xd2, 0x78, 0xe2, 0xa8, 0x8c, 0x7a, 0x6e, 0xd8, 0x07, 0x50, 0x91, 0x4d,
	0x11, 0x40, 0x4d, 0x17, 0xec, 0xfc, 0x87, 0xea, 0x50, 0x1e, 0x4f, 0x9c, 0x4e, 0xc9, 0x7e, 0x03,
	0xf0, 0xd0, 0x05, 0xbd, 0x04, 0x23, 0xa2, 0xb7, 0xc2, 0x5b, 0xf0, 0x38, 0xff, 0x26, 0xcd, 0xf5,
	0x37, 0xb9, 0x72, 0x70, 0x5d, 0x92, 0xef, 0x79, 0x6c, 0xdf, 0x57, 0x60, 0x55, 0x1e, 0x9d, 0xc0,
	0xb6, 0x9c, 0xc4, 0x63, 0x01, 0x8d, 0x04, 0x9b, 0x31, 0x9a, 0xa8, 0x54, 0x13, 0xb7, 0x25, 0x7c,
	0xb5, 0x46, 0x37, 0x8a, 0x6f, 0xfd, 0xbb, 0x38, 0x7a, 0x06, 0x10, 0x72, 0x9f, 0x84, 0x5e, 0x9c,
	0xd0, 0x99, 0x12, 0x68, 0xe2, 0x86, 0x42, 0xe4, 0x5a, 0xd0, 0x19, 0xd4, 0x47, 0xae, 0x16, 0x5f,
	0x51, 0x1b, 0xb1, 0x0a, 0xe2, 0x47, 0xae, 0x9c, 0xc9, 0xa5, 0xf3, 0x25, 0x8d, 0x04, 0xae, 0x69,
	0x17, 0x3d, 0x81, 0x1a, 0x4f, 0xd8, 0x9c, 0x45, 0x56, 0x55, 0x55, 0xcb, 0x3d, 0xd4, 0x81, 0xf2,
	0xc7, 0xcb, 0x0b, 0xab, 0xa6, 0x40, 0x69, 0x22, 0x04, 0x95, 0xcb, 0xf1, 0xc4, 0xb1, 0xea, 0xdd,
	0x52, 0xcf, 0xc0, 0xca, 0x46, 0xc7, 0xd0, 0x96, 0x9f, 0xbb, 0xa0, 0xcf, 0x50, 0x09, 0xe6, 0x78,
	0xe2, 0x14, 0xe4, 0x1d, 0x41, 0x2d, 0xe5, 0x59, 0xe2, 0x53, 0xab, 0xf1, 0xb7, 0xb8, 0x9c, 0x42,
	0x5d, 0x68, 0xfa, 0x7c, 0xb9, 0xcc, 0x22, 0x26, 0x18, 0x4d, 0x2d, 0xe8, 0x96, 0x7b, 0x26, 0x2e,
	0x42, 0xe8, 0x1d, 0xec, 0x84, 0x24, 0x99, 0x53, 0xaf, 0x18, 0xd7, 0x54, 0x42, 0x9f, 0x16, 0x84,
	0x5e, 0xcb, 0x98, 0xf3, 0x3c, 0xe4, 0x0e, 0x77, 0xc2, 0xa2, 0x2f, 0xeb, 0x1c, 0x81, 0xa9, 0x55,
	0x12, 0xc1, 0x13, 0x8f, 0x05, 0x56, 0x4b, 0x0d, 0xdd, 0x7a, 0x00, 0xaf, 0x02, 0x74, 0x08, 0x2d,
	0x3f, 0xcc, 0x52, 0x41, 0x13, 0x2f, 0x64, 0xa9, 0xb0, 0xcc, 0x7c, 0x1e, 0x8d, 0x5d, 0xb3, 0x54,
	0xa0, 0x0f, 0x80, 0xb2, 0xe8, 0x7b, 0xc4, 0x7f, 0x46, 0x1e, 0x11, 0x22, 0x61, 0xd3, 0x4c, 0xd0,
	0xd4, 0x6a, 0xab, 0x81, 0x0e, 0x0a, 0x03, 0x7d, 0xd5, 0x41, 0xa3, 0x55, 0x0c, 0xde, 0xc9, 0x1e,
	0x21, 0xa9, 0x7d, 0x01, 0xe6, 0xc6, 0x82, 0xd0, 0x0b, 0x68, 0x8e, 0x5c, 0x2f, 0xa5, 0x3f, 0x32,
	0x1a, 0xf9, 0xfa, 0x97, 0x31, 0x30, 0x8c, 0x5c, 0x37, 0x47, 0xe4, 0x3e, 0x46, 0xee, 0x27, 0xfd,
	0xef, 0x99, 0x58, 0xd9, 0xf6, 0xaf, 0x12, 0xb4, 0x37, 0xe5, 0xa3, 0x33, 0xd8, 0x9d, 0x87, 0x7c,
	0x4a, 0x42, 0x8f, 0x04, 0x4b, 0x16, 0xb1, 0x54, 0x24, 0x52, 0x61, 0x7e, 0x88, 0xff, 0x6b, 0x6e,
	0x54, 0xa4, 0xe4, 0x95, 0x05, 0x44, 0x10, 0x2f, 0x26, 0x89, 0x38, 0x53, 0xf7, 0x68, 0xe2, 0x86,
	0x44, 0x1c, 0x09, 0x6c, 0xd0, 0xc3, 0xd5, 0x11, 0xae, 0xe8, 0xa1, 0xfd, 0xbb, 0x04, 0x9d, 0xc7,
	0x8a, 0xd1, 0x3e, 0x18, 0x3c, 0x16, 0x8c, 0x47, 0x24, 0xcc, 0xa5, 0xac, 0x7d, 0xf4, 0x1c, 0x40,
	0x24, 0x24, 0x4a, 0x99, 0x60, 0x37, 0x54, 0xb5, 0x33, 0x70, 0x01, 0x41, 0x16, 0xd4, 0x65, 0x2b,
	0x46, 0x42, 0xd5, 0xcc, 0xc0, 0x2b, 0x17, 0x1d, 0x40, 0x43, 0xbe, 0x16, 0x9e, 0xcf, 0x03, 0x6a,
	0x55, 0xd4, 0x20, 0x86, 0x04, 0xce, 0x79, 0x40, 0xd1, 0x2e, 0x54, 0x6f, 0x48, 0x98, 0x51, 0x75,
	0xd8, 0x2d, 0xac, 0x9d, 0xf1, 0xab, 0x6f, 0x27, 0x73, 0x26, 0x16, 0xd9, 0xb4, 0xef, 0xf3, 0xe5,
	0x60, 0xca, 0xf8, 0xa9, 0xdc, 0x11, 0x8b, 0xe6, 0xda, 0x0e, 0x06, 0xeb, 0x37, 0x73, 0x5a, 0x53,
	0x0f, 0xe2, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xfd, 0xef, 0x89, 0x47, 0x05, 0x00,
	0x00,
}
