package glyph

import (
	"strings"

	"github.com/elliotchance/orderedmap"
)

func (g *Glyph) addCoordinateLines() error {
	opts := orderedmap.NewOrderedMap()
	opts.Set("stroke", "#DFDFDF")
	opts.Set("stroke-linecap", "round")
	opts.Set("stroke-linejoin", "round")
	opts.Set("stroke-width", "0.5")
	opts.Set("fill-opacity", "0.0")
	lines := []*Line{}
	pointSets := []string{
		"a0 f0 k0 k5 fa a5 a0",
		"a1 g0",
		"a2 h0",
		"a3 i0",
		"a4 j0",
		"a5 k0",
		"a0 k5",
		"a1 j6",
		"a2 i7",
		"a3 h8",
		"a4 g9",
		"k1 e0",
		"k2 d0",
		"k3 c0",
		"k4 b0",
		"k1 b6",
		"k2 c7",
		"k3 d8",
		"k4 e9",
		"b0 b6",
		"c0 c7",
		"d0 d8",
		"e0 e9",
		"f0 fa",
		"g0 g9",
		"h0 h8",
		"i0 i7",
		"j0 j6",
	}
	ps := GetPoints()
	for _, psets := range pointSets {
		l := &Line{
			opts: opts,
		}
		for _, k := range strings.Split(psets, " ") {
			p, err := ps.Get(k)
			if err != nil {
				return err
			}
			l.points = append(l.points, p)
		}
		lines = append(lines, l)
	}
	g.lines = append(lines, g.lines...)
	return nil
}

func (g *Glyph) addCoordinateTexts() error {
	opts := orderedmap.NewOrderedMap()
	opts.Set("font-size", "3")
	opts.Set("text-anchor", "middle")
	opts.Set("font-color", "#333333")
	texts := []*Text{}
	ps := GetPoints()
	for _, k := range ps.Keys() {
		p, err := ps.Get(k)
		if err != nil {
			return err
		}
		texts = append(texts, &Text{
			point: p,
			text:  k,
			opts:  opts,
		})
	}
	g.texts = append(g.texts, texts...)
	return nil
}
