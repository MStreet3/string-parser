package calculator

import "github.com/mstreet3/calculator/pkg/parser"

type Calculator interface {
	Calculate(expr string) int
}

type basicCalculator struct {
	parser parser.Interface
}

// NewBasicCalculator returns a calculator that evaluates expressions assuming that
// each rune is separated by white space.
func NewBasicCalculator() Calculator {
	return &basicCalculator{
		parser: &parser.BasicParser{},
	}
}

// Calculate accepts a string of addition / subtraction operations
// and also parentheses to indicate order of operations.
func (bc *basicCalculator) Calculate(expr string) int {
	ast := bc.parser.Parse(expr)
	return ast.Evaluate()
}
