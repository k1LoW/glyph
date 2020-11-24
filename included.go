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
	"db": NewSubGlyph([]LineAndOpts{
		"b0 d0 h0 j0 j6 h8 d8 b6 b0",
		"b0 d2 h2 j0",
		"b2 d4 h4 j2",
		"b4 d6 h6 j4",
	}),
	"doc": NewSubGlyph([]LineAndOpts{
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
		"c2 f5",
		"f5 f8",
		"g4 j1",
	}),
	"lb": NewSubGlyph([]LineAndOpts{
		"b4 a4 a5 b6 c6 c5 b4",
		"g0 f1 f2 g2 h1 h0 g0",
		"j3 i4 i5 j5 k4 k3 j3",
		"c5 d5 d2 f2",
		"d5 g7 i5",
	}),
	"terminal": NewSubGlyph([]LineAndOpts{
		LineAndOpts(fmt.Sprintf("d0 d5 i7 i2 d0 fill:%s", defaultColor)),
		LineAndOpts(fmt.Sprintf("e2 f4 e4 stroke:%s", defaultFillColor)),
		LineAndOpts(fmt.Sprintf("f5 g5 stroke:%s", defaultFillColor)),
	}),
	"proxy": NewSubGlyph([]LineAndOpts{
		"d3 d5 f7 h5 h3 f3 d3",
		"d3 f5 f6",
		"b4 d4",
		"b5 d5",
		"b6 e6",
		"f3 f2 h0",
		"h5 i5 k3",
	}),
	"cloud": NewSubGlyph([]LineAndOpts{
		"a3 a4 b5 j5 k4 k2 j2 i1 h1 f2 d3 c2 b2 b3 a3",
	}),
	"metrics": NewSubGlyph([]LineAndOpts{
		"c0 c4 i7 i3",
		"c1 d4 e3 f6 g4 h1 i5 i7 c4 c1",
		"c2 d4 e4 f3 g6 h4 i6",
	}),
}
