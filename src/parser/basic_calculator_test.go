package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicCalculator(t *testing.T) {
	tt := []struct {
		expr string
		want int
	}{
		{
			expr: "1 - 2 + 3",
			want: 2,
		},
		{
			expr: "10 + 5 - ( 5 + 10 )",
			want: 0,
		},
		{
			expr: "1 - ( 2 + 3 )",
			want: -4,
		},
		{
			expr: "1 + 2",
			want: 3,
		},
		{
			expr: "( 1 )",
			want: 1,
		},
		{
			expr: "( 1 - 2 ) + ( 3 + 3 )",
			want: 5,
		},
		{
			expr: "0",
			want: 0,
		},
		{
			expr: "( ( 1 - 5 ) + 4 ) + ( 4 - 1 )",
			want: 3,
		},
		{
			expr: "( ( 1 - 5 ) + ( 4 + ( 3 ) ) ) + ( 4 - ( ( 1 ) ) )",
			want: 6,
		},
		{
			expr: "( 99 + 1 )", // 99 problems
			want: 100,
		},
	}

	calculator := NewBasicCalculator()
	for i, tc := range tt {
		tc := tc
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := calculator.Calculate(tc.expr)
			require.Equal(t, tc.want, got)
		})
	}
}
