package assetfs

/*
Thanks Jaana Burcu Doğan https://github.com/rakyll/statik/blob/master/fs/fs.go
*/
import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/akmyazilim/assetmanager"
)

// New assetfs
func New(manager *assetmanager.AssetManager) http.FileSystem {
	return &statikFS{manager: manager}
}

// file holds unzipped read-only file contents and file metadata.
type file struct {
	name string
	//os.FileInfo
	data []byte
}

type statikFS struct {
	manager *assetmanager.AssetManager
}

// Open returns a file matching the given file name, or os.ErrNotExists if
// no file matching the given file name is found in the archive.
// If a directory is requested, Open returns the file named "index.html"
// in the requested directory, if that file exists.
func (fs *statikFS) Open(name string) (http.File, error) {
	// litter.Dump(name)

	name = strings.Replace(name, "//", "/", -1)
	if name[0] == '/' {
		name = name[1:]
	}
	// litter.Dump(name)

	f, err := fs.manager.Get(name)
	//litter.Dump(err.Error())
	if err == nil {
		return newHTTPFile(file{data: f, name: name}, false), nil
	}
	// The file doesn't match, but maybe it's a directory,
	// thus we should look for index.html
	indexName := strings.Replace(name+"/index.html", "//", "/", -1)
	f, err = fs.manager.Get(indexName)
	// litter.Dump(indexName)
	if err != nil {
		return nil, os.ErrNotExist
	}
	return newHTTPFile(file{data: f, name: indexName}, true), nil
}

func newHTTPFile(file file, isDir bool) *httpFile {
	return &httpFile{
		file:   file,
		reader: bytes.NewReader(file.data),
		isDir:  isDir,
	}
}

// httpFile represents an HTTP file and acts as a bridge
// between file and http.File.
type httpFile struct {
	file

	reader *bytes.Reader
	isDir  bool
}

// Read reads bytes into p, returns the number of read bytes.
func (f *httpFile) Name() string {
	return f.file.name
}

// Read reads bytes into p, returns the number of read bytes.
func (f *httpFile) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}

// Seek seeks to the offset.
func (f *httpFile) Seek(offset int64, whence int) (ret int64, err error) {
	return f.reader.Seek(offset, whence)
}

// Stat stats the file.
func (f *httpFile) Stat() (os.FileInfo, error) {
	return f, nil
}

// IsDir returns true if the file location represents a directory.
func (f *httpFile) IsDir() bool {
	return f.isDir
}

// Readdir returns an empty slice of files, directory
// listing is disabled.
func (f *httpFile) Readdir(count int) ([]os.FileInfo, error) {
	// directory listing is disabled.
	return make([]os.FileInfo, 0), nil
}

func (f *httpFile) Close() error {
	return nil
}
func (f *httpFile) Mode() os.FileMode {
	return 0777
}

func (f *httpFile) ModTime() time.Time {
	return time.Now()
}
func (f *httpFile) Size() int64 {
	return int64(len(f.file.data))
}
func (f *httpFile) Sys() interface{} {
	return nil
}
