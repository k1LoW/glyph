package glyph

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"strconv"
	"strings"

	svgo "github.com/ajstarks/svgo/float"
	"github.com/elliotchance/orderedmap"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

var defaultLineOpts = []string{
	"stroke:#4B75B9",
	"stroke-linecap:round",
	"stroke-linejoin:round",
	"fill:#FFFFFF",
	"fill-opacity:0.0",
}

type Line struct {
	points []*Point
	opts   *orderedmap.OrderedMap
}

type Text struct {
	point *Point
	text  string
	opts  *orderedmap.OrderedMap
}

type Glyph struct {
	w, h, minx, miny, vw, vh, lw float64
	lineOpts                     *orderedmap.OrderedMap
	lines                        []*Line
	texts                        []*Text
	showCoordinates              bool
}

type Option func(*Glyph) error

func Witdh(w float64) Option {
	return func(g *Glyph) error {
		g.w = w
		return nil
	}
}

func Height(h float64) Option {
	return func(g *Glyph) error {
		g.h = h
		return nil
	}
}

func Color(c string) Option {
	return func(g *Glyph) error {
		g.lineOpts.Set("stroke", c)
		return nil
	}
}

func FillColor(c string) Option {
	return func(g *Glyph) error {
		g.lineOpts.Set("fill", c)
		return nil
	}
}

func LineWitdh(lw float64) Option {
	return func(g *Glyph) error {
		g.lw = lw
		return nil
	}
}

func New(opts ...Option) (*Glyph, error) {
	g := &Glyph{
		w:               110.0,
		h:               110.0,
		minx:            0.0,
		miny:            0.0,
		vw:              110.0,
		vh:              110.0,
		lw:              4.0,
		lineOpts:        orderedmap.NewOrderedMap(),
		showCoordinates: false,
	}
	for _, opt := range defaultLineOpts {
		splited := strings.Split(opt, ":")
		g.lineOpts.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
	}

	for _, opt := range opts {
		if err := opt(g); err != nil {
			return nil, err
		}
	}

	// Set stroke-width using g.lw
	g.lineOpts.Set("stroke-width", strconv.FormatFloat(g.lw, 'f', -1, 64))

	return g, nil
}

func (g *Glyph) AddLine(points []string, opts ...string) error {
	m := orderedmap.NewOrderedMap()
	for _, k := range g.lineOpts.Keys() {
		v, _ := g.lineOpts.Get(k)
		m.Set(k, v.(string))
	}
	l := &Line{}
	ps := GetPoints()
	for _, k := range points {
		p, err := ps.Get(k)
		if err != nil {
			return err
		}
		l.points = append(l.points, p)
	}
	if l.points[0].X == l.points[len(l.points)-1].X && l.points[0].Y == l.points[len(l.points)-1].Y {
		// polygon
		m.Set("fill-opacity", "1.0")
	} else {
		// line, polyline
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

func (g *Glyph) AddText(point, text string, opts ...string) error {
	ps := GetPoints()
	p, err := ps.Get(point)
	if err != nil {
		return err
	}
	m := orderedmap.NewOrderedMap()
	for _, opt := range opts {
		splited := strings.Split(opt, ":")
		m.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
	}
	t := &Text{
		point: p,
		text:  text,
		opts:  m,
	}
	g.texts = append(g.texts, t)
	return nil
}

func (g *Glyph) ShowCoordinates() error {
	g.showCoordinates = true
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
	svg.StartviewUnit(g.w, g.h, "px", g.minx, g.miny, g.vw, g.vh)
	if g.showCoordinates {
		if err := g.addCoordinateLines(); err != nil {
			return err
		}
		if err := g.addCoordinateTexts(); err != nil {
			return err
		}
	}

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
		} else if len(l.points) == 2 {
			// line
			svg.Line(l.points[0].X, l.points[0].Y, l.points[1].X, l.points[1].Y, strings.Join(opts, "; "))
		} else {
			// polyline
			svg.Polyline(x, y, strings.Join(opts, "; "))
		}
	}
	for _, t := range g.texts {
		opts := []string{}
		for _, k := range t.opts.Keys() {
			v, _ := t.opts.Get(k)
			opts = append(opts, fmt.Sprintf("%s:%s", k, v.(string)))
		}
		svg.Text(t.point.X, t.point.Y, t.text, strings.Join(opts, "; "))
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
