package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func regexSingle(regex, input string) string {
	re := regexp.MustCompile(regex)
	vals := re.FindStringSubmatch(input)
	if len(vals) < 2 {
		return ""
	}
	return vals[1]
}

func regexMultiple(regex, input string) []string {
	re := regexp.MustCompile(regex)
	vals := re.FindStringSubmatch(input)
	if len(vals) < 1 {
		return []string{}
	}
	return vals[1:]
}

func IntInString(s string) int {
	in := regexSingle(`([0-9]+)`, s)
	num, _ := strconv.Atoi(in)
	return num
}

func FloatInString(s string) float64 {
	in := regexSingle(`([0-9]+)`, s)
	num, _ := strconv.ParseFloat(in, 10)
	return num
}

func skipLines(start int, data []string) int {
	for i := start; i < len(data); i++ {
		val := data[i]
		if strings.Contains(val, "}") {
			return i
		}
	}
	return 0
}

func saveFile(file, data string) {
	err := ioutil.WriteFile(file, []byte(data), 0644) // For read access.
	if err != nil {
		panic(err)
	}
}

func removeSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func openFile(file string) string {
	scssFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Cannot open file: %v\n", file)
		panic(err)
	}
	return string(scssFile)
}

func StringInt(a string) int {
	val, _ := strconv.Atoi(a)
	return val
}

func DownloadRemote(file string) string {
	response, err := http.Get(file)
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			return ""
		}
		return string(contents)
	}

	return ""
}
