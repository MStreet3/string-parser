package entities

type Evaluator interface {
	Evaluate() (int, error)
}

type Calculator interface {
	Calculate(expr string) int
}

type Parser interface {
	Parse(string) (Evaluator, error)
}
