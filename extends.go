package main

import (
	"regexp"
	"strings"
)

func CollectExtended(i int, rows []string) int {
	var colls []string
	var count int
	for i := i; i <= len(rows)-1; i++ {
		mix := rows[i]
		colls = append(colls, mix)
		count++
		if strings.Contains(mix, "}") {
			break
		}
	}

	compiledMixin := strings.Join(colls, "\n")

	re := regexp.MustCompile(`\%(.|\n)*\{`)
	exName := re.FindStringSubmatch(compiledMixin)
	extendName := exName[0][1 : len(exName[0])-2]

	re = regexp.MustCompile(`\{(.|\n)*\}`)
	mixStyle := re.FindStringSubmatch(compiledMixin)
	mixCorrected := mixStyle[0][2 : len(mixStyle[0])-2]

	extendsFuncs[extendName] = mixCorrected

	return count
}

func ParseExtended(i int, rows []string) int {
	var muxColl []string
	var count int
	for i := i; i <= len(rows)-1; i++ {

		mix := rows[i]
		//fmt.Println(mix)
		muxColl = append(muxColl, mix)
		count++
	}

	comp := strings.Join(muxColl, "\n")

	re := regexp.MustCompile(`\%(.*)\;`)

	extName := re.FindStringSubmatch(comp)

	extendedName := extName[0][1 : len(extName[0])-1]

	baseCss = append(baseCss, extendsFuncs[extendedName])

	return count

}
