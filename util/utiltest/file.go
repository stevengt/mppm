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
		SetContentsFromHexString("1f8b08082fabe15e000366616b652e616c732e786d6c00b3714cca492dc9cfb3b3f1c92c4b0d4e2db1b3d147b06092007ceffd5b26000000")
}

func GetFakeUncompressedAbletonLiveSetFileBuilder() *MockFileBuilder {
	return NewMockFileBuilder().
		SetFilePath("fake-ableton-live-set.als.xml").
		SetContentsFromString("<Ableton><LiveSet></LiveSet></Ableton>")
}

func GetFakeAbletonLiveClipFileBuilder() *MockFileBuilder {
	return NewMockFileBuilder().
		SetFilePath("fake-ableton-live-clip.alc").
		SetContentsFromHexString("1f8b0808ecb2e15e000366616b652d616c632e786d6c006d8e4b0a803010438f34171806449776552fa075c0c1d26a5b3dbfa2d41f6e42022179587496937784b5acac391136a135632454d2cbe1092b5ec57039b4b217d5ae9ae7859de140585a99b4f5a99698ee44083ff6ecc067015ef3f0f885cc02171d64e00de4b13c5fba000000")
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
	Files                       map[string]*MockFile
	FileNamesAndContentsAsBytes map[string][]byte // Use this if you want the builder to create MockFile instances for you.
	FilesAsList                 []*MockFile       // A list of files with non-empty FilePath fields.
	MockFileBuilders            []*MockFileBuilder
	UseDefaultOpenFileError     bool
	UseDefaultCreateFileError   bool
	UseDefaultRemoveFileError   bool
	UseDefaultWalkFilePathError bool
	UseDefaultUserHomeDirError  bool
}

func NewMockFileSystemDelegaterBuilder() *MockFileSystemDelegaterBuilder {
	return &MockFileSystemDelegaterBuilder{
		Files:                       make(map[string]*MockFile),
		FileNamesAndContentsAsBytes: make(map[string][]byte),
		FilesAsList:                 make([]*MockFile, 0),
		MockFileBuilders:            make([]*MockFileBuilder, 0),
	}
}

func (builder *MockFileSystemDelegaterBuilder) SetFiles(files map[string]*MockFile) *MockFileSystemDelegaterBuilder {
	builder.Files = files
	return builder
}

func (builder *MockFileSystemDelegaterBuilder) SetFileNamesAndContentsAsBytes(fileNamesAndContentsAsBytes map[string][]byte) *MockFileSystemDelegaterBuilder {
	builder.FileNamesAndContentsAsBytes = fileNamesAndContentsAsBytes
	return builder
}

func (builder *MockFileSystemDelegaterBuilder) SetFilesAsList(filesAsList []*MockFile) *MockFileSystemDelegaterBuilder {
	builder.FilesAsList = filesAsList
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

	if builder.Files != nil {
		mockFileSystemDelegater.Files = builder.Files
	}

	if builder.FileNamesAndContentsAsBytes != nil {
		for fileName, fileContents := range builder.FileNamesAndContentsAsBytes {
			mockFileSystemDelegater.Files[fileName] = NewMockFileFromBytes(fileName, fileContents)
		}
	}

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
	err = mockFileSystemDelegater.OpenFileError
	if err == nil {
		var doesFileExist bool
		file, doesFileExist = mockFileSystemDelegater.Files[fileName]
		if !doesFileExist {
			err = errors.New("Unable to open file " + fileName)
		}
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) CreateFile(fileName string) (file io.ReadWriteCloser, err error) {
	err = mockFileSystemDelegater.CreateFileError
	if err == nil {
		if mockFileSystemDelegater.Files[fileName] != nil {
			err = mockFileSystemDelegater.RemoveFile(fileName)
			if err != nil {
				return
			}
		}
		fileContents := make([]byte, 0)
		mockFileSystemDelegater.Files[fileName] = NewMockFileFromBytes(fileName, fileContents)
		file = mockFileSystemDelegater.Files[fileName]
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) RemoveFile(fileName string) (err error) {
	err = mockFileSystemDelegater.RemoveFileError
	if err == nil {
		delete(mockFileSystemDelegater.Files, fileName)
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) WalkFilePath(root string, walkFn filepath.WalkFunc) (err error) {
	err = mockFileSystemDelegater.WalkFilePathError
	if err == nil {
		for fileName, _ := range mockFileSystemDelegater.Files {
			err = walkFn(fileName, nil, nil)
			if err != nil {
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
