// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: supplychain/newcontract.proto

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

type Newcontract struct {
	Index         string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	DealId        string `protobuf:"bytes,2,opt,name=dealId,proto3" json:"dealId,omitempty"`
	ContractId    string `protobuf:"bytes,3,opt,name=contractId,proto3" json:"contractId,omitempty"`
	Consumer      string `protobuf:"bytes,4,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Desc          string `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	Ownereta      uint64 `protobuf:"varint,6,opt,name=ownereta,proto3" json:"ownereta,omitempty"`
	Vendoreta     uint64 `protobuf:"varint,7,opt,name=vendoreta,proto3" json:"vendoreta,omitempty"`
	Status        string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Fees          string `protobuf:"bytes,9,opt,name=fees,proto3" json:"fees,omitempty"`
	Expiry        string `protobuf:"bytes,10,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Shippingdelay uint64 `protobuf:"varint,11,opt,name=shippingdelay,proto3" json:"shippingdelay,omitempty"`
	Starttime     string `protobuf:"bytes,12,opt,name=starttime,proto3" json:"starttime,omitempty"`
	Deliverydelay uint64 `protobuf:"varint,13,opt,name=deliverydelay,proto3" json:"deliverydelay,omitempty"`
}

func (m *Newcontract) Reset()         { *m = Newcontract{} }
func (m *Newcontract) String() string { return proto.CompactTextString(m) }
func (*Newcontract) ProtoMessage()    {}
func (*Newcontract) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9b40b922fae5c37, []int{0}
}
func (m *Newcontract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Newcontract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Newcontract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Newcontract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Newcontract.Merge(m, src)
}
func (m *Newcontract) XXX_Size() int {
	return m.Size()
}
func (m *Newcontract) XXX_DiscardUnknown() {
	xxx_messageInfo_Newcontract.DiscardUnknown(m)
}

var xxx_messageInfo_Newcontract proto.InternalMessageInfo

func (m *Newcontract) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Newcontract) GetDealId() string {
	if m != nil {
		return m.DealId
	}
	return ""
}

func (m *Newcontract) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *Newcontract) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

func (m *Newcontract) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Newcontract) GetOwnereta() uint64 {
	if m != nil {
		return m.Ownereta
	}
	return 0
}

func (m *Newcontract) GetVendoreta() uint64 {
	if m != nil {
		return m.Vendoreta
	}
	return 0
}

func (m *Newcontract) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Newcontract) GetFees() string {
	if m != nil {
		return m.Fees
	}
	return ""
}

func (m *Newcontract) GetExpiry() string {
	if m != nil {
		return m.Expiry
	}
	return ""
}

func (m *Newcontract) GetShippingdelay() uint64 {
	if m != nil {
		return m.Shippingdelay
	}
	return 0
}

func (m *Newcontract) GetStarttime() string {
	if m != nil {
		return m.Starttime
	}
	return ""
}

func (m *Newcontract) GetDeliverydelay() uint64 {
	if m != nil {
		return m.Deliverydelay
	}
	return 0
}

func init() {
	proto.RegisterType((*Newcontract)(nil), "upalchowdhury.supplychain.supplychain.Newcontract")
}

func init() { proto.RegisterFile("supplychain/newcontract.proto", fileDescriptor_e9b40b922fae5c37) }

var fileDescriptor_e9b40b922fae5c37 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x31, 0x6e, 0x32, 0x31,
	0x10, 0x85, 0x59, 0x7e, 0xe0, 0x07, 0x13, 0x1a, 0x2b, 0x8a, 0xac, 0x28, 0xb1, 0x50, 0x94, 0x48,
	0x54, 0x50, 0xa4, 0x49, 0x9d, 0x8e, 0x26, 0x05, 0x65, 0x3a, 0x63, 0x4f, 0x58, 0x4b, 0x8b, 0x6d,
	0xd9, 0x5e, 0x60, 0x6f, 0x91, 0x63, 0xa5, 0xa4, 0x4c, 0x19, 0xc1, 0x11, 0x72, 0x81, 0xc8, 0x76,
	0x02, 0xbb, 0xdd, 0xbc, 0x6f, 0xde, 0xbc, 0x19, 0x69, 0xd0, 0xad, 0x2b, 0x8d, 0x29, 0x2a, 0x9e,
	0x33, 0xa9, 0x66, 0x0a, 0xb6, 0x5c, 0x2b, 0x6f, 0x19, 0xf7, 0x53, 0x63, 0xb5, 0xd7, 0xf8, 0xa1,
	0x34, 0xac, 0xe0, 0xb9, 0xde, 0x8a, 0xbc, 0xb4, 0xd5, 0xb4, 0x66, 0xae, 0xd7, 0x77, 0xdf, 0x6d,
	0x34, 0x7c, 0x39, 0x0f, 0xe3, 0x4b, 0xd4, 0x95, 0x4a, 0xc0, 0x8e, 0x64, 0xe3, 0x6c, 0x32, 0x58,
	0x24, 0x81, 0xaf, 0x50, 0x4f, 0x00, 0x2b, 0xe6, 0x82, 0xb4, 0x23, 0xfe, 0x55, 0x98, 0x22, 0xf4,
	0x37, 0x39, 0x17, 0xe4, 0x5f, 0xec, 0xd5, 0x08, 0xbe, 0x46, 0x7d, 0xae, 0x95, 0x2b, 0xd7, 0x60,
	0x49, 0x27, 0x76, 0x4f, 0x1a, 0x63, 0xd4, 0x11, 0xe0, 0x38, 0xe9, 0x46, 0x1e, 0xeb, 0xe0, 0xd7,
	0x5b, 0x05, 0x16, 0x3c, 0x23, 0xbd, 0x71, 0x36, 0xe9, 0x2c, 0x4e, 0x1a, 0xdf, 0xa0, 0xc1, 0x06,
	0x94, 0xd0, 0xb1, 0xf9, 0x3f, 0x36, 0xcf, 0x20, 0x5c, 0xe8, 0x3c, 0xf3, 0xa5, 0x23, 0xfd, 0x74,
	0x61, 0x52, 0x61, 0xcb, 0x1b, 0x80, 0x23, 0x83, 0xb4, 0x25, 0xd4, 0xc1, 0x0b, 0x3b, 0x23, 0x6d,
	0x45, 0x50, 0xf2, 0x26, 0x85, 0xef, 0xd1, 0xc8, 0xe5, 0xd2, 0x18, 0xa9, 0x56, 0x02, 0x0a, 0x56,
	0x91, 0x61, 0xdc, 0xd2, 0x84, 0xe1, 0x0e, 0xe7, 0x99, 0xf5, 0x5e, 0xae, 0x81, 0x5c, 0xc4, 0x80,
	0x33, 0x08, 0x19, 0x02, 0x0a, 0xb9, 0x01, 0x5b, 0xa5, 0x8c, 0x51, 0xca, 0x68, 0xc0, 0xe7, 0xc5,
	0xc7, 0x81, 0x66, 0xfb, 0x03, 0xcd, 0xbe, 0x0e, 0x34, 0x7b, 0x3f, 0xd2, 0xd6, 0xfe, 0x48, 0x5b,
	0x9f, 0x47, 0xda, 0x7a, 0x7d, 0x5a, 0x49, 0x9f, 0x97, 0xcb, 0x29, 0xd7, 0xeb, 0x59, 0xe3, 0x83,
	0xb3, 0xfa, 0xbb, 0x77, 0x0d, 0xe5, 0x2b, 0x03, 0x6e, 0xd9, 0x8b, 0x7f, 0x7f, 0xfc, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xcf, 0x11, 0x4b, 0x45, 0x18, 0x02, 0x00, 0x00,
}

func (m *Newcontract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Newcontract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Newcontract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Deliverydelay != 0 {
		i = encodeVarintNewcontract(dAtA, i, uint64(m.Deliverydelay))
		i--
		dAtA[i] = 0x68
	}
	if len(m.Starttime) > 0 {
		i -= len(m.Starttime)
		copy(dAtA[i:], m.Starttime)
		i = encodeVarintNewcontract(dAtA, i, uint64(len(m.Starttime)))
		i--
		dAtA[i] = 0x62
	}
	if m.Shippingdelay != 0 {
		i = encodeVarintNewcontract(dAtA, i, uint64(m.Shippingdelay))
		i--
		dAtA[i] = 0x58
	}
	if len(m.Expiry) > 0 {
		i -= len(m.Expiry)
		copy(dAtA[i:], m.Expiry)
		i = encodeVarintNewcontract(dAtA, i, uint64(len(m.Expiry)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Fees) > 0 {
		i -= len(m.Fees)
		copy(dAtA[i:], m.Fees)
		i = encodeVarintNewcontract(dAtA, i, uint64(len(m.Fees)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintNewcontract(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x42
	}
	if m.Vendoreta != 0 {
		i = encodeVarintNewcontract(dAtA, i, uint64(m.Vendoreta))
		i--
		dAtA[i] = 0x38
	}
	if m.Ownereta != 0 {
		i = encodeVarintNewcontract(dAtA, i, uint64(m.Ownereta))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintNewcontract(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Consumer) > 0 {
		i -= len(m.Consumer)
		copy(dAtA[i:], m.Consumer)
		i = encodeVarintNewcontract(dAtA, i, uint64(len(m.Consumer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ContractId) > 0 {
		i -= len(m.ContractId)
		copy(dAtA[i:], m.ContractId)
		i = encodeVarintNewcontract(dAtA, i, uint64(len(m.ContractId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DealId) > 0 {
		i -= len(m.DealId)
		copy(dAtA[i:], m.DealId)
		i = encodeVarintNewcontract(dAtA, i, uint64(len(m.DealId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintNewcontract(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNewcontract(dAtA []byte, offset int, v uint64) int {
	offset -= sovNewcontract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Newcontract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovNewcontract(uint64(l))
	}
	l = len(m.DealId)
	if l > 0 {
		n += 1 + l + sovNewcontract(uint64(l))
	}
	l = len(m.ContractId)
	if l > 0 {
		n += 1 + l + sovNewcontract(uint64(l))
	}
	l = len(m.Consumer)
	if l > 0 {
		n += 1 + l + sovNewcontract(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovNewcontract(uint64(l))
	}
	if m.Ownereta != 0 {
		n += 1 + sovNewcontract(uint64(m.Ownereta))
	}
	if m.Vendoreta != 0 {
		n += 1 + sovNewcontract(uint64(m.Vendoreta))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovNewcontract(uint64(l))
	}
	l = len(m.Fees)
	if l > 0 {
		n += 1 + l + sovNewcontract(uint64(l))
	}
	l = len(m.Expiry)
	if l > 0 {
		n += 1 + l + sovNewcontract(uint64(l))
	}
	if m.Shippingdelay != 0 {
		n += 1 + sovNewcontract(uint64(m.Shippingdelay))
	}
	l = len(m.Starttime)
	if l > 0 {
		n += 1 + l + sovNewcontract(uint64(l))
	}
	if m.Deliverydelay != 0 {
		n += 1 + sovNewcontract(uint64(m.Deliverydelay))
	}
	return n
}

func sovNewcontract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNewcontract(x uint64) (n int) {
	return sovNewcontract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Newcontract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNewcontract
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
			return fmt.Errorf("proto: Newcontract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Newcontract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
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
				return ErrInvalidLengthNewcontract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewcontract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DealId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
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
				return ErrInvalidLengthNewcontract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewcontract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DealId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
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
				return ErrInvalidLengthNewcontract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewcontract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
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
				return ErrInvalidLengthNewcontract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewcontract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Consumer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
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
				return ErrInvalidLengthNewcontract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewcontract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ownereta", wireType)
			}
			m.Ownereta = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ownereta |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vendoreta", wireType)
			}
			m.Vendoreta = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vendoreta |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
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
				return ErrInvalidLengthNewcontract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewcontract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
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
				return ErrInvalidLengthNewcontract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewcontract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
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
				return ErrInvalidLengthNewcontract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewcontract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Expiry = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shippingdelay", wireType)
			}
			m.Shippingdelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Shippingdelay |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Starttime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
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
				return ErrInvalidLengthNewcontract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewcontract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Starttime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deliverydelay", wireType)
			}
			m.Deliverydelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewcontract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deliverydelay |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNewcontract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNewcontract
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
func skipNewcontract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNewcontract
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
					return 0, ErrIntOverflowNewcontract
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
					return 0, ErrIntOverflowNewcontract
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
				return 0, ErrInvalidLengthNewcontract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNewcontract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNewcontract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNewcontract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNewcontract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNewcontract = fmt.Errorf("proto: unexpected end of group")
)
