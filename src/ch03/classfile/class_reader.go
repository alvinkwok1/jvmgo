package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// 读取u1
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// 读取u2
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// 读取u4
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

// 读取u8
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// 读取u2表
func (self *ClassReader) readUint16s() []uint16 {
	num := self.readUint16()
	s := make([]uint16, num)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(size uint32) []byte {
	bytes := self.data[:size]
	self.data = self.data[size:]
	return bytes
}
