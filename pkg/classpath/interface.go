package classpath

import (
	"fmt"
	"strings"
)

const (
	pathListSeperator = ";"
)

var (
	zipSuffixes = []string{".jar", ".JAR", ".zip", ".ZIP"}
)

type ErrClassNotFound struct {
	className string
}

func (self *ErrClassNotFound) Error() string {
	return fmt.Sprintf("class not found: %s", self.className)
}

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func NewEntry(dirPath string) Entry {
	if strings.Contains(dirPath, pathListSeperator) {
		return NewCompositeEntry(dirPath)
	}

	if strings.HasSuffix(dirPath, "*") {
		return NewWildcardEntry(dirPath)
	}

	if isZipFile(dirPath) {
		return NewZipEntry(dirPath)
	}

	return NewDirEntry(dirPath)
}
