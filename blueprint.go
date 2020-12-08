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

type Blueprint struct {
	RawLines []LineAndOpts `json:"lines" yaml:"lines" toml:"lines"`
	RawTexts []TextAndOpts `json:"texts" yaml:"texts" toml:"texts"`
}

func (b Blueprint) ToGlyph() (*Glyph, error) {
	g, err := New()
	if err != nil {
		return nil, err
	}
	for _, l := range b.RawLines {
		points, opts, err := l.Parse()
		if err != nil {
			return nil, err
		}
		if err := g.Line(points, opts...); err != nil {
			return nil, err
		}
	}
	for _, t := range b.RawTexts {
		text, point, opts, err := t.Parse()
		if err != nil {
			return nil, err
		}
		if err := g.Text(text, point, opts...); err != nil {
			return nil, err
		}
	}
	return g, nil
}

// NewBlueprint create new Blueprint
func NewBlueprint() *Blueprint {
	return &Blueprint{}
}

func (b *Blueprint) Lines(lines []LineAndOpts) *Blueprint {
	b.RawLines = lines
	return b
}

func (b *Blueprint) Texts(texts []TextAndOpts) *Blueprint {
	b.RawTexts = texts
	return b
}
