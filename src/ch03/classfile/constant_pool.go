package classfile

type ConstantPool []ConstantInfo

// 常量池解析
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		// TODO 需要处理Long和Double
	}
	return cp
}

// 访问指定索引的常量信息
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index !")
}

//
//func (self ConstantPool) getClassName(index uint16) string {
//	classInfo := self.getConstantInfo(index).(*C)
//}

// 从常量池中查找指定索引的utf8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
