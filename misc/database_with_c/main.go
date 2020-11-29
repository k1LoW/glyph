// +build ignore

package main

import (
	"os"

	"github.com/k1LoW/glyph"
)

func main() {
	g, _ := glyph.Included().Get("database")
	_ = glyph.Width(400.0)(g)
	_ = glyph.Height(400.0)(g)
	_ = g.ShowCoordinates()
	if err := g.Write(os.Stdout); err != nil {
		panic(err)
	}
}
