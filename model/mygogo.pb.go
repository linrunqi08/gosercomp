// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mygogo.proto

package model

import (
	fmt "fmt"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type GogoProtoColorGroup struct {
	Id                   *int32   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name                 *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Colors               []string `protobuf:"bytes,3,rep,name=colors" json:"colors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GogoProtoColorGroup) Reset()         { *m = GogoProtoColorGroup{} }
func (m *GogoProtoColorGroup) String() string { return proto.CompactTextString(m) }
func (*GogoProtoColorGroup) ProtoMessage()    {}
func (*GogoProtoColorGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd40715f6243e6fe, []int{0}
}
func (m *GogoProtoColorGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GogoProtoColorGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GogoProtoColorGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GogoProtoColorGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GogoProtoColorGroup.Merge(m, src)
}
func (m *GogoProtoColorGroup) XXX_Size() int {
	return m.Size()
}
func (m *GogoProtoColorGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_GogoProtoColorGroup.DiscardUnknown(m)
}

var xxx_messageInfo_GogoProtoColorGroup proto.InternalMessageInfo

func (m *GogoProtoColorGroup) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GogoProtoColorGroup) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GogoProtoColorGroup) GetColors() []string {
	if m != nil {
		return m.Colors
	}
	return nil
}

func init() {
	proto.RegisterType((*GogoProtoColorGroup)(nil), "model.GogoProtoColorGroup")
}

func init() { proto.RegisterFile("mygogo.proto", fileDescriptor_dd40715f6243e6fe) }

var fileDescriptor_dd40715f6243e6fe = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xad, 0x4c, 0xcf,
	0x4f, 0xcf, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x51,
	0x0a, 0xe4, 0x12, 0x76, 0xcf, 0x4f, 0xcf, 0x0f, 0x00, 0x89, 0x39, 0xe7, 0xe7, 0xe4, 0x17, 0xb9,
	0x17, 0xe5, 0x97, 0x16, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x69, 0xb0,
	0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30,
	0x69, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x62, 0x5c, 0x6c, 0xc9, 0x20, 0x1d, 0xc5, 0x12, 0xcc, 0x0a,
	0xcc, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x93, 0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x03, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x60,
	0xde, 0x44, 0x7a, 0x00, 0x00, 0x00,
}

func (m *GogoProtoColorGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GogoProtoColorGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GogoProtoColorGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Colors) > 0 {
		for iNdEx := len(m.Colors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Colors[iNdEx])
			copy(dAtA[i:], m.Colors[iNdEx])
			i = encodeVarintMygogo(dAtA, i, uint64(len(m.Colors[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Name == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i -= len(*m.Name)
		copy(dAtA[i:], *m.Name)
		i = encodeVarintMygogo(dAtA, i, uint64(len(*m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		i = encodeVarintMygogo(dAtA, i, uint64(*m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMygogo(dAtA []byte, offset int, v uint64) int {
	offset -= sovMygogo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GogoProtoColorGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		n += 1 + sovMygogo(uint64(*m.Id))
	}
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovMygogo(uint64(l))
	}
	if len(m.Colors) > 0 {
		for _, s := range m.Colors {
			l = len(s)
			n += 1 + l + sovMygogo(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMygogo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMygogo(x uint64) (n int) {
	return sovMygogo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GogoProtoColorGroup) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMygogo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GogoProtoColorGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GogoProtoColorGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMygogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Id = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMygogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMygogo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMygogo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Name = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Colors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMygogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMygogo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMygogo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Colors = append(m.Colors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMygogo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMygogo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMygogo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMygogo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMygogo
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
					return 0, ErrIntOverflowMygogo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMygogo
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
			if length < 0 {
				return 0, ErrInvalidLengthMygogo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMygogo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMygogo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMygogo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMygogo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMygogo = fmt.Errorf("proto: unexpected end of group")
)
