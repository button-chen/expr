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

func TestBatchEval(t *testing.T) {
	// test.data数据格式: 结果:表达式，表达式由GenerateExpression函数生成，结果由python函数eval计算得到
	data, err := os.ReadFile("test.data")
	if err != nil {
		t.Error("read test.data failed")
	}
	n := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		tmp := strings.Split(line, ":")
		if len(tmp) != 2 {
			continue
		}
		result, expr := tmp[0], strings.TrimSpace(tmp[1])
		r, _ := strconv.ParseFloat(result, 64)
		if Eval(expr) != r {
			t.Error("calc failed: ", line)
		}
		n++
	}
	fmt.Println("success: ", n)
}
