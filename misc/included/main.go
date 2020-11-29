package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/k1LoW/glyph"
)

const ts = `# Included Icon Set

| Name | Icon (SVG) | Icon (PNG) |
| ---- | ---- | ---- |
{{- range $_, $p := .Included }}
| {{ index $p 0 }}{{ if ne (index $p 3) "" }} ( {{ index $p 3 }} ){{ end }} | ![{{ index $p 0 }}]({{ index $p 1 }}) | ![{{ index $p 0 }}]({{ index $p 2 }}) |
{{- end }}
`

func main() {
	m := glyph.NewMapWithIncluded(glyph.Width(200.0), glyph.Height(200.0))
	included := [][]string{}
	aliases := glyph.IncludedAliases()
	for _, k := range m.Keys() {
		a, ok := aliases[k]
		if !ok {
			continue
		}
		g, err := m.Get(k)
		if err != nil {
			panic(err)
		}
		p := filepath.Join("img", "included", fmt.Sprintf("%s.svg", k))
		_, _ = fmt.Fprintf(os.Stderr, "create %s\n", p)
		i, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) // #nosec
		if err != nil {
			panic(err)
		}
		if err := g.Write(i); err != nil {
			_ = i.Close()
			panic(err)
		}
		_ = i.Close()

		pp := filepath.Join("img", "included", fmt.Sprintf("%s.png", k))
		_, _ = fmt.Fprintf(os.Stderr, "create %s\n", pp)
		ip, err := os.OpenFile(pp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) // #nosec
		if err != nil {
			panic(err)
		}
		if err := g.WriteImage(ip); err != nil {
			_ = ip.Close()
			panic(err)
		}
		_ = ip.Close()

		included = append(included, []string{k, p, pp, strings.Join(a, ", ")})
	}
	tmpl := template.Must(template.New("included").Parse(ts))
	tmplData := map[string]interface{}{
		"Included": included,
	}
	_, _ = fmt.Fprintf(os.Stderr, "create %s\n", "included.md")
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
