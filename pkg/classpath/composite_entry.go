package classpath

import (
	"path/filepath"
	"strings"
)

type CompositoryEntry []Entry

func NewCompositeEntry(paths string) CompositoryEntry {
	pathList := strings.Split(paths, pathListSeperator)
	entries := make([]Entry, len(pathList))

	for i, each := range pathList {
		entry := NewEntry(each)
		entries[i] = entry
	}

	return entries
}

func (self CompositoryEntry) String() string {
	paths := make([]string, len(self))
	for i, entry := range self {
		paths[i] = entry.String()
	}

	return strings.Join(paths, string(filepath.Separator))
}

func (self CompositoryEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}

	return nil, nil, &ErrClassNotFound{className: className}
}
