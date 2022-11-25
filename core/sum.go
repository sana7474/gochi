package core

import (
	"strings"

	"gochi/core/numbers"
)

func SumProp(num int, sGender, sCase string) string {
	if num > 999999999999 {
		return ""
	}
	text := ""
	for _, part := range PartNum(num) {
		text += " " + numbers.Load(part.Num, part.Rate, sGender, sCase)
	}
	return strings.ReplaceAll(text, "  ", " ")
}
