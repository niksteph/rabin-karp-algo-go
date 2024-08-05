package rabinkarp

import (
	"testing"
)

func TestFindSad(t *testing.T) {
    got := find("sadbutsad", "sad")
    if got != 0 {
        t.Errorf("got %v, want 0", got)
    }
}

func TestFindLeeto(t *testing.T) {
    got := find("leetcode", "leeto")
    if got != -1 {
        t.Errorf("got %v, want -1", got)
    }
}

func TestFindIssipi(t *testing.T) {
    got := find("mississippi", "issipi")
    if got != -1 {
        t.Errorf("got %v, want -1", got)
    }
}

func TestFindAaaa(t *testing.T) {
    got := find("aaa", "aaaa")
    if got != -1 {
        t.Errorf("got %v, want -1", got)
    }
}

func TestFindA(t *testing.T) {
    got := find("a", "a")
    if got != 0 {
        t.Errorf("got %v, want 0", got)
    }
}

func TestFindLl(t *testing.T) {
    want := 2
    got := find("hello", "ll")
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}

func TestFindAbabcaabc(t *testing.T) {
    want := 6
    got := find("ababcaababcaabc", "ababcaabc")
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}