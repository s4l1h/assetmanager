package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/akmyazilim/assetmanager"
)

var asset *assetmanager.AssetManager

// this function generate assetFile "generatedAsset.go"
// and call with "go run main.go --build yes"
func generateAsset() {
	fmt.Println("generateAsset called")
	assetmanager.Generate(
		assetmanager.GenerateOPT{
			File:      "./generatedAsset.go",
			Namespace: "main",
			Asset:     asset,
			CacheKey:  "mainAsset", // you can use multiple build file (mainAsset)
		},
	)
}
func init() {
	fmt.Println("main init")
	//Normal asset manager
	asset = assetmanager.New()
	// replacer function
	asset.AddReplacer("renamer", func(name string) string {
		return strings.Replace(name, "../test/", "", -1)
	})
	// add test directory to assetmanager
	asset.AddDir("../test")

	//fmt.Println(assetmanager.GeneratedCache)
	// if assetmanager have cached object use this. (mainAsset)
	// if have generatedAsset.go file
	if val, ok := assetmanager.GeneratedCache["mainAsset"]; ok {
		asset.Copy(val)
	}

}

func main() {

	build := flag.String("build", "no", "is build")
	flag.Parse()
	// go run main.go --build yes
	// if build == yes
	if *build == "yes" {
		generateAsset()
	}

	fmt.Println(asset.Files)

}
