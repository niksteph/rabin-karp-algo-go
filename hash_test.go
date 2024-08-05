package rabinkarp

import (
	"strconv"
	"testing"
	"slices"
)

const formatHash string = "hash(%q) base %d mod %d = %d, but got %d"

func TestHashA(t *testing.T) {
	s := "a"
	base, mod := 256, 19
	want := int('a' % 19)
	got := hash(s, base, mod)
	if got != want {
		t.Errorf(formatHash, s, base, mod, want, got)
	}
}

func TestHash2a(t *testing.T) {
	s := "aa"
	base, mod := 256, 19
	want := int('a'+256*'a') % 19
	got := hash(s, base, mod)
	if got != want {
		t.Errorf(formatHash, s, base, mod, want, got)
	}
}

func TestHash3a(t *testing.T) {
	s := "aaa"
	base, mod := 256, 19
	want := int('a'+256*('a'+256*'a')) % 19
	got := hash(s, base, mod)
	if got != want {
		t.Errorf(formatHash, s, base, mod, want, got)
	}
}

func TestHash9a(t *testing.T) {
	s := "aaaaaaaaa"
	base, mod := 256, 19
	want := 0
	got := hash(s, base, mod)
	if got != want {
		t.Errorf(formatHash, s, base, mod, want, got)
	}
}

func TestHashNums(t *testing.T) {
	s := string([]byte{3, 1, 4, 1, 5})
	base, mod := 10, 13
	want := 7
	got := hash(s, base, mod)
	if got != want {
		t.Errorf(formatHash, s, base, mod, want, got)
	}
}

func TestHashNums2(t *testing.T) {
	s := string([]byte{1, 4, 1, 5, 2})
	base, mod := 10, 13
	want := 8
	got := hash(s, base, mod)
	if got != want {
		t.Errorf(formatHash, s, base, mod, want, got)
	}
}

func TestHashAbabcaabc(t *testing.T) {
	s := "ababcaabc"
	base, mod := 256, 19
	want := 5
	got := hash(s, base, mod)
	if got != want {
		t.Errorf(formatHash, s, base, mod, want, got)
	}
}

// ---

const formatRollHash string = "rollHash(%v[%v]%v) base %d mod %d = %d, but got %d"

func TestRollHashNums(t *testing.T) {
	mid := "1415" // example from clsr algo 32.4
	prev := 7
	var oldHigh, newLow byte = 3, 2
	base, mod, order := 10, 13, 5
	want := 8
	got := rollHash(prev, oldHigh, newLow, base, mod, order)
	if got != want {
		t.Errorf(formatRollHash, strconv.Itoa(int(oldHigh)), mid, strconv.Itoa(int(newLow)), base, mod, want, got)
	}
}

func TestRollHashSameAsHash(t *testing.T) {
	s := "ababcaababcaabc"
	n := len(s)
	m := 9
	base, mod := 256, 19
	hashes := make([]int, n-m)
	for i, _ := range s[:n-m] {
		hashes[i] = hash(s[i:i+m], base, mod)
	}
	rollHashes := make([]int, n-m)
	rollHashes[0] = hash(s[:m], base, mod)
	for i := 1; i < n-m; i++ {
		rollHashes[i] = rollHash(rollHashes[i-1], s[i-1], s[i-1+m], base, mod, m)
	}
	if !slices.Equal(hashes, rollHashes) {
		t.Logf(" direct hashes: %v", hashes)
		t.Logf("rolling hashes: %v", rollHashes)
		t.Fail()
	}
}
