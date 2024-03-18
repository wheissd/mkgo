package parse

import (
	"os"
	"strings"
)

const maxLevel = 10

func GetModLevel(path string, level int) (string, int) {
	pkgFile, err := os.ReadFile(strings.Repeat("../", level) + "go.mod")
	if err != nil {
		if level > maxLevel {
			panic(err)
		}

		return GetModLevel(path, level+1)
	}

	splittedModFile := strings.SplitN(string(pkgFile), "\n", 2)
	pkg := strings.TrimPrefix(splittedModFile[0], "module ")

	return pkg, level
}

func GetMod(path string) string {
	res, _ := GetModLevel(path, 0)
	return res
}
