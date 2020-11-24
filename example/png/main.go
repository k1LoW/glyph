// +build ignore

package main

import (
	"os"
	"strings"

	"github.com/k1LoW/glyph"
)

func main() {
	g, _ := glyph.New(glyph.Witdh(500.0), glyph.Height(500.0))
	_ = g.Line(strings.Split("b0 d0 h0 j0 j6 h8 d8 b6 b0", " "))
	_ = g.Line(strings.Split("b0 d2 h2 j0", " "))
	_ = g.Line(strings.Split("b2 d4 h4 j2", " "))
	_ = g.Line(strings.Split("b4 d6 h6 j4", " "))
	if err := g.WriteImage(os.Stdout); err != nil {
		panic(err)
	}
}
