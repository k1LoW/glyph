package glyph

import (
	"fmt"
	"sort"
)

// Map is icon (glyph) map
type Map struct {
	opts   []Option
	glyphs map[string]*Glyph
}

func (m *Map) Keys() []string {
	keys := []string{}
	for k := range m.glyphs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (m *Map) Get(k string) (*Glyph, error) {
	g, ok := m.glyphs[k]
	if !ok {
		return nil, fmt.Errorf("invalid key: %s", k)
	}
	for _, opt := range m.opts {
		if err := opt(g); err != nil {
			return nil, err
		}
	}
	return g, nil
}

// NewMap return *Map
func NewMap(opts ...Option) *Map {
	return &Map{
		opts:   opts,
		glyphs: make(map[string]*Glyph),
	}
}

// NewMapWithIncluded *Map with Included icons
func NewMapWithIncluded(opts ...Option) *Map {
	m := NewMap(opts...)
	for k, sg := range Included() {
		g, err := sg.ToGlyph()
		if err != nil {
			panic(err)
		}
		m.glyphs[k] = g
	}
	return m
}
