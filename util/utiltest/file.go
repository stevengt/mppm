package utiltest

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"io"
	"path/filepath"
	"strings"

	"github.com/stevengt/mppm/util"
)

// ------------------------------------------------------------------------------

var DefaultOpenFileError error = errors.New("There was a problem opening the file.")

var DefaultCreateFileError error = errors.New("There was a problem creating the file.")

var DefaultRenameFileError error = errors.New("There was a problem renaming the file.")

var DefaultRemoveFileError error = errors.New("There was a problem removing the file.")

var DefaultWalkFilePathError error = errors.New("There was a problem walking the file path.")

var DefaultUserHomeDirError error = errors.New("There was a problem getting the user's home directory.")

// ------------------------------------------------------------------------------

func GetPlainTextFileBuilder() *MockFileBuilder {
	return NewMockFileBuilder().
		SetFilePath("plain-text-file.txt").
		SetContentsFromString("file 1 contents")
}

func GetGzippedPlainTextFileBuilder() *MockFileBuilder {
	return NewMockFileBuilder().
		SetFilePath("plain-text-file.txt.gz").
		SetContentsFromHexString("1f8b08000000000000ff4acbcc4955305448cecf2b49cd2b2906040000ffff4eb0a0e30f000000")
}

func GetBinaryContaining0xdeadbeefFileBuilder() *MockFileBuilder {
	return NewMockFileBuilder().
		SetFilePath("deadbeef.bin").
		SetContentsFromHexString("deadbeef")
}

func GetEmptyFileBuilder() *MockFileBuilder {
	return NewMockFileBuilder().
		SetFilePath("empty-file.bin").
		SetContentsFromBytes(make([]byte, 0))
}

func GetFakeAbletonLiveSetFileBuilder() *MockFileBuilder {
	return NewMockFileBuilder().
		SetFilePath("fake-ableton-live-set.als").
		SetContentsFromHexString("1f8b08000000000000ffb2714cca492dc9cfb3b3f1c92c4b0d4e2db1b3d147b0609280000000ffff7ceffd5b26000000")
}

func GetFakeUncompressedAbletonLiveSetFileBuilder() *MockFileBuilder {
	return NewMockFileBuilder().
		SetFilePath("fake-ableton-live-set.als.xml").
		SetContentsFromString("<Ableton><LiveSet></LiveSet></Ableton>")
}

func GetFakeAbletonLiveClipFileBuilder() *MockFileBuilder {
	return NewMockFileBuilder().
		SetFilePath("fake-ableton-live-clip.alc").
		SetContentsFromHexString("1f8b08000000000000ff6c8e4b0ac24010448fd417281a242e93d57881712cb0709868d2e6fc82127f64f71a9a570fbb63658ccdd16b6162380e532e97d931e8a4273bf65c54d89db39a63c86a89b73b5be1e4e8aaaea98ed16b8ecfe5b00d7cfdd89fc17ef4f6b56b6b8bbdeb6c0d7e040000ffffe4b13c5fba000000")
}

func GetFakeUncompressedAbletonLiveClipFileBuilder() *MockFileBuilder {
	return NewMockFileBuilder().
		SetFilePath("fake-ableton-live-clip.alc.xml").
		SetContentsFromString("<Ableton><LiveSet><Tracks><MidiTrack><DeviceChain><MainSequencer><ClipSlotList><ClipSlot></ClipSlot></ClipSlotList></MainSequencer></DeviceChain></MidiTrack></Tracks></LiveSet></Ableton>")
}

// ------------------------------------------------------------------------------

