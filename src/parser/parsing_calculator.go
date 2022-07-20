package parser

type BasicCalculator struct {
	parser *BasicParser
}

// Calculate accepts a string of addition / subtraction operations
// and also parentheses to indicate order of operations.
func (pc *BasicCalculator) Calculate(expr string) int {
	ast := pc.parser.Parse(expr)
	return ast.Root.Evaluate()
}

// NewBasicCalculator returns a calculator that evaluates expressions assuming that
// each rune is separated by white space.
func NewBasicCalculator() Calculator {
	return &BasicCalculator{
		parser: &BasicParser{},
	}
}
