package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/k1LoW/glyph"
)

const ts = `# Preset Icons

| Name | Icon |
| ---- | ---- |
{{- range $_, $p := .Preset }}
| {{ index $p 0 }} | ![{{ index $p 0 }}]({{ index $p 1 }}) |
{{- end }}
`

func main() {
	m := glyph.NewMapWithPreset()
	preset := [][]string{}
	for _, k := range m.Keys() {
		g, err := m.Get(k)
		if err != nil {
			panic(err)
		}
		p := filepath.Join("img", "preset", fmt.Sprintf("%s.svg", k))
		i, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) // #nosec
		if err != nil {
			panic(err)
		}
		if err := g.Write(i); err != nil {
			_ = i.Close()
			panic(err)
		}
		_ = i.Close()
		preset = append(preset, []string{k, p})
	}
	tmpl := template.Must(template.New("preset").Parse(ts))
	tmplData := map[string]interface{}{
		"Preset": preset,
	}
	md, err := os.OpenFile("preset.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) // #nosec
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(md, tmplData); err != nil {
		_ = md.Close()
		panic(err)
	}
	_ = md.Close()
}
