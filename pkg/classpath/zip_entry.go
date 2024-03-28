package classpath

import (
	"archive/zip"
	"io"
	"strings"
)

type zipEntry struct {
	zipFilename string
	zipReader   *zip.ReadCloser
}

func NewZipEntry(zipFilename string) *zipEntry {
	return &zipEntry{
		zipFilename: zipFilename,
	}
}

func (self *zipEntry) String() string {
	return self.zipFilename
}

func (self *zipEntry) readClass(className string) ([]byte, Entry, error) {
	var err error
	if self.zipReader == nil {
		if self.zipReader, err = zip.OpenReader(self.zipFilename); err != nil {
			return nil, nil, err
		}
	}

	file, err := self.zipReader.Open(className)
	if err != nil {
		return nil, nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}

	buffer := make([]byte, fileInfo.Size())
	_, err = file.Read(buffer)
	if err == io.EOF {
		err = nil
	}
	return buffer, self, err
}

func isZipFile(dirPath string) bool {
	for _, suffix := range zipSuffixes {
		if strings.HasSuffix(dirPath, suffix) {
			return true
		}
	}

	return false
}
