package shader

import (
	"bytes"
	"text/template"
)

type Uniform interface{}

type Shader struct {
	Uniforms map[string]Uniform
	Vertex   string
	Fragment string
}

func Template(source *template.Template, data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := source.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
