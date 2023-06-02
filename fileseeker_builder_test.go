package fileseeker

import (
	"testing"
)

func TestNewFileSeekerBuilder(t *testing.T) {
	folderPath := "/path/to/folder"

	builder := NewFileSeekerBuilder(folderPath)

	if builder.(*fileSeekerBuilder).fsc.folderPath != folderPath {
		t.Errorf("Expected folderPath to be %s, but got %s", folderPath, builder.(*fileSeekerBuilder).fsc.folderPath)
	}
	if builder.(*fileSeekerBuilder).fsc.includeSubdirs != true {
		t.Error("Expected includeSubdirs to be true, but got false")
	}
	if len(builder.(*fileSeekerBuilder).fsc.patterns) != 0 {
		t.Errorf("Expected patterns to be empty, but got %v", builder.(*fileSeekerBuilder).fsc.patterns)
	}
	if len(builder.(*fileSeekerBuilder).fsc.fileExtensions) != 0 {
		t.Errorf("Expected fileExtensions to be empty, but got %v", builder.(*fileSeekerBuilder).fsc.fileExtensions)
	}
	if builder.(*fileSeekerBuilder).fsc.useRegExp != false {
		t.Error("Expected useRegExp to be false, but got true")
	}
}

func TestFileSeekerBuilder_Patterns(t *testing.T) {
	builder := &fileSeekerBuilder{}

	patterns := []string{"^file[0-9].txt$", "test.*"}
	builder.Patterns(patterns)

	if len(builder.fsc.patterns) != len(patterns) {
		t.Errorf("Expected %d patterns, but got %d", len(patterns), len(builder.fsc.patterns))
	}
	for i := range patterns {
		if builder.fsc.patterns[i] != patterns[i] {
			t.Errorf("Expected pattern: %s, but got: %s", patterns[i], builder.fsc.patterns[i])
		}
	}
	if builder.fsc.useRegExp != true {
		t.Error("Expected useRegExp to be true, but got false")
	}
}

func TestFileSeekerBuilder_FileExtensions(t *testing.T) {
	builder := &fileSeekerBuilder{}

	extensions := []string{"txt", "jpg"}
	builder.FileExtensions(extensions)

	if len(builder.fsc.fileExtensions) != len(extensions) {
		t.Errorf("Expected %d file extensions, but got %d", len(extensions), len(builder.fsc.fileExtensions))
	}
	for i := range extensions {
		if builder.fsc.fileExtensions[i] != extensions[i] {
			t.Errorf("Expected file extension: %s, but got: %s", extensions[i], builder.fsc.fileExtensions[i])
		}
	}
}

func TestFileSeekerBuilder_ExcludeSubdirs(t *testing.T) {
	builder := &fileSeekerBuilder{}

	builder.ExcludeSubdirs()

	if builder.fsc.includeSubdirs != false {
		t.Error("Expected includeSubdirs to be false, but got true")
	}
}

func TestFileSeekerBuilder_Build(t *testing.T) {
	folderPath := "/path/to/folder"
	builder := &fileSeekerBuilder{
		fileSeekerConfig{folderPath: folderPath,
			patterns:       []string{"^file[0-9].txt$"},
			fileExtensions: []string{"jpg"},
			useRegExp:      true,
			includeSubdirs: true,
		},
	}

	fileSeeker := builder.Build()

	fs, ok := fileSeeker.(*fileSeekerImpl)
	if !ok {
		t.Error("Expected fileSeeker to be of type *fileSeekerImpl")
	}
	if fs.fsc.folderPath != folderPath {
		t.Errorf("Expected folderPath to be %s, but got %s", folderPath, fs.fsc.folderPath)
	}
	if len(fs.fsc.patterns) != len(builder.fsc.patterns) {
		t.Errorf("Expected %d patterns, but got %d", len(builder.fsc.patterns), len(fs.fsc.patterns))
	}
	for i := range builder.fsc.patterns {
		if fs.fsc.patterns[i] != builder.fsc.patterns[i] {
			t.Errorf("Expected pattern: %s, but got: %s", builder.fsc.patterns[i], fs.fsc.patterns[i])
		}
	}
	if len(fs.fsc.fileExtensions) != len(builder.fsc.fileExtensions) {
		t.Errorf("Expected %d file extensions, but got %d", len(builder.fsc.fileExtensions), len(fs.fsc.fileExtensions))
	}
	for i := range builder.fsc.fileExtensions {
		if fs.fsc.fileExtensions[i] != builder.fsc.fileExtensions[i] {
			t.Errorf("Expected file extension: %s, but got: %s", builder.fsc.fileExtensions[i], fs.fsc.fileExtensions[i])
		}
	}
	if fs.fsc.useRegExp != builder.fsc.useRegExp {
		t.Errorf("Expected useRegExp to be %t, but got %t", builder.fsc.useRegExp, fs.fsc.useRegExp)
	}
	if fs.fsc.includeSubdirs != builder.fsc.includeSubdirs {
		t.Errorf("Expected includeSubdirs to be %t, but got %t", builder.fsc.includeSubdirs, fs.fsc.includeSubdirs)
	}
}
