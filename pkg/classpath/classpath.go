package classpath

import (
	"fmt"
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	appClasspath  Entry
}

const (
	jreDir = "./jre"
)

var (
	ErrNoJreDirFound = "no any JRE dir found"
)

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}

	cp.parseBootAndExtClasspath(jreOption)
	cp.parseAppClasspath(cpOption)

	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = fmt.Sprintf("%s.class", className)
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.appClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.appClasspath.String()
}

func (self *Classpath) parseAppClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	self.appClasspath = NewEntry(cpOption)
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = NewWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = NewWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" {
		if _, err := os.Stat(jreOption); err == nil {
			return jreOption
		}
	}

	if _, err := os.Stat(jreDir); err == nil {
		return jreDir
	}

	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, jreDir)
	}

	return ErrNoJreDirFound
}
