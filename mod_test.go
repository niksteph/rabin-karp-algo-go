package rabinkarp

import (
	"testing"
)

const formatMod string = "%d mod %d = %d, but got %d"

func TestModSmaller(t *testing.T) {
	a, n := 5, 13
	want := 5
	got := mod(a,n)
	if got != want {
		t.Errorf(formatMod, a, n, want, got)
	}
}

func TestModGreater(t *testing.T) {
	a, n := 15, 13
	want := 2
	got := mod(a,n)
	if got != want {
		t.Errorf(formatMod, a, n, want, got)
	}
}

func TestModZero(t *testing.T) {
	a, n := 26, 13
	want := 0
	got := mod(a,n)
	if got != want {
		t.Errorf(formatMod, a, n, want, got)
	}
}

func TestModNegative(t *testing.T) {
	a, n := -5, 13
	want := 8
	got := mod(a,n)
	if got != want {
		t.Errorf(formatMod, a, n, want, got)
	}
}

func TestModGreaterNegative(t *testing.T) {
	a, n := -15, 13
	want := 11
	got := mod(a,n)
	if got != want {
		t.Errorf(formatMod, a, n, want, got)
	}
}

func TestModEvenGreaterNegative(t *testing.T) {
	a, n := -421278, 19
	want := 9
	got := mod(a,n)
	if got != want {
		t.Errorf(formatMod, a, n, want, got)
	}
}