func GetTestFileNamesAndContents() map[string][]byte {
	return map[string][]byte{
		"plain-text-file.txt": []byte("file 1 contents"),
		"plain-text-file.txt.gz": []byte{
			0x1f, 0x8b, 0x08, 0x08, 0x8a, 0xab, 0xd4, 0x5e, 0x00, 0x03,
			0x74, 0x65, 0x73, 0x74, 0x00, 0x4b, 0xcb, 0xcc, 0x49, 0x55,
			0x30, 0x54, 0x48, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x29,
			0xe6, 0x02, 0x00, 0xb4, 0xca, 0x50, 0xa3, 0x10, 0x00, 0x00, 0x00,
		},
		"file2.bin":      []byte{0xDE, 0xAD, 0xBE, 0xEF},
		"empty-file.bin": make([]byte, 0),
		"does-not-exist.gz": []byte{
			0x1f, 0x8b, 0x08, 0x08, 0x96, 0xa8, 0xd4, 0x5e, 0x00, 0x03,
			0x74, 0x65, 0x73, 0x74, 0x00, 0x4b, 0xc9, 0x4f, 0x2d, 0xd6,
			0xcd, 0xcb, 0x2f, 0xd1, 0x4d, 0xad, 0xc8, 0x2c, 0x2e, 0xe1,
			0x02, 0x00, 0x2a, 0x53, 0xd8, 0x28, 0x0f, 0x00, 0x00, 0x00,
		},
		"ableton-project.als": []byte{
			0x1f, 0x8b, 0x08, 0x08, 0x8a, 0xab, 0xd4, 0x5e, 0x00, 0x03,
			0x74, 0x65, 0x73, 0x74, 0x00, 0x4b, 0xcb, 0xcc, 0x49, 0x55,
			0x30, 0x54, 0x48, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x29,
			0xe6, 0x02, 0x00, 0xb4, 0xca, 0x50, 0xa3, 0x10, 0x00, 0x00, 0x00,
		},
	}
}

func GetMockFileSystemDelegaterFromBuilderOrNil(mockFileSystemDelegaterBuilder *MockFileSystemDelegaterBuilder) *MockFileSystemDelegater {
	if mockFileSystemDelegaterBuilder != nil {
		return mockFileSystemDelegaterBuilder.Build()
	} else {
		return NewMockFileSystemDelegater()
	}
}

// ------------------------------------------------------------------------------

type MockFileSystemDelegaterBuilder struct {
	FilesAsList                 []*MockFile // A list of files with non-empty FilePath fields.
	MockFileBuilders            []*MockFileBuilder
	UseDefaultOpenFileError     bool
	UseDefaultCreateFileError   bool
	UseDefaultRenameFileError   bool
	UseDefaultRemoveFileError   bool
	UseDefaultWalkFilePathError bool
	UseDefaultUserHomeDirError  bool
}

func NewMockFileSystemDelegaterBuilder() *MockFileSystemDelegaterBuilder {
	return &MockFileSystemDelegaterBuilder{
		FilesAsList:      make([]*MockFile, 0),
		MockFileBuilders: make([]*MockFileBuilder, 0),
	}
}

func (builder *MockFileSystemDelegaterBuilder) SetFilesAsList(files ...*MockFile) *MockFileSystemDelegaterBuilder {
	builder.FilesAsList = files
	return builder
}

func (builder *MockFileSystemDelegaterBuilder) SetMockFileBuilders(mockFileBuilders ...*MockFileBuilder) *MockFileSystemDelegaterBuilder {
	builder.MockFileBuilders = mockFileBuilders
	return builder
}

func (builder *MockFileSystemDelegaterBuilder) SetUseDefaultOpenFileError(useDefaultOpenFileError bool) *MockFileSystemDelegaterBuilder {
	builder.UseDefaultOpenFileError = useDefaultOpenFileError
	return builder
}

func (builder *MockFileSystemDelegaterBuilder) SetUseDefaultCreateFileError(useDefaultCreateFileError bool) *MockFileSystemDelegaterBuilder {
	builder.UseDefaultCreateFileError = useDefaultCreateFileError
	return builder
}

func (builder *MockFileSystemDelegaterBuilder) SetUseDefaultRenameFileError(useDefaultRenameFileError bool) *MockFileSystemDelegaterBuilder {
	builder.UseDefaultRenameFileError = useDefaultRenameFileError
	return builder
}

func (builder *MockFileSystemDelegaterBuilder) SetUseDefaultRemoveFileError(useDefaultRemoveFileError bool) *MockFileSystemDelegaterBuilder {
	builder.UseDefaultRemoveFileError = useDefaultRemoveFileError
	return builder
}

func (builder *MockFileSystemDelegaterBuilder) SetUseDefaultWalkFilePathError(useDefaultWalkFilePathError bool) *MockFileSystemDelegaterBuilder {
	builder.UseDefaultWalkFilePathError = useDefaultWalkFilePathError
	return builder
}

