package classfile

import "fmt"

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	supperClass  uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
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
	// 读取常量池
	self.readConstantPool(reader)
	// 读取访问标志
	self.accessFlags = reader.readUint16()
	// 读取类索引
	self.thisClass = reader.readUint16()
	// 读取父类索引
	self.supperClass = reader.readUint16()
	// 读取接口索引
	self.interfaces = reader.readUint16s()
	// 读取字段信息
	self.fields = readMembers(reader, self.constantPool)
	// 读取方法信息
	self.methods = readMembers(reader, self.constantPool)
	// 读取属性信息
	self.attributes = readAttributes(reader, self.constantPool)
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

// 读取常量池
func (self *ClassFile) readConstantPool(reader *ClassReader) {

}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	return self.constantPool.getClassName(self.supperClass)
}

func (self *ClassFile) Interfaces() []string {
	if self.interfaces == nil {
		return nil
	}
	interfaces := make([]string, len(self.interfaces))
	for i, index := range self.interfaces {
		interfaces[i] = self.constantPool.getClassName(index)
	}
	return interfaces
}
