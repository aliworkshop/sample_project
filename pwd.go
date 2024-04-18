package sample_project

import (
	"path"
	"runtime"
)

func AppRootPath() string {
	_, fullFilename, _, _ := runtime.Caller(0)
	return path.Dir(fullFilename)
}