func (builder *MockFileSystemDelegaterBuilder) SetUseDefaultUserHomeDirError(useDefaultUserHomeDirError bool) *MockFileSystemDelegaterBuilder {
	builder.UseDefaultUserHomeDirError = useDefaultUserHomeDirError
	return builder
}

func (builder *MockFileSystemDelegaterBuilder) Build() *MockFileSystemDelegater {

	mockFileSystemDelegater := NewMockFileSystemDelegater()

	if builder.FilesAsList != nil {
		for _, mockFile := range builder.FilesAsList {
			mockFileSystemDelegater.Files[mockFile.FilePath] = mockFile
		}
	}

	if builder.MockFileBuilders != nil {
		for _, mockFileBuilder := range builder.MockFileBuilders {
			mockFile := mockFileBuilder.Build()
			mockFileSystemDelegater.Files[mockFile.FilePath] = mockFile
		}
	}

	if builder.UseDefaultOpenFileError {
		mockFileSystemDelegater.OpenFileError = DefaultOpenFileError
	}

	if builder.UseDefaultCreateFileError {
		mockFileSystemDelegater.CreateFileError = DefaultCreateFileError
	}

	if builder.UseDefaultRenameFileError {
		mockFileSystemDelegater.RenameFileError = DefaultRenameFileError
	}

	if builder.UseDefaultRemoveFileError {
		mockFileSystemDelegater.RemoveFileError = DefaultRemoveFileError
	}

	if builder.UseDefaultWalkFilePathError {
		mockFileSystemDelegater.WalkFilePathError = DefaultWalkFilePathError
	}

	if builder.UseDefaultUserHomeDirError {
		mockFileSystemDelegater.UserHomeDirError = DefaultUserHomeDirError
	}

	return mockFileSystemDelegater

}

// ------------------------------------------------------------------------------

type MockFileSystemDelegater struct {
	Files             map[string]*MockFile // Map of file names to mocked file instances.
	OpenFileError     error
	CreateFileError   error
	RenameFileError   error
	RemoveFileError   error
	WalkFilePathError error
	UserHomeDirError  error
}

func NewMockFileSystemDelegater() *MockFileSystemDelegater {
	return &MockFileSystemDelegater{
		Files: make(map[string]*MockFile),
	}
}

func (mockFileSystemDelegater *MockFileSystemDelegater) Init() {
	util.FileSystemProxy = mockFileSystemDelegater
}

func (mockFileSystemDelegater *MockFileSystemDelegater) InitFiles(fileNamesAndContents map[string][]byte) {
	files := make(map[string]*MockFile)
	for fileName, fileContents := range fileNamesAndContents {
		files[fileName] = NewMockFileFromBytes(fileName, fileContents)
	}
	mockFileSystemDelegater.Files = files
}

func (mockFileSystemDelegater *MockFileSystemDelegater) GetMockFileAndContentsIfFileExistsElseReturnNil(fileName string) (file *MockFile, contents []byte) {
	if mockFileSystemDelegater.DoesFileExist(fileName) {
		file = mockFileSystemDelegater.Files[fileName]
		contents = file.Contents
		return
	}
	return nil, nil
}

