package bsass

import (
	"strings"
)

func ScanAllVars(contents string) {
	baseLines := strings.Split(contents, "\n")
	for _, v := range baseLines {
		varLess := regexMultiple(`\$(.*):(.*)\;`, v)
		if len(varLess) < 2 {
			continue
		}
		varValue := varLess[1]
		varName := removeSpaces(varLess[0])
		scssVars[varName] = varValue
	}
}

func ReplaceVarArray(arr []string) []string {
	var inParams []string
	for _, p := range arr {
		if p[:1] == "$" {
			p = scssVars[p[1:]]
		}
		inParams = append(inParams, p)
	}
	return inParams
}
