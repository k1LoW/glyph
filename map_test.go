package glyph

import (
	"testing"
)

func TestMap(t *testing.T) {
	{
		m := NewMap()
		got := len(m.Keys())
		want := 0
		if got != want {
			t.Errorf("got %v\nwant %v", got, want)
		}
	}

	{
		m := NewMapWithIncluded()
		got := len(m.Keys())
		want := len(Included())
		if got != want {
			t.Errorf("got %v\nwant %v", got, want)
		}
	}

	{
		m := NewMapWithIncluded()
		g, _ := New()
		m.Set("additional-glyph", g)
		got := len(m.Keys())
		want := len(Included()) + 1
		if got != want {
			t.Errorf("got %v\nwant %v", got, want)
		}
	}

	{
		m := NewMapWithIncluded()
		m.Delete("db")
		got := len(m.Keys())
		want := len(Included()) - 1
		if got != want {
			t.Errorf("got %v\nwant %v", got, want)
		}
	}
}
