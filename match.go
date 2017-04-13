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
	"regexp"
)

func identifier(data string) (ret Token, succ bool) {
	regexIdentifier := regexp.MustCompile("^[a-zA-Z\u4e00-\u9fa5]+")

	if str := regexIdentifier.FindString(data); str != "" {
		return Token{Type: "identity", Value: str}, true
	}
	return Token{}, false
}

func number(data string) (ret Token, succ bool) {
	regexNumber := regexp.MustCompile("^[1-9][0-9]*")

	if str := regexNumber.FindString(data); str != "" {
		return Token{Type: "numeric", Value: str}, true
	}
	return Token{}, false
}

func operator(data string) (ret Token, succ bool) {
	if _, ok := terminals[data[0]]; ok {
		return Token{Type: "operator", Value: string(data[0])}, true
	}
	return Token{}, false
}
