package bsass

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/fatih/color"
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
		VERSION = "0.16"
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
		ThrowError(err)
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
		//fmt.Printf("    EXTEND:%v\n", extendName)
	}
	return out
}

func ScanAll(filename string) string {
	scssFile, err := ioutil.ReadFile(filename)
	if err != nil {
		ThrowError(err)
	}
	scssData := string(scssFile)

	ScanAllImports(scssData)

	ScanAllVars(scssData)

	ScanAllMixins(scssData)

	ScanAllExtends(scssData)

	Log("Scan Complete. %v vars | %v mixins | %v extends\n", len(scssVars), len(mixins), len(extends))

	Log("Beginning replacement process now...\n")

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

		if strings.Contains(v, "//") {
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
			params := strings.Split(removeSpaces(mixParams), ",")
			inParams := ReplaceVarArray(params)
			mixData := ReplaceMixins(mixName, inParams)
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
		funcVar := regexMultiple(`\:(.*) (.*)\((.*)\)(.*)\;`, v)
		if len(funcVar) != 0 {
			function := removeSpaces(funcVar[1])
			if function != "darken" && function != "lighten" {
				continue
			}
			exFunc := fmt.Sprintf(`:(.*)%v\((.*)\)(.*)\;`, function)
			otherElems := regexMultiple(exFunc, v)
			prependStyle := otherElems[0]

			switch function {

			case "darken":
				funcParams := regexSingle(`\((.*)\);`, removeSpaces(v))
				splitParams := strings.Split(funcParams, ",")
				inParams := ReplaceVarArray(splitParams)
				color := darken(removeSpaces(inParams[0]), FloatInString(inParams[1]))
				cssEntry := strings.Split(v, ":")
				out := fmt.Sprintf("%v:%v%v;", cssEntry[0], prependStyle, color)
				RecompiledCss = append(RecompiledCss, out)
				continue
			case "lighten":
				funcParams := regexSingle(`\((.*)\);`, removeSpaces(v))
				splitParams := strings.Split(funcParams, ",")
				inParams := ReplaceVarArray(splitParams)
				color := lighten(removeSpaces(inParams[0]), FloatInString(inParams[1]))
				cssEntry := strings.Split(v, ":")
				out := fmt.Sprintf("%v:%v%v;", cssEntry[0], prependStyle, color)
				RecompiledCss = append(RecompiledCss, out)
				continue
			}

		}

		// search for $variable in line
		variable := regexMultiple(`\$(.*?)[;|\s]`, v)
		if len(variable) != 0 {
			stringLine := v

			for _, va := range variable {
				if va[len(va)-1:] == "," {
					continue
				}
				if scssVars[va] == "" {
					err := fmt.Sprintf("missing variable %v %v %v\n", va, errLine, onLine)
					ThrowError(err)
				}
				stringLine = strings.Replace(stringLine, "$"+va, scssVars[va], 2)
			}

			mathCheck := regexMultiple(`\:(.*) -|\+|\/|\* (.*)\;`, stringLine)
			if len(mathCheck) > 0 {
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

func ShowHeader() {
	c := color.New(color.FgCyan)
	c.Println(" _                       ")
	c.Println("| |__  ___  __ _ ___ ___ ")
	c.Println("| '_ \\/ __|/ _` / __/ __|")
	c.Println("| |_) \\__ \\ (_| \\__ \\__ \\    💁")
	c.Printf("|_.__/|___/\\__,_|___/___/  v%v\n", VERSION)
	c.Println("  It's basically sass...")
}

func Log(msg string, data ...interface{}) {
	c := color.New(color.FgHiGreen)
	c.Printf(msg, data...)
}

func ThrowError(err interface{}) {
	c := color.New(color.FgHiRed)
	msg := fmt.Sprintf("\n  Line #%v %v\n  Issue: %v\n", onLine+1, onFile, err)
	c.Printf(msg)
	panic(err)
	os.Exit(2)
}
