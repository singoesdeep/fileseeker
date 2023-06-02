package fileseeker

import (
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type FileSeeker interface {
	SeekFiles() ([]File, error)
}

type fileSeekerImpl struct {
	fsc fileSeekerConfig
}

func (fs *fileSeekerImpl) SeekFiles() ([]File, error) {
	var filesTemp []File
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
			filesTemp = append(filesTemp, subfolderFiles...)
		} else {
			if len(fs.fsc.patterns) > 0 {
				filePath := filepath.Join(fs.fsc.folderPath, entry.Name())
				if fs.matchesPattern(filePath) {
					file := NewFile(filePath)
					filesTemp = append(filesTemp, file)
				}
			} else {
				filePath := filepath.Join(fs.fsc.folderPath, entry.Name())
				file := NewFile(filePath)
				filesTemp = append(filesTemp, file)
			}
		}
	}
	for _, file := range filesTemp {
		if filterTimeRange(file.ModificationDate, fs.fsc.modificationDateRange) &&
			filterSizeRange(file.Size, fs.fsc.sizeRange) {
			files = append(files, file)
		}
	}
	return files, nil
}

func (fs *fileSeekerImpl) matchesPattern(filePath string) bool {
	if len(fs.fsc.patterns) == 0 {
		return false
	}

	for _, pattern := range fs.fsc.patterns {
		if matchPattern(pattern, filePath) {
			return true
		}
	}

	return false
}

func matchPattern(pattern, filePath string) bool {
	regExp, err := regexp.Compile(pattern)

	if err != nil {
		return false
	}

	return regExp.MatchString(filePath)
}

func filterTimeRange(ft time.Time, t [2]time.Time) bool {
	if ft.After(t[0]) && ft.Before(t[1]) {
		return true
	}
	return false
}

func filterSizeRange(fs int64, s [2]int64) bool {
	if fs >= s[0] && fs <= s[1] {
		return true
	}
	return false
}
