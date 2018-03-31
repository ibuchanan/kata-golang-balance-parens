package balance

import (
	"testing"
)

func testIsBalanced(t *testing.T) {
	if !isBalanced() {
		t.Error("isBalanced() was false.")
	}
}
