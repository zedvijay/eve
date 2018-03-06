// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DsType int32

const (
	DsType_DsUnknown DsType = 0
	DsType_DsHttp    DsType = 1
	DsType_DsHttps   DsType = 2
	DsType_DsS3      DsType = 3
	DsType_DsSFTP    DsType = 4
)

var DsType_name = map[int32]string{
	0: "DsUnknown",
	1: "DsHttp",
	2: "DsHttps",
	3: "DsS3",
	4: "DsSFTP",
}
var DsType_value = map[string]int32{
	"DsUnknown": 0,
	"DsHttp":    1,
	"DsHttps":   2,
	"DsS3":      3,
	"DsSFTP":    4,
}

func (x DsType) String() string {
	return proto.EnumName(DsType_name, int32(x))
}
func (DsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type Format int32

const (
	Format_FmtUnknown Format = 0
	Format_Raw        Format = 1
	Format_QCOW       Format = 2
	Format_QCOW2      Format = 3
	Format_VHD        Format = 4
	Format_VMDK       Format = 5
	Format_OVA        Format = 6
	Format_VHDX       Format = 7
)

var Format_name = map[int32]string{
	0: "FmtUnknown",
	1: "Raw",
	2: "QCOW",
	3: "QCOW2",
	4: "VHD",
	5: "VMDK",
	6: "OVA",
	7: "VHDX",
}
var Format_value = map[string]int32{
	"FmtUnknown": 0,
	"Raw":        1,
	"QCOW":       2,
	"QCOW2":      3,
	"VHD":        4,
	"VMDK":       5,
	"OVA":        6,
	"VHDX":       7,
}

func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}
func (Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type Target int32

const (
	Target_TgtUnknown Target = 0
	Target_Disk       Target = 1
	Target_Kernel     Target = 2
	Target_Initrd     Target = 3
	Target_RamDisk    Target = 4
)

var Target_name = map[int32]string{
	0: "TgtUnknown",
	1: "Disk",
	2: "Kernel",
	3: "Initrd",
	4: "RamDisk",
}
var Target_value = map[string]int32{
	"TgtUnknown": 0,
	"Disk":       1,
	"Kernel":     2,
	"Initrd":     3,
	"RamDisk":    4,
}

func (x Target) String() string {
	return proto.EnumName(Target_name, int32(x))
}
func (Target) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

type DriveType int32

const (
	DriveType_Unclassified DriveType = 0
	DriveType_CDROM        DriveType = 1
	DriveType_HDD          DriveType = 2
	DriveType_NET          DriveType = 3
)

var DriveType_name = map[int32]string{
	0: "Unclassified",
	1: "CDROM",
	2: "HDD",
	3: "NET",
}
var DriveType_value = map[string]int32{
	"Unclassified": 0,
	"CDROM":        1,
	"HDD":          2,
	"NET":          3,
}

func (x DriveType) String() string {
	return proto.EnumName(DriveType_name, int32(x))
}
func (DriveType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

type SignatureInfo struct {
	Intercertsurl string `protobuf:"bytes,1,opt,name=intercertsurl" json:"intercertsurl,omitempty"`
	Signercerturl string `protobuf:"bytes,2,opt,name=signercerturl" json:"signercerturl,omitempty"`
	Signature     []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignatureInfo) Reset()                    { *m = SignatureInfo{} }
func (m *SignatureInfo) String() string            { return proto.CompactTextString(m) }
func (*SignatureInfo) ProtoMessage()               {}
func (*SignatureInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *SignatureInfo) GetIntercertsurl() string {
	if m != nil {
		return m.Intercertsurl
	}
	return ""
}

func (m *SignatureInfo) GetSignercerturl() string {
	if m != nil {
		return m.Signercerturl
	}
	return ""
}

func (m *SignatureInfo) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type DatastoreConfig struct {
	Id       string `protobuf:"bytes,100,opt,name=id" json:"id,omitempty"`
	DType    DsType `protobuf:"varint,1,opt,name=dType,enum=DsType" json:"dType,omitempty"`
	Fqdn     string `protobuf:"bytes,2,opt,name=fqdn" json:"fqdn,omitempty"`
	ApiKey   string `protobuf:"bytes,3,opt,name=apiKey" json:"apiKey,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	// depending on datastore types, it could be bucket or path
	Dpath string `protobuf:"bytes,5,opt,name=dpath" json:"dpath,omitempty"`
}

func (m *DatastoreConfig) Reset()                    { *m = DatastoreConfig{} }
func (m *DatastoreConfig) String() string            { return proto.CompactTextString(m) }
func (*DatastoreConfig) ProtoMessage()               {}
func (*DatastoreConfig) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *DatastoreConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DatastoreConfig) GetDType() DsType {
	if m != nil {
		return m.DType
	}
	return DsType_DsUnknown
}

func (m *DatastoreConfig) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *DatastoreConfig) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *DatastoreConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *DatastoreConfig) GetDpath() string {
	if m != nil {
		return m.Dpath
	}
	return ""
}

type Image struct {
	Id             string          `protobuf:"bytes,100,opt,name=id" json:"id,omitempty"`
	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion" json:"uuidandversion,omitempty"`
	// it could be relative path/name as well
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Sha256  string `protobuf:"bytes,3,opt,name=sha256" json:"sha256,omitempty"`
	Size    int64  `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
	Iformat Format `protobuf:"varint,4,opt,name=iformat,enum=Format" json:"iformat,omitempty"`
	// if its signed image
	Siginfo *SignatureInfo `protobuf:"bytes,5,opt,name=siginfo" json:"siginfo,omitempty"`
	DsId    string         `protobuf:"bytes,6,opt,name=dsId" json:"dsId,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func (m *Image) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Image) GetIformat() Format {
	if m != nil {
		return m.Iformat
	}
	return Format_FmtUnknown
}

func (m *Image) GetSiginfo() *SignatureInfo {
	if m != nil {
		return m.Siginfo
	}
	return nil
}

func (m *Image) GetDsId() string {
	if m != nil {
		return m.DsId
	}
	return ""
}

type Drive struct {
	Image    *Image    `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Maxsize  int64     `protobuf:"varint,2,opt,name=maxsize" json:"maxsize,omitempty"`
	Readonly bool      `protobuf:"varint,5,opt,name=readonly" json:"readonly,omitempty"`
	Preserve bool      `protobuf:"varint,6,opt,name=preserve" json:"preserve,omitempty"`
	Drvtype  DriveType `protobuf:"varint,8,opt,name=drvtype,enum=DriveType" json:"drvtype,omitempty"`
	Target   Target    `protobuf:"varint,9,opt,name=target,enum=Target" json:"target,omitempty"`
}

func (m *Drive) Reset()                    { *m = Drive{} }
func (m *Drive) String() string            { return proto.CompactTextString(m) }
func (*Drive) ProtoMessage()               {}
func (*Drive) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *Drive) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Drive) GetMaxsize() int64 {
	if m != nil {
		return m.Maxsize
	}
	return 0
}

func (m *Drive) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

func (m *Drive) GetPreserve() bool {
	if m != nil {
		return m.Preserve
	}
	return false
}

func (m *Drive) GetDrvtype() DriveType {
	if m != nil {
		return m.Drvtype
	}
	return DriveType_Unclassified
}

func (m *Drive) GetTarget() Target {
	if m != nil {
		return m.Target
	}
	return Target_TgtUnknown
}

func init() {
	proto.RegisterType((*SignatureInfo)(nil), "SignatureInfo")
	proto.RegisterType((*DatastoreConfig)(nil), "DatastoreConfig")
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*Drive)(nil), "Drive")
	proto.RegisterEnum("DsType", DsType_name, DsType_value)
	proto.RegisterEnum("Format", Format_name, Format_value)
	proto.RegisterEnum("Target", Target_name, Target_value)
	proto.RegisterEnum("DriveType", DriveType_name, DriveType_value)
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x94, 0x51, 0x6f, 0x22, 0x37,
	0x10, 0xc7, 0xb3, 0x0b, 0xbb, 0x0b, 0x93, 0x40, 0x2c, 0xab, 0xaa, 0x56, 0x51, 0xa2, 0xa4, 0x28,
	0x0f, 0x88, 0x87, 0x8d, 0x44, 0xd4, 0x56, 0xea, 0x53, 0xdb, 0x6c, 0x29, 0x28, 0x4a, 0xd3, 0x6e,
	0x80, 0x9e, 0x4e, 0xf7, 0xe2, 0xac, 0xcd, 0xc6, 0x0a, 0x6b, 0x73, 0xb6, 0x21, 0x47, 0x3e, 0xcc,
	0x7d, 0x94, 0xfb, 0x4e, 0xf7, 0x0d, 0x4e, 0xf6, 0x02, 0xb9, 0xdc, 0xbd, 0xcd, 0xff, 0x3f, 0xc3,
	0xf8, 0x37, 0x33, 0x00, 0xb4, 0xb4, 0x91, 0x8a, 0x14, 0x2c, 0x59, 0x28, 0x69, 0xe4, 0xd1, 0x21,
	0x65, 0xab, 0x5c, 0x96, 0xa5, 0x14, 0x95, 0xd1, 0x59, 0x43, 0xeb, 0x8e, 0x17, 0x82, 0x98, 0xa5,
	0x62, 0x23, 0x31, 0x93, 0xf8, 0x1c, 0x5a, 0x5c, 0x18, 0xa6, 0x72, 0xa6, 0x8c, 0x5e, 0xaa, 0x79,
	0xec, 0x9d, 0x79, 0xdd, 0x66, 0xf6, 0xda, 0xb4, 0x55, 0x9a, 0x17, 0xa2, 0x72, 0x6c, 0x95, 0x5f,
	0x55, 0xbd, 0x32, 0xf1, 0x31, 0x34, 0xf5, 0xb6, 0x79, 0x5c, 0x3b, 0xf3, 0xba, 0x07, 0xd9, 0x8b,
	0xd1, 0xf9, 0xe8, 0xc1, 0x61, 0x4a, 0x0c, 0xb1, 0x84, 0xec, 0x4a, 0x8a, 0x19, 0x2f, 0x70, 0x1b,
	0x7c, 0x4e, 0x63, 0xea, 0x9a, 0xf9, 0x9c, 0xe2, 0x13, 0x08, 0xe8, 0x78, 0xbd, 0x60, 0x8e, 0xa2,
	0xdd, 0x8f, 0x92, 0x54, 0x5b, 0x99, 0x55, 0x2e, 0xc6, 0x50, 0x9f, 0xbd, 0xa7, 0x62, 0xf3, 0xba,
	0x8b, 0xf1, 0x8f, 0x10, 0x92, 0x05, 0xbf, 0x66, 0x6b, 0xf7, 0x62, 0x33, 0xdb, 0x28, 0x7c, 0x04,
	0x8d, 0x05, 0xd1, 0xfa, 0x49, 0x2a, 0x1a, 0xd7, 0x5d, 0x66, 0xa7, 0xf1, 0x0f, 0x10, 0xd0, 0x05,
	0x31, 0x0f, 0x71, 0xe0, 0x12, 0x95, 0xe8, 0x7c, 0xf6, 0x20, 0x18, 0x95, 0xa4, 0x60, 0xdf, 0x61,
	0xfd, 0x0a, 0xed, 0xe5, 0x92, 0x53, 0x22, 0xe8, 0x8a, 0x29, 0xcd, 0xa5, 0x70, 0x7c, 0xfb, 0xfd,
	0xc3, 0x64, 0x32, 0x19, 0xa5, 0x44, 0xd0, 0x69, 0x65, 0x67, 0xdf, 0x94, 0x59, 0x60, 0x41, 0x4a,
	0xb6, 0x05, 0xb6, 0xb1, 0x05, 0xd6, 0x0f, 0xa4, 0xff, 0xf3, 0x2f, 0x5b, 0xe0, 0x4a, 0xd9, 0x5a,
	0xcd, 0x9f, 0x59, 0x1c, 0x9d, 0x79, 0xdd, 0x5a, 0xe6, 0x62, 0xfc, 0x13, 0x44, 0x7c, 0x26, 0x55,
	0x49, 0x8c, 0x9b, 0xc1, 0x6e, 0x64, 0xe0, 0x64, 0xb6, 0xf5, 0x71, 0x17, 0x22, 0xcd, 0x0b, 0x2e,
	0x66, 0xd2, 0x4d, 0xb3, 0xdf, 0x6f, 0x27, 0xaf, 0x2e, 0x9c, 0x6d, 0xd3, 0xf6, 0x01, 0xaa, 0x47,
	0x34, 0x0e, 0x2b, 0x18, 0x1b, 0x77, 0x3e, 0x79, 0x10, 0xa4, 0x8a, 0xaf, 0x18, 0x3e, 0x86, 0x80,
	0xdb, 0xe1, 0x37, 0xa3, 0x85, 0x89, 0x5b, 0x45, 0x56, 0x99, 0x38, 0x86, 0xa8, 0x24, 0x1f, 0x1c,
	0x9f, 0xef, 0xf8, 0xb6, 0xd2, 0xee, 0x59, 0x31, 0x42, 0xa5, 0x98, 0xaf, 0x1d, 0x40, 0x23, 0xdb,
	0x69, 0x77, 0x03, 0xc5, 0x34, 0x53, 0x2b, 0xe6, 0x5e, 0x6d, 0x64, 0x3b, 0x8d, 0xcf, 0x21, 0xa2,
	0x6a, 0x65, 0xec, 0xb1, 0x1b, 0x6e, 0x34, 0x48, 0x1c, 0x88, 0xbb, 0xf7, 0x36, 0x85, 0x4f, 0x21,
	0x34, 0x44, 0x15, 0xcc, 0xc4, 0xcd, 0xcd, 0xfc, 0x63, 0x27, 0xb3, 0x8d, 0xdd, 0x1b, 0x40, 0x58,
	0x7d, 0x47, 0x70, 0x0b, 0x9a, 0xa9, 0x9e, 0x88, 0x47, 0x21, 0x9f, 0x04, 0xda, 0xc3, 0x60, 0x13,
	0x43, 0x63, 0x16, 0xc8, 0xc3, 0xfb, 0x10, 0x55, 0xb1, 0x46, 0x3e, 0x6e, 0x40, 0x3d, 0xd5, 0x77,
	0x97, 0xa8, 0x56, 0x95, 0xdc, 0x0d, 0xc6, 0xff, 0xa2, 0x7a, 0xef, 0x1d, 0x84, 0xd5, 0x66, 0x71,
	0x1b, 0x60, 0x50, 0x9a, 0x97, 0x46, 0x11, 0xd4, 0x32, 0xf2, 0x84, 0x3c, 0xfb, 0xc1, 0xff, 0xae,
	0x6e, 0xff, 0x47, 0x3e, 0x6e, 0x42, 0x60, 0xa3, 0x3e, 0xaa, 0xd9, 0xec, 0x74, 0x98, 0xa2, 0xba,
	0xcd, 0x4e, 0x6f, 0xd2, 0x6b, 0x14, 0x58, 0xeb, 0x76, 0xfa, 0x07, 0x0a, 0x9d, 0x35, 0x4c, 0xdf,
	0xa0, 0xa8, 0xf7, 0x37, 0x84, 0x15, 0xb7, 0xed, 0x3e, 0x2e, 0xbe, 0xea, 0x6e, 0x69, 0xb8, 0x7e,
	0x44, 0x9e, 0xa5, 0xb9, 0x66, 0x4a, 0xb0, 0x39, 0xf2, 0x6d, 0x3c, 0x12, 0xdc, 0x28, 0x8a, 0x6a,
	0x16, 0x3e, 0x23, 0xa5, 0x2b, 0xaa, 0xf7, 0x7e, 0x83, 0xe6, 0x6e, 0x4b, 0x18, 0xc1, 0xc1, 0x44,
	0xe4, 0x73, 0xa2, 0x35, 0x9f, 0x71, 0x46, 0xd1, 0x9e, 0x05, 0xbb, 0x4a, 0xb3, 0xdb, 0x1b, 0xe4,
	0x59, 0x8a, 0x61, 0x9a, 0x22, 0xdf, 0x06, 0xff, 0xfc, 0x35, 0x46, 0xb5, 0x3f, 0x7f, 0x87, 0xd3,
	0x5c, 0x96, 0xc9, 0x33, 0xa3, 0x8c, 0x92, 0x24, 0x9f, 0xcb, 0x25, 0x4d, 0x96, 0xf6, 0x18, 0x3c,
	0xdf, 0xfc, 0x5f, 0xbc, 0x3d, 0x29, 0xb8, 0x79, 0x58, 0xde, 0x27, 0xb9, 0x2c, 0x2f, 0xaa, 0xba,
	0x0b, 0xb2, 0xe0, 0x17, 0xcf, 0xb9, 0xfb, 0xb9, 0xde, 0x87, 0xae, 0xea, 0xf2, 0x4b, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x32, 0xa5, 0x51, 0xca, 0x66, 0x04, 0x00, 0x00,
}
