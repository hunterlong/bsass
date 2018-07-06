package main

import (
	"fmt"
	"regexp"
	"strings"
)

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

	re = regexp.MustCompile(`\(.*\)`)
	params := re.FindStringSubmatch(compiledMixin)

	//muxinFuncs[mixName] = compiledMixin

	//re = regexp.MustCompile(`/*(.*?)\{`)
	//class := re.FindStringSubmatch(compiledMixin)
	//fmt.Println(class)
	muxRender := RenderMixin(muxinFuncs[mixName], params[0])

	baseCss = append(baseCss, muxRender)

	fmt.Printf("mux: '%v' has %v parameters on %v rows\n", mixName, 0, count)

	return count

}
