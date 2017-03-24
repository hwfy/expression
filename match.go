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
