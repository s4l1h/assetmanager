package assetmanagerpongo

import (
	"io"

	"github.com/s4l1h/assetmanager"
	"github.com/flosch/pongo2"
)

//New return new pongo2 template loader
func New(assetManager *assetmanager.AssetManager) pongo2.TemplateLoader {
	return &Pongo2Loader{AssetManager: assetManager}
}

// Pongo2Loader pongo2 template loader
type Pongo2Loader struct {
	AssetManager *assetmanager.AssetManager
}

// Abs pongo2 loader interface
func (m Pongo2Loader) Abs(base, name string) string {
	return name
}

// Get file
func (m Pongo2Loader) Get(path string) (io.Reader, error) {
	return m.AssetManager.File(path)
}
