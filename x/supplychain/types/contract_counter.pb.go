// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: supplychain/contract_counter.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ContractCounter struct {
	DealId  string `protobuf:"bytes,1,opt,name=dealId,proto3" json:"dealId,omitempty"`
	IdValue uint64 `protobuf:"varint,2,opt,name=idValue,proto3" json:"idValue,omitempty"`
}

func (m *ContractCounter) Reset()         { *m = ContractCounter{} }
func (m *ContractCounter) String() string { return proto.CompactTextString(m) }
func (*ContractCounter) ProtoMessage()    {}
func (*ContractCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_486b4ae796099ea9, []int{0}
}
func (m *ContractCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCounter.Merge(m, src)
}
func (m *ContractCounter) XXX_Size() int {
	return m.Size()
}
func (m *ContractCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCounter.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCounter proto.InternalMessageInfo

func (m *ContractCounter) GetDealId() string {
	if m != nil {
		return m.DealId
	}
	return ""
}

func (m *ContractCounter) GetIdValue() uint64 {
	if m != nil {
		return m.IdValue
	}
	return 0
}

func init() {
	proto.RegisterType((*ContractCounter)(nil), "upalchowdhury.supplychain.supplychain.ContractCounter")
}

func init() {
	proto.RegisterFile("supplychain/contract_counter.proto", fileDescriptor_486b4ae796099ea9)
}

var fileDescriptor_486b4ae796099ea9 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x2e, 0x2d, 0x28,
	0xc8, 0xa9, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e,
	0x89, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52,
	0x2d, 0x2d, 0x48, 0xcc, 0x49, 0xce, 0xc8, 0x2f, 0x4f, 0xc9, 0x28, 0x2d, 0xaa, 0xd4, 0x43, 0xd2,
	0x81, 0xcc, 0x56, 0x72, 0xe6, 0xe2, 0x77, 0x86, 0x1a, 0xe0, 0x0c, 0xd1, 0x2f, 0x24, 0xc6, 0xc5,
	0x96, 0x92, 0x9a, 0x98, 0xe3, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x09,
	0x49, 0x70, 0xb1, 0x67, 0xa6, 0x84, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xb0,
	0x04, 0xc1, 0xb8, 0x4e, 0x41, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65,
	0x91, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0xe2, 0x20, 0x7d, 0x64,
	0x2f, 0x54, 0xa0, 0xf0, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xde, 0x30, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x50, 0x0b, 0x6d, 0x45, 0xec, 0x00, 0x00, 0x00,
}

func (m *ContractCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IdValue != 0 {
		i = encodeVarintContractCounter(dAtA, i, uint64(m.IdValue))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DealId) > 0 {
		i -= len(m.DealId)
		copy(dAtA[i:], m.DealId)
		i = encodeVarintContractCounter(dAtA, i, uint64(len(m.DealId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContractCounter(dAtA []byte, offset int, v uint64) int {
	offset -= sovContractCounter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DealId)
	if l > 0 {
		n += 1 + l + sovContractCounter(uint64(l))
	}
	if m.IdValue != 0 {
		n += 1 + sovContractCounter(uint64(m.IdValue))
	}
	return n
}

func sovContractCounter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContractCounter(x uint64) (n int) {
	return sovContractCounter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContractCounter
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
			return fmt.Errorf("proto: ContractCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DealId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCounter
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
				return ErrInvalidLengthContractCounter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractCounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DealId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdValue", wireType)
			}
			m.IdValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractCounter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IdValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipContractCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContractCounter
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
func skipContractCounter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContractCounter
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
					return 0, ErrIntOverflowContractCounter
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
					return 0, ErrIntOverflowContractCounter
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
				return 0, ErrInvalidLengthContractCounter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContractCounter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContractCounter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContractCounter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContractCounter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContractCounter = fmt.Errorf("proto: unexpected end of group")
)