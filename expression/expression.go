package expression

import (
	token "Lox/Token"
	"fmt"
)

type Expression interface {
	eval() int
	String() string
}

type Binary struct {
	Right    Expression
	Operator token.Token
	Left     Expression
}

func (b *Binary) eval() int {
	return 0
}

func (b *Binary) String() string {
	return fmt.Sprintf("Left: (%v), Operator: %v, Right:(%v)", b.Left, b.Operator, b.Right)
}

type Unary struct {
	Operator token.Token
	Right    Expression
}

func (u *Unary) eval() int {
	return 0
}

func (u *Unary) String() string {
	return fmt.Sprintf("Unary: %v, Right: %v", u.Operator, u.Right)
}

type Literal struct {
	Object token.Token
}

func (l *Literal) eval() int {
	return 0
}

func (l *Literal) String() string {
	return fmt.Sprintf("Literal: %v", l.Object)
}

type Grouping struct {
	Exp Expression
}

func (g *Grouping) eval() int {
	return 0
}

func (g *Grouping) String() string {
	return fmt.Sprintf("Expression:(%v)", g.Exp)
}
