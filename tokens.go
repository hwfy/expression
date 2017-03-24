package expression

type (
	tokenType int

	Token struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	searchFunc func(data string) (ret Token, succ bool)
)

const (
	LBRACK tokenType = iota
	RBRACK
	COMMA
	MINUS
	PLUS
	MULTIPLY
	DIVIDED
	PRECE
)

var terminals = map[byte]tokenType{
	'(': LBRACK,
	')': RBRACK,
	',': COMMA,
	'-': MINUS,
	'+': PLUS,
	'*': MULTIPLY,
	'/': DIVIDED,
	'%': PRECE,
}

var matchFunctions = []searchFunc{
	identifier,
	number,
	operator,
}
