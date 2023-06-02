package fileseeker

import (
	"time"
)

var (
	MIN_FILE_SIZE int64     = -9_223_372_036_854_775_808
	MAX_FILE_SIZE int64     = 9_223_372_036_854_775_807
	MIN_DATE_TIME time.Time = time.Date(0001, time.January, 1, 0, 0, 0, 0, time.UTC)
	MAX_DATE_TIME time.Time = time.Date(2262, time.April, 11, 23, 0, 0, 0, time.UTC)
)

type FileSeekerBuilder interface {
	Patterns([]string) FileSeekerBuilder
	ExcludeSubdirs() FileSeekerBuilder
	Build() FileSeeker
	SizeRangeFilter([2]int64) FileSeekerBuilder
	ModificationDateRangeFilter([2]time.Time) FileSeekerBuilder
}

type fileSeekerBuilder struct {
	fsc fileSeekerConfig
}

func NewFileSeekerBuilder(folderPath string) FileSeekerBuilder {
	return &fileSeekerBuilder{
		fileSeekerConfig{
			folderPath:     folderPath,
			includeSubdirs: true,
			sizeRange: [2]int64{
				MIN_FILE_SIZE,
				MAX_FILE_SIZE,
			},
			modificationDateRange: [2]time.Time{
				MIN_DATE_TIME,
				MAX_DATE_TIME,
			},
		},
	}
}

func (fsb *fileSeekerBuilder) Patterns(patterns []string) FileSeekerBuilder {
	fsb.fsc.patterns = patterns
	return fsb
}

func (fsb *fileSeekerBuilder) ExcludeSubdirs() FileSeekerBuilder {
	fsb.fsc.includeSubdirs = false
	return fsb
}

func (fsb *fileSeekerBuilder) SizeRangeFilter(sizeRange [2]int64) FileSeekerBuilder {
	fsb.fsc.sizeRange = sizeRange
	return fsb
}

func (fsb *fileSeekerBuilder) ModificationDateRangeFilter(dateRange [2]time.Time) FileSeekerBuilder {
	fsb.fsc.modificationDateRange = dateRange
	return fsb
}

func (fsb *fileSeekerBuilder) Build() FileSeeker {
	return &fileSeekerImpl{
		fileSeekerConfig{
			folderPath:            fsb.fsc.folderPath,
			patterns:              fsb.fsc.patterns,
			includeSubdirs:        fsb.fsc.includeSubdirs,
			sizeRange:             fsb.fsc.sizeRange,
			modificationDateRange: fsb.fsc.modificationDateRange,
		},
	}
}
