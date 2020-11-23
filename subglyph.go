package glyph

import (
	"fmt"
	"strings"

	"github.com/elliotchance/orderedmap"
)

type LineAndOpts string

func (l LineAndOpts) Parse() (*Line, error) {
	line := &Line{
		opts: orderedmap.NewOrderedMap(),
	}
	ps := GetPoints()
	for _, s := range strings.Split(string(l), " ") {
		p, err := ps.Get(s)
		if err == nil {
			line.points = append(line.points, p)
			continue
		}
		splited := strings.Split(s, ":")
		if len(splited) != 2 {
			return nil, fmt.Errorf("invalid option: %s", s)
		}
		line.opts.Set(strings.Trim(splited[0], " ;"), strings.Trim(splited[1], " ;"))
	}
	return line, nil
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
		line, err := l.Parse()
		if err != nil {
			return nil, err
		}
		g.lines = append(g.lines, line)
	}
	return g, nil
}

// NewSubGlyph create new SubGlyph
func NewSubGlyph(lines []LineAndOpts) SubGlyph {
	return SubGlyph{
		lines: lines,
	}
}
