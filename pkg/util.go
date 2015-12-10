package util

import "path/filepath"

func GetWorkDir(mainFilePath string) (workdir string, err error) {
	dir, err := filepath.Abs(filepath.Dir(mainFilePath))
	if err != nil {
		return "", err
	}
	return dir, nil
}
