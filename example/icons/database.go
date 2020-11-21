package main

import (
	"os"
	"strings"

	"github.com/k1LoW/glyph"
)

func main() {
	g := glyph.New()
	_ = g.AddLine(strings.Split("b0 d0 h0 j0 h2 d2 b0", " "))
	_ = g.AddLine(strings.Split("b0 b6 d8 h8 j6 j0", " "))
	_ = g.AddLine(strings.Split("b0 d2 h2 j0", " "))
	_ = g.AddLine(strings.Split("b2 d4 h4 j2", " "))
	_ = g.AddLine(strings.Split("b4 d6 h6 j4", " "))
	if err := g.Write(os.Stdout); err != nil {
		panic(err)
	}
}
