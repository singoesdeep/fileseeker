package fileseeker

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileSeeker interface {
	SeekFiles() ([]string, error)
}

type fileSeekerImpl struct {
	folderPath     string
	patterns       []string
	fileExtensions []string
	useRegExp      bool
	includeSubdirs bool
}

func (fs *fileSeekerImpl) SeekFiles() ([]string, error) {
	var files []string

	entries, err := os.ReadDir(fs.folderPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() && fs.includeSubdirs {
			subfolderPath := filepath.Join(fs.folderPath, entry.Name())
			subfolderFiles, err := NewFileSeekerBuilder(subfolderPath).
				Patterns(fs.patterns).
				FileExtensions(fs.fileExtensions).
				Build().
				SeekFiles()
			if err != nil {
				return nil, err
			}
			files = append(files, subfolderFiles...)
		} else {
			if len(fs.patterns) == 0 && len(fs.fileExtensions) == 0 {
				files = append(files, filepath.Join(fs.folderPath, entry.Name()))
			} else {
				filePath := filepath.Join(fs.folderPath, entry.Name())
				if fs.matchesPattern(filePath) || fs.matchesExtension(entry.Name()) {
					files = append(files, filePath)
				}
			}
		}
	}

	return files, nil
}

func (fs *fileSeekerImpl) matchesPattern(filePath string) bool {
	if !fs.useRegExp || len(fs.patterns) == 0 {
		return false
	}

	for _, pattern := range fs.patterns {
		if fs.matchPattern(pattern, filePath) {
			return true
		}
	}

	return false
}

func (fs *fileSeekerImpl) matchPattern(pattern, filePath string) bool {
	regExp, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return regExp.MatchString(filePath)
}

func (fs *fileSeekerImpl) matchesExtension(fileName string) bool {
	if len(fs.fileExtensions) == 0 {
		return false
	}

	extension := strings.TrimPrefix(filepath.Ext(fileName), ".")

	for _, ext := range fs.fileExtensions {
		if extension == ext {
			return true
		}
	}

	return false
}
