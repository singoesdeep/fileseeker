package fileseeker

type fileSeekerConfig struct {
	folderPath     string
	patterns       []string
	fileExtensions []string
	useRegExp      bool
	includeSubdirs bool
}
