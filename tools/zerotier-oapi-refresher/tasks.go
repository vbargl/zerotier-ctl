package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	oapigen "github.com/vbargl/oapi-codegen/v2/pkg/codegen"
	oapiutil "github.com/vbargl/oapi-codegen/v2/pkg/util"
)

var (
	OAPIConfigurationTemplate = oapigen.Configuration{
		Generate: oapigen.GenerateOptions{
			Models:       true,
			Client:       true,
			EmbeddedSpec: true,
		},
		Compatibility: oapigen.CompatibilityOptions{
			AlwaysPrefixEnumValues: true,
		},
		OutputOptions: oapigen.OutputOptions{
			NullableType:        true,
			GenAnonymousObjects: true,
		},
	}
)

func downloadSpec(file string) error {
	resp, err := http.Get(OAPILink)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateSpec(specFile string, genFile string) error {
	f, err := oapiutil.LoadSwagger(specFile)
	if err != nil {
		return fmt.Errorf("error loading spec: %v", err)
	}

	cfg := OAPIConfigurationTemplate
	cfg.PackageName = filepath.Base(filepath.Dir(genFile))
	if cfg.PackageName == "." {
		cwd, _ := os.Getwd()
		cfg.PackageName = filepath.Base(cwd)
	}

	content, err := oapigen.Generate(f, cfg)
	if err != nil {
		return fmt.Errorf("error generating code: %v", err)
	}

	err = os.WriteFile(genFile, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error saving code: %v", err)
	}

	return nil
}
