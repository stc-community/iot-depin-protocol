// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iotdepinprotocol/iotdepinprotocol/event_pb.proto

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

type EventPb struct {
	PubId   string `protobuf:"bytes,1,opt,name=pubId,proto3" json:"pubId,omitempty"`
	Topic   string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	PubType string `protobuf:"bytes,3,opt,name=pubType,proto3" json:"pubType,omitempty"`
	Payload string `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	PubTime uint64 `protobuf:"varint,5,opt,name=pubTime,proto3" json:"pubTime,omitempty"`
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *EventPb) Reset()         { *m = EventPb{} }
func (m *EventPb) String() string { return proto.CompactTextString(m) }
func (*EventPb) ProtoMessage()    {}
func (*EventPb) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5711a30d9e13a8, []int{0}
}
func (m *EventPb) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPb.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPb.Merge(m, src)
}
func (m *EventPb) XXX_Size() int {
	return m.Size()
}
func (m *EventPb) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPb.DiscardUnknown(m)
}

var xxx_messageInfo_EventPb proto.InternalMessageInfo

func (m *EventPb) GetPubId() string {
	if m != nil {
		return m.PubId
	}
	return ""
}

func (m *EventPb) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *EventPb) GetPubType() string {
	if m != nil {
		return m.PubType
	}
	return ""
}

func (m *EventPb) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *EventPb) GetPubTime() uint64 {
	if m != nil {
		return m.PubTime
	}
	return 0
}

func (m *EventPb) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*EventPb)(nil), "stccommunity.iotdepinprotocol.iotdepinprotocol.EventPb")
}

func init() {
	proto.RegisterFile("iotdepinprotocol/iotdepinprotocol/event_pb.proto", fileDescriptor_4e5711a30d9e13a8)
}

var fileDescriptor_4e5711a30d9e13a8 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc8, 0xcc, 0x2f, 0x49,
	0x49, 0x2d, 0xc8, 0xcc, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0xc7, 0x10, 0x48,
	0x2d, 0x4b, 0xcd, 0x2b, 0x89, 0x2f, 0x48, 0xd2, 0x03, 0x8b, 0x08, 0xe9, 0x15, 0x97, 0x24, 0x27,
	0xe7, 0xe7, 0xe6, 0x96, 0xe6, 0x65, 0x96, 0x54, 0xea, 0xa1, 0xab, 0xc6, 0x10, 0x50, 0x9a, 0xcb,
	0xc8, 0xc5, 0xee, 0x0a, 0x32, 0x22, 0x20, 0x49, 0x48, 0x84, 0x8b, 0xb5, 0xa0, 0x34, 0xc9, 0x33,
	0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x01, 0x89, 0x96, 0xe4, 0x17, 0x64, 0x26,
	0x4b, 0x30, 0x41, 0x44, 0xc1, 0x1c, 0x21, 0x09, 0x2e, 0xf6, 0x82, 0xd2, 0xa4, 0x90, 0xca, 0x82,
	0x54, 0x09, 0x66, 0xb0, 0x38, 0x8c, 0x0b, 0x96, 0x49, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x91, 0x60,
	0x81, 0xca, 0x40, 0xb8, 0x30, 0x3d, 0x99, 0xb9, 0xa9, 0x12, 0xac, 0x0a, 0x8c, 0x1a, 0x2c, 0x41,
	0x30, 0x2e, 0x48, 0x26, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0xbf, 0x48, 0x82, 0x0d, 0xa2, 0x07, 0xca,
	0x75, 0x8a, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x97, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe2, 0x92, 0x64, 0x5d, 0xb8, 0xaf, 0x41,
	0x61, 0xa4, 0x0b, 0xf6, 0xa5, 0x2e, 0x3c, 0x94, 0x2a, 0x30, 0x03, 0xae, 0xa4, 0xb2, 0x20, 0xb5,
	0x38, 0x89, 0x0d, 0xcc, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x42, 0x4f, 0xa6, 0x6a,
	0x01, 0x00, 0x00,
}

func (m *EventPb) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPb) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPb) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEventPb(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if m.PubTime != 0 {
		i = encodeVarintEventPb(dAtA, i, uint64(m.PubTime))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintEventPb(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PubType) > 0 {
		i -= len(m.PubType)
		copy(dAtA[i:], m.PubType)
		i = encodeVarintEventPb(dAtA, i, uint64(len(m.PubType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Topic) > 0 {
		i -= len(m.Topic)
		copy(dAtA[i:], m.Topic)
		i = encodeVarintEventPb(dAtA, i, uint64(len(m.Topic)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PubId) > 0 {
		i -= len(m.PubId)
		copy(dAtA[i:], m.PubId)
		i = encodeVarintEventPb(dAtA, i, uint64(len(m.PubId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEventPb(dAtA []byte, offset int, v uint64) int {
	offset -= sovEventPb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventPb) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubId)
	if l > 0 {
		n += 1 + l + sovEventPb(uint64(l))
	}
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovEventPb(uint64(l))
	}
	l = len(m.PubType)
	if l > 0 {
		n += 1 + l + sovEventPb(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovEventPb(uint64(l))
	}
	if m.PubTime != 0 {
		n += 1 + sovEventPb(uint64(m.PubTime))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEventPb(uint64(l))
	}
	return n
}

func sovEventPb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEventPb(x uint64) (n int) {
	return sovEventPb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventPb) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventPb
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
			return fmt.Errorf("proto: EventPb: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPb: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventPb
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
				return ErrInvalidLengthEventPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventPb
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
				return ErrInvalidLengthEventPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventPb
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
				return ErrInvalidLengthEventPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventPb
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
				return ErrInvalidLengthEventPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubTime", wireType)
			}
			m.PubTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PubTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventPb
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
				return ErrInvalidLengthEventPb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEventPb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEventPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEventPb
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
func skipEventPb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEventPb
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
					return 0, ErrIntOverflowEventPb
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
					return 0, ErrIntOverflowEventPb
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
				return 0, ErrInvalidLengthEventPb
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEventPb
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEventPb
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEventPb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEventPb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEventPb = fmt.Errorf("proto: unexpected end of group")
)