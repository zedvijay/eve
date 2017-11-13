// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fw.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ACEMatch struct {
	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ACEMatch) Reset()                    { *m = ACEMatch{} }
func (m *ACEMatch) String() string            { return proto.CompactTextString(m) }
func (*ACEMatch) ProtoMessage()               {}
func (*ACEMatch) Descriptor() ([]byte, []int) { return fileDescriptorFw, []int{0} }

func (m *ACEMatch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ACEMatch) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ACEAction struct {
	Drop       bool   `protobuf:"varint,1,opt,name=drop,proto3" json:"drop,omitempty"`
	Limit      bool   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Limitrate  uint32 `protobuf:"varint,3,opt,name=limitrate,proto3" json:"limitrate,omitempty"`
	Limitunit  string `protobuf:"bytes,4,opt,name=limitunit,proto3" json:"limitunit,omitempty"`
	Limitburst uint32 `protobuf:"varint,5,opt,name=limitburst,proto3" json:"limitburst,omitempty"`
}

func (m *ACEAction) Reset()                    { *m = ACEAction{} }
func (m *ACEAction) String() string            { return proto.CompactTextString(m) }
func (*ACEAction) ProtoMessage()               {}
func (*ACEAction) Descriptor() ([]byte, []int) { return fileDescriptorFw, []int{1} }

func (m *ACEAction) GetDrop() bool {
	if m != nil {
		return m.Drop
	}
	return false
}

func (m *ACEAction) GetLimit() bool {
	if m != nil {
		return m.Limit
	}
	return false
}

func (m *ACEAction) GetLimitrate() uint32 {
	if m != nil {
		return m.Limitrate
	}
	return 0
}

func (m *ACEAction) GetLimitunit() string {
	if m != nil {
		return m.Limitunit
	}
	return ""
}

func (m *ACEAction) GetLimitburst() uint32 {
	if m != nil {
		return m.Limitburst
	}
	return 0
}

type ACE struct {
	Matches []*ACEMatch  `protobuf:"bytes,1,rep,name=matches" json:"matches,omitempty"`
	Actions []*ACEAction `protobuf:"bytes,2,rep,name=actions" json:"actions,omitempty"`
}

func (m *ACE) Reset()                    { *m = ACE{} }
func (m *ACE) String() string            { return proto.CompactTextString(m) }
func (*ACE) ProtoMessage()               {}
func (*ACE) Descriptor() ([]byte, []int) { return fileDescriptorFw, []int{2} }

func (m *ACE) GetMatches() []*ACEMatch {
	if m != nil {
		return m.Matches
	}
	return nil
}

