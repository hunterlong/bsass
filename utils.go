package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func saveFile(file, data string) {
	err := ioutil.WriteFile(file, []byte(data), 0644) // For read access.
	if err != nil {
		panic(err)
	}
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
