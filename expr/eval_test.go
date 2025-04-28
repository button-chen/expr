package expr

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestEval(t *testing.T) {
	v := Eval("-1.5+(3.2*2)/2*(-3)-10.7*3+1e-5")
	fmt.Println("result: ", v)
}

