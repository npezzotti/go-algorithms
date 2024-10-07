package accounts_merge

import (
	"reflect"
	"sort"
	"testing"
)

func TestAccountsMerge(t *testing.T) {
	tcases := []struct {
		input  [][]string
		output [][]string
	}{
		{
			input: [][]string{
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
			output: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
		{
			input: [][]string{
				{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
				{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
				{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
				{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
				{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
			},
			output: [][]string{
				{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
				{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
				{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
				{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
			},
		},
	}

	for _, tc := range tcases {
		output := accountsMerge(tc.input)

		// accountsMerge does not gurantee order of returned
		// account slices. Sort account slices by account name, 
		// falling back on first email if account names are equal.
		sort.Slice(output, func(i, j int) bool {
			if output[i][0] == output[j][0] {
				return output[i][1] < output[j][1]
			}

			return output[i][0] < output[j][0]
		})

		if !reflect.DeepEqual(output, tc.output) {
			t.Fatalf("expected %v, got %v", tc.output, output)
		}
	}
}
