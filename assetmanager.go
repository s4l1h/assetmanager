package assetmanager

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReplaceFunc file name replacer function
type ReplaceFunc func(name string) string

//AssetManager asset manager
type AssetManager struct {
	Files         map[string][]byte
	AllowedExt    map[string]bool
	DisallowedExt map[string]bool
	Replacers     map[string]ReplaceFunc
}

// New asset manager
func New() *AssetManager {
	return &AssetManager{
		Files:         make(map[string][]byte),
		AllowedExt:    make(map[string]bool),
		DisallowedExt: make(map[string]bool),
		Replacers:     make(map[string]ReplaceFunc),
	}
}

// AddReplacer add replacer function
func (manager *AssetManager) AddReplacer(name string, fn ReplaceFunc) {
	manager.Replacers[name] = fn
}

// RemoveReplacer remove replacer function
func (manager *AssetManager) RemoveReplacer(name string) {
	delete(manager.Replacers, name)
}

// ExistsReplacer check replacer function exists
func (manager *AssetManager) ExistsReplacer(name string) bool {
	_, ok := manager.Replacers[name]
	return ok
}

// AddAllowedExt add allowed extension
func (manager *AssetManager) AddAllowedExt(ext ...string) {
	for _, e := range ext {
		manager.AllowedExt[e] = true
	}
}

// RemoveAllowedExt remove allowed extension
func (manager *AssetManager) RemoveAllowedExt(ext string) {
	delete(manager.AllowedExt, ext)
}

// AddDisallowExt add disallow extension
func (manager *AssetManager) AddDisallowExt(ext ...string) {
	for _, e := range ext {
		manager.DisallowedExt[e] = true
	}
}

// RemoveDisallowExt remove disallow extension
func (manager *AssetManager) RemoveDisallowExt(ext string) {
	delete(manager.DisallowedExt, ext)
}

// CheckAllowed check file allowed and trusted file.
func (manager *AssetManager) CheckAllowed(ext string) bool {
	if len(manager.AllowedExt) == 0 {
		return true
	}
	if _, ok := manager.AllowedExt[ext]; ok {
		return true
	}
	return false
}

// CheckDisallowed check file disallowed and trusted file.
func (manager *AssetManager) CheckDisallowed(ext string) bool {
	if len(manager.DisallowedExt) == 0 {
		return true
	}
	if _, ok := manager.DisallowedExt[ext]; ok {
		return false
	}
	return true
}

// GetExt return file extension
func (manager *AssetManager) GetExt(file string) string {
	return filepath.Ext(file)
}

// CheckAddFile check file.
func (manager *AssetManager) CheckAddFile(file string) bool {
	ext := manager.GetExt(file)
	if manager.CheckAllowed(ext) && manager.CheckDisallowed(ext) {
		return true
	}
	return false
}

// AddFile add file
func (manager *AssetManager) AddFile(files ...string) {
	for _, file := range files {
		//logrus.Warn(file)
		if !manager.CheckAddFile(file) {
			continue
		}
		if b, er := ioutil.ReadFile(file); er == nil {
			manager.add(file, b)
			//manager.Files[file] = b
		}
	}
}

// AddFileString add file
func (manager *AssetManager) AddFileString(name, content string) {
	if manager.CheckAddFile(name) {
		manager.add(name, []byte(content))
		//manager.Files[name] = []byte(content)
	}
}

func (manager *AssetManager) add(name string, content []byte) {
	//logrus.Warn(len(manager.Replacers))
	if len(manager.Replacers) != 0 {
		for _, replacer := range manager.Replacers {
			//logrus.Warnf("%T", replacer)
			name = replacer(name)
		}
	}
	manager.Files[name] = content
}

// AddDir add directories
func (manager *AssetManager) AddDir(dirs ...string) {
	fileList := []string{}
	for _, dir := range dirs {
		//logrus.Warn("Add Dir", dir)
		if _, err := os.Stat(dir); err == nil {
			filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
				if !f.IsDir() {
					//logrus.Warnf("Add File %s ", path)
					fileList = append(fileList, path)
				}
				return nil
			})
		}
	}
	manager.AddFile(fileList...)
	//logrus.Warn(fileList)
}

// Get get file
func (manager *AssetManager) Get(file string) ([]byte, error) {
	if val, ok := manager.Files[file]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("File Not Found %s", file)
}

// GetString get file string
func (manager *AssetManager) GetString(file string) (string, error) {
	content, err := manager.Get(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Exists check file exists
func (manager *AssetManager) Exists(file string) bool {
	_, ok := manager.Files[file]
	return ok
}

// Merge another assetmanager
func (manager *AssetManager) Merge(asset *AssetManager) *AssetManager {
	if len(asset.Files) != 0 {
		for name, file := range asset.Files {
			manager.Files[name] = file
		}
	}
	return manager
}

// MergeAndRunReplacer merge another assetmanager and run replacer
func (manager *AssetManager) MergeAndRunReplacer(asset *AssetManager) *AssetManager {
	if len(asset.Files) != 0 {
		for name, file := range asset.Files {
			manager.add(name, file)
		}
	}
	return manager
}

// File io.Reader implement
func (manager *AssetManager) File(name string) (*bytes.Buffer, error) {
	b, e := manager.Get(name)
	return bytes.NewBuffer(b), e
}
