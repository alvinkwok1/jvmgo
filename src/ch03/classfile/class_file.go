package classfile

import "fmt"

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	// 读取魔数
	self.readAndCheckMagic(reader)
	// 读取版本号信息
	self.readAndCheckVersion(reader)
}

// 读取并校验魔数
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	self.magic = reader.readUint32()
	if self.magic != 0xCAFEBABE {
		panic("java.lang.ClassFile: Magic mismatch")
	}
}

// 读取版本信息
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 读取次版本号
	self.minorVersion = reader.readUint16()
	// 读取主版本号
	self.majorVersion = reader.readUint16()
	// 校验版本信息
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError")
}
