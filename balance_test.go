package balance

import "testing"

func Test_parens(t *testing.T) {
	t.Run("all paren characters", func(t *testing.T) {
		allParens := "()[]{}"
		if got := parens(); got != allParens {
			t.Errorf("isBalanced() = %v, want %v", got, allParens)
		}
	})
}

func Test_top(t *testing.T) {
	tests := []struct {
		name  string
		stack []rune
		top   rune
		last  int
	}{
		{"empty array", make([]rune, 0), 0, 0},
		{"one element array", []rune{'a'}, 'a', 0},
		{"two element array", []rune{'a', 'b'}, 'b', 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if top, last := top(tt.stack); top != tt.top || last != tt.last {
				t.Errorf("top() = (%v, %v), want (%v, %v)", top, last, tt.top, tt.last)
			}
		})
	}
}

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		expression string
		want       bool
	}{
		{"no parens", true},
		{"(", false},
		{"()", true},
		{"(()", false},
		{"())", false},
		{"[", false},
		{"[]", true},
		{"[[]", false},
		{"[]]", false},
		{"{", false},
		{"{}", true},
		{"{{}", false},
		{"{}}", false},
		{"{{)(}}", false},
		{"({)}", false},
		{"[({})]", true},
		{"{}([])", true},
		{"{()}[[{}]]", true},
	}
	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			if got := isBalanced(tt.expression); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
