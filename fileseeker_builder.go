package fileseeker

type FileSeekerBuilder interface {
	Patterns([]string) FileSeekerBuilder
	FileExtensions([]string) FileSeekerBuilder
	ExcludeSubdirs() FileSeekerBuilder
	Build() FileSeeker
}

type fileSeekerBuilder struct {
	folderPath     string
	patterns       []string
	fileExtensions []string
	useRegExp      bool
	includeSubdirs bool
}

func NewFileSeekerBuilder(folderPath string) FileSeekerBuilder {
	return &fileSeekerBuilder{
		folderPath:     folderPath,
		includeSubdirs: true,
	}
}

func (fsb *fileSeekerBuilder) Patterns(patterns []string) FileSeekerBuilder {
	fsb.patterns = patterns
	fsb.useRegExp = true
	return fsb
}

func (fsb *fileSeekerBuilder) FileExtensions(extensions []string) FileSeekerBuilder {
	fsb.fileExtensions = extensions
	return fsb
}

func (fsb *fileSeekerBuilder) ExcludeSubdirs() FileSeekerBuilder {
	fsb.includeSubdirs = false
	return fsb
}

func (fsb *fileSeekerBuilder) Build() FileSeeker {
	return &fileSeekerImpl{
		folderPath:     fsb.folderPath,
		patterns:       fsb.patterns,
		fileExtensions: fsb.fileExtensions,
		useRegExp:      fsb.useRegExp,
		includeSubdirs: fsb.includeSubdirs,
	}
}
