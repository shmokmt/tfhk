package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func main() {
	recursive := false
	flag.BoolVar(&recursive, "recursive", false, "Also process files in subdirectories. By default, only the given directroy (or current directroy) is processed.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-recursive] [target]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	dirEntry := flag.Arg(0)

	if dirEntry == "" {
		dirEntry = "."
	}

	err := filepath.Walk(dirEntry, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path == dirEntry {
			return nil
		}
		if !recursive && info.IsDir() {
			return filepath.SkipDir
		}
		if !strings.HasSuffix(path, ".tf") {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Failed to read file: %s\n", err)
			return nil
		}

		hclFile, diags := hclwrite.ParseConfig(content, path, hcl.InitialPos)
		if diags.HasErrors() {
			fmt.Printf("Failed to parse HCL: %s\n", diags.Error())
			return nil
		}
		removeBlocks(hclFile.Body())
		err = os.WriteFile(path, hclFile.Bytes(), 0644)
		if err != nil {
			fmt.Printf("Failed to write file: %s\n", err)
			return nil
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func removeBlocks(body *hclwrite.Body) {
	for _, block := range body.Blocks() {
		switch block.Type() {
		case "moved", "import", "removed":
			body.RemoveBlock(block)
		}
	}
}
