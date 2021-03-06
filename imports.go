package bsass

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ScanAllImports(contents string) {
	baseLines := strings.Split(contents, "\n")

	for _, v := range baseLines {
		if len(v) < 8 {
			continue
		}

		importData := regexMultiple(`\@import \'(.*)\'\;`, v)
		if len(importData) == 1 {
			importFile := importData[0]

			if strings.Contains(importFile, "http") {
				remoteData := DownloadRemote(importFile)
				imports[importFile] = remoteData
				continue
			}

			scanFile := fmt.Sprintf("%v/%v.scss", PathJoin, importFile)

			fileContents, err := ioutil.ReadFile(scanFile)
			if err != nil {
				ThrowError(err)
			}

			imports[importFile] = string(fileContents)

			Log("Scanning $variable in %v...\n", scanFile)
			ScanAllVars(string(fileContents))

			Log("Scanning @mixin in %v...\n", scanFile)
			ScanAllMixins(string(fileContents))

			Log("Scanning &extends in %v...\n", scanFile)
			ScanAllExtends(string(fileContents))

		}

	}
}

func ReplaceImport(name string) string {
	var data []string
	lines := strings.Split(imports[name], "\n")
	for i := 0; i < len(lines); i++ {
		val := lines[i]
		if len(val) == 0 {
			continue
		}
		if val[:1] == "$" {
			continue
		}
		if val[:1] == "@" {
			i = skipLines(i, lines)
			continue
		}
		if val[:1] == "%" {
			i = skipLines(i, lines)
			continue
		}
		data = append(data, val)
	}
	if len(data) == 0 {
		return ""
	}
	return strings.Join(data, "\n")
}
