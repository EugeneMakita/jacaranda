package interpretor

import (
	"Lox/expression"
	"fmt"
)

type Interpretor struct {
}

func CreateInterpretor() *Interpretor {
	return &Interpretor{}
}

func (i *Interpretor) Evaluate(exp expression.Expression) {
	fmt.Println(exp.Visit())
}

func (i *Interpretor) EvaluateUnitaryExpression(exp expression.Unary) {
	exp.Visit()
}

func (i *Interpretor) Emak() {
	fmt.Println("james")
}
