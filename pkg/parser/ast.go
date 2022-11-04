package parser

type Node interface {
	Evaluate() int
}

type ASTree struct {
	root Node
}

func NewASTree(n Node) *ASTree {
	return &ASTree{
		root: n,
	}
}

func (ast *ASTree) Evaluate() int {
	return ast.root.Evaluate()
}

type ASTNode struct {
	Type     GrammarProduction
	Value    int
	Operator string
	left     *ASTNode
	right    *ASTNode
}

// Evaluate returns the evaluated expression of the AST
func (n *ASTNode) Evaluate() int {
	if n.Type == NumericLiteral {
		return n.Value
	} else if n.Operator == "+" {
		return n.left.Evaluate() + n.right.Evaluate()
	} else {
		return n.left.Evaluate() - n.right.Evaluate()
	}
}
