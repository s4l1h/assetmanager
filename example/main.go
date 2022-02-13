package main

import (
	"fmt"

	"github.com/s4l1h/assetmanager"
	"github.com/s4l1h/assetmanager/example/modules/t1"
	"github.com/s4l1h/assetmanager/example/modules/t2"
)

func main() {

	asset := assetmanager.New()

	asset.AddDir("./views")
	//fmt.Println(asset.Files)
	asset.Merge(t1.Asset)
	asset.Merge(t2.Asset)

	for name := range asset.Files {
		fmt.Printf("File Name : %s \n", name)
		c, _ := asset.GetString(name)
		fmt.Printf("File Contents: %s \n", c)
	}

	if content, err := asset.GetString("views/index.html"); err == nil {
		//fmt.Println(err)
		fmt.Println(content)
	}
	if content, err := asset.GetString("t1:/views/index.html"); err == nil {
		//fmt.Println(err)
		fmt.Println(content)
	}
	if content, err := asset.GetString("t2:/views/index.html"); err == nil {
		//fmt.Println(err)
		fmt.Println(content)
	}

	// Output:
	/*
		File Name : t2:/views/add.html
		File Contents: t2 add.html
		File Name : views/index.html
		File Contents: main index.html
		File Name : t1:/views/add.html
		File Contents: t1 add.html
		File Name : t1:/views/index.html
		File Contents: t1 index.html
		File Name : t2:/views/index.html
		File Contents: t2 index.html
		File Name : t2:/views/list.html
		File Contents: t2 list.html
		main index.html
		t1 index.html
		t2 index.html
	*/

}
