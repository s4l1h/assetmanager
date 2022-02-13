s4l1hs4l1hpackage main

import (
	"fmt"

	"github.com/akmyazilim/assetmanager"
	"github.com/akmyazilim/assetmanager/assetmanagerpongo"
	"github.com/flosch/pongo2"
)

func main() {
	// Create new assetmanager
	asset := assetmanager.New()
	// Virtual file name
	fileName := "module://index.html"
	// Add file to assetmanager from string
	asset.AddFileString(fileName, `{{hello}} from {{filename}}`)

	// Create new pongo2 asset manager loader
	loader := assetmanagerpongo.New(asset)
	// Add to pongo set
	pongoSet := pongo2.NewSet("assetmanager", loader)

	tmpl, err := pongoSet.FromCache(fileName)
	if err != nil {
		panic(err)
	}

	out, err := tmpl.Execute(pongo2.Context{"hello": "Hello Man", "filename": fileName})
	if err != nil {
		panic(err)
	}

	fmt.Println(out)

	// Output:
	/*
		Hello Man from module://index.html
	*/
}
