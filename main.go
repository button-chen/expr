package main

import (
	"bufio"
	"expr/expr"
	"fmt"
	"os"
	"time"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}
		var source string
		source += scanner.Text()
		if source == "" {
			continue
		}
		if source == "quit()" {
			break
		}
		tm := time.Now()
		v := expr.Eval(source)

		fmt.Printf("%#v  (%v ms)\n", v, time.Since(tm).Milliseconds())
	}
}
