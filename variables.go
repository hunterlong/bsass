package main

import (
	"fmt"
	"strings"
)

func CollectVariables(file string) map[string]string {

	if strings.Contains(file, "http") {
		remote := DownloadRemote(file)
		baseCss = append(baseCss, remote)
		return scssVars
	}

	location := fmt.Sprintf("%v/%v.scss", pathJoin, file)
	importedData := openFile(location)
	importedRows := strings.Split(string(importedData), "\n")
	for i := 0; i <= len(importedRows)-1; i++ {
		obj := importedRows[i]
		if len(obj) == 0 {
			continue
		}

		// importing extends
		if obj[:1] == "%" {
			skip := CollectExtended(i, importedRows)
			i += skip
			continue
		}

		// importing imports
		if obj[:1] == "@" {
			skip := CollectMixin(i, importedRows)
			i += skip
			continue
		}

		// importing variables
		if obj[:1] != "$" {
			baseCss = append(baseCss, obj)
			continue
		}

		varPass := strings.Split(obj, ": ")
		scssVars[varPass[0]] = varPass[1][0 : len(varPass[1])-1]
	}
	return scssVars
}
