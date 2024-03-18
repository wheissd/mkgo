package parse

import (
	"os"
	"strings"

	"github.com/samber/lo"
)

type PkgInfo struct {
	Pkg     string
	RootDir string
}

func GetPkgInfo() PkgInfo {
	wd, _ := os.Getwd()
	wdSplitted := strings.Split(wd, "/")
	pkg, lvl := GetModLevel(wd, 0)
	rootPkg := strings.Join(lo.Filter(wdSplitted, func(item string, i int) bool {
		return i > len(wdSplitted)-lvl-1
	}), "/")
	return PkgInfo{
		Pkg:     pkg,
		RootDir: rootPkg,
	}
}
