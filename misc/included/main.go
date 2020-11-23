package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/k1LoW/glyph"
)

const ts = `# Included Icon Set

| Name | Icon |
| ---- | ---- |
{{- range $_, $p := .Included }}
| {{ index $p 0 }} | ![{{ index $p 0 }}]({{ index $p 1 }}) |
{{- end }}
`

func main() {
	m := glyph.NewMapWithIncluded()
	included := [][]string{}
	for _, k := range m.Keys() {
		g, err := m.Get(k)
		if err != nil {
			panic(err)
		}
		p := filepath.Join("img", "included", fmt.Sprintf("%s.svg", k))
		i, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) // #nosec
		if err != nil {
			panic(err)
		}
		if err := g.Write(i); err != nil {
			_ = i.Close()
			panic(err)
		}
		_ = i.Close()
		included = append(included, []string{k, p})
	}
	tmpl := template.Must(template.New("included").Parse(ts))
	tmplData := map[string]interface{}{
		"Included": included,
	}
	md, err := os.OpenFile("included.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) // #nosec
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(md, tmplData); err != nil {
		_ = md.Close()
		panic(err)
	}
	_ = md.Close()
}
