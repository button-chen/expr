package expr

import (
	"fmt"
	"math/rand"
	"strconv"
)

var operators = []string{"+", "-", "*", "/"}

func GenerateExpression(depth int) string {
	if depth == 0 {
		return strconv.FormatFloat(rand.Float64()+float64(rand.Int()%777), 'f', -1, 64)
	}

	left := GenerateExpression(depth - 1)
	right := GenerateExpression(depth - 1)

	operator := operators[rand.Intn(len(operators))]

	if rand.Intn(2) == 0 {
		return fmt.Sprintf("(%s %s %s)", left, operator, right)
	}

	return fmt.Sprintf("%s %s %s", left, operator, right)
}
