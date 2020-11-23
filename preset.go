package glyph

import (
	"fmt"
)

type Set map[string]SubGlyph

func (s Set) Get(k string) (*Glyph, error) {
	sg, ok := s[k]
	if !ok {
		return nil, fmt.Errorf("invalid key: %s", k)
	}
	return sg.ToGlyph()
}

var Preset = Set{
	"database": NewSubGlyph([]LineAndOpts{
		"b0 d0 h0 j0 j6 h8 d8 b6 b0",
		"b0 d2 h2 j0",
		"b2 d4 h4 j2",
		"b4 d6 h6 j4",
	}),
}
