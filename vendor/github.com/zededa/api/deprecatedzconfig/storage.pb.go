// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage.proto

/*
	Package deprecatedzconfig is a generated protocol buffer package.

	It is generated from these files:
		storage.proto

	It has these top-level messages:
		SignatureInfo
		StorageConfig
*/
package deprecatedzconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignatureInfo struct {
	Intermediatecertspem []byte `protobuf:"bytes,1,opt,name=intermediatecertspem,proto3" json:"intermediatecertspem,omitempty"`
	Signercertpem        []byte `protobuf:"bytes,2,opt,name=signercertpem,proto3" json:"signercertpem,omitempty"`
	Signature            []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignatureInfo) Reset()                    { *m = SignatureInfo{} }
func (m *SignatureInfo) String() string            { return proto.CompactTextString(m) }
func (*SignatureInfo) ProtoMessage()               {}
func (*SignatureInfo) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{0} }

func (m *SignatureInfo) GetIntermediatecertspem() []byte {
	if m != nil {
		return m.Intermediatecertspem
	}
	return nil
}

func (m *SignatureInfo) GetSignercertpem() []byte {
	if m != nil {
		return m.Signercertpem
	}
	return nil
}

func (m *SignatureInfo) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type StorageConfig struct {
	Downloadurl string         `protobuf:"bytes,1,opt,name=downloadurl,proto3" json:"downloadurl,omitempty"`
	Maxsize     uint32         `protobuf:"varint,2,opt,name=maxsize,proto3" json:"maxsize,omitempty"`
	Siginfo     *SignatureInfo `protobuf:"bytes,3,opt,name=siginfo" json:"siginfo,omitempty"`
	Imagesha256 string         `protobuf:"bytes,4,opt,name=imagesha256,proto3" json:"imagesha256,omitempty"`
	Readonly    bool           `protobuf:"varint,5,opt,name=readonly,proto3" json:"readonly,omitempty"`
	Preserve    bool           `protobuf:"varint,6,opt,name=preserve,proto3" json:"preserve,omitempty"`
	Format      string         `protobuf:"bytes,7,opt,name=format,proto3" json:"format,omitempty"`
	Devtype     string         `protobuf:"bytes,8,opt,name=devtype,proto3" json:"devtype,omitempty"`
	Target      string         `protobuf:"bytes,9,opt,name=target,proto3" json:"target,omitempty"`
}

func (m *StorageConfig) Reset()                    { *m = StorageConfig{} }
func (m *StorageConfig) String() string            { return proto.CompactTextString(m) }
func (*StorageConfig) ProtoMessage()               {}
func (*StorageConfig) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{1} }

func (m *StorageConfig) GetDownloadurl() string {
	if m != nil {
		return m.Downloadurl
	}
	return ""
}

func (m *StorageConfig) GetMaxsize() uint32 {
	if m != nil {
		return m.Maxsize
	}
	return 0
}

func (m *StorageConfig) GetSiginfo() *SignatureInfo {
	if m != nil {
		return m.Siginfo
	}
	return nil
}

func (m *StorageConfig) GetImagesha256() string {
	if m != nil {
		return m.Imagesha256
	}
	return ""
}

func (m *StorageConfig) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

func (m *StorageConfig) GetPreserve() bool {
	if m != nil {
		return m.Preserve
	}
	return false
}

func (m *StorageConfig) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *StorageConfig) GetDevtype() string {
	if m != nil {
		return m.Devtype
	}
	return ""
}

func (m *StorageConfig) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func init() {
	proto.RegisterType((*SignatureInfo)(nil), "SignatureInfo")
	proto.RegisterType((*StorageConfig)(nil), "StorageConfig")
}
func (m *SignatureInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Intermediatecertspem) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Intermediatecertspem)))
		i += copy(dAtA[i:], m.Intermediatecertspem)
	}
	if len(m.Signercertpem) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Signercertpem)))
		i += copy(dAtA[i:], m.Signercertpem)
	}
	if len(m.Signature) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	return i, nil
}

func (m *StorageConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StorageConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Downloadurl) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Downloadurl)))
		i += copy(dAtA[i:], m.Downloadurl)
	}
	if m.Maxsize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.Maxsize))
	}
	if m.Siginfo != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.Siginfo.Size()))
		n1, err := m.Siginfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Imagesha256) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Imagesha256)))
		i += copy(dAtA[i:], m.Imagesha256)
	}
	if m.Readonly {
		dAtA[i] = 0x28
		i++
		if m.Readonly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Preserve {
		dAtA[i] = 0x30
		i++
		if m.Preserve {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Format) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Format)))
		i += copy(dAtA[i:], m.Format)
	}
	if len(m.Devtype) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Devtype)))
		i += copy(dAtA[i:], m.Devtype)
	}
	if len(m.Target) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Target)))
		i += copy(dAtA[i:], m.Target)
	}
	return i, nil
}