func (m *ACE) GetActions() []*ACEAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*ACEMatch)(nil), "ACEMatch")
	proto.RegisterType((*ACEAction)(nil), "ACEAction")
	proto.RegisterType((*ACE)(nil), "ACE")
}
func (m *ACEMatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ACEMatch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFw(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFw(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *ACEAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ACEAction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Drop {
		dAtA[i] = 0x8
		i++
		if m.Drop {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Limit {
		dAtA[i] = 0x10
		i++
		if m.Limit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Limitrate != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintFw(dAtA, i, uint64(m.Limitrate))
	}
	if len(m.Limitunit) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintFw(dAtA, i, uint64(len(m.Limitunit)))
		i += copy(dAtA[i:], m.Limitunit)
	}
	if m.Limitburst != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintFw(dAtA, i, uint64(m.Limitburst))
	}
	return i, nil
}

func (m *ACE) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ACE) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Matches) > 0 {
		for _, msg := range m.Matches {
			dAtA[i] = 0xa
			i++
			i = encodeVarintFw(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Actions) > 0 {
		for _, msg := range m.Actions {
			dAtA[i] = 0x12
			i++
			i = encodeVarintFw(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintFw(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ACEMatch) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovFw(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovFw(uint64(l))
	}
	return n
}

func (m *ACEAction) Size() (n int) {
	var l int
	_ = l
	if m.Drop {
		n += 2
	}
	if m.Limit {
		n += 2
	}
	if m.Limitrate != 0 {
		n += 1 + sovFw(uint64(m.Limitrate))
	}
	l = len(m.Limitunit)
	if l > 0 {
		n += 1 + l + sovFw(uint64(l))
	}
	if m.Limitburst != 0 {
		n += 1 + sovFw(uint64(m.Limitburst))
	}
	return n
}

func (m *ACE) Size() (n int) {
	var l int
	_ = l
	if len(m.Matches) > 0 {
		for _, e := range m.Matches {
			l = e.Size()
			n += 1 + l + sovFw(uint64(l))
		}
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovFw(uint64(l))
		}
	}
	return n
}

func sovFw(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFw(x uint64) (n int) {
	return sovFw(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ACEMatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFw
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
			return fmt.Errorf("proto: ACEMatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ACEMatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFw
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
				return ErrInvalidLengthFw
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFw
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
				return ErrInvalidLengthFw
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFw(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFw
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
func (m *ACEAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFw
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
			return fmt.Errorf("proto: ACEAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ACEAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Drop", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFw
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
			m.Drop = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFw
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
			m.Limit = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limitrate", wireType)
			}
			m.Limitrate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limitrate |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limitunit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFw
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
				return ErrInvalidLengthFw
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Limitunit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limitburst", wireType)
			}
			m.Limitburst = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFw
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limitburst |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFw(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFw
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
func (m *ACE) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFw
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
			return fmt.Errorf("proto: ACE: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ACE: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Matches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFw
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
				return ErrInvalidLengthFw
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Matches = append(m.Matches, &ACEMatch{})
			if err := m.Matches[len(m.Matches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFw
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
				return ErrInvalidLengthFw
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, &ACEAction{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFw(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFw
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
func skipFw(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFw
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
					return 0, ErrIntOverflowFw
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
					return 0, ErrIntOverflowFw
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
				return 0, ErrInvalidLengthFw
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFw
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
				next, err := skipFw(dAtA[start:])
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
	ErrInvalidLengthFw = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFw   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("fw.proto", fileDescriptorFw) }

var fileDescriptorFw = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0x3d, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0xfd, 0xa1, 0xc9, 0x43, 0x2c, 0x16, 0x83, 0x07, 0x08, 0x55, 0x61, 0xe8, 0xe4, 0x4a,
	0xc0, 0x05, 0x42, 0xd5, 0x11, 0xa9, 0xf2, 0xc8, 0xe6, 0x38, 0x6e, 0x6b, 0x29, 0x89, 0xa3, 0xc4,
	0x2e, 0xa2, 0xb7, 0x60, 0xe3, 0x48, 0x8c, 0x1c, 0x01, 0x85, 0x8b, 0xa0, 0xbc, 0x28, 0x2d, 0xdb,
	0xf7, 0x2b, 0xfb, 0x7b, 0x10, 0x6c, 0xde, 0x78, 0x59, 0x59, 0x67, 0x67, 0x4f, 0x10, 0xc4, 0xcb,
	0xd5, 0x8b, 0x74, 0x6a, 0x47, 0x29, 0x8c, 0xdc, 0x7b, 0xa9, 0x19, 0x99, 0x92, 0x79, 0x28, 0x10,
	0xd3, 0x2b, 0x18, 0xef, 0x65, 0xe6, 0x35, 0x1b, 0xa0, 0xd8, 0x91, 0xd9, 0x07, 0x81, 0x30, 0x5e,
	0xae, 0x62, 0xe5, 0x8c, 0x2d, 0xda, 0x5e, 0x5a, 0xd9, 0x12, 0x7b, 0x81, 0x40, 0xdc, 0xf6, 0x32,
	0x93, 0x1b, 0x87, 0xbd, 0x40, 0x74, 0x84, 0x5e, 0x43, 0x88, 0xa0, 0x92, 0x4e, 0xb3, 0xe1, 0x94,
	0xcc, 0x2f, 0xc5, 0x49, 0x38, 0xba, 0xbe, 0x30, 0x8e, 0x8d, 0xf0, 0xbd, 0x93, 0x40, 0x23, 0x00,
	0x24, 0x89, 0xaf, 0x6a, 0xc7, 0xc6, 0x58, 0xfe, 0xa7, 0xcc, 0xd6, 0x30, 0x8c, 0x97, 0x2b, 0x7a,
	0x07, 0x93, 0xbc, 0x5d, 0xa3, 0x6b, 0x46, 0xa6, 0xc3, 0xf9, 0xc5, 0x43, 0xc8, 0xfb, 0x81, 0xa2,
	0x77, 0xe8, 0x3d, 0x4c, 0x24, 0xfe, 0xbd, 0x66, 0x03, 0x0c, 0x01, 0x3f, 0xce, 0x11, 0xbd, 0xf5,
	0xbc, 0xfe, 0x6a, 0x22, 0xf2, 0xdd, 0x44, 0xe4, 0xa7, 0x89, 0xc8, 0xe7, 0x6f, 0x74, 0x06, 0xb7,
	0xca, 0xe6, 0xfc, 0xa0, 0x53, 0x9d, 0x4a, 0xae, 0x32, 0xeb, 0x53, 0xee, 0x6b, 0x5d, 0xed, 0x8d,
	0xd2, 0xdd, 0x39, 0x5f, 0x6f, 0xb6, 0xc6, 0xed, 0x7c, 0xc2, 0x95, 0xcd, 0x17, 0x5d, 0x6e, 0x21,
	0x4b, 0xb3, 0x38, 0x28, 0x5b, 0x6c, 0xcc, 0x36, 0x39, 0xc7, 0xd4, 0xe3, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x87, 0xcf, 0xe0, 0xaa, 0x80, 0x01, 0x00, 0x00,
}