package util

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetFilenameFromPath(t *testing.T) {
	testCases := []struct {
		inputPath        string
		expectedFilename string
	}{
		{"/path/to/file.txt", "file.txt"},
		{"dir/subdir/image.jpg", "image.jpg"},
		{"file", "file"},
		{"", ""},
	}

	for _, tc := range testCases {
		actualFilename := GetFilenameFromPath(tc.inputPath)
		if actualFilename != tc.expectedFilename {
			t.Errorf("Expected filename %q, but got %q for path %q", tc.expectedFilename, actualFilename, tc.inputPath)
		}
	}
}

func TestFilenameWithoutExt(t *testing.T) {
	testCases := []struct {
		filename         string
		expectedFilename string
	}{
		{"file.txt", "file"},
		{"dir/subdir/image.jpg", "image"},
		{"file", "file"},
		{".gitignore", ".gitignore"},
		{"", ""},
	}

	for _, tc := range testCases {
		actualFilename := GetFilenameWithoutExt(tc.filename)
		if actualFilename != tc.expectedFilename {
			t.Errorf("Expected filename %q, but got %q for filename %q", tc.expectedFilename, actualFilename, tc.filename)
		}
	}
}

func TestListAllFolder(t *testing.T) {
	// Create temporary test directory and files
	tempDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("Failed to create temporary test directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, f := range testFiles {
		if _, err := os.Create(filepath.Join(tempDir, f)); err != nil {
			t.Fatalf("Failed to create test file %q: %v", f, err)
		}
	}

	testSubdir := filepath.Join(tempDir, "subdir")
	if err := os.Mkdir(testSubdir, 0755); err != nil {
		t.Fatalf("Failed to create test subdir: %v", err)
	}

	// Define test cases
	testCases := []struct {
		path            string
		expectedFolders []string
		expectedError   error
	}{
		{tempDir, []string{"subdir"}, nil},
		{testSubdir, []string{}, nil},
	}

	// Run tests
	for _, tc := range testCases {
		actualFolders, actualError := ListAllFolder(tc.path)

		if len(actualFolders) == 0 && len(tc.expectedFolders) == 0 {
			continue
		}

		if !reflect.DeepEqual(actualFolders, tc.expectedFolders) {
			t.Errorf("Expected folders %v, but got %v for path %q", tc.expectedFolders, actualFolders, tc.path)
		}

		if !errors.Is(actualError, tc.expectedError) {
			t.Errorf("Expected error %v, but got %v for path %q", tc.expectedError, actualError, tc.path)
		}
	}
}

func TestListAllFiles(t *testing.T) {
	// Create temporary test directory and files
	tempDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("Failed to create temporary test directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
	for _, f := range testFiles {
		if _, err := os.Create(filepath.Join(tempDir, f)); err != nil {
			t.Fatalf("Failed to create test file %q: %v", f, err)
		}
	}

	testSubdir := filepath.Join(tempDir, "subdir")
	if err := os.Mkdir(testSubdir, 0755); err != nil {
		t.Fatalf("Failed to create test subdir: %v", err)
	}

	// Define test cases
	testCases := []struct {
		path          string
		expectedFiles []string
		expectedError error
	}{
		{tempDir, []string{"file1.txt", "file2.txt", "file3.txt"}, nil},
		{testSubdir, []string{}, nil},
		{"nonexistentdir", nil, errors.New("open nonexistentdir: no such file or directory")},
	}

	// Run tests
	for _, tc := range testCases {
		actualFiles, actualError := ListAllFiles(tc.path)

		if len(actualFiles) == 0 && len(tc.expectedFiles) == 0 {
			continue
		}

		if !reflect.DeepEqual(actualFiles, tc.expectedFiles) {
			t.Errorf("Expected files %v, but got %v for path %q", tc.expectedFiles, actualFiles, tc.path)
		}

		if !errors.Is(actualError, tc.expectedError) {
			t.Errorf("Expected error %v, but got %v for path %q", tc.expectedError, actualError, tc.path)
		}
	}
}
