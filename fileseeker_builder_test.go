package fileseeker

import (
	"testing"
)

func TestNewFileSeekerBuilder(t *testing.T) {
	folderPath := "/path/to/folder"

	builder := NewFileSeekerBuilder(folderPath)

	if builder.(*fileSeekerBuilder).folderPath != folderPath {
		t.Errorf("Expected folderPath to be %s, but got %s", folderPath, builder.(*fileSeekerBuilder).folderPath)
	}
	if builder.(*fileSeekerBuilder).includeSubdirs != true {
		t.Error("Expected includeSubdirs to be true, but got false")
	}
	if len(builder.(*fileSeekerBuilder).patterns) != 0 {
		t.Errorf("Expected patterns to be empty, but got %v", builder.(*fileSeekerBuilder).patterns)
	}
	if len(builder.(*fileSeekerBuilder).fileExtensions) != 0 {
		t.Errorf("Expected fileExtensions to be empty, but got %v", builder.(*fileSeekerBuilder).fileExtensions)
	}
	if builder.(*fileSeekerBuilder).useRegExp != false {
		t.Error("Expected useRegExp to be false, but got true")
	}
}

func TestFileSeekerBuilder_Patterns(t *testing.T) {
	builder := &fileSeekerBuilder{}

	patterns := []string{"^file[0-9].txt$", "test.*"}
	builder.Patterns(patterns)

	if len(builder.patterns) != len(patterns) {
		t.Errorf("Expected %d patterns, but got %d", len(patterns), len(builder.patterns))
	}
	for i := range patterns {
		if builder.patterns[i] != patterns[i] {
			t.Errorf("Expected pattern: %s, but got: %s", patterns[i], builder.patterns[i])
		}
	}
	if builder.useRegExp != true {
		t.Error("Expected useRegExp to be true, but got false")
	}
}

func TestFileSeekerBuilder_FileExtensions(t *testing.T) {
	builder := &fileSeekerBuilder{}

	extensions := []string{"txt", "jpg"}
	builder.FileExtensions(extensions)

	if len(builder.fileExtensions) != len(extensions) {
		t.Errorf("Expected %d file extensions, but got %d", len(extensions), len(builder.fileExtensions))
	}
	for i := range extensions {
		if builder.fileExtensions[i] != extensions[i] {
			t.Errorf("Expected file extension: %s, but got: %s", extensions[i], builder.fileExtensions[i])
		}
	}
}

func TestFileSeekerBuilder_ExcludeSubdirs(t *testing.T) {
	builder := &fileSeekerBuilder{}

	builder.ExcludeSubdirs()

	if builder.includeSubdirs != false {
		t.Error("Expected includeSubdirs to be false, but got true")
	}
}

func TestFileSeekerBuilder_Build(t *testing.T) {
	folderPath := "/path/to/folder"
	builder := &fileSeekerBuilder{
		folderPath:     folderPath,
		patterns:       []string{"^file[0-9].txt$"},
		fileExtensions: []string{"jpg"},
		useRegExp:      true,
		includeSubdirs: true,
	}

	fileSeeker := builder.Build()

	fs, ok := fileSeeker.(*fileSeekerImpl)
	if !ok {
		t.Error("Expected fileSeeker to be of type *fileSeekerImpl")
	}
	if fs.folderPath != folderPath {
		t.Errorf("Expected folderPath to be %s, but got %s", folderPath, fs.folderPath)
	}
	if len(fs.patterns) != len(builder.patterns) {
		t.Errorf("Expected %d patterns, but got %d", len(builder.patterns), len(fs.patterns))
	}
	for i := range builder.patterns {
		if fs.patterns[i] != builder.patterns[i] {
			t.Errorf("Expected pattern: %s, but got: %s", builder.patterns[i], fs.patterns[i])
		}
	}
	if len(fs.fileExtensions) != len(builder.fileExtensions) {
		t.Errorf("Expected %d file extensions, but got %d", len(builder.fileExtensions), len(fs.fileExtensions))
	}
	for i := range builder.fileExtensions {
		if fs.fileExtensions[i] != builder.fileExtensions[i] {
			t.Errorf("Expected file extension: %s, but got: %s", builder.fileExtensions[i], fs.fileExtensions[i])
		}
	}
	if fs.useRegExp != builder.useRegExp {
		t.Errorf("Expected useRegExp to be %t, but got %t", builder.useRegExp, fs.useRegExp)
	}
	if fs.includeSubdirs != builder.includeSubdirs {
		t.Errorf("Expected includeSubdirs to be %t, but got %t", builder.includeSubdirs, fs.includeSubdirs)
	}
}
