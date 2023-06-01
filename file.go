package fileseeker

import (
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Path      string
	Name      string
	Extension string
	Size      int64
}

func NewFile(filePath string) File {
	fileInfo, _ := os.Stat(filePath)

	name := fileInfo.Name()
	extension := strings.TrimPrefix(filepath.Ext(name), ".")

	return File{
		Path:      filePath,
		Name:      name,
		Extension: extension,
		Size:      fileInfo.Size(),
	}
}
