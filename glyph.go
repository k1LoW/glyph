package glyph

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"strconv"
	"strings"

	svgo "github.com/ajstarks/svgo/float"
	"github.com/beta/freetype/truetype"
	"github.com/elliotchance/orderedmap"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/math/fixed"
)

const defaultColor = "#4B75B9"
const defaultFillColor = "#FFFFFF"
const defaultFontSize = "15.0"

var defaultLineOpts = []string{
	fmt.Sprintf("stroke:%s", defaultColor),
	"stroke-width:4.0",
	"stroke-linecap:round",
	"stroke-linejoin:round",
	fmt.Sprintf("fill:%s", defaultFillColor),
}

var defaultTextOpts = []string{
	fmt.Sprintf("fill:%s", defaultColor),
	"font-anchor:middle",
	"font-weight:bold",
	"font-family:Arial,sans-serif",
	fmt.Sprintf("font-size:%s", defaultFontSize),
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
	w, h, minx, miny, vw, vh float64
	lineOpts                 *orderedmap.OrderedMap
	lines                    []*Line
	texts                    []*Text
	textOpts                 *orderedmap.OrderedMap
	showCoordinates          bool
}

type Option func(*Glyph) error

// Width set SVG width
func Width(w float64) Option {
	return func(g *Glyph) error {
		g.w = w
		return nil
	}
}

// Height set SVG height
func Height(h float64) Option {
	return func(g *Glyph) error {
		g.h = h
		return nil
	}
}

// Color set SVG line 'stroke'
func Color(c string) Option {
	return func(g *Glyph) error {
		g.lineOpts.Set("stroke", c)
		return nil
	}
}

// FillColor set SVG line 'fill'
func FillColor(c string) Option {
	return func(g *Glyph) error {
		g.lineOpts.Set("fill", c)
		return nil
	}
}

// TextColor set SVG text 'fill'
func TextColor(c string) Option {
	return func(g *Glyph) error {
		g.textOpts.Set("fill", c)
		return nil
	}
}

// LineOpt set SVG line option
func LineOpt(opt string) Option {
	return func(g *Glyph) error {
		splited := strings.Split(opt, ":")
		g.lineOpts.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
		return nil
	}
}

// TextOpt set SVG text option
func TextOpt(opt string) Option {
	return func(g *Glyph) error {
		splited := strings.Split(opt, ":")
		g.textOpts.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
		return nil
	}
}

// LineOpts set SVG line options
func LineOpts(opts []string) Option {
	return func(g *Glyph) error {
		for _, opt := range opts {
			splited := strings.Split(opt, ":")
			g.lineOpts.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
		}
		return nil
	}
}

// TextOpts set SVG text options
func TextOpts(opts []string) Option {
	return func(g *Glyph) error {
		for _, opt := range opts {
			splited := strings.Split(opt, ":")
			g.textOpts.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
		}
		return nil
	}
}

// New return *Glyph
func New(opts ...Option) (*Glyph, error) {
	g := &Glyph{
		w:               110.0,
		h:               110.0,
		minx:            0.0,
		miny:            0.0,
		vw:              110.0,
		vh:              110.0,
		lineOpts:        orderedmap.NewOrderedMap(),
		textOpts:        orderedmap.NewOrderedMap(),
		showCoordinates: false,
	}
	for _, opt := range defaultLineOpts {
		splited := strings.Split(opt, ":")
		g.lineOpts.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
	}
	for _, opt := range defaultTextOpts {
		splited := strings.Split(opt, ":")
		g.textOpts.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
	}

	for _, opt := range opts {
		if err := opt(g); err != nil {
			return nil, err
		}
	}

	return g, nil
}

// Line draw line
func (g *Glyph) Line(points []string, opts ...string) error {
	m := orderedmap.NewOrderedMap()
	l := &Line{}
	ps := GetPoints()
	for _, k := range points {
		p, err := ps.Get(k)
		if err != nil {
			return err
		}
		l.points = append(l.points, p)
	}
	for _, opt := range opts {
		splited := strings.Split(opt, ":")
		m.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
	}
	l.opts = m
	g.lines = append(g.lines, l)
	return nil
}

