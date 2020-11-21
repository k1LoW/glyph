package glyph

import (
	"testing"
)

func TestNew(t *testing.T) {
	g, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if want := 110.0; g.w != want {
		t.Errorf("got %v\nwant %v", g.w, want)
	}
	if want := 6; g.lineOpts.Len() != want {
		t.Errorf("got %v\nwant %v", g.lineOpts.Len(), want)
	}
}
