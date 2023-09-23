package expression

import (
	token "Lox/Token"
	"fmt"
	"strconv"
)

type Object interface {
	NumericValue() (float64, bool)
}

type FloatObject float64

func (f FloatObject) NumericValue() (float64, bool) {
	return float64(f), true
}

type IntObject int

func (i IntObject) NumericValue() (float64, bool) {
	return float64(i), true
}

type StringObject string

func (s StringObject) NumericValue() (float64, bool) {
	return 0, false
}

type BoolObject bool

func (b BoolObject) NumericValue() (float64, bool) {
	if b {
		return 1, true
	}
	return 0, true
}

type Expression interface {
	eval() int
	String() string
	Visit() Object
}

type Binary struct {
	Right    Expression
	Operator token.Token
	Left     Expression
}

func (b *Binary) eval() int {
	return 0
}

func (b *Binary) Visit() Object {
	right := b.Right.Visit()
	left := b.Left.Visit()
	rightValue, ok1 := right.NumericValue()
	leftValue, ok2 := left.NumericValue()

	switch b.Operator.Type {
	case token.ADDITION:
		if ok1 && ok2 {
			return FloatObject(leftValue + rightValue)
		}
	case token.SUBRACT:
		if ok1 && ok2 {
			return FloatObject(leftValue - rightValue)
		}
	case token.DIVIDE:
		if ok1 && ok2 {
			return FloatObject(leftValue / rightValue)
		}
	case token.MULTiPLY:
		if ok1 && ok2 {
			return FloatObject(leftValue * rightValue)
		}
	case token.MODULUS:
		if ok1 && ok2 {
			return IntObject(int64(leftValue) % int64(rightValue))
		}
	case token.LESS_THAN:
		if ok1 && ok2 {
			return BoolObject(leftValue < rightValue)
		}
	case token.LESS_THAN_OR_EQUAL:
		if ok1 && ok2 {
			return BoolObject(leftValue <= rightValue)
		}
	case token.GREATER_THAN:
		if ok1 && ok2 {
			return BoolObject(leftValue > rightValue)
		}
	case token.GREATER_THAN_OR_EQUAL:
		if ok1 && ok2 {
			return BoolObject(leftValue >= rightValue)
		}
	case token.DOUBLE_EQUAL:
		if ok1 && ok2 {
			return BoolObject(leftValue == rightValue)
		}
	case token.NOT_EQUAL:
		if ok1 && ok2 {
			return BoolObject(leftValue != rightValue)
		}
	}

	return BoolObject(false)
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

func (u *Unary) Visit() Object {
	switch u.Operator.Type {
	case token.SUBRACT:
		right, ok := u.Right.Visit().NumericValue()
		if ok {
			return FloatObject(-right)
		}

	}
	return BoolObject(false)
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
	return fmt.Sprintf("Literals: %v", l.Object)
}
func (l *Literal) Visit() Object {
	switch l.Object.Type {
	case token.INTEGER:
		val, err := strconv.ParseInt(l.Object.Char, 10, 64)
		if err != nil {
			panic(err)
		}
		return IntObject(val)
	case token.FLOAT:
		val, err := strconv.ParseFloat(l.Object.Char, 64)
		if err != nil {
			panic(err)
		}
		return FloatObject(val)
	case token.STRING:
		return StringObject(l.Object.Char)
	case token.FALSE:
		return BoolObject(false)
	case token.TRUE:
		return BoolObject(true)
	}

	fmt.Println("I don't know what we have here lol")
	return BoolObject(false)
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

func (g *Grouping) Visit() Object {
	res := g.Exp.Visit()
	val, ok := res.NumericValue()
	if ok {
		return FloatObject(val)
	}
	return BoolObject(false)
}
