package glyph

import (
	"fmt"
	"strings"
)

type LineAndOpts string

func (l LineAndOpts) Parse() ([]string, []string, error) {
	points := []string{}
	opts := []string{}
	ps := GetPoints()
	for _, s := range strings.Split(string(l), " ") {
		_, err := ps.Get(s)
		if err == nil {
			points = append(points, s)
			continue
		}
		if strings.Count(s, ":") != 1 {
			return nil, nil, fmt.Errorf("invalid option: %s", s)
		}
		opts = append(opts, s)
	}
	return points, opts, nil
}

type SubGlyph struct {
	lines []LineAndOpts
}

func (s SubGlyph) ToGlyph() (*Glyph, error) {
	g, err := New()
	if err != nil {
		return nil, err
	}
	for _, l := range s.lines {
		points, opts, err := l.Parse()
		if err != nil {
			return nil, err
		}
		if err := g.AddLine(points, opts...); err != nil {
			return nil, err
		}
	}
	return g, nil
}

// NewSubGlyph create new SubGlyph
func NewSubGlyph(lines []LineAndOpts) SubGlyph {
	return SubGlyph{
		lines: lines,
	}
}
