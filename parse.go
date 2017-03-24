package expression

import (
	"fmt"
	"strings"
)

// Parse parses the expression into a variable, an operator, a value
// eg. "age +20"  ->  [{identity age} {operator +} {numeric 20}]
func Parse(data string) (tokens []Token, err error) {
	data = strings.Trim(data, " ")
	for len(data) > 0 {
		matched := false
		for _, matchFunc := range matchFunctions {
			if ret, succ := matchFunc(data); succ {
				data = strings.Trim(data[len(ret.Value):], " ")
				tokens = append(tokens, ret)
				matched = true
				break
			}
		}
		if !matched {
			return nil, fmt.Errorf("Error: parse %s unknown token %s", data, string(data[0]))
		}
	}
	return tokens, nil
}

//log.Print("skipping unknown token: " + string(data[0]))
//data = data[1:]
