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

type TextAndOpts string

func (t TextAndOpts) Parse() (string, string, []string, error) {
	splited := strings.Split(string(t), " ")
	if len(splited) < 2 {
		return "", "", []string{}, fmt.Errorf("invalid text and opts: %s", t)
	}
	text := splited[0]
	point := splited[1]
	opts := splited[2:]
	ps := GetPoints()
	_, err := ps.Get(point)
	if err != nil {
		return "", "", []string{}, err
	}
	for _, s := range opts {
		if strings.Count(s, ":") != 1 {
			return "", "", []string{}, fmt.Errorf("invalid option: %s", s)
		}
	}
	return text, point, opts, nil
}

type SubGlyph struct {
	lines []LineAndOpts
	texts []TextAndOpts
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
	for _, t := range s.texts {
		text, point, opts, err := t.Parse()
		if err != nil {
			return nil, err
		}
		if err := g.AddText(text, point, opts...); err != nil {
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

func (s SubGlyph) Texts(texts []TextAndOpts) SubGlyph {
	s.texts = texts
	return s
}
