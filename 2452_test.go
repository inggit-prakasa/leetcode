package main

import "testing"

func Test_twoEditWords(t *testing.T) {
	tests := []struct {
		name       string
		queries    []string
		dictionary []string
		expected   []string
	}{
		{
			name:       "example 1",
			queries:    []string{"word", "note", "ants", "wood"},
			dictionary: []string{"wood", "joke", "moat"},
			expected:   []string{"word", "note", "wood"},
		},
		{
			name:       "example 2",
			queries:    []string{"yes"},
			dictionary: []string{"not"},
			expected:   []string{},
		},
		{
			name:       "example 3",
			queries:    []string{"tsl", "sri", "yyy", "rbc", "dda", "qus", "hyb", "ilu", "ahd"},
			dictionary: []string{"uyj", "bug", "dba", "xbe", "blu", "wuo", "tsf", "tga"},
			expected:   []string{"tsl", "yyy", "rbc", "dda", "qus", "hyb", "ilu"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoEditWords(tt.queries, tt.dictionary)
			if !compareStringSlices(result, tt.expected) {
				t.Errorf("twoEditWords(%v, %v) = %v; want %v", tt.queries, tt.dictionary, result, tt.expected)
			}
		})
	}
}

func compareStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
