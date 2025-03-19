package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 负责搜寻和加载class文件
	readClass(className string) ([]byte, Entry, error)
	// 相当于toString，用于调试
	String() string
}

func newEntry(path string) Entry {
	// 对于路径中包含分隔符的情况，使用CompositeEntry
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	// 对于路径中包含通配符的情况，使用WildcardEntry
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	// 对于路径中包含zip/jar文件的情况，使用ZipEntry
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	// 对于普通路径的情况，使用DirEntry
	return newDirEntry(path)
}
