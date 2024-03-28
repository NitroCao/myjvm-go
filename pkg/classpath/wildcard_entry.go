package classpath

import (
	"io/fs"
	"path/filepath"
)

func NewWildcardEntry(dirPath string) CompositoryEntry {
	baseDir := filepath.Dir(dirPath)
	entries := []Entry{}

	filepath.Walk(baseDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return fs.SkipDir
		}

		if isZipFile(path) {
			jarEntry := NewZipEntry(path)
			entries = append(entries, jarEntry)
		}

		return nil
	})

	return entries
}
