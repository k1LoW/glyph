// +build ignore

package main

import (
	"os"
	"strings"

	"github.com/k1LoW/glyph"
)

func main() {
	g, _ := glyph.New(glyph.Width(400.0), glyph.Height(400.0))
	_ = g.Line(strings.Split("j1 h1 f2 b2 d3", " "))
	_ = g.Line(strings.Split("c6 g4", " "))
	_ = g.Line(strings.Split("g4 h7 h4 g4", " "))
	_ = g.Line(strings.Split("d3 g2 j2 g4 d3", " "))
	_ = g.Line(strings.Split("f3 f4 g4 h3 h2 g2 f3", " "))
	_ = g.Line(strings.Split("c5 b4 b5 c6", " "))
	_ = g.Write(os.Stdout)
}
