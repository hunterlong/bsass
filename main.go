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
	muxinFuncs  map[string]string
	baseCss     []string
	compiledCss string
	directory   string
)

func init() {
	scssVars = make(map[string]string)
	muxinFuncs = make(map[string]string)
	var err error
	directory, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
}

func CollectMixin(i int, rows []string) int {

	var muxColl []string
	var count int
	for i := i; i <= len(rows)-1; i++ {

		mix := rows[i]
		//fmt.Println(mix)
		muxColl = append(muxColl, mix)
		count++
	}

	compiledMixin := strings.Join(muxColl, "\n")

	re := regexp.MustCompile(`\@mixin (.*?)\(`)
	mixinFunc := re.FindStringSubmatch(compiledMixin)
	mixName := mixinFunc[1]

	muxinFuncs[mixName] = compiledMixin

	re = regexp.MustCompile(`\((.*?)\)`)
	params := re.FindStringSubmatch(compiledMixin)
	params = strings.Split(params[1], ",")

	fmt.Printf("mux: '%v' has %v parameters on %v rows\n", mixName, len(params), count)

	return count

}

func CollectVariables(file string) map[string]string {
	location := fmt.Sprintf("%v.scss", file)
	importedData := openFile(location)
	importedRows := strings.Split(string(importedData), "\n")
	for i := 0; i <= len(importedRows)-1; i++ {
		obj := importedRows[i]
		if len(obj) == 0 {
			continue
		}

		if obj[:1] == "@" {
			skip := CollectMixin(i, importedRows)
			i += skip
			continue
		}

		if obj[:1] != "$" {
			baseCss = append(baseCss, obj)
			continue
		}

		varPass := strings.Split(obj, ": ")
		scssVars[varPass[0]] = varPass[1][0 : len(varPass[1])-1]
	}
	fmt.Printf("+ Imported %v variables from %v.scss\n", len(scssVars), file)
	return scssVars
}

func ParseInclude(i int, rows []string) int {

	var muxColl []string
	var count int
	for i := i; i <= len(rows)-1; i++ {

		mix := rows[i]
		//fmt.Println(mix)
		muxColl = append(muxColl, mix)
		count++
	}

	compiledMixin := strings.Join(muxColl, "\n")

	re := regexp.MustCompile(`\@include (.*?)\(`)
	mixinFunc := re.FindStringSubmatch(compiledMixin)
	mixName := mixinFunc[1]

	fmt.Println("include found:", mixName)

	re = regexp.MustCompile(`\(.*\)`)
	params := re.FindStringSubmatch(compiledMixin)

	fmt.Println(params[0])

	//muxinFuncs[mixName] = compiledMixin

	re = regexp.MustCompile(`/*(.*?)\{`)
	class := re.FindStringSubmatch(compiledMixin)
	fmt.Println(class)
	muxRender := RenderMixin(muxinFuncs[mixName], params[0])

	baseCss = append(baseCss, muxRender)

	fmt.Printf("mux: '%v' has %v parameters on %v rows\n", mixName, 0, count)

	return count

}

func RenderMixin(mix, vars string) string {
	re := regexp.MustCompile(`\{(.|\n)*\}`)
	mixStyle := re.FindStringSubmatch(mix)
	mixCorrected := mixStyle[0][2 : len(mixStyle[0])-2]
	out := strings.Replace(mixCorrected, "$property", vars[1:len(vars)-1], 4)
	return out
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

	for k, v := range baseLines {

		imports := strings.Contains(v, "@import")
		includes := strings.Contains(v, "@include")

		if imports {

			importFile := strings.Split(v, "'")

			CollectVariables(pathJoin + "/" + importFile[1])

		} else if includes {

			ParseInclude(k, baseLines)
			//baseCss = append(baseCss, "FUNCTION: "+v)

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

	if len(s) < 2 {
		return 2, s
	}

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
