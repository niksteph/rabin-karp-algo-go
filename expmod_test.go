package rabinkarp

import (
	"testing"
)

const formatExpMod string = "%d^%d mod %d = %d, but got %d"

func TestExpMod(t *testing.T) {
	a, b, n := 7, 560, 561
	want := 1
	got := expMod(a,b,n)
	if got != want {
		t.Errorf(formatExpMod, a, b, n, want, got)
	}
}
