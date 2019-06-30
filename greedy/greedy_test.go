package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit_ShouldReturnValidSplits(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  []string
	}{
		{"no_split", "car", []string{"car"}},
		{"by_lower_to_upper_case", "getString", []string{"get", "String"}},
		{"by_upper_to_lower_case", "GPSstate", []string{"GPS", "state"}},
		{"with_upper_case_and_softword_starting_with_upper_case", "ASTVisitor", []string{"AST", "Visitor"}},
		{"lowercase_softword", "notype", []string{"no", "type"}},
	}

	dicc := []string{"get", "string", "gps", "state", "ast", "visitor", "no", "type"}
	list := NewListBuilder().Dicctionary(dicc).Build()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Split(tt.token, list)

			assert.Equal(t, tt.want, got, "elements should match in number and order")
		})
	}
}

func BenchmarkGreedySplitting(b *testing.B) {
	/*emtpyList := make(map[string]interface{})
	greedy := NewGreedy(&emtpyList, &emtpyList, &emtpyList)
	for i := 0; i < b.N; i++ {
		Split("spongebob_squarePants", *greedy)
	}*/
}
