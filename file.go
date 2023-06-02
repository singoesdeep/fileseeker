package fileseeker

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type File struct {
	Path             string
	Name             string
	Extension        string
	Size             int64
	ModificationDate time.Time
	Permissions      os.FileMode
}

func NewFile(filePath string) File {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return File{}
	}

	name := fileInfo.Name()
	extension := filepath.Ext(name)
	if fileInfo.IsDir() {
		extension = "dir"
	} else {
		extension = strings.TrimPrefix(extension, ".")
	}

	return File{
		Path:             filePath,
		Name:             name,
		Extension:        extension,
		Size:             fileInfo.Size(),
		ModificationDate: fileInfo.ModTime(),
		Permissions:      fileInfo.Mode().Perm(),
	}
}

func (f *File) String() []string {
	return []string{
		f.Path,
		f.Name,
		f.Extension,
		f.sizeString(),
		f.modTimeString(),
		f.permissionString(),
	}
}

func (f *File) sizeString() string {
	return strconv.FormatInt(f.Size, 10)
}

func (f *File) modTimeString() string {
	return f.ModificationDate.Format("2006-01-02 15:04:05")
}

func (f *File) permissionString() string {
	const permFormat = "rwxrwxrwx"

	dirPermMap := map[os.FileMode]string{
		0400: "r--------",
		0440: "r--r-----",
		0444: "r--r--r--",
		0500: "r-x------",
		0550: "r-xr-x---",
		0555: "r-xr-xr-x",
		0600: "rw-------",
		0640: "rw-r-----",
		0644: "rw-r--r--",
		0700: "rwx------",
		0750: "rwxr-x---",
		0755: "rwxr-xr-x",
		0777: "rwxrwxrwx",
	}

	if f.Permissions&os.ModeDir != 0 {
		if dirPerm, ok := dirPermMap[f.Permissions.Perm()]; ok {
			return dirPerm
		}
	}

	var result []byte
	for i := 0; i < 9; i++ {
		if f.Permissions&(1<<(8-i)) != 0 {
			result = append(result, permFormat[i])
		} else {
			result = append(result, '-')
		}
	}

	return string(result)
}
