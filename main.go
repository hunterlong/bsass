package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	VERSION string
	scssVars     map[string]string
	muxinFuncs   map[string]string
	muxinParams  map[string][]string
	extendsFuncs map[string]string
	baseCss      []string
	compiledCss  string
	directory    string
	pathJoin     string
)

func init() {
	scssVars = make(map[string]string)
	muxinFuncs = make(map[string]string)
	muxinParams = make(map[string][]string)
	extendsFuncs = make(map[string]string)
	var err error
	directory, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Not enough parameters!\n")
		os.Exit(2)
	}

	scss := os.Args[1]
	css := os.Args[2]

	fmt.Printf("Using %v and exporting to %v\n", scss, css)
	pathDir := strings.Split(scss, "/")
	pathJoin = strings.Join(pathDir[:len(pathDir)-1], "/")

	scssFile, err := ioutil.ReadFile(scss)
	if err != nil {
		panic(err)
	}

	baseLines := strings.Split(string(scssFile), "\n")

	for k, v := range baseLines {

		imports := strings.Contains(v, "@import")
		includes := strings.Contains(v, "@include")
		extends := strings.Contains(v, "@extend")

		if imports {

			importFile := strings.Split(v, "'")

			CollectVariables(importFile[1])

		} else if includes {

			ParseInclude(k, baseLines)

		} else if extends {

			ParseExtended(k, baseLines)

		} else {

			baseCss = append(baseCss, v)

		}

	}

	//renderCss := strings.Join(baseCss, "\n")

	//var renderCss string

	var unvarredScss []string
	for _, b := range baseCss {
		if strings.Contains(b, "$") {
			var objVar string
			obj := strings.Split(b, "$")
			sobj := strings.Split(obj[1], " ")
			objVar = sobj[0]
			if sobj[0][len(sobj[0])-1:] == ";" {
				objVar = objVar[0 : len(sobj[0])-1]
			}
			objReplace := fmt.Sprintf("$%v", objVar)
			out := strings.Replace(b, objReplace, scssVars[objReplace], 1)
			unvarredScss = append(unvarredScss, out)
			continue
		}
		unvarredScss = append(unvarredScss, b)
	}

	compiled := strings.Join(unvarredScss, "\n")

	compiled = ScanRows(compiled)

	saveFile(css, compiled)
	fmt.Printf("Saved rendered CSS file to: %v\n", css)

}

func ScanRows(base string) string {

	var rendered []string

	splitBase := strings.Split(base, "\n")

	for _, l := range splitBase {

		cut := strings.Split(l, ": ")
		if len(cut) <= 1 {
			rendered = append(rendered, l)
			continue
		}

		// check for any math operations
		prevAmount := len(rendered)
		rendered = checkOperations(cut, rendered)
		if prevAmount != len(rendered) {
			continue
		}

		rendered = append(rendered, l)
	}

	return strings.Join(rendered, "\n")

}
