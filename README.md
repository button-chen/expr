golang implements simple mathematical expression calculation

Example:

```go
import (
	"fmt"

	"github.com/button-chen/expr/expr"
)

func main() {
	v := expr.Eval("-1.5+(3.2*2)/2*(-3)-10.7*3+1e-5")
	fmt.Println("result: ", v)
}

// output: result: -43.19998999999999
```






 
