package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func openFile(file string) string {
	scssFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Cannot open file: %v\n", file)
		panic(err)
	}
	return string(scssFile)
}

func saveFile(file, data string) {
	err := ioutil.WriteFile(file, []byte(data), 0644) // For read access.
	if err != nil {
		panic(err)
	}
}

var (
	renderedCss []string
	scssVars    map[string]string
	baseCss     []string
	compiledCss string
	directory   string
)

func init() {
	scssVars = make(map[string]string)
	var err error
	directory, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
}

func CollectVariables(file string) map[string]string {
	location := fmt.Sprintf("%v.scss", file)
	importedData := openFile(location)
	importedRows := strings.Split(string(importedData), "\n")
	for _, i := range importedRows {
		varPass := strings.Split(i, ": ")
		scssVars[varPass[0]] = varPass[1][0 : len(varPass[1])-1]
	}
	fmt.Printf("+ Imported %v variables from %v.scss\n", len(scssVars), file)
	return scssVars
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
	pathJoin := strings.Join(pathDir[:len(pathDir)-1], "/")

	scssFile, err := ioutil.ReadFile(scss)
	if err != nil {
		panic(err)
	}

	baseLines := strings.Split(string(scssFile), "\n")

	for _, v := range baseLines {

		imports := strings.Contains(v, "@import")

		if imports {

			importFile := strings.Split(v, "'")

			CollectVariables(pathJoin + "/" + importFile[1])

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

func TrueValue(s string) (int, string) {

	s = strings.TrimSpace(s)

	if s[len(s)-1:] == ";" {
		s = s[:len(s)-1]
	}

	re := regexp.MustCompile("[0-9]+")
	trueN1 := re.FindString(s)

	symbol := s[len(s)-(len(s)-len(trueN1)):]
	return StringInt(trueN1), symbol
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

		div := strings.Contains(cut[1], "/")
		plus := strings.Contains(cut[1], "+")
		multi := strings.Contains(cut[1], "*")
		sub := strings.Contains(cut[1], "-")

		if multi {
			subCut := strings.Split(cut[1], "*")
			trueN1, vv1 := TrueValue(subCut[0])
			trueN2, vv2 := TrueValue(subCut[1])
			if vv2 == "" {
				vv2 = vv1
			}
			mathIt := trueN1 * trueN2
			new := fmt.Sprintf("%v: %v%v;", cut[0], mathIt, vv2)
			rendered = append(rendered, new)
			continue
		}

		if sub {
			subCut := strings.Split(cut[1], "-")
			trueN1, vv1 := TrueValue(subCut[0])
			trueN2, vv2 := TrueValue(subCut[1])
			if vv2 == "" {
				vv2 = vv1
			}
			mathIt := trueN1 - trueN2
			new := fmt.Sprintf("%v: %v%v;", cut[0], mathIt, vv2)
			rendered = append(rendered, new)
			continue
		}

		if plus {
			subCut := strings.Split(cut[1], "+")
			trueN1, vv1 := TrueValue(subCut[0])
			trueN2, vv2 := TrueValue(subCut[1])
			if vv2 == "" {
				vv2 = vv1
			}
			mathIt := trueN1 + trueN2
			new := fmt.Sprintf("%v: %v%v;", cut[0], mathIt, vv2)
			rendered = append(rendered, new)
			continue
		}

		if div {
			subCut := strings.Split(cut[1], "/")
			trueN1, vv1 := TrueValue(subCut[0])
			trueN2, vv2 := TrueValue(subCut[1])
			if vv2 == "" {
				vv2 = vv1
			}
			mathIt := trueN1 / trueN2
			new := fmt.Sprintf("%v: %v%v;", cut[0], mathIt, vv2)
			rendered = append(rendered, new)
			continue
		}

		rendered = append(rendered, l)
	}

	return strings.Join(rendered, "\n")

}

func StringInt(a string) int {
	val, _ := strconv.Atoi(a)
	return val
}

func ReplaceVariables(vars map[string]string, full string) string {
	for scssK, scssV := range vars {
		search := fmt.Sprintf(`\%v`, scssK)
		fmt.Println(search)
		var re = regexp.MustCompile(search)
		full = re.ReplaceAllString(full, scssV)
	}
	return full
}
