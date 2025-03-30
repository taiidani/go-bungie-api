package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/taiidani/go-bungie-api/generate/internal"
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

	if err = doc.Validate(loader.Context, openapi3.DisableExamplesValidation()); err != nil {
		slog.Warn(err.Error())
	}

	for key, schema := range doc.Components.Schemas {
		if err := internal.RenderComponents(ctx, tmp, key, *schema); err != nil {
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
