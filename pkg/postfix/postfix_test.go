package postfix

import (
	"fmt"
	"testing"

	"github.com/mstreet3/rdp/pkg/tokenizer"
	"github.com/stretchr/testify/require"
)

func TestPostfixEvaluator(t *testing.T) {
	tt := []struct {
		expr    string
		want    int
		wantErr error
	}{
		{
			expr: "1 2 +",
			want: 3,
		},
		{
			expr: "10 2 + 5 -",
			want: 7,
		},
		{
			expr: "1 2 + 3 -",
			want: 0,
		},
		{
			expr:    "1 2 * 3 -",
			wantErr: ErrUnexpectedInput,
		},
		{
			expr: "",
			want: 0,
		},
		{
			expr:    "1 -",
			wantErr: ErrInvalidPostfixExpression,
		},
	}

	for i, tc := range tt {
		tc := tc
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			e := &expression{
				tokenizer: tokenizer.NewBasicTokenizer(tc.expr),
			}

			got, err := e.Evaluate()

			if tc.wantErr == nil {
				require.Equal(t, tc.want, got)
				require.NoError(t, err)
			}

			if tc.wantErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.wantErr)
			}
		})
	}
}

func TestFromInfixTransform(t *testing.T) {
	tt := []struct {
		want    string
		expr    string
		wantErr error
	}{
		{
			expr: "1 + 2",
			want: "1 2 +",
		},
		{
			expr: "10 + 2 - 5",
			want: "10 2 + 5 -",
		},
		{
			expr: "1 + 2 - 3",
			want: "1 2 + 3 -",
		},
		{
			expr: "1 + ( 2 - 3 )",
			want: "1 2 3 - +",
		},
		{
			expr: "1 + ( 2 - ( 3 + 1 ) )",
			want: "1 2 3 1 + - +",
		},
	}

	for i, tc := range tt {
		tc := tc
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got, err := parse(tc.expr)

			if tc.wantErr == nil {
				require.Equal(t, tc.want, got)
				require.NoError(t, err)
			}

			if tc.wantErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.wantErr)
			}
		})
	}

}
