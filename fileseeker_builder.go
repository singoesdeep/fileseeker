package fileseeker

type FileSeekerBuilder interface {
	Patterns([]string) FileSeekerBuilder
	FileExtensions([]string) FileSeekerBuilder
	ExcludeSubdirs() FileSeekerBuilder
	Build() FileSeeker
}

type fileSeekerBuilder struct {
	fsc fileSeekerConfig
}

func NewFileSeekerBuilder(folderPath string) FileSeekerBuilder {
	return &fileSeekerBuilder{
		fileSeekerConfig{
			folderPath:     folderPath,
			includeSubdirs: true,
		},
	}
}

func (fsb *fileSeekerBuilder) Patterns(patterns []string) FileSeekerBuilder {
	fsb.fsc.patterns = patterns
	fsb.fsc.useRegExp = true
	return fsb
}

func (fsb *fileSeekerBuilder) FileExtensions(extensions []string) FileSeekerBuilder {
	fsb.fsc.fileExtensions = extensions
	return fsb
}

func (fsb *fileSeekerBuilder) ExcludeSubdirs() FileSeekerBuilder {
	fsb.fsc.includeSubdirs = false
	return fsb
}

func (fsb *fileSeekerBuilder) Build() FileSeeker {
	return &fileSeekerImpl{
		fileSeekerConfig{
			folderPath:     fsb.fsc.folderPath,
			patterns:       fsb.fsc.patterns,
			fileExtensions: fsb.fsc.fileExtensions,
			useRegExp:      fsb.fsc.useRegExp,
			includeSubdirs: fsb.fsc.includeSubdirs,
		},
	}
}
