//Copyright (c) 2017, hwfy

//Permission to use, copy, modify, and/or distribute this software for any
//purpose with or without fee is hereby granted, provided that the above
//copyright notice and this permission notice appear in all copies.

//THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
//WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
//MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
//ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
//WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
//ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
//OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

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
