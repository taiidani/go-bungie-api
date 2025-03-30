package internal

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type Component struct {
	Enums  []Enum
	Object *Object
}

func RenderComponents(ctx context.Context, dir string, key string, schema openapi3.SchemaRef) error {
	filename := "component_" + strings.ReplaceAll(key, ".", "_") + ".go"
	f, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		return fmt.Errorf("could not create file %q: %w", filename, err)
	}
	defer f.Close()

	enums, err := renderEnums(ctx, key, schema)
	if err != nil {
		return fmt.Errorf("could not render enums in %q: %w", filename, err)
	}

	cmp := Component{
		Enums: enums,
	}

	obj, err := renderObject(ctx, key, schema)
	if err != nil {
		return fmt.Errorf("could not render object in %q: %w", filename, err)
	} else if obj != nil {
		cmp.Object = obj
	}

	tmpl := template.Must(template.ParseFS(templates, "templates/**"))
	return tmpl.ExecuteTemplate(f, "component.go.tmpl", cmp)
}

type Enum struct {
	Name         string
	Format       string
	Key          string
	Identifier   string
	NumericValue string
	Description  []string
}

func renderEnums(_ context.Context, key string, schema openapi3.SchemaRef) ([]Enum, error) {
	if !schema.Value.Type.Is("integer") {
		return nil, nil
	}

	format := schema.Value.Format

	enumValues, ok := schema.Value.Extensions["x-enum-values"]
	if !ok {
		slog.Debug("No enums", "filename", key)
		return nil, nil
	}

	valArray, ok := enumValues.([]any)
	if !ok {
		return nil, fmt.Errorf("invalid x-enum-values specification")
	}

	ret := []Enum{}
	for i, en := range valArray {
		add := Enum{
			Format: format,
			Key:    key,
		}

		enum, ok := en.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid x-enum-values %q enum %d", key, i)
		}

		identifier, ok := enum["identifier"]
		if !ok {
			return nil, fmt.Errorf("no identifier on %q enum %d", key, i)
		}
		add.Identifier = identifier.(string)
		add.Name = fmt.Sprintf("%s%s", strings.ReplaceAll(key, ".", "_"), identifier)

		numericValue, ok := enum["numericValue"]
		if !ok {
			return nil, fmt.Errorf("no numericValue on %q enum %d", key, i)
		}
		add.NumericValue = numericValue.(string)

		if description, ok := enum["description"]; ok {
			add.Description = strings.Split(description.(string), "\n")
		}

		ret = append(ret, add)
	}

	return ret, nil
}

type Object struct {
	Name       string
	Properties []Property
}

type Property struct {
	Name        string
	Identifier  string
	Type        string
	Key         string
	Description []string
}

func renderObject(_ context.Context, key string, schema openapi3.SchemaRef) (*Object, error) {
	if !schema.Value.Type.Is("object") {
		return nil, nil
	}

	ret := &Object{
		Name:       strings.ReplaceAll(key, ".", "_"),
		Properties: []Property{},
	}

	for id, prop := range schema.Value.Properties {
		add := Property{
			Name:        strings.ToUpper(string(id[0])) + id[1:],
			Identifier:  id,
			Key:         key,
			Type:        resolveTypeFormat(prop),
			Description: strings.Split(prop.Value.Description, "\n"),
		}

		ret.Properties = append(ret.Properties, add)
	}

	return ret, nil
}

func resolveTypeFormat(prop *openapi3.SchemaRef) string {
	switch prop.Value.Format {
	case "date-time":
		// TODO Can we get more specific?
		return "string"
	case "byte", "int16", "int32", "uint32", "int64":
		return prop.Value.Format
	case "float":
		return "float32"
	case "double":
		return "float64"
	default:
		if prop.Value.Type.Is("string") {
			return "string"
		} else if prop.Value.Type.Is("boolean") {
			return "bool"
		} else if prop.Value.Type.Is("integer") {
			// We couldn't find a more specific type in either Type or Format
			return "int"
		} else if prop.Value.Type.Is("array") {
			return "[]" + resolveTypeFormat(prop.Value.Items)
		} else if prop.Ref != "" {
			// #/components/schemas/Destiny.Config.ImagePyramidEntry
			ref := strings.TrimPrefix(prop.Ref, "#/components/schemas/")
			return strings.ReplaceAll(ref, ".", "_")
		} else if prop.Value.Type.Is("object") {
			// TODO
			return "any"
		} else {
			slog.Error(fmt.Sprintf("'%q' property type and %q format not supported", prop.Value.Type, prop.Value.Format))
			return ""
		}
	}
}
