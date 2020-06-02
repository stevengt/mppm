package util_test

import (
	"errors"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/util/utiltest"
)

func TestCopyFile(t *testing.T) {

	testCases := []*utiltest.CopyFileTestCase{
		&utiltest.CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "file2.bin",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      nil,
		},
		&utiltest.CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "file2.bin",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        errors.New("There was a problem opening the file."),
			CreateFileError:      nil,
		},
		&utiltest.CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "file2.bin",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      errors.New("There was a problem creating the file."),
		},
		&utiltest.CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "new-file",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      nil,
		},
		&utiltest.CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "new-file",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        errors.New("There was a problem opening the file."),
			CreateFileError:      nil,
		},
		&utiltest.CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "new-file",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      errors.New("There was a problem creating the file."),
		},
		&utiltest.CopyFileTestCase{
			SourceFileName:       "does-not-exist",
			TargetFileName:       "new-file",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      nil,
		},
		&utiltest.CopyFileTestCase{
			SourceFileName:       "empty-file.bin",
			TargetFileName:       "new-file",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      nil,
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGzipFile(t *testing.T) {

	testCases := []*utiltest.GzipFileTestCase{
		&utiltest.GzipFileTestCase{
			FileName:                       "file1.txt",
			ExpectedCompressedFileContents: []byte{0x4e, 0xb0, 0xa0, 0xe3, 0xf, 0x0, 0x0, 0x0},
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                nil,
		},
		&utiltest.GzipFileTestCase{
			FileName:                       "file1.txt",
			ExpectedCompressedFileContents: nil,
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  errors.New("There was a problem opening the file."),
			CreateFileError:                nil,
		},
		&utiltest.GzipFileTestCase{
			FileName:                       "file1.txt",
			ExpectedCompressedFileContents: nil,
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                errors.New("There was a problem creating the file."),
		},
		&utiltest.GzipFileTestCase{
			FileName:                       "file2.bin",
			ExpectedCompressedFileContents: []byte{0x5a, 0xa3, 0x9c, 0x7c, 0x4, 0x0, 0x0, 0x0},
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                nil,
		},
		&utiltest.GzipFileTestCase{
			FileName:                       "file2.bin",
			ExpectedCompressedFileContents: nil,
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  errors.New("There was a problem opening the file."),
			CreateFileError:                nil,
		},
		&utiltest.GzipFileTestCase{
			FileName:                       "file2.bin",
			ExpectedCompressedFileContents: nil,
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                errors.New("There was a problem creating the file."),
		},
		&utiltest.GzipFileTestCase{
			FileName:                       "empty-file.bin",
			ExpectedCompressedFileContents: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                nil,
		},
		&utiltest.GzipFileTestCase{
			FileName:                       "does-not-exist",
			ExpectedCompressedFileContents: nil,
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                nil,
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGunzipFile(t *testing.T) {

	testCases := []*utiltest.GunzipFileTestCase{
		&utiltest.GunzipFileTestCase{
			FileName:                         "file1.txt.gz",
			ExpectedUncompressedFileContents: append([]byte("file 1 contents"), 0xa),
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  nil,
		},
		&utiltest.GunzipFileTestCase{
			FileName:                         "file1.txt.gz",
			ExpectedUncompressedFileContents: nil,
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    errors.New("There was a problem opening the file."),
			CreateFileError:                  nil,
		},
		&utiltest.GunzipFileTestCase{
			FileName:                         "file1.txt.gz",
			ExpectedUncompressedFileContents: nil,
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  errors.New("There was a problem creating the file."),
		},
		&utiltest.GunzipFileTestCase{
			FileName:                         "does-not-exist.gz",
			ExpectedUncompressedFileContents: append([]byte("does-not-exist"), 0xa),
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  nil,
		},
		&utiltest.GunzipFileTestCase{
			FileName:                         "does-not-exist.gz",
			ExpectedUncompressedFileContents: nil,
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    errors.New("There was a problem opening the file."),
			CreateFileError:                  nil,
		},
		&utiltest.GunzipFileTestCase{
			FileName:                         "does-not-exist.gz",
			ExpectedUncompressedFileContents: nil,
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  errors.New("There was a problem creating the file."),
		},
		&utiltest.GunzipFileTestCase{
			FileName:                         "empty-file.bin",
			ExpectedUncompressedFileContents: make([]byte, 0),
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  nil,
		},
		&utiltest.GunzipFileTestCase{
			FileName:                         "does-not-exist",
			ExpectedUncompressedFileContents: nil,
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  nil,
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGetAllFileNamesWithExtension(t *testing.T) {

	fileExtensionsAndExpectedFileNames := map[string][]string{
		"txt":  []string{"file1.txt"},
		"bin":  []string{"file2.bin", "empty-file.bin"},
		"fake": []string{},
	}

	for fileExtension, expectedFileNames := range fileExtensionsAndExpectedFileNames {

		mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{}
		mockFileSystemDelegater.InitFiles(utiltest.GetTestFileNamesAndContents())
		util.FileSystemProxy = mockFileSystemDelegater

		actualFileNames, err := util.GetAllFileNamesWithExtension(fileExtension)
		sort.Strings(expectedFileNames)
		sort.Strings(actualFileNames)

		assert.Nil(t, err)
		assert.Exactly(t, expectedFileNames, actualFileNames)
	}

	for fileExtension, _ := range fileExtensionsAndExpectedFileNames {

		expectedFileNames := make([]string, 0)
		expectedError := errors.New("There was a problem while walking the filepath.")

		mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{
			WalkFilePathError: expectedError,
		}
		mockFileSystemDelegater.InitFiles(utiltest.GetTestFileNamesAndContents())
		util.FileSystemProxy = mockFileSystemDelegater

		actualFileNames, actualError := util.GetAllFileNamesWithExtension(fileExtension)
		assert.NotNil(t, actualError)
		assert.Exactly(t, expectedFileNames, actualFileNames)
		assert.Exactly(t, expectedError, actualError)

	}

}
