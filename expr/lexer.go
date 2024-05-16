package expr

import "bytes"

const (
	TOKEN_INVALID = iota + 1
	TOKEN_NUM
	TOKEN_OP
	TOKEN_SCOPE_BEGIN
	TOKEN_SCOPE_END
	TOKEN_EOF
)

const (
	PRIORITY_LOWEST = iota
	PRIORITY_ADDSUB // + -
	PRIORITY_MULQUO // * /
	PRIORITY_SCOPE  // ()
)

const INVALID = "INVALID"

type Token struct {
	Ty       int
	Lit      string
	Priority int
}

func (t Token) String() string {
	return t.Lit
}

type Lexer struct {
	tokens []Token
	data   []byte
	pos    int
}

func NewLexerFromString(data string) *Lexer {
	lex := &Lexer{
		tokens: make([]Token, 0),
		data:   []byte(data),
	}
	lex.Parse()
	return lex
}

func NewLexerFromTokenSlice(tokens []Token) *Lexer {
	lex := &Lexer{
		tokens: tokens,
	}
	return lex
}

func (lx *Lexer) Parse() {
	lx.data = bytes.TrimSpace(lx.data)
	for i := 0; i < len(lx.data); i++ {
		c := lx.data[i]
		if c == ' ' || c == '\t' {
			continue
		}
		if c >= '0' && c <= '9' || (i == 0 && (c == '-' || c == '+')) {
			num, j := lx.parseNum(lx.data[i:])
			lx.tokens = append(lx.tokens, Token{Lit: num, Ty: TOKEN_NUM, Priority: PRIORITY_LOWEST})
			i += j
			continue
		}
		switch c {
		case '+', '-':
			lx.tokens = append(lx.tokens, Token{Lit: string(c), Ty: TOKEN_OP, Priority: PRIORITY_ADDSUB})
		case '*', '/':
			lx.tokens = append(lx.tokens, Token{Lit: string(c), Ty: TOKEN_OP, Priority: PRIORITY_MULQUO})
		case '(':
			lx.tokens = append(lx.tokens, Token{Lit: string(c), Ty: TOKEN_SCOPE_BEGIN, Priority: PRIORITY_SCOPE})
		case ')':
			lx.tokens = append(lx.tokens, Token{Lit: string(c), Ty: TOKEN_SCOPE_END, Priority: PRIORITY_LOWEST})
		}
	}
}

func (lx *Lexer) GetToken() Token {
	if lx.pos == len(lx.tokens) {
		return Token{Ty: TOKEN_EOF}
	}
	t := lx.tokens[lx.pos]
	lx.pos++
	return t
}

func (lx *Lexer) PeekNextOp() Token {
	for i := lx.pos; i < len(lx.tokens); i++ {
		if lx.tokens[i].Ty == TOKEN_OP {
			return lx.tokens[i]
		}
	}
	return Token{Ty: TOKEN_INVALID, Priority: PRIORITY_LOWEST}
}

func (lx *Lexer) parseNum(data []byte) (string, int) {
	num := make([]byte, 0)
	num = append(num, data[0])
	var lastChar byte
	j := 1
	for ; j < len(data); j++ {
		c := data[j]
		if lastChar == 'e' {
			if c == '+' || c == '-' {
				num = append(num, c)
				continue
			}
		}
		if c >= '0' && c <= '9' || c == '.' || c == 'e' {
			lastChar = c
			num = append(num, c)
			continue
		}
		break
	}
	return string(num), j - 1
}
