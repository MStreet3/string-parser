package parser

type ASTree struct {
	Root *ASTNode
}

type ASTNode struct {
	Type     GrammarProduction
	Value    int
	Operator string
	left     *ASTNode
	right    *ASTNode
}

/* Evaluate returns the evaluated expression of the AST */
func (n *ASTNode) Evaluate() int {
	if n.Type == NumericLiteral {
		return n.Value
	} else if n.Operator == "+" {
		return n.left.Evaluate() + n.right.Evaluate()
	} else {
		return n.left.Evaluate() - n.right.Evaluate()
	}
}
