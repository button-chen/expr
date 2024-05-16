package expr

import (
	"fmt"
	"math/big"
	"strconv"
)

type Parser struct {
	lx *Lexer
}

func NewParser(lx *Lexer) *Parser {
	return &Parser{
		lx: lx,
	}
}

func (p *Parser) Eval() string {
	result := "INVALID"
	p.eval(&result)
	return result
}

func (p *Parser) eval(result *string) {
	toks := make([]Token, 0)
	for {
		t := p.lx.GetToken()
		if *result != INVALID {
			toks = append(toks, Token{Ty: TOKEN_NUM, Lit: *result})
			*result = INVALID
		}
		if t.Ty == TOKEN_EOF {
			break
		}
		if t.Ty == TOKEN_SCOPE_BEGIN {
			p.eval(result)
			continue
		}
		if t.Ty == TOKEN_SCOPE_END {
			break
		}
		toks = append(toks, t)
	}
	v := strconv.FormatFloat(p.execute(p.build(NewLexerFromTokenSlice(toks), false)), 'f', -1, 64)
	*result = v
}

func (p *Parser) build(lx *Lexer, IsEnterHighPriority bool) Node {
	var left *Node
	var preToken Token
	for {
		t := lx.GetToken()
		if t.Ty == TOKEN_EOF {
			break
		}
		if t.Ty != TOKEN_OP {
			preToken = t
			continue
		}
		node := Node{Val: t}
		if left != nil {
			node.Left = left
			left = nil
		} else {
			node.Left = &Node{Val: preToken}
		}

		nextOp := lx.PeekNextOp()
		if nextOp.Priority > t.Priority {
			n := p.build(lx, true)
			node.Right = &n
		} else if nextOp.Priority < t.Priority {
			node.Right = &Node{Val: lx.GetToken()}
			if IsEnterHighPriority {
				return node
			}
		} else {
			node.Right = &Node{Val: lx.GetToken()}
		}
		left = &node
	}
	if left == nil {
		left = &Node{Val: preToken}
	}
	return *left
}

func (p *Parser) execute(node Node) float64 {
	if node.Val.Ty == TOKEN_OP {
		a := p.execute(*node.Left)
		b := p.execute(*node.Right)
		switch node.Val.Lit {
		case "+":
			lh := new(big.Float).SetFloat64(a)
			rh := new(big.Float).SetFloat64(b)
			f, _ := new(big.Float).Add(lh, rh).Float64()
			return f
		case "-":
			lh := new(big.Float).SetFloat64(a)
			rh := new(big.Float).SetFloat64(b)
			f, _ := new(big.Float).Sub(lh, rh).Float64()
			return f
		case "*":
			lh := new(big.Float).SetFloat64(a)
			rh := new(big.Float).SetFloat64(b)
			f, _ := new(big.Float).Mul(lh, rh).Float64()
			return f
		case "/":
			if b == 0 {
				panic(ErrDivisionZero)
			}
			lh := new(big.Float).SetFloat64(a)
			rh := new(big.Float).SetFloat64(b)
			f, _ := new(big.Float).Quo(lh, rh).Float64()
			return f
		default:
			panic(ErrUnsupportedArithmetic)
		}
	}
	v, _ := strconv.ParseFloat(node.Val.Lit, 64)
	return v
}

var ErrDivisionZero = fmt.Errorf("division by zero error")
var ErrUnsupportedArithmetic = fmt.Errorf("unsupported arithmetic")
