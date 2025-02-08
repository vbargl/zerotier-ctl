package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	version  = "v1.1.1"
	OAPILink = fmt.Sprintf("https://github.com/zerotier/zerotier-one-api-spec/releases/download/%s/openapi.yaml", version)
)

var (
	specFile = flag.String("output-spec-file", "", "Path to the OpenAPI spec file (default: temp file)")
	genFile  = flag.String("output-gen-file", "zerotieroapi_gen.go", "Path to the OpenAPI spec file (default: zerotieroapi_gen.go)")
)

func main() {
	flag.Parse()

	if *specFile == "" {
		b := make([]byte, 3)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println("could not create temporary file: ", err)
			os.Exit(1)
		}

		*specFile = filepath.Join(os.TempDir(), fmt.Sprintf("zerotier-openapi-%s.yaml", hex.EncodeToString(b)))
	}

	err := downloadSpec(*specFile)
	if err != nil {
		fmt.Println("error downloading spec: ", err)
		os.Exit(1)
	}

	err = generateSpec(*specFile, *genFile)
	if err != nil {
		fmt.Println("error generating spec", err)
		os.Exit(1)
	}
}
