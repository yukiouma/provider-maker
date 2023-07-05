package providermaker

import (
	"os"
	"path/filepath"
)

func toAbs(path string) (s string, err error) {
	if filepath.IsAbs(path) {
		s = path
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	s = filepath.Join(pwd, path)
	return
}
