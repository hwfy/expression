# expression


The expression is parsed as a variable, an operator, and a value
# Installation


> go get github.com/hwfy/expression
# Example


```go
package main
import (
	"encoding/json"
	"fmt"
	
	"github.com/hwfy/expression"
)

func main() {
	tokens, err := expression.Parse("(age+10)*2")
	if err != nil {
		panic(err)
	}
	data, _ := json.Marshal(tokens)

	fmt.Println(string(data))
}
//Output:
[
    {
        "type": "operator",
        "value": "("
    },
    {
        "type": "identity",
        "value": "age"
    },
    {
        "type": "operator",
        "value": "+"
    },
    {
        "type": "numeric",
        "value": "10"
    },
    {
        "type": "operator",
        "value": ")"
    },
    {
        "type": "operator",
        "value": "*"
    },
    {
        "type": "numeric",
        "value": "2"
    }
]
```
