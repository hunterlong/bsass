package bsass

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	VERSION        string
	scssVars       map[string]string
	imports        map[string]string
	includes       map[string]string
	includesParams map[string][]string
	mixins         map[string]string
	mixinParams    map[string][]string
	extends        map[string]string
	RecompiledCss  []string
	directory      string
	PathJoin       string
	onLine         int
	onFile         string
	errLine        string
)

func init() {
	if VERSION == "" {
		VERSION = "0.14"
	}
	scssVars = make(map[string]string)
	imports = make(map[string]string)
	includes = make(map[string]string)
	includesParams = make(map[string][]string)
	mixins = make(map[string]string)
	mixinParams = make(map[string][]string)
	extends = make(map[string]string)

	var err error
	directory, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
}

func ScanAllExtends(contents string) map[string]string {
	baseLines := strings.Split(contents, "\n")
	var out map[string]string
	for _, v := range baseLines {
		extendName := regexSingle(`\%(.*?) \{`, v)
		if extendName == "" {
			continue
		}
		extendSprint := fmt.Sprintf(`%v {([^}]*)}`, extendName)
		fullExtend := regexSingle(extendSprint, contents)
		extends[extendName] = fullExtend[1 : len(fullExtend)-1]
		fmt.Printf("    EXTEND:%v\n", extendName)
	}
	return out
}

func ScanAll(filename string) string {
	scssFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scssData := string(scssFile)

	ScanAllImports(scssData)

	ScanAllVars(scssData)

	ScanAllMixins(scssData)

	ScanAllExtends(scssData)

	fmt.Printf("Scan Complete. %v vars | %v mixins | %v extends\n", len(scssVars), len(mixins), len(extends))

	fmt.Println("Beginning replacement process now...")

	return scssData
}

func SassReplacement(filename string) {
	onFile = filename
	scssFile, err := ioutil.ReadFile(filename)
	if err != nil {
		ThrowError(err)
	}
	scssData := string(scssFile)
	baseLines := strings.Split(scssData, "\n")

	for k, v := range baseLines {
		onLine = k
		errLine = v
		if len(v) == 0 {
			continue
		}

		// search for @import in line
		importName := regexSingle(`\@import '(.*?)'\;`, v)
		if importName != "" {
			RecompiledCss = append(RecompiledCss, ReplaceImport(importName))
			continue
		}

		// search for @include in line
		included := regexSingle(`\@include (.*?)\;`, v)
		if included != "" {
			mixName := regexSingle(`(.*?)\(`, included)
			mixParams := regexSingle(`\((.*?)\)`, included)
			params := strings.Split(mixParams, ",")
			mixData := ReplaceMixins(mixName, params)
			RecompiledCss = append(RecompiledCss, mixData)
			continue
		}

		// search for @extend in line
		extended := regexSingle(`\@extend %(.*?)[;|\s]`, v)
		if extended != "" {
			extendData := extends[extended]
			RecompiledCss = append(RecompiledCss, extendData)
			continue
		}

		//search for functions like 'darken'
		function := regexSingle(`\:(.*?)\(\$`, removeSpaces(v))
		if len(function) != 0 {
			switch function {

			case "darken":
				funcParams := regexSingle(`\$(.*?)\)`, removeSpaces(v))
				splitParams := strings.Split(funcParams, ",")
				color := darken(scssVars[splitParams[0]], FloatInString(splitParams[1]))
				cssEntry := strings.Split(v, ":")
				out := fmt.Sprintf("%v: %v;", cssEntry[0], color)
				RecompiledCss = append(RecompiledCss, out)
				continue
			case "lighten":
				funcParams := regexSingle(`\$(.*?)\)`, removeSpaces(v))
				splitParams := strings.Split(funcParams, ",")
				color := lighten(scssVars[splitParams[0]], FloatInString(splitParams[1]))
				cssEntry := strings.Split(v, ":")
				out := fmt.Sprintf("%v: %v;", cssEntry[0], color)
				RecompiledCss = append(RecompiledCss, out)
				continue
			}

		}

		// search for $variable in line
		variable := regexMultiple(`\$(.*?)[;|\s]`, v)
		if len(variable) != 0 {
			stringLine := v

			for _, va := range variable {
				stringLine = strings.Replace(stringLine, "$"+va, scssVars[va], 1)
			}

			if strings.ContainsAny(stringLine, "-+*/") {
				math := regexSingle(`\:(.*)`, stringLine)
				math = removeSpaces(math)
				reg := regexp.MustCompile(`[^0-9|/|+|\-|*]+`)
				mathProblem := reg.ReplaceAllString(math, "")
				varType := regexSingle(`([a-zA-Z]+)`, math)
				expression, err := govaluate.NewEvaluableExpression(mathProblem)
				if err != nil {
					ThrowError(err)
					continue
				}
				result, err := expression.Evaluate(nil)
				if err != nil {
					ThrowError(err)
					continue
				}
				cssEntry := strings.Split(stringLine, ":")
				stringLine = fmt.Sprintf("%v: %v%v;", cssEntry[0], result, varType)
			}

			RecompiledCss = append(RecompiledCss, stringLine)
			continue
		}

		RecompiledCss = append(RecompiledCss, v)
	}

}

func ThrowError(err error) {
	fmt.Printf("\nError in '%v', line #%v, %v\nIssue: %v\n", onFile, onLine+1, err, errLine)
	os.Exit(2)
}
