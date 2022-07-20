package parser

type Node interface {
	Evaluate() int
}

type Parser interface {
	Parse(expr string) *ASTree
}

type Calculator interface {
	Calculate(expr string) int
}
