package fileseeker

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileSeekerImpl_SeekFiles(t *testing.T) {
	testFolder := createTestFolder(t, "fileseeker_test")
	testSubDirOne := createTestFolder(t, "fileseeker_test/subdir")
	testSubDirTwo := createTestFolder(t, "fileseeker_test/subdir2")
	defer removeTestFolder(t, testFolder)

	createTestFile(t, testFolder, "file1.txt")
	createTestFile(t, testFolder, "file2.txt")
	createTestFile(t, testFolder, "file3.doc")
	createTestFile(t, testSubDirOne, "file4.txt")
	createTestFile(t, testSubDirOne, "file5.jpg")
	createTestFile(t, testSubDirTwo, "file6.txt")

	fileSeeker := &fileSeekerImpl{
		fileSeekerConfig{
			folderPath:     testFolder,
			patterns:       []string{"file[0-9].txt"},
			fileExtensions: []string{"jpg"},
			useRegExp:      true,
			includeSubdirs: true,
		},
	}

	files, err := fileSeeker.SeekFiles()
	if err != nil {
		t.Fatalf("Error seeking files: %v", err)
	}

	expectedFiles := []File{
		NewFile(testFolder + "/file1.txt"),
		NewFile(testFolder + "/file2.txt"),
		NewFile(testSubDirOne + "/file4.txt"),
		NewFile(testSubDirOne + "/file5.jpg"),
		NewFile(testSubDirTwo + "/file6.txt"),
	}
	assertFilesEqual(t, expectedFiles, files)
}

func assertFilesEqual(t *testing.T, expected []File, actual []File) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Errorf("Expected %d files, but got %d", len(expected), len(actual))
		return
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Expected file: %v, but got: %v", expected[i], actual[i])
		}
	}
}

func createTestFolder(t *testing.T, folderName string) string {
	t.Helper()
	folderPath := "./testdata/" + folderName
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Error creating test folder: %v", err)
	}
	return folderPath
}

func createTestFile(t *testing.T, folderPath, fileName string) {
	t.Helper()
	filePath := filepath.Join(folderPath, fileName)
	_, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
}

func removeTestFolder(t *testing.T, folderPath string) {
	t.Helper()
	err := os.RemoveAll(folderPath)
	if err != nil {
		t.Fatalf("Error removing test folder: %v", err)
	}
}
