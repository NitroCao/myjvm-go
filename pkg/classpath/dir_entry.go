package classpath

import (
	"os"
	"path/filepath"
)

type dirEntry struct {
	absDir string
}

func NewDirEntry(absPath string) *dirEntry {
	return &dirEntry{
		absDir: absPath,
	}
}

func (self *dirEntry) String() string {
	return self.absDir
}

func (self *dirEntry) readClass(className string) ([]byte, Entry, error) {
	classFilename := filepath.Join(self.absDir, className)
	data, err := os.ReadFile(classFilename)
	return data, self, err
}
