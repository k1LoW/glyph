package glyph

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	g, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if want := 5; g.lineOpts.Len() != want {
		t.Errorf("got %v\nwant %v", g.lineOpts.Len(), want)
	}
}

func TestOptionWidth(t *testing.T) {
	g, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if want := 110.0; g.w != want {
		t.Errorf("got %v\nwant %v", g.w, want)
	}
	opt := Width(120.0)
	if err := opt(g); err != nil {
		t.Fatal(err)
	}
	if want := 120.0; g.w != want {
		t.Errorf("got %v\nwant %v", g.w, want)
	}
	if want := 5; g.lineOpts.Len() != want {
		t.Errorf("got %v\nwant %v", g.lineOpts.Len(), want)
	}
}

func TestOptionColor(t *testing.T) {
	g, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if err := g.Line([]string{"a0", "a1"}); err != nil {
		t.Fatal(err)
	}
	{
		want := defaultColor
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
	{
		want := "#FFFFFF"
		opt := Color(want)
		if err := opt(g); err != nil {
			t.Fatal(err)
		}
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
		if want := 5; g.lineOpts.Len() != want {
			t.Errorf("got %v\nwant %v", g.lineOpts.Len(), want)
		}
	}
}
