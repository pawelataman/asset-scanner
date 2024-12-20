// Package main provides an asset scanner CLI tool
package main

import (
	"flag"
	"fmt"
	"github.com/pawelataman/asset-scanner/internal/asset-scanner"
)

func main() {
	// CLI flags
	var withRemove bool
	var fileExt string
	flag.BoolVar(&withRemove, "remove", false, "Remove assets that were not found in target directory")
	flag.StringVar(&fileExt, "ext", "", "Asset file extension")

	flag.Parse()

	// Get required directory paths from arguments
	assetPath := flag.Arg(0)
	targetPath := flag.Arg(1)

	// Validate required arguments
	if assetPath == "" {
		panic(fmt.Errorf("asset path must be provided"))
	}
	if targetPath == "" {
		panic(fmt.Errorf("target path must be provided"))
	}

	// Run the scanner
	err := asset_scanner.ScanAssets(assetPath, targetPath, fileExt, withRemove)
	if err != nil {
		panic(err)
	}
}
