// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/campaign_shared_set.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// CampaignSharedSets are used for managing the shared sets associated with a
// campaign.
type CampaignSharedSet struct {
	// The resource name of the campaign shared set.
	// Campaign shared set resource names have the form:
	//
	// `customers/{customer_id}/campaignSharedSets/{campaign_id}~{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign to which the campaign shared set belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The shared set associated with the campaign. This may be a negative keyword
	// shared set of another customer. This customer should be a manager of the
	// other customer, otherwise the campaign shared set will exist but have no
	// serving effect. Only negative keyword shared sets can be associated with
	// Shopping campaigns. Only negative placement shared sets can be associated
	// with Display mobile app campaigns.
	SharedSet *wrappers.StringValue `protobuf:"bytes,4,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
	// The status of this campaign shared set. Read only.
	Status               enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus `protobuf:"varint,2,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *CampaignSharedSet) Reset()         { *m = CampaignSharedSet{} }
func (m *CampaignSharedSet) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSet) ProtoMessage()    {}
func (*CampaignSharedSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_799fb2d64e95762f, []int{0}
}

func (m *CampaignSharedSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSet.Unmarshal(m, b)
}
func (m *CampaignSharedSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSet.Marshal(b, m, deterministic)
}
func (m *CampaignSharedSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSet.Merge(m, src)
}
func (m *CampaignSharedSet) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSet.Size(m)
}
func (m *CampaignSharedSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSet.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSet proto.InternalMessageInfo

func (m *CampaignSharedSet) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignSharedSet) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignSharedSet) GetSharedSet() *wrappers.StringValue {
	if m != nil {
		return m.SharedSet
	}
	return nil
}

func (m *CampaignSharedSet) GetStatus() enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus {
	if m != nil {
		return m.Status
	}
	return enums.CampaignSharedSetStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignSharedSet)(nil), "google.ads.googleads.v2.resources.CampaignSharedSet")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/campaign_shared_set.proto", fileDescriptor_799fb2d64e95762f)
}

var fileDescriptor_799fb2d64e95762f = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdf, 0x4a, 0xf3, 0x30,
	0x1c, 0xa5, 0xdd, 0xc7, 0xf8, 0x96, 0xef, 0x0f, 0xd8, 0x0b, 0x29, 0x63, 0xc8, 0xa6, 0x0c, 0x76,
	0x95, 0x42, 0xbd, 0x91, 0x0c, 0x84, 0x4e, 0x64, 0xe0, 0x85, 0x8c, 0x16, 0x7a, 0x21, 0x85, 0x92,
	0xad, 0x31, 0x56, 0xd6, 0xa4, 0x24, 0xe9, 0x7c, 0x00, 0x1f, 0xc3, 0x3b, 0x2f, 0x7d, 0x14, 0x1f,
	0xc5, 0xa7, 0x90, 0xb5, 0x4d, 0xbc, 0x98, 0x53, 0xef, 0x4e, 0xfb, 0x3b, 0xe7, 0xfc, 0xce, 0x49,
	0x02, 0xa6, 0x94, 0x73, 0xba, 0x26, 0x1e, 0xce, 0xa4, 0xd7, 0xc0, 0x2d, 0xda, 0xf8, 0x9e, 0x20,
	0x92, 0x57, 0x62, 0x45, 0xa4, 0xb7, 0xc2, 0x45, 0x89, 0x73, 0xca, 0x52, 0x79, 0x87, 0x05, 0xc9,
	0x52, 0x49, 0x14, 0x2c, 0x05, 0x57, 0xdc, 0x19, 0x35, 0x0a, 0x88, 0x33, 0x09, 0x8d, 0x18, 0x6e,
	0x7c, 0x68, 0xc4, 0xfd, 0xf3, 0x7d, 0xfe, 0x84, 0x55, 0xc5, 0xa7, 0xde, 0xa9, 0x54, 0x58, 0x55,
	0xb2, 0x59, 0xd1, 0x3f, 0x6a, 0xf5, 0xf5, 0xd7, 0xb2, 0xba, 0xf5, 0x1e, 0x04, 0x2e, 0x4b, 0x22,
	0xf4, 0x7c, 0xa0, 0xfd, 0xcb, 0xdc, 0xc3, 0x8c, 0x71, 0x85, 0x55, 0xce, 0x59, 0x3b, 0x3d, 0x7e,
	0xb2, 0xc1, 0xc1, 0x45, 0xbb, 0x22, 0xaa, 0x37, 0x44, 0x44, 0x39, 0x27, 0xe0, 0x9f, 0x0e, 0x98,
	0x32, 0x5c, 0x10, 0xd7, 0x1a, 0x5a, 0x93, 0x5e, 0xf8, 0x57, 0xff, 0xbc, 0xc6, 0x05, 0x71, 0xce,
	0xc0, 0x6f, 0x1d, 0xce, 0xed, 0x0c, 0xad, 0xc9, 0x1f, 0x7f, 0xd0, 0x76, 0x84, 0x3a, 0x0b, 0x8c,
	0x94, 0xc8, 0x19, 0x8d, 0xf1, 0xba, 0x22, 0xa1, 0x61, 0x3b, 0x53, 0x00, 0x3e, 0xda, 0xb8, 0xbf,
	0x7e, 0xa0, 0xed, 0x49, 0x93, 0xed, 0x1e, 0x74, 0x9b, 0xfe, 0xae, 0x3d, 0xb4, 0x26, 0xff, 0xfd,
	0x10, 0xee, 0x3b, 0xe3, 0xfa, 0x00, 0xe1, 0x4e, 0xbb, 0xa8, 0x56, 0x5f, 0xb2, 0xaa, 0xd8, 0x37,
	0x0b, 0xdb, 0x0d, 0xb3, 0x47, 0x1b, 0x8c, 0x57, 0xbc, 0x80, 0xdf, 0xde, 0xe2, 0xec, 0x70, 0xc7,
	0x6a, 0xb1, 0xed, 0xb1, 0xb0, 0x6e, 0xae, 0x5a, 0x31, 0xe5, 0x6b, 0xcc, 0x28, 0xe4, 0x82, 0x7a,
	0x94, 0xb0, 0xba, 0xa5, 0xbe, 0xef, 0x32, 0x97, 0x5f, 0x3c, 0xaf, 0xa9, 0x41, 0xcf, 0x76, 0x67,
	0x1e, 0x04, 0x2f, 0xf6, 0x68, 0xde, 0x58, 0x06, 0x99, 0x84, 0x0d, 0xdc, 0xa2, 0xd8, 0x87, 0xa1,
	0x66, 0xbe, 0x6a, 0x4e, 0x12, 0x64, 0x32, 0x31, 0x9c, 0x24, 0xf6, 0x13, 0xc3, 0x79, 0xb3, 0xc7,
	0xcd, 0x00, 0xa1, 0x20, 0x93, 0x08, 0x19, 0x16, 0x42, 0xb1, 0x8f, 0x90, 0xe1, 0x2d, 0xbb, 0x75,
	0xd8, 0xd3, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xa6, 0x90, 0x5c, 0x0a, 0x03, 0x00, 0x00,
}
