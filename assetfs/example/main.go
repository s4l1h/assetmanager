package main

import (
	"net/http"
	"strings"

	"github.com/akmyazilim/assetmanager"
	"github.com/akmyazilim/assetmanager/assetfs"
)

// Then, run the main program and visit http://localhost:8080/static/views/ or http://localhost:8080/all/views/
func main() {

	// Create new asset manager
	asset := assetmanager.New()
	// add views and config directory
	asset.AddDir("./views")
	asset.AddDir("./config")

	// Create new asset manager
	views := assetmanager.New()
	// Add replacer
	views.AddReplacer("onlyViews", func(name string) string {
		//if filename doesn't contains a views.Remove it.
		if !strings.Contains(name, "views") {
			return "" // return empty name and automatic remove from assetmanager files
		}
		// if file name contains ".go". Remove it.
		if strings.Contains(name, ".go") {
			return "" // return empty name and automatic remove from assetmanager files
		}
		return name
	})

	// copy all files from asset to views
	views.Copy(asset) // add all files and run "onlyViews" replacer...

	assetFS := assetfs.New(asset)
	http.Handle("/all/", http.StripPrefix("/all/", http.FileServer(assetFS)))

	assetFSStatic := assetfs.New(views)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(assetFSStatic)))

	http.ListenAndServe(":8080", nil)

}
