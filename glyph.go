package glyph

import (
	"io"
	"strings"

	svgo "github.com/ajstarks/svgo/float"
)

var defaultOpt = []string{"stroke:#4B75B9", "stroke-width:4", "stroke-linecap:round", "stroke-linejoin:round"}

type Line struct {
	points []*Point
}

type Glyph struct {
	w, h, minx, miny, vw, vh float64
	lines                    []*Line
}

func New() *Glyph {
	return &Glyph{
		w:    110.0,
		h:    110.0,
		minx: 0.0,
		miny: 0.0,
		vw:   110.0,
		vh:   110.0,
	}
}

func (g *Glyph) Draw(strp []string) error {
	l := &Line{}
	ps := points()
	for _, k := range strp {
		p, err := ps.Get(k)
		if err != nil {
			return err
		}
		l.points = append(l.points, p)
	}
	g.lines = append(g.lines, l)
	return nil
}

func (g *Glyph) Write(w io.Writer) error {
	svg := svgo.New(w)
	svg.Startview(g.w, g.h, g.minx, g.miny, g.vw, g.vh)
	for _, l := range g.lines {
		x := []float64{}
		y := []float64{}
		for _, p := range l.points {
			x = append(x, p.X)
			y = append(y, p.Y)
		}
		if l.points[0].X == l.points[len(l.points)-1].X && l.points[0].Y == l.points[len(l.points)-1].Y {
			// polygon
			opt := append(defaultOpt, "fill:#FFFFFF")
			svg.Polygon(x, y, strings.Join(opt, "; "))
		} else {
			// polyline
			opt := append(defaultOpt, "fill-opacity:0.0")
			svg.Polyline(x, y, strings.Join(opt, "; "))
		}
	}
	svg.End()
	return nil
}
