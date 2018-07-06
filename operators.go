package main

import (
	"fmt"
	"regexp"
	"strings"
)

func checkOperations(cut []string, arr []string) []string {

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
		arr = append(arr, new)
		return arr
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
		arr = append(arr, new)
		return arr
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
		arr = append(arr, new)
		return arr
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
		arr = append(arr, new)
		return arr
	}

	return arr

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
