package classfile

// 注意,字符串在class文件中是以MUTF-8(Modified UTF-8)方  式编码的。但为什么没有用标准的UTF-8编码方式,笔者没有找到
// 明确的原因 [1] 。MUTF-8编码方式和UTF-8大致相同,但并不兼容。
//
// 差别有两点:一是null字符(代码点U+0000)会被编码成2字节:  0xC0、0x80;二是补充字符(Supplementary Characters,
// 代码点大于  U+FFFF的Unicode字符)是按UTF-16拆分为代理对(Surrogate Pair)  分别编码的。
// 具体细节超出了本章的讨论范围,有兴趣的读者可以  阅读Java虚拟机规范和Unicode规范的相关章节
// TODO 需要研究下怎么解析的UTF8
// CONSTANT_Utf8_info { u1 tag; u2 length; u1 bytes[length]; }
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	// 字符串变量看来长度最大只有2字节，最多是2^16-1长度
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
