package glyph

import (
	"fmt"
	"sort"
)

var included = []struct {
	key     string
	lines   []string
	texts   []string
	aliases []string
}{
	{
		"db",
		[]string{
			"b0 d0 h0 j0 j6 h8 d8 b6 b0",
			"b0 d2 h2 j0",
			"b2 d4 h4 j2",
			"b4 d6 h6 j4",
		},
		[]string{},
		[]string{"database"},
	},
	{
		"doc",
		[]string{
			"b0 b5 f9 f4 b0",
			"c0 c1 f4 f8 g8 g3 c0",
			"d0 d1 g3 g7 h7 h2 d0",
			"d0 d1 g3 g7 h7 h2 d0",
			"f1 f2 h2 h6 j6 j1 f1",
			"c2 d3",
			"c3 e5",
			"c4 e6",
			"d6 e7",
		},
		[]string{},
		[]string{"document"},
	},
	{
		"browser",
		[]string{
			"d0 d5 i7 i2 d0",
			"d1 i3",
			"e1 e2",
			"f2 f3",
			"g2 g3",
		},
		[]string{},
		[]string{},
	},
	{
		"cube",
		[]string{
			"f1 b1 b5 f9 j5 j1 f1",
			"c2 f5",
			"f5 f8",
			"g4 j1",
		},
		[]string{},
		[]string{},
	},
	{
		"lb",
		[]string{
			"a3 a4 c6 e6 e5 c3 a3",
			"f1 f2 g2 h1 h0 g0 f1",
			"h2 h3 i3 j2 j1 i1 h2",
			"i4 i5 j5 k4 k3 j3 i4",
			"d4 e4 e2 f2",
			"e5 h3",
			"e6 f6 h6 i5",
		},
		[]string{},
		[]string{},
	},
	{
		"terminal",
		[]string{
			fmt.Sprintf("d0 d5 i7 i2 d0 fill:%s", defaultColor),
			fmt.Sprintf("e2 f4 e4 stroke:%s", defaultFillColor),
			fmt.Sprintf("f5 g5 stroke:%s", defaultFillColor),
		},
		[]string{},
		[]string{},
	},
	{
		"proxy",
		[]string{
			"d3 d4 f6 h4 h3 f3 d3",
			"a2 a3 b4 c4 c3 b2 a2",
			"c5 c6 d7 e7 e6 d5 c5",
			"g1 g2 h2 i1 i0 h0 g1",
			"i2 i3 j3 k2 k1 j1 i2",
			"c4 d4",
			"d5 e5",
			"g3 h2",
			"h3 i2",
		},
		[]string{},
		[]string{},
	},
	{
		"cloud",
		[]string{
			"a3 a4 b5 j5 k4 k2 j2 i1 h1 f2 d3 c2 b2 b3 a3",
		},
		[]string{},
		[]string{},
	},
	{
		"metrics",
		[]string{
			"c0 c4 i7 i3",
			"c1 d4 e3 f6 g4 h1 i5 i7 c4 c1",
			"c2 d4 e4 f3 g6 h4 i6 i7 c4 c2",
			"c1 d4 e3 f6 g4 h1 i5",
		},
		[]string{},
		[]string{},
	},
	{
		"globe",
		[]string{
			"a1 a4 c7 e9 g9 i7 k4 k1 i0 g0 e0 c0 a1",
			"e0 b2 b4 d8",
			"e0 d2 d6 e9",
			"g0 h2 h6 g9",
			"g0 j2 j4 h8",
			"c0 c1 e3 g3 i1 i0",
			"a1 b3 d5 h5 j3 k1",
			"a3 b5 e8 g8 j5 k3",
		},
		[]string{},
		[]string{},
	},
	{
		"lb-l4",
		[]string{
			"a3 a4 c6 e6 e5 c3 a3",
			"f1 f2 g2 h1 h0 g0 f1",
			"h2 h3 i3 j2 j1 i1 h2",
			"i4 i5 j5 k4 k3 j3 i4",
			"d4 e4 e2 f2",
			"e5 h3",
			"e6 f6 h6 i5",
		},
		[]string{
			"L4 a1 font-size:22.0",
		},
		[]string{},
	},
	{
		"lb-l7",
		[]string{
			"a3 a4 c6 e6 e5 c3 a3",
			"f1 f2 g2 h1 h0 g0 f1",
			"h2 h3 i3 j2 j1 i1 h2",
			"i4 i5 j5 k4 k3 j3 i4",
			"d4 e4 e2 f2",
			"e5 h3",
			"e6 f6 h6 i5",
		},
		[]string{
			"L7 a1 font-size:22.0",
		},
		[]string{},
	},
	{
		"shield",
		[]string{
			"c1 c5 d7 f9 h7 i5 i1 g2 f2 e2 c1",
		},
		[]string{},
		[]string{},
	},
	{
		"blocks",
		[]string{
			"b1 b5 f9 j5 j1 f1 b1",
			"b3 f7 f9",
			"f7 j3",
			"b1 d3 d5",
			"d3 f3 f5",
			"d1 h3 h5",
			"h3 j1",
			"d5 f5 h5",
		},
		[]string{},
		[]string{},
	},
	{
		"cubes",
		[]string{
			"b3 b5 d7 f7 h7 j5 j3 h3 h1 f1 d1 d3 b3",
			"d1 f3 h1",
			"d3 f5 h3",
			"b3 d5 f5 h5 j3",
			"f3 f7",
			"d5 d7",
			"h5 h7",
		},
		[]string{},
		[]string{},
	},
	{
		"cubes2",
		[]string{
			"b1 b3 d5 d7 f9 h7 h5 j3 j1 h1 f3 d1 b1",
			"b1 d3 f3 h3 j1",
			"d3 d5 f5 h5 h3",
			"f3 f5",
			"d5 f7 h5",
			"f7 f9",
		},
		[]string{},
		[]string{},
	},
	{
		"gear",
		[]string{
			"c1 b2 c3 c4 b4 c6 d6 e7 e8 g8 g7 h6 i6 j4 i4 i3 j2 i1 h2 g2 g1 e1 e2 d2 c1",
			"e4 e5 f6 g5 g4 f4 e4",
		},
		[]string{},
		[]string{},
	},
	{
		"hex",
		[]string{
			"f1 b1 b5 f9 j5 j1 f1",
		},
		[]string{},
		[]string{"hexagon"},
	},
	{
		"hex2",
		[]string{
			"f1 b1 b5 f9 j5 j1 f1",
			fmt.Sprintf("d3 d5 f7 h5 h3 f3 d3 fill:%s", defaultColor),
		},
		[]string{},
		[]string{"hexagon2"},
	},
	{
		"cube2",
		[]string{
			"c2 c5 f8 i5 i2 f2 c2",
			"d3 f5 f7",
			"g4 i2",
		},
		[]string{},
		[]string{},
	},
	{
		"cube3",
		[]string{
			"d3 d5 f7 h5 h3 f3 d3",
			"d3 f5 h3",
			"f5 f7",
		},
		[]string{},
		[]string{},
	},
	{
		"cube4",
		[]string{
			"f1 b1 b5 f9 j5 j1 f1",
			"c2 c5 f8 i5 i2 f2 c2",
			"d3 f5 f7",
			"g4 i2",
		},
		[]string{},
		[]string{},
	},
	{
		"cube5",
		[]string{
			"f1 b1 b5 f9 j5 j1 f1",
			"d3 d5 f7 h5 h3 f3 d3",
			"d3 f5 h3",
			"f5 f7",
		},
		[]string{},
		[]string{},
	},
	{
		"text",
		[]string{
			"d1 d6 h8 h3 d1",
			"e3 f4",
			"e4 g5",
			"e5 g6",
			"f7 g7",
		},
		[]string{},
		[]string{"page"},
	},
}

var cSet Set

type Set map[string]*SubGlyph

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

// Included return included icon set
func Included() Set {
	if len(cSet) > 0 {
		return cSet
	}
	s := Set{}
	for _, g := range included {
		lo := []LineAndOpts{}
		for _, l := range g.lines {
			lo = append(lo, LineAndOpts(l))
		}
		to := []TextAndOpts{}
		for _, t := range g.texts {
			to = append(to, TextAndOpts(t))
		}
		sg := NewSub().Lines(lo).Texts(to)
		s[g.key] = sg
		for _, a := range g.aliases {
			s[a] = sg
		}
	}
	cSet = s
	return s
}

// IncludedAliases return included icon aliases
func IncludedAliases() map[string][]string {
	aliases := map[string][]string{}
	for _, g := range included {
		aliases[g.key] = g.aliases
	}
	return aliases
}
