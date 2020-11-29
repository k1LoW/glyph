// +build ignore

package main

import (
	"os"

	"github.com/k1LoW/glyph"
)

func main() {
	g, _ := glyph.New(glyph.Width(400.0), glyph.Height(400.0))
	_ = g.ShowCoordinates()
	if err := g.Write(os.Stdout); err != nil {
		panic(err)
	}
}