func (mockFileSystemDelegater *MockFileSystemDelegater) OpenFile(fileName string) (file io.ReadWriteCloser, err error) {
	if err = mockFileSystemDelegater.OpenFileError; err == nil {
		var doesFileExist bool
		file, doesFileExist = mockFileSystemDelegater.Files[fileName]
		if !doesFileExist {
			err = errors.New("Unable to open file " + fileName)
		}
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) CreateFile(fileName string) (file io.ReadWriteCloser, err error) {
	if err = mockFileSystemDelegater.CreateFileError; err == nil {
		if mockFileSystemDelegater.Files[fileName] != nil {
			if err = mockFileSystemDelegater.RemoveFile(fileName); err != nil {
				return
			}
		}
		fileContents := make([]byte, 0)
		mockFileSystemDelegater.Files[fileName] = NewMockFileFromBytes(fileName, fileContents)
		file = mockFileSystemDelegater.Files[fileName]
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) RenameFile(fileName string, newFileName string) (err error) {
	if err = mockFileSystemDelegater.RenameFileError; err == nil {
		mockFileSystemDelegater.Files[newFileName] = mockFileSystemDelegater.Files[fileName]
		delete(mockFileSystemDelegater.Files, fileName)
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) RemoveFile(fileName string) (err error) {
	if err = mockFileSystemDelegater.RemoveFileError; err == nil {
		delete(mockFileSystemDelegater.Files, fileName)
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) WalkFilePath(root string, walkFn filepath.WalkFunc) (err error) {
	if err = mockFileSystemDelegater.WalkFilePathError; err == nil {
		for fileName, _ := range mockFileSystemDelegater.Files {
			if err = walkFn(fileName, nil, nil); err != nil {
				return
			}
		}
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) UserHomeDir() (string, error) {
	err := mockFileSystemDelegater.UserHomeDirError
	if err == nil {
		return "/home/testuser", err
	}
	return "", err
}

func (mockFileSystemDelegater *MockFileSystemDelegater) JoinFilePath(elem ...string) string {
	return strings.Join(elem, "/")
}

func (mockFileSystemDelegater *MockFileSystemDelegater) DoesFileExist(filePath string) bool {
	var doesFileExist bool
	_, doesFileExist = mockFileSystemDelegater.Files[filePath]
	return doesFileExist
}

// ------------------------------------------------------------------------------

type MockFileBuilder struct {
	FilePath  string
	Contents  []byte
	WasClosed bool
}

func NewMockFileBuilder() *MockFileBuilder {
	return &MockFileBuilder{
		WasClosed: false,
	}
}

func (builder *MockFileBuilder) SetFilePath(filePath string) *MockFileBuilder {
	builder.FilePath = filePath
	return builder
}

func (builder *MockFileBuilder) SetContentsFromBytes(contents []byte) *MockFileBuilder {
	builder.Contents = contents
	return builder
}

func (builder *MockFileBuilder) SetContentsFromString(contents string) *MockFileBuilder {
	builder.Contents = []byte(contents)
	return builder
}

func (builder *MockFileBuilder) SetContentsFromHexString(contents string) *MockFileBuilder {
	contentsAsBytes, err := hex.DecodeString(contents)
	if err != nil {
		util.ExitWithError(err)
	}
	builder.Contents = contentsAsBytes
	return builder
}

func (builder *MockFileBuilder) SetWasClosed(wasClosed bool) *MockFileBuilder {
	builder.WasClosed = wasClosed
	return builder
}

func (builder *MockFileBuilder) Build() *MockFile {
	mockFile := &MockFile{
		FilePath:  builder.FilePath,
		Contents:  builder.Contents,
		WasClosed: builder.WasClosed,
	}
	mockFile.resetBuffer()
	return mockFile
}

// ------------------------------------------------------------------------------

type MockFile struct {
	FilePath         string
	Contents         []byte
	bufferReadWriter *bufio.ReadWriter
	WasClosed        bool
}

func NewMockFileFromBytes(filePath string, contents []byte) *MockFile {
	mockFile := &MockFile{
		FilePath:  filePath,
		Contents:  contents,
		WasClosed: false,
	}
	mockFile.resetBuffer()
	return mockFile
}

func NewMockFileFromString(filePath string, contents string) *MockFile {
	return NewMockFileFromBytes(filePath, []byte(contents))
}

func NewMockFileFromHexString(filePath string, contents string) *MockFile {
	contentsAsBytes, err := hex.DecodeString(contents)
	if err != nil {
		util.ExitWithError(err)
	}
	return NewMockFileFromBytes(filePath, contentsAsBytes)
}

func (mockFile *MockFile) Read(p []byte) (n int, err error) {
	return mockFile.bufferReadWriter.Read(p)
}

func (mockFile *MockFile) Write(p []byte) (n int, err error) {
	mockFile.Contents = append(mockFile.Contents, p...)
	return mockFile.bufferReadWriter.Write(p)
}

func (mockFile *MockFile) Close() error {
	mockFile.WasClosed = true
	mockFile.resetBuffer()
	return nil
}

func (mockFile *MockFile) resetBuffer() {
	buffer := bytes.NewBuffer(mockFile.Contents)
	bufferReader := bufio.NewReader(buffer)
	bufferWriter := bufio.NewWriter(buffer)
	mockFile.bufferReadWriter = bufio.NewReadWriter(bufferReader, bufferWriter)
}
