package expr

import (
	"strconv"
)

func Eval(expr string) float64 {
	v, _ := strconv.ParseFloat(NewParser(NewLexerFromString(expr)).Eval(), 64)
	return v
}
