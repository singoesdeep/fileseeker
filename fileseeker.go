package fileseeker

import (
	"os"
	"path/filepath"
	"regexp"
)

type FileSeeker interface {
	SeekFiles() ([]File, error)
}

type fileSeekerImpl struct {
	fsc fileSeekerConfig
}

func (fs *fileSeekerImpl) SeekFiles() ([]File, error) {
	var files []File

	entries, err := os.ReadDir(fs.fsc.folderPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() && fs.fsc.includeSubdirs {
			subfolderPath := filepath.Join(fs.fsc.folderPath, entry.Name())
			subfolderFiles, err := NewFileSeekerBuilder(subfolderPath).
				Patterns(fs.fsc.patterns).
				Build().
				SeekFiles()
			if err != nil {
				return nil, err
			}
			files = append(files, subfolderFiles...)
		} else {
			if len(fs.fsc.patterns) > 0 {
				filePath := filepath.Join(fs.fsc.folderPath, entry.Name())
				if fs.matchesPattern(filePath) {
					file := NewFile(filePath)
					files = append(files, file)
				}
			} else {
				filePath := filepath.Join(fs.fsc.folderPath, entry.Name())
				file := NewFile(filePath)
				files = append(files, file)
			}
		}
	}

	return files, nil
}

func (fs *fileSeekerImpl) matchesPattern(filePath string) bool {
	if len(fs.fsc.patterns) == 0 {
		return false
	}

	for _, pattern := range fs.fsc.patterns {
		if fs.matchPattern(pattern, filePath) {
			return true
		}
	}

	return false
}

func (fsi *fileSeekerImpl) matchPattern(pattern, filePath string) bool {
	regExp, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return regExp.MatchString(filePath)
}
