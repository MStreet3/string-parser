package parser

type Calculator interface {
	Calculate(expr string) int
}

type BasicCalculator struct {
	parser *BasicParser
}

// NewBasicCalculator returns a calculator that evaluates expressions assuming that
// each rune is separated by white space.
func NewBasicCalculator() Calculator {
	return &BasicCalculator{
		parser: &BasicParser{},
	}
}

// Calculate accepts a string of addition / subtraction operations
// and also parentheses to indicate order of operations.
func (bc *BasicCalculator) Calculate(expr string) int {
	ast := bc.parser.Parse(expr)
	return ast.Evaluate()
}
