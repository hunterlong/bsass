package bsass

import (
	"fmt"
	"strings"
)

func ScanAllMixins(contents string) {
	baseLines := strings.Split(contents, "\n")

	for _, v := range baseLines {
		mixinName := regexSingle(`\@mixin (.*?)\(`, v)
		if mixinName == "" {
			continue
		}

		funcReg := fmt.Sprintf(`%v\((.*?)\)`, mixinName)

		params := regexSingle(funcReg, v)

		paramSplit := strings.Split(params, ",")

		escapedParams := strings.Replace(params, "$", "\\$", 4)

		mixinSprint := fmt.Sprintf(`%v\(%v\) {([^}]*)}`, mixinName, escapedParams)
		fullMixin := regexSingle(mixinSprint, contents)

		mixins[mixinName] = fullMixin[1 : len(fullMixin)-1]
		mixinParams[mixinName] = paramSplit

		fmt.Printf("    MIXIN:%v=%v\n", mixinName, paramSplit)

	}

}

func ReplaceMixins(name string, params []string) string {
	mix := mixins[name]
	mixParams := mixinParams[name]
	for k, m := range mixParams {
		if strings.Contains(params[k], "(") {
			params[k] += ")"
		}
		mix = strings.Replace(mix, m, params[k], 4)
	}
	return mix
}
