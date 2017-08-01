package main

import (
	"testing"
)

func TestParse(t *testing.T) {

	sample := []struct {
		src string
		exp []int
	}{
		{"Uhura", []int{0xF8E5, 0xF8D6, 0xF8E5, 0xF8E1, 0xF8D0}},
	}
	for _, v := range sample {
		tks, err := parse(v.src)
		if err != nil {
			t.Fatal(err)
		}
		if len(tks) != len(v.exp) {
			t.Errorf("expected %d got %d", len(v.exp), len(tks))
		}
		for k, tv := range tks {
			if tv.Hex() != v.exp[k] {
				t.Errorf("expected %d got %d", v.exp[k], tv.Hex())
			}
		}
	}
}
