package fileseeker

import (
	"time"
)

type fileSeekerConfig struct {
	folderPath            string
	patterns              []string
	includeSubdirs        bool
	sizeRange             [2]int64
	modificationDateRange [2]time.Time
}
