package validators

import (
	"fmt"
	"os"
	"path/filepath"
)

type argFileValidator struct {
	file string
}

func NewArgFileValidator(file string) argFileValidator {
	return argFileValidator{file}
}

func (v argFileValidator) Validate() error {
	if !v.fileExists() {
		return fmt.Errorf("error: file %s not found", v.file)
	}
	if filepath.Ext(v.file) != ".json" {
		return fmt.Errorf("error: only .json extension is supported")
	}
	return nil
}

func (v argFileValidator) fileExists() bool {
	info, err := os.Stat(v.file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
