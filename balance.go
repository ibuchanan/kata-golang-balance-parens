package balance

import "strings"

var pairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func parens() string {
	var p []rune
	for k, v := range pairs {
		p = append(p, k, v)
	}
	return string(p)
}

func top(stack []rune) (rune, int) {
	var last int
	var top rune
	if len(stack) > 0 {
		last = len(stack) - 1
		top = stack[last]
	}
	return top, last
}

func isBalanced(expression string) bool {
	var stack []rune
	for _, c := range expression {
		top, last := top(stack)
		if strings.Contains(parens(), string(c)) {
			if c == pairs[top] {
				stack = stack[:last]
			} else {
				stack = append(stack, c)
			}
		}
	}
	return (len(stack) == 0)
}
