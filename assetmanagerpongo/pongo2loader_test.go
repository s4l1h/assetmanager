package assetmanagerpongo

import (
	"testing"

	"github.com/akmyazilim/assetmanager"
	"github.com/flosch/pongo2"
)

func Test(t *testing.T) {
	expected := "from test.html"
	// Create new assetmanager
	asset := assetmanager.New()
	// Virtual file name
	fileName := "test.html"
	// Add file to assetmanager from string
	asset.AddFileString(fileName, `from {{filename}}`)

	// Create new pongo2 asset manager loader
	loader := New(asset)
	// Add to pongo set
	pongoSet := pongo2.NewSet("assetmanager", loader)

	tmpl, err := pongoSet.FromCache(fileName)
	if err != nil {
		panic(err)
	}

	out, err := tmpl.Execute(pongo2.Context{"filename": fileName})
	if err != nil {
		t.Error(err)
	}
	if out != expected {
		t.Errorf("%s not equal expected : %s", out, expected)
	}
}
