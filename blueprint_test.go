package glyph

import (
	"testing"

	"github.com/goccy/go-yaml"
)

func TestBlueprintYAMLUnmarshal(t *testing.T) {
	tests := []struct {
		in          string
		wantKey     string
		wantLineLen int
		wantTextLen int
	}{
		{
			`---
key: db
lines:
  - b0 d0 h0 j0 j6 h8 d8 b6 b0
  - b0 d2 h2 j0
  - b2 d4 h4 j2
  - b4 d6 h6 j4
`,
			"db",
			4,
			0,
		},
	}
	for _, tt := range tests {
		b := NewBlueprint()
		yaml.Unmarshal([]byte(tt.in), b)
		g, key, err := b.ToGlyphAndKey()
		if err != nil {
			t.Fatal(err)
		}
		if got := key; got != tt.wantKey {
			t.Errorf("got %v\nwant %v", got, tt.wantKey)
		}
		if got := len(g.lines); got != tt.wantLineLen {
			t.Errorf("got %v\nwant %v", got, tt.wantLineLen)
		}
		if got := len(g.texts); got != tt.wantTextLen {
			t.Errorf("got %v\nwant %v", got, tt.wantTextLen)
		}
	}
}
