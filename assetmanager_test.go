s4l1hs4l1hpackage assetmanager_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/akmyazilim/assetmanager"
)

func ExampleAssetManager_File() {

	asset := assetmanager.New()
	asset.AddDir("./test")
	buffer, err := asset.File("test/index.html")
	if err != nil {
		fmt.Println(err)
	}
	r, _ := ioutil.ReadAll(buffer)
	fmt.Println(string(r))
	// Output: index.html
}
func ExampleAssetManager_AddReplacer() {
	asset := assetmanager.New()
	// Add Replacer
	asset.AddReplacer("testReplacer", func(name string) string {
		return fmt.Sprintf("main:%s", name)
	})
	// Add Another Replacer.
	asset.AddReplacer("convergohtml", func(name string) string {
		return strings.Replace(name, ".html", ".gohtml", -1)
	})
	// add file from string
	asset.AddFileString("index.html", "test content")
	// get file content
	c, _ := asset.GetString("main:index.gohtml")
	fmt.Println(c)
	// Output: test content
}
func replaceMe(name string) string {
	return name
}

func ExampleAssetManager_Merge() {
	asset := assetmanager.New()
	asset.AddFileString("index.html", "index.html content")
	asset.AddFileString("asset1.html", "asset1.html content")
	asset2 := assetmanager.New()
	asset2.AddFileString("asset2.html", "asset2.html content")
	asset.Merge(asset2)
	c, _ := asset.GetString("asset2.html")
	fmt.Println(c)
	// Output: asset2.html content
}
func ExampleAssetManager_MergeAndRunReplacer() {
	asset := assetmanager.New()
	asset.AddReplacer("ToUpper", func(name string) string {
		return strings.ToUpper(name)
	})
	asset.AddFileString("index.html", "index.html content")
	asset.AddFileString("asset1.html", "asset1.html content")

	asset2 := assetmanager.New()
	asset2.AddFileString("asset2.html", "asset2.html content")
	asset.MergeAndRunReplacer(asset2)
	c, _ := asset.GetString("ASSET2.HTML")
	fmt.Println(c)
	// Output: asset2.html content
}
func TestReplacer(t *testing.T) {

	fileName := "test/index.html"
	fileContent := "index.html content"

	newFileName := "main::test/index.html"

	asset := assetmanager.New()
	// Add Replace Function
	asset.AddReplacer("replacerTestName", replaceMe)

	if asset.ExistsReplacer("replacerTestName") == false {
		t.Error("AddReplacer Error")
	}
	asset.RemoveReplacer("replacerTestName")

	if asset.ExistsReplacer("replacerTestName") == true {
		t.Error("RemoveReplacer Error")
	}

	// Add another Replacer Function
	asset.AddReplacer("testToMain", func(name string) string {
		name = strings.Replace(name, "test/", "main::test/", -1)
		return name
	})

	// Add File
	asset.AddFileString(fileName, fileContent)
	// Get it with the new name.
	r, e := asset.GetString(newFileName)
	if e != nil {
		t.Error(e)
	}
	if r != fileContent {
		t.Errorf("File content not equal %s", fileContent)
	}
}
func TestExt(t *testing.T) {
	asset := assetmanager.New()
	asset.AddAllowedExt(".html", ".tpl")
	asset.AddDisallowExt(".php", ".jpg")

	ext1 := asset.GetExt("index.html")
	ext2 := asset.GetExt("index.php")

	if ext1 != ".html" {
		t.Errorf("getExt Error ")
	}
	if !asset.CheckAllowed(ext1) {
		t.Errorf("CheckAllowed Error ")
	}
	if asset.CheckAllowed(ext2) {
		t.Errorf("CheckAllowed Error ")
	}

	if !asset.CheckDisallowed(ext1) {
		t.Errorf("CheckDisallowed Error: CheckDisallowed ")
	}
	if asset.CheckDisallowed(ext2) {
		t.Errorf("CheckDisallowed Error: CheckDisallowed ")
	}

	if asset.CheckAddFile("index.php") {
		t.Errorf("canAddFile Error: canAddFile index.php")
	}
	if !asset.CheckAddFile("index.html") {
		t.Errorf("canAddFile Error: canAddFile index.html")
	}

}

