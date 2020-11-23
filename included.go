package glyph

import (
	"fmt"
	"sort"
)

type Set map[string]SubGlyph

func (s Set) Get(k string) (*Glyph, error) {
	sg, ok := s[k]
	if !ok {
		return nil, fmt.Errorf("invalid key: %s", k)
	}
	return sg.ToGlyph()
}

func (s Set) Keys() []string {
	keys := []string{}
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

var Included = Set{
	"database": NewSubGlyph([]LineAndOpts{
		"b0 d0 h0 j0 j6 h8 d8 b6 b0",
		"b0 d2 h2 j0",
		"b2 d4 h4 j2",
		"b4 d6 h6 j4",
	}),
	"documents": NewSubGlyph([]LineAndOpts{
		"b0 b5 f9 f4 b0",
		"c0 c1 f4 f8 g8 g3 c0",
		"d0 d1 g3 g7 h7 h2 d0",
		"d0 d1 g3 g7 h7 h2 d0",
		"f1 f2 h2 h6 j6 j1 f1",
	}),
	"browser": NewSubGlyph([]LineAndOpts{
		"d0 d5 i7 i2 d0",
		"d1 i3",
		"e1 e2",
		"f2 f3",
		"g2 g3",
	}),
	"cube": NewSubGlyph([]LineAndOpts{
		"f1 b1 b5 f9 j5 j1 f1",
		"b1 f5",
		"f5 f9",
		"f5 j1",
	}),
}
