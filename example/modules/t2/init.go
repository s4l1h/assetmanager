s4l1hpackage t2

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/akmyazilim/assetmanager"
)

// Asset assetmanager
var Asset *assetmanager.AssetManager

func init() {

	_, file, _, _ := runtime.Caller(0)
	// Get Full package dir
	path := filepath.Dir(file)
	views := fmt.Sprintf("%s/views", path)
	//logrus.Warn(path)
	Asset = assetmanager.New()
	Asset.AddReplacer("renamer", func(name string) string {
		return strings.Replace(name, path, "t2:", -1)
	})
	Asset.AddDir(views)
}
