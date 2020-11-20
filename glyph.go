package glyph

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"strings"

	svgo "github.com/ajstarks/svgo/float"
	"github.com/elliotchance/orderedmap"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

var defaultLineOpt = []string{"stroke:#4B75B9", "stroke-width:4", "stroke-linecap:round", "stroke-linejoin:round"}

type Line struct {
	points []*Point
	opts   *orderedmap.OrderedMap
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

func (g *Glyph) AddLine(strp []string, opts ...string) error {
	m := orderedmap.NewOrderedMap()
	for _, opt := range defaultLineOpt {
		splited := strings.Split(opt, ":")
		m.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
	}
	l := &Line{}
	ps := points()
	for _, k := range strp {
		p, err := ps.Get(k)
		if err != nil {
			return err
		}
		l.points = append(l.points, p)
	}
	if l.points[0].X == l.points[len(l.points)-1].X && l.points[0].Y == l.points[len(l.points)-1].Y {
		// polygon
		m.Set("fill", "#FFFFFF")
	} else {
		// polyline
		m.Set("fill-opacity", "0.0")
	}
	for _, opt := range opts {
		splited := strings.Split(opt, ":")
		m.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
	}
	l.opts = m

	g.lines = append(g.lines, l)
	return nil
}

func (g *Glyph) Write(w io.Writer) error {
	return g.writeSVG(w)
}

func (g *Glyph) WriteImage(w io.Writer) error {
	return g.writePNG(w)
}

func (g *Glyph) writeSVG(w io.Writer) error {
	svg := svgo.New(w)
	svg.Startview(g.w, g.h, g.minx, g.miny, g.vw, g.vh)
	for _, l := range g.lines {
		x := []float64{}
		y := []float64{}
		for _, p := range l.points {
			x = append(x, p.X)
			y = append(y, p.Y)
		}
		opts := []string{}
		for _, k := range l.opts.Keys() {
			v, _ := l.opts.Get(k)
			opts = append(opts, fmt.Sprintf("%s:%s", k, v.(string)))
		}
		if l.points[0].X == l.points[len(l.points)-1].X && l.points[0].Y == l.points[len(l.points)-1].Y {
			// polygon
			svg.Polygon(x, y, strings.Join(opts, "; "))
		} else {
			// polyline
			svg.Polyline(x, y, strings.Join(opts, "; "))
		}
	}
	svg.End()
	return nil
}

func (g *Glyph) writePNG(w io.Writer) error {
	svgbuf := new(bytes.Buffer)
	if err := g.Write(svgbuf); err != nil {
		return err
	}
	icon, err := oksvg.ReadIconStream(svgbuf)
	if err != nil {
		return err
	}
	icon.SetTarget(0, 0, g.w, g.h)
	rgba := image.NewRGBA(image.Rect(0, 0, round(g.w), round(g.h)))
	icon.Draw(rasterx.NewDasher(round(g.w), round(g.h), rasterx.NewScannerGV(round(g.w), round(g.h), rgba, rgba.Bounds())), 1)
	return png.Encode(w, rgba)
}

func round(v float64) int {
	if v < 0 {
		return int(v - 0.5)
	}
	return int(v + 0.5)
}
