package classfile

// attribute_info {
// u2 attribute_name_index;
// u4 attribute_length;
// u1 info[attribute_length];
// }
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributeCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
	panic("Invalid constant pool index !")
}
