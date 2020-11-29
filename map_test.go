package glyph

import (
	"bytes"
	"fmt"
	"strings"
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

func TestMapOption(t *testing.T) {
	want := "#FFFFFF"
	m := NewMapWithIncluded(Color(want))
	g, _ := m.Get("db")
	got, _ := g.lineOpts.Get("stroke")
	if got != want {
		t.Errorf("got %v\nwant %v", got, want)
	}
	buf := new(bytes.Buffer)
	if err := g.Write(buf); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), fmt.Sprintf("stroke:%s", want)) {
		t.Errorf("can not change color\n%v", buf.String())
	}
}
