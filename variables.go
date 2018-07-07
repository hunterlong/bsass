package bsass

import (
	"regexp"
	"strings"
)

func ScanAllVars(contents string) {
	baseLines := strings.Split(contents, "\n")

	for _, v := range baseLines {
		re := regexp.MustCompile(`\$(.*)*\;`)
		varRegex := re.FindStringSubmatch(v)
		if len(varRegex) <= 1 {
			continue
		}
		re = regexp.MustCompile(`\$(.*?)\:`)
		regName := re.FindStringSubmatch(varRegex[0])
		re = regexp.MustCompile(`\:(.*?)\;`)
		regVal := re.FindStringSubmatch(varRegex[0])

		if len(regVal) <= 1 {
			continue
		}

		varValue := removeSpaces(regVal[1])
		varName := removeSpaces(regName[1])

		scssVars[varName] = varValue

		//fmt.Printf("    VAR:%v=%v\n", varName, scssVars[varName])

	}

}
