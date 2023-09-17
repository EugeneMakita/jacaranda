package parser

import (
	token "Lox/Token"
	"Lox/expression"
)

type Parser struct {
	Tokens  []*token.Token
	Current int
}

func CreateParser(tokens []*token.Token) *Parser {
	return &Parser{
		Tokens: tokens,
	}
}

func (p *Parser) Equality() expression.Expression {
	// comparison ((!= | ==)comparison)*
	var exp expression.Expression = p.Comparison()
	for p.Match(token.DOUBLE_EQUAL, token.NOT_EQUAL) {
		op := p.Previous()
		right := p.Comparison()
		exp = &expression.Binary{
			Operator: *op,
			Right:    right,
			Left:     exp,
		}

	}

	return exp
}

func (p *Parser) Comparison() expression.Expression {
	// term ((>= | <= | > | <)term)*
	var exp expression.Expression = p.Term()
	for p.Match(token.GREATER_THAN, token.GREATER_THAN_OR_EQUAL, token.LESS_THAN_OR_EQUAL, token.LESS_THAN) {
		op := p.Previous()
		right := p.Term()
		exp = &expression.Binary{
			Operator: *op,
			Right:    right,
			Left:     exp,
		}
	}

	return exp
}

func (p *Parser) Term() expression.Expression {
	// factor ((+ | -)factor)*
	var exp expression.Expression = p.Factor()
	for p.Match(token.SUBRACT, token.ADDITION) {
		op := p.Previous()
		right := p.Factor()
		exp = &expression.Binary{
			Operator: *op,
			Right:    right,
			Left:     exp,
		}

	}

	return exp
}

func (p *Parser) Factor() expression.Expression {
	// unary ((* | /)unary)*
	var exp expression.Expression = p.Unary()
	for p.Match(token.MULTiPLY, token.DIVIDE, token.MODULUS) {
		prev := p.Previous()
		right := p.Unary()
		exp = &expression.Binary{
			Operator: *prev,
			Left:     exp,
			Right:    right,
		}
	}

	return exp
}

func (p *Parser) Unary() expression.Expression {
	// (! | -)unary | primary
	if p.Match(token.NEGATION, token.SUBRACT) {
		operator := p.Previous()
		unary := p.Unary()
		return &expression.Unary{
			Operator: *operator,
			Right:    unary,
		}
	}
	return p.Primary()
}

func (p *Parser) Primary() expression.Expression {
	//nill |false| true| string | float| integer | (" expression")
	if p.Match(token.TRUE, token.FALSE, token.INTEGER, token.FLOAT, token.STRING) {
		return &expression.Literal{Object: *p.Previous()}
	}

	if p.Match(token.LEFT_BRACKET) {
		exp := p.Equality()
		if !p.Peek(token.RIGHT_BRACKET) {
			panic("closing bracket not found")
		}
		return &expression.Grouping{
			Exp: exp,
		}
	}

	return nil
}

func (p *Parser) Check(tokenType token.Token_type) bool {
	return p.GetCurrent().Type == tokenType
}

func (p *Parser) GetCurrent() token.Token {
	return *p.Tokens[p.Current]
}

func (p *Parser) MoveCursor() *token.Token {
	token := p.Tokens[p.Current]
	p.Current++
	return token
}

func (p *Parser) Match(tokenTypes ...token.Token_type) bool {
	for _, tokenType := range tokenTypes {
		if p.Check(tokenType) {
			p.MoveCursor()
			return true
		}
	}

	return false
}

func (p *Parser) IsNotAtTheEnd() bool {
	return p.Tokens[p.Current].Char == "EOF"
}

func (p *Parser) Previous() *token.Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) Peek(tokenType token.Token_type) bool {
	if p.Check(tokenType) {
		p.MoveCursor()
		return true
	}

	return false
}
