package glyph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPoints(t *testing.T) {
	got := points()
	if diff := cmp.Diff(got["a0"], &Point{X: 60.0 - float64(5)*dx, Y: 35.0}, nil); diff != "" {
		t.Errorf("%s", diff)
	}

	if diff := cmp.Diff(got["a5"], &Point{X: 60.0 - float64(5)*dx, Y: 85.0}, nil); diff != "" {
		t.Errorf("%s", diff)
	}

	if _, ok := got["a6"]; ok {
		t.Errorf("%v", got["a6"])
	}

	if diff := cmp.Diff(got["f0"], &Point{X: 60.0, Y: 10.0}, nil); diff != "" {
		t.Errorf("%s", diff)
	}

	if diff := cmp.Diff(got["fa"], &Point{X: 60.0, Y: 110.0}, nil); diff != "" {
		t.Errorf("%s", diff)
	}

	if _, ok := got["fb"]; ok {
		t.Errorf("%v", got["fb"])
	}

	if diff := cmp.Diff(got["k0"], &Point{X: 60.0 + float64(5)*dx, Y: 35.0}, nil); diff != "" {
		t.Errorf("%s", diff)
	}

	if diff := cmp.Diff(got["k5"], &Point{X: 60.0 + float64(5)*dx, Y: 85.0}, nil); diff != "" {
		t.Errorf("%s", diff)
	}

	if _, ok := got["k6"]; ok {
		t.Errorf("%v", got["k6"])
	}

	if want := 91; len(got) != want {
		t.Errorf("len(got) %v\nwant %v", len(got), want)
	}
}