// Line draw text
func (g *Glyph) Text(text, point string, opts ...string) error {
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

func (g *Glyph) HideCoordinates() error {
	g.showCoordinates = false
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
		m := orderedmap.NewOrderedMap()
		for _, k := range g.lineOpts.Keys() {
			v, _ := g.lineOpts.Get(k)
			m.Set(k, v.(string))
		}
		for _, k := range l.opts.Keys() {
			v, _ := l.opts.Get(k)
			m.Set(k, v.(string))
		}
		if l.points[0].X == l.points[len(l.points)-1].X && l.points[0].Y == l.points[len(l.points)-1].Y {
			// polygon
			if _, exist := m.Get("fill-opacity"); !exist {
				m.Set("fill-opacity", "1.0")
			}
		} else {
			// line, polyline
			if _, exist := m.Get("fill-opacity"); !exist {
				m.Set("fill-opacity", "0.0")
			}
			_ = m.Delete("fill")
		}

		// generate SVG options
		opts := []string{}
		for _, k := range m.Keys() {
			v, _ := m.Get(k)
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
		m := orderedmap.NewOrderedMap()
		for _, k := range g.textOpts.Keys() {
			v, _ := g.textOpts.Get(k)
			m.Set(k, v.(string))
		}
		for _, k := range t.opts.Keys() {
			v, _ := t.opts.Get(k)
			m.Set(k, v.(string))
		}
		opts := []string{}
		for _, k := range m.Keys() {
			v, _ := m.Get(k)
			opts = append(opts, fmt.Sprintf("%s:%s", k, v.(string)))
		}
		svg.Text(t.point.X, t.point.Y, t.text, strings.Join(opts, "; "))
	}

	svg.End()
	return nil
}

func (g *Glyph) writePNG(w io.Writer) error {
	svgbuf := new(bytes.Buffer)

	// oksvg workaround: stroke-width
	sw, exist := g.lineOpts.Get("stroke-width")
	if exist {
		fsw, err := strconv.ParseFloat(sw.(string), 64)
		if err != nil {
			return err
		}
		g.lineOpts.Set("stroke-width", strconv.FormatFloat(fsw*(g.w/g.vw), 'f', -1, 64))
	}

	if err := g.Write(svgbuf); err != nil {
		return err
	}

	// revert oksvg workaround: stroke-width
	if exist {
		g.lineOpts.Set("stroke-width", sw.(string))
	}

	icon, err := oksvg.ReadIconStream(svgbuf)
	if err != nil {
		return err
	}
	icon.SetTarget(0, 0, g.w, g.h)
	rgba := image.NewRGBA(image.Rect(0, 0, round(g.w), round(g.h)))
	icon.Draw(rasterx.NewDasher(round(g.w), round(g.h), rasterx.NewScannerGV(round(g.w), round(g.h), rgba, rgba.Bounds())), 1)

	// oksvg workaround: text
	for _, t := range g.texts {
		m := orderedmap.NewOrderedMap()
		for _, k := range g.textOpts.Keys() {
			v, _ := g.textOpts.Get(k)
			m.Set(k, v.(string))
		}
		for _, k := range t.opts.Keys() {
			v, _ := t.opts.Get(k)
			m.Set(k, v.(string))
		}
		size := defaultFontSize
		sizei, exist := m.Get("font-size")
		if exist {
			size = sizei.(string)
		}
		clr := defaultColor
		clri, exist := m.Get("fill")
		if exist {
			clr = clri.(string)
		}
		if err := g.addTextToImage(rgba, t.point.X, t.point.Y, t.text, size, clr); err != nil {
			return err
		}
	}

	return png.Encode(w, rgba)
}

func (g *Glyph) addTextToImage(img *image.RGBA, x, y float64, text, size, clr string) error {
	size64, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return err
	}

	c, err := parseColor(clr)
	if err != nil {
		return err
	}

	to := &truetype.Options{
		Size:              size64 * (g.w / g.vw),
		DPI:               0,
		Hinting:           font.HintingNone,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}
	f, err := truetype.Parse(gobold.TTF)
	if err != nil {
		return err
	}
	face := truetype.NewFace(f, to)

	dr := &font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{c},
		Face: face,
		Dot: fixed.Point26_6{
			X: fixed.I(round(x * (g.w / g.vw))),
			Y: fixed.I(round(y * (g.w / g.vw))),
		},
	}

	dr.DrawString(text)

	return nil
}

func round(v float64) int {
	if v < 0 {
		return int(v - 0.5)
	}
	return int(v + 0.5)
}

func parseColor(s string) (color.RGBA, error) {
	c := color.RGBA{
		A: 0xff,
	}
	switch len(s) {
	case 7:
		if _, err := fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B); err != nil {
			return c, fmt.Errorf("invalid color: %s", s)
		}
	case 4:
		if _, err := fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B); err != nil {
			return c, fmt.Errorf("invalid color: %s", s)
		}
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		return c, fmt.Errorf("invalid color: %s", s)
	}
	return c, nil
}
