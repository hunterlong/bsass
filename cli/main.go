package main

import (
	"errors"
	"fmt"
	"github.com/hunterlong/bsass"
	"os"
	"strings"
)

func runCLI() error {
	if len(os.Args) == 2 {
		method := os.Args[1]
		if method == "version" {
			fmt.Printf("bsass v%v\n", bsass.VERSION)
			return nil
		}
	}
	if len(os.Args) < 3 {
		return errors.New("flags not found to run bsass")
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		err := runCLI()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}
	bsass.ShowHeader()
	scss := os.Args[1]
	css := os.Args[2]

	bsass.Log("Using %v and exporting to %v\n", scss, css)

	pathDir := strings.Split(scss, "/")
	bsass.PathJoin = strings.Join(pathDir[:len(pathDir)-1], "/")

	bsass.Log("Scanning file %v...\n", scss)

	bsass.ScanAll(scss)

	bsass.SassReplacement(scss)

	bsass.SaveFile(css, strings.Join(bsass.RecompiledCss, "\n"))

	bsass.Log("Saved rendered CSS file to: %v\n", css)

}
