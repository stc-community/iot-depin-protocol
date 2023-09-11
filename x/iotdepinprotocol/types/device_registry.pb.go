// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iotdepinprotocol/iotdepinprotocol/device_registry.proto

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

type DeviceRegistry struct {
	Mid      string `protobuf:"bytes,1,opt,name=mid,proto3" json:"mid,omitempty"`
	MetaData string `protobuf:"bytes,2,opt,name=metaData,proto3" json:"metaData,omitempty"`
	Creator  string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *DeviceRegistry) Reset()         { *m = DeviceRegistry{} }
func (m *DeviceRegistry) String() string { return proto.CompactTextString(m) }
func (*DeviceRegistry) ProtoMessage()    {}
func (*DeviceRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e624334e96d3ade6, []int{0}
}
func (m *DeviceRegistry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceRegistry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceRegistry.Merge(m, src)
}
func (m *DeviceRegistry) XXX_Size() int {
	return m.Size()
}
func (m *DeviceRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceRegistry proto.InternalMessageInfo

func (m *DeviceRegistry) GetMid() string {
	if m != nil {
		return m.Mid
	}
	return ""
}

func (m *DeviceRegistry) GetMetaData() string {
	if m != nil {
		return m.MetaData
	}
	return ""
}

func (m *DeviceRegistry) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*DeviceRegistry)(nil), "stccommunity.iotdepinprotocol.iotdepinprotocol.DeviceRegistry")
}

func init() {
	proto.RegisterFile("iotdepinprotocol/iotdepinprotocol/device_registry.proto", fileDescriptor_e624334e96d3ade6)
}

var fileDescriptor_e624334e96d3ade6 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xcf, 0xcc, 0x2f, 0x49,
	0x49, 0x2d, 0xc8, 0xcc, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0xc7, 0x10, 0x48,
	0x49, 0x2d, 0xcb, 0x4c, 0x4e, 0x8d, 0x2f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x03,
	0x4b, 0x08, 0xe9, 0x15, 0x97, 0x24, 0x27, 0xe7, 0xe7, 0xe6, 0x96, 0xe6, 0x65, 0x96, 0x54, 0xea,
	0xa1, 0x6b, 0xc2, 0x10, 0x50, 0x8a, 0xe0, 0xe2, 0x73, 0x01, 0x1b, 0x14, 0x04, 0x35, 0x47, 0x48,
	0x80, 0x8b, 0x39, 0x37, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14, 0x92,
	0xe2, 0xe2, 0xc8, 0x4d, 0x2d, 0x49, 0x74, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x02, 0x0b, 0xc3, 0xf9,
	0x42, 0x12, 0x5c, 0xec, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xf9, 0x45, 0x12, 0xcc, 0x60, 0x29, 0x18,
	0xd7, 0x29, 0xee, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x5c, 0xd2, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x8b, 0x4b, 0x92, 0x75, 0xe1, 0xee, 0x05,
	0x79, 0x52, 0x17, 0xec, 0x3e, 0x5d, 0xb8, 0x37, 0x2b, 0x30, 0x7d, 0x5e, 0x52, 0x59, 0x90, 0x5a,
	0x9c, 0xc4, 0x06, 0xe6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x11, 0xd2, 0x88, 0x2b,
	0x01, 0x00, 0x00,
}

func (m *DeviceRegistry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceRegistry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceRegistry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDeviceRegistry(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MetaData) > 0 {
		i -= len(m.MetaData)
		copy(dAtA[i:], m.MetaData)
		i = encodeVarintDeviceRegistry(dAtA, i, uint64(len(m.MetaData)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Mid) > 0 {
		i -= len(m.Mid)
		copy(dAtA[i:], m.Mid)
		i = encodeVarintDeviceRegistry(dAtA, i, uint64(len(m.Mid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeviceRegistry(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeviceRegistry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceRegistry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Mid)
	if l > 0 {
		n += 1 + l + sovDeviceRegistry(uint64(l))
	}
	l = len(m.MetaData)
	if l > 0 {
		n += 1 + l + sovDeviceRegistry(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDeviceRegistry(uint64(l))
	}
	return n
}

func sovDeviceRegistry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeviceRegistry(x uint64) (n int) {
	return sovDeviceRegistry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceRegistry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceRegistry
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
			return fmt.Errorf("proto: DeviceRegistry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceRegistry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceRegistry
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
				return ErrInvalidLengthDeviceRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceRegistry
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
				return ErrInvalidLengthDeviceRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetaData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceRegistry
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
				return ErrInvalidLengthDeviceRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceRegistry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeviceRegistry
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
func skipDeviceRegistry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceRegistry
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
					return 0, ErrIntOverflowDeviceRegistry
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
					return 0, ErrIntOverflowDeviceRegistry
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
				return 0, ErrInvalidLengthDeviceRegistry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeviceRegistry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeviceRegistry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeviceRegistry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceRegistry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeviceRegistry = fmt.Errorf("proto: unexpected end of group")
)