package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExprCalculator(t *testing.T) {
	/* define a new calculator to evaluate the expressions */
	calculator := NewBasicParsingCalculator()

	/* define some test cases */
	cases := []string{
		"1 - 2 + 3",
		"10 + 5 - ( 5 + 10 )",
		"1 - ( 2 + 3 )",
		"1 + 2",
		"( 1 )",
		"( 1 - 2 ) + ( 3 + 3 )",
		"0",
		"( ( 1 - 5 ) + 4 ) + ( 4 - 1 )",
		"( ( 1 - 5 ) + ( 4 + ( 3 ) ) ) + ( 4 - ( ( 1 ) ) )",
		"( 99 + 1 )", // 99 problems
	}
	expected := []int{
		2,
		0,
		-4,
		3,
		1,
		5,
		0,
		3,
		6,
		100,
	}

	for i, expression := range cases {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			result := calculator.Calculate(expression)
			require.Equal(t, expected[i], result)
		})
	}
}
