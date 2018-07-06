package main

import (
	"regexp"
	"strings"
)

func RenderMixin(mix, vars string) string {
	re := regexp.MustCompile(`\{(.|\n)*\}`)
	mixStyle := re.FindStringSubmatch(mix)
	mixCorrected := mixStyle[0][2 : len(mixStyle[0])-2]
	out := strings.Replace(mixCorrected, "$property", vars[1:len(vars)-1], 4)
	return out
}

func CollectMixin(i int, rows []string) int {
	var muxColl []string
	var count int
	for i := i; i <= len(rows)-1; i++ {
		mix := rows[i]
		muxColl = append(muxColl, mix)
		count++
		if strings.Contains(mix, "}") {
			break
		}
	}

	compiledMixin := strings.Join(muxColl, "\n")

	re := regexp.MustCompile(`\@mixin (.*?)\(`)
	mixinFunc := re.FindStringSubmatch(compiledMixin)
	mixName := mixinFunc[1]

	muxinFuncs[mixName] = compiledMixin

	re = regexp.MustCompile(`\((.*?)\)`)
	params := re.FindStringSubmatch(compiledMixin)
	params = strings.Split(params[1], ",")

	muxinParams[mixName] = params

	return count

}