func encodeVarintStorage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SignatureInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Intermediatecertspem)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Signercertpem)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *StorageConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Downloadurl)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	if m.Maxsize != 0 {
		n += 1 + sovStorage(uint64(m.Maxsize))
	}
	if m.Siginfo != nil {
		l = m.Siginfo.Size()
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Imagesha256)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	if m.Readonly {
		n += 2
	}
	if m.Preserve {
		n += 2
	}
	l = len(m.Format)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Devtype)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func sovStorage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStorage(x uint64) (n int) {
	return sovStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignatureInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignatureInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignatureInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Intermediatecertspem", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Intermediatecertspem = append(m.Intermediatecertspem[:0], dAtA[iNdEx:postIndex]...)
			if m.Intermediatecertspem == nil {
				m.Intermediatecertspem = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signercertpem", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signercertpem = append(m.Signercertpem[:0], dAtA[iNdEx:postIndex]...)
			if m.Signercertpem == nil {
				m.Signercertpem = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StorageConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StorageConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StorageConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Downloadurl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Downloadurl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Maxsize", wireType)
			}
			m.Maxsize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Maxsize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Siginfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Siginfo == nil {
				m.Siginfo = &SignatureInfo{}
			}
			if err := m.Siginfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Imagesha256", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Imagesha256 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Readonly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Readonly = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Preserve", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Preserve = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Format = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Devtype", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Devtype = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStorage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStorage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStorage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStorage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("storage.proto", fileDescriptorStorage) }

var fileDescriptorStorage = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x4e, 0xc3, 0x30,
	0x18, 0xc5, 0x71, 0x81, 0xfe, 0x71, 0x09, 0x12, 0x16, 0x42, 0x16, 0x42, 0xa1, 0xaa, 0x18, 0x3a,
	0x65, 0x28, 0x82, 0x03, 0x00, 0x0b, 0x6b, 0xba, 0xb1, 0x99, 0xf8, 0x8b, 0xb1, 0x94, 0xd8, 0x91,
	0xed, 0x14, 0xda, 0x0b, 0x70, 0x05, 0xb8, 0x11, 0x23, 0x47, 0x40, 0xe5, 0x22, 0xc8, 0x4e, 0x03,
	0xad, 0xc4, 0xf8, 0xde, 0xfb, 0x3e, 0xff, 0x9e, 0x6d, 0x1c, 0x59, 0xa7, 0x0d, 0x13, 0x90, 0x54,
	0x46, 0x3b, 0x3d, 0x7e, 0x45, 0x38, 0x9a, 0x49, 0xa1, 0x98, 0xab, 0x0d, 0xdc, 0xab, 0x5c, 0x93,
	0x29, 0x3e, 0x96, 0xca, 0x81, 0x29, 0x81, 0x4b, 0xe6, 0x20, 0x03, 0xe3, 0x6c, 0x05, 0x25, 0x45,
	0x23, 0x34, 0x39, 0x48, 0xff, 0xcd, 0xc8, 0x05, 0x8e, 0xac, 0x14, 0x0a, 0x8c, 0x77, 0xfc, 0x70,
	0x27, 0x0c, 0x6f, 0x9b, 0xe4, 0x0c, 0x0f, 0x6c, 0x8b, 0xa2, 0xbb, 0x61, 0xe2, 0xcf, 0x18, 0xbf,
	0x77, 0x70, 0x34, 0x6b, 0xba, 0xdd, 0x6a, 0x95, 0x4b, 0x41, 0x46, 0x78, 0xc8, 0xf5, 0xb3, 0x2a,
	0x34, 0xe3, 0xb5, 0x29, 0x42, 0x81, 0x41, 0xba, 0x69, 0x11, 0x8a, 0x7b, 0x25, 0x7b, 0xb1, 0x72,
	0x09, 0x81, 0x18, 0xa5, 0xad, 0x24, 0x13, 0xdc, 0xb3, 0x52, 0x48, 0x95, 0xeb, 0x40, 0x1a, 0x4e,
	0x0f, 0x93, 0xad, 0x6b, 0xa6, 0x6d, 0xec, 0x29, 0xb2, 0x64, 0x02, 0xec, 0x13, 0x9b, 0x5e, 0x5d,
	0xd3, 0xbd, 0x86, 0xb2, 0x61, 0x91, 0x53, 0xdc, 0x37, 0xc0, 0xb8, 0x56, 0xc5, 0x82, 0xee, 0x8f,
	0xd0, 0xa4, 0x9f, 0xfe, 0x6a, 0x9f, 0x55, 0x06, 0x2c, 0x98, 0x39, 0xd0, 0x6e, 0x93, 0xb5, 0x9a,
	0x9c, 0xe0, 0x6e, 0xae, 0x4d, 0xc9, 0x1c, 0xed, 0x85, 0x43, 0xd7, 0xca, 0xb7, 0xe6, 0x30, 0x77,
	0x8b, 0x0a, 0x68, 0x3f, 0x04, 0xad, 0xf4, 0x1b, 0x8e, 0x19, 0x01, 0x8e, 0x0e, 0x9a, 0x8d, 0x46,
	0xdd, 0xdc, 0x7d, 0xac, 0x62, 0xf4, 0xb9, 0x8a, 0xd1, 0xd7, 0x2a, 0x46, 0x6f, 0xdf, 0xf1, 0x0e,
	0x3e, 0xcf, 0x74, 0x99, 0x2c, 0x81, 0x03, 0x67, 0x49, 0x56, 0xe8, 0x9a, 0x27, 0xb5, 0x87, 0xca,
	0x6c, 0xfd, 0xb1, 0x0f, 0x47, 0x1c, 0x2a, 0x03, 0x19, 0x73, 0xc0, 0x97, 0x59, 0x78, 0xcf, 0xc7,
	0x6e, 0x48, 0x2e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0x8c, 0x25, 0x9b, 0x03, 0x02, 0x00,
	0x00,
}