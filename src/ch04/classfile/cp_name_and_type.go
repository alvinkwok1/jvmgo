package classfile

// CONSTANT_NameAndType_info { u1 tag; u2 name_index; u2 descriptor_index; }

type ConstantNameAndTypeInfo struct {
	cp              ConstantPool
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}

func (self *ConstantNameAndTypeInfo) ClassName() string {
	return self.cp.getClassName(self.nameIndex)
}

func (self *ConstantNameAndTypeInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.descriptorIndex)
}
