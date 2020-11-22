// +build ignore

package main

import (
	"os"
	"strings"

	"github.com/k1LoW/glyph"
)

func main() {
	g, _ := glyph.New(glyph.Witdh(500.0), glyph.Height(500.0), glyph.LineWitdh(0.5), glyph.Color("#DFDFDF"))
	_ = g.AddLine(strings.Split("a0 f0 k0 k5 fa a5 a0", " "))
	_ = g.AddLine(strings.Split("a1 g0", " "))
	_ = g.AddLine(strings.Split("a2 h0", " "))
	_ = g.AddLine(strings.Split("a3 i0", " "))
	_ = g.AddLine(strings.Split("a4 j0", " "))
	_ = g.AddLine(strings.Split("a5 k0", " "))
	_ = g.AddLine(strings.Split("a0 k5", " "))
	_ = g.AddLine(strings.Split("a1 j6", " "))
	_ = g.AddLine(strings.Split("a2 i7", " "))
	_ = g.AddLine(strings.Split("a3 h8", " "))
	_ = g.AddLine(strings.Split("a4 g9", " "))
	_ = g.AddLine(strings.Split("k1 e0", " "))
	_ = g.AddLine(strings.Split("k2 d0", " "))
	_ = g.AddLine(strings.Split("k3 c0", " "))
	_ = g.AddLine(strings.Split("k4 b0", " "))
	_ = g.AddLine(strings.Split("k1 b6", " "))
	_ = g.AddLine(strings.Split("k2 c7", " "))
	_ = g.AddLine(strings.Split("k3 d8", " "))
	_ = g.AddLine(strings.Split("k4 e9", " "))
	_ = g.AddLine(strings.Split("b0 b6", " "))
	_ = g.AddLine(strings.Split("c0 c7", " "))
	_ = g.AddLine(strings.Split("d0 d8", " "))
	_ = g.AddLine(strings.Split("e0 e9", " "))
	_ = g.AddLine(strings.Split("f0 fa", " "))
	_ = g.AddLine(strings.Split("g0 g9", " "))
	_ = g.AddLine(strings.Split("h0 h8", " "))
	_ = g.AddLine(strings.Split("i0 i7", " "))
	_ = g.AddLine(strings.Split("j0 j6", " "))

	points := glyph.GetPoints()
	for k, _ := range points {
		_ = g.AddText(k, k, "font-size:3", "text-anchor:middle", "font-color:#333333")
	}

	if err := g.Write(os.Stdout); err != nil {
		panic(err)
	}
}