func TestAddRemoveAllowed(t *testing.T) {
	asset := assetmanager.New()
	asset.AddAllowedExt(".html", ".tpl", ".php")

	if asset.CheckAllowed(".html") != true {
		t.Errorf("RemoveAllowedExt Error: CheckAllowed ")
	}

	asset.RemoveAllowedExt(".html")

	if asset.CheckAllowed(".html") != false {
		t.Errorf("RemoveAllowedExt Error: CheckAllowed ")
	}
}
func TestAddRemoveDisallowed(t *testing.T) {
	asset := assetmanager.New()
	asset.AddDisallowExt(".html", ".tpl", ".php")
	//t.Error(asset.DisallowedExt)
	if asset.CheckDisallowed(".html") != false {
		t.Errorf("RemoveDisallowExt Error: CheckAllowed ")
	}
	asset.RemoveDisallowExt(".html")
	//t.Error(asset.DisallowedExt)
	if asset.CheckDisallowed(".html") != true {
		t.Errorf("RemoveDisallowExt Error: CheckAllowed ")
	}
}

func TestAddFile(t *testing.T) {

	asset := assetmanager.New()
	file := "test/index.html"
	asset.AddFile(file)
	if !asset.Exists(file) {
		t.Errorf("File Not Found %s", file)
	}
	content, err := asset.GetString(file)
	if err != nil {
		t.Error("Get File Error", err)
	}
	if content != "index.html" {
		t.Error("Get File content not equal index.html")
	}

}

func TestAddFileString(t *testing.T) {

	asset := assetmanager.New()
	file := "test/index.html"
	content := "index.html"
	asset.AddFileString(file, content)
	if !asset.Exists(file) {
		t.Errorf("File Not Found %s", file)
	}
	content, err := asset.GetString(file)
	if err != nil {
		t.Error("GetString File Error", err)
	}
	if content != "index.html" {
		t.Error("GetString File content not equal index.html")
	}

}

func TestAddDir(t *testing.T) {

	asset := assetmanager.New()
	asset.AddAllowedExt(".html", ".tpl")
	asset.AddDisallowExt(".php", ".tpl")
	//asset.AddDir(helpers.DirName())
	asset.AddDir("./test")
	//t.Error(asset.Files)
	content, err := asset.Get("test/index.html")
	if err != nil {
		t.Error("Get File Error test/index.html ", err)
	}
	if string(content) != string([]byte("index.html")) {
		t.Error("Get File content not equal index.html")
	}

	contentString, errString := asset.GetString("test/index.html")
	if errString != nil {
		t.Error("GetString Error test/index.html ", errString)
	}
	if contentString != "index.html" {
		t.Error("Get File content not equal index.html")
	}

	r, e := asset.GetString("test/randomFile.html")
	if e == nil {
		t.Error("GetString need return error  ", e)
	}
	if r != "" {
		t.Error("GetString need return empty")
	}
}
func TestDelete(t *testing.T) {

	asset := assetmanager.New()
	file := "index.html"
	content := "index.html file content"
	asset.AddFileString(file, content)
	if !asset.Exists(file) {
		t.Errorf("File Not Found %s", file)
	}
	asset.Delete(file)
	if asset.Exists(file) {
		t.Errorf("Delete error %s", file)
	}

}

func TestCopy(t *testing.T) {

	asset := assetmanager.New()
	file := "index.html"
	content := "index.html file content"
	asset.AddFileString(file, content)

	assetCopy := assetmanager.New()
	assetCopy.Copy(asset)

	if !asset.Exists(file) {
		t.Errorf("File Not Found %s", file)
	}
	if !assetCopy.Exists(file) {
		t.Errorf("asset copy error File Not Found %s", file)
	}

}
func TestEmptyNameReplacer(t *testing.T) {

	file := "index.html"
	content := "index.html file content"

	asset := assetmanager.New()

	// Add Replacer
	asset.AddReplacer("testReplacer", func(name string) string {
		if name == file {
			return "" // empt
		}
		return name
	})

	asset.AddFileString(file, content)

	if asset.Exists(file) {
		t.Errorf("Empty Name Replacer Error File Found %s", file)
	}

}
