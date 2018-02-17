package splitters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConservSplitting(t *testing.T) {
	cases := []struct {
		token    string
		expected []string
	}{
		{"spongebob_squarePants", []string{"spongebob", "square", "Pants"}},
		{"Extraordinaire", []string{"Extraordinaire"}},
		{"extraordinairE", []string{"extraordinair", "E"}},
		{"extraordinaire", []string{"extraordinaire"}},
		{"extra_ordinaire", []string{"extra", "ordinaire"}},
		{"leto2nd", []string{"leto", "2", "nd"}},
		{"brooklyn99", []string{"brooklyn", "99"}},
		{"mySQL", []string{"my", "SQL"}},
		{"mySql", []string{"my", "Sql"}},
		{"mySQl", []string{"my", "S", "Ql"}},
		{"9999", []string{"9999"}},
		{"", []string{""}},
	}

	conserv := new(Conserv)
	for _, c := range cases {
		got, err := conserv.Split(c.token)
		if err != nil {
			assert.Fail(t, "we shouldn't get any errors at this point", err)
		}

		assert.Equal(t, c.expected, got, "elements should match in number and order")
	}
}

func BenchmarkConservSplitting(b *testing.B) {
	conserv := new(Conserv)
	for i := 0; i < b.N; i++ {
		conserv.Split("spongebob_squarePants")
	}
}