package calculator

import (
	"github.com/mstreet3/rdp/pkg/entities"
)

type calculator struct {
	parser entities.Parser
}

// NewCalculator returns a calculator that evaluates expressions.  Expression
// parsing is defined by the parser.
func NewCalculator(p entities.Parser) entities.Calculator {
	return &calculator{
		parser: p,
	}
}

// Calculate accepts a string of addition / subtraction operations
// and also parentheses to indicate order of operations.
func (bc *calculator) Calculate(expr string) int {
	e, err := bc.parser.Parse(expr)
	if err != nil {
		panic(err)
	}

	val, err := e.Evaluate()
	if err != nil {
		panic(err)
	}

	return val
}
