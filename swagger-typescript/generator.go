package main

import (
	"bytes"
	"fmt"
	"github.com/go-openapi/spec"
)

type Generator struct {
	buf *bytes.Buffer
}

func NewGenerator() *Generator {
	g := &Generator{}
	g.buf = bytes.NewBufferString("")

	return g
}

func (g *Generator) P(format string, a ...interface{}) {
	g.buf.WriteString(fmt.Sprintf(format, a...))
}

func (g *Generator) Pn(format string, a ...interface{}) {
	g.buf.WriteString(fmt.Sprintf(format+"\n", a...))
}

func (g *Generator) genSchemaField(name string, schema spec.Schema) (err error) {
	g.P("    %s", name)

	required := false
	if schema.Required != nil {
		for _, v := range schema.Required {
			if name == v {
				required = true
				break
			}
		}
	}
	if !required {
		g.P("?")
	}

	g.P(":")

	if schema.Type != nil {
		if len(schema.Type) == 1 {
			switch schema.Type[0] {
			case "string":
				break
			case "integer":
				break
			case "number":
				break
			case "boolean":
				break
			default:
				return fmt.Errorf("unknown schema.Type " + schema.Type[0])
			}
		} else {
			return fmt.Errorf("len(schema.Type) == 1 else")
		}
	} else {
		return fmt.Errorf("schema.Type != nil else")
	}

	g.Pn(";")

	return nil
}

func (g *Generator) genDefinition(name string, schema spec.Schema) (err error) {
	g.Pn("export interface %s{", name)
	if schema.Properties != nil {
		for fieldName, field := range schema.Properties {
			err = g.genSchemaField(fieldName, field)
			if err != nil {
				return err
			}
		}
	}
	g.P("}")
	g.Pn("")

	return nil
}

func (g *Generator) Gen(swagger *spec.Swagger) (code string, err error) {
	host := swagger.Host
	if host == "" {
		host = "https://localhost"
	}

	g.Pn("const BASE_PATH = \"%s\".replace(/\\/+$/, \"\");", swagger.Host+swagger.BasePath)
	g.Pn("")

	if swagger.Definitions != nil { //todo sort
		for k, v := range swagger.Definitions {
			err = g.genDefinition(k, v)
			if err != nil {
				return "", err
			}
		}
	}

	return g.buf.String(), nil
}
