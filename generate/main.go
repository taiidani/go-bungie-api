package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

const BungieAPISpecURL = "https://raw.githubusercontent.com/Bungie-net/api/refs/heads/master/openapi.json"

func main() {
	// Interruptible context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// One temp directory to hold the API for parsing
	tmp, teardown := tempDirectory()
	defer teardown()

	err := downloadAPISpec()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(2)
	}

	loader := openapi3.NewLoader()
	loader.Context = ctx
	loader.IsExternalRefsAllowed = true

	doc, err := loader.LoadFromFile("openapi.json")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(2)
	}

	if err = doc.Validate(loader.Context); err != nil {
		slog.Warn(err.Error())
	}

	for key, schema := range doc.Components.Schemas {
		if err := renderComponents(ctx, tmp, key, *schema); err != nil {
			slog.Error(err.Error())
			os.Exit(3)
		}
	}

	// Copy the generated spec into the output directory
	os.RemoveAll("../api")
	if err := os.CopyFS("../api", os.DirFS(tmp)); err != nil {
		slog.Error(err.Error())
		os.Exit(4)
	}
}

func tempDirectory() (string, func()) {
	tmp, err := os.MkdirTemp("", "go-bungie-api")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return tmp, func() {
		_ = os.RemoveAll(tmp)
	}
}

func downloadAPISpec() error {
	const filename = "openapi.json"
	if _, err := os.Stat(filename); err == nil {
		return nil
	}

	resp, err := http.Get(BungieAPISpecURL)
	if err != nil {
		return fmt.Errorf("could not download spec: %w", err)
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create local spec file: %w", err)
	}

	_, err = io.Copy(f, resp.Body)
	return err
}

func renderComponents(ctx context.Context, dir string, key string, schema openapi3.SchemaRef) error {
	// TODO Add more type support
	if !schema.Value.Type.Is("integer") {
		return nil
	}

	filename := "component_" + strings.ReplaceAll(key, ".", "_") + ".go"
	f, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		slog.Error(err.Error(), "filename", key)
		return nil
	}
	defer f.Close()

	fmt.Fprintf(f, "// This file was generated by https://github.com/taiidani/go-bungie-api/tree/main/generate.\n// DO NOT EDIT THIS FILE.\npackage api\n\n")

	return renderEnums(ctx, f, key, schema)
}

func renderEnums(_ context.Context, f io.Writer, key string, schema openapi3.SchemaRef) error {
	enumValues, ok := schema.Value.Extensions["x-enum-values"]
	if !ok {
		slog.Debug("No enums", "filename", key)
		return nil
	}

	valArray, ok := enumValues.([]any)
	if !ok {
		return fmt.Errorf("invalid x-enum-values specification")
	}

	if len(valArray) == 0 {
		return nil
	}

	fmt.Fprintf(f, "const (\n")
	for i, en := range valArray {
		enum, ok := en.(map[string]any)
		if !ok {
			return fmt.Errorf("invalid x-enum-values %q enum %d", key, i)
		}

		identifier, ok := enum["identifier"]
		if !ok {
			return fmt.Errorf("no identifier on %q enum %d", key, i)
		}

		numericValue, ok := enum["numericValue"]
		if !ok {
			return fmt.Errorf("no numericValue on %q enum %d", key, i)
		}

		fqIdentifier := fmt.Sprintf("%s%s", strings.ReplaceAll(key, ".", "_"), identifier)
		fmt.Fprintf(f, "    // %s represents the %s enum in the %s component.\n", fqIdentifier, identifier, key)

		if description, ok := enum["description"]; ok {
			fmt.Fprint(f, commentize(description.(string)))
		}
		fmt.Fprintf(f, "    %s = %s\n\n", fqIdentifier, numericValue)
	}
	fmt.Fprintf(f, ")\n\n")

	return nil
}

func commentize(in string) string {
	strs := strings.Split(in, "\n")
	ret := ""
	for _, str := range strs {
		ret += "    //\n"
		ret += "    // " + str + "\n"
	}

	return ret
}
