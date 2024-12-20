// Package asset_scanner provides functionality to scan for asset usage in target directories
package asset_scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// AssetOccurrence tracks how many times an asset is used
type AssetOccurrence struct {
	Name       string
	Occurrence int
}

// ScanAssets scans target directory for asset usage and optionally removes unused assets
func ScanAssets(assetsPath string, targetPath string, assetFileExt string, withRemove bool) error {
	// Initialize map of assets and their occurrences
	assets, err := initAssetsOccurrence(assetsPath, assetFileExt)
	if err != nil {
		return err
	}

	// Search for asset usage in target directory
	err = searchForUsed(assets, targetPath)
	if err != nil {
		return err
	}

	// Handle unused assets
	for assetPath, assetOccurrence := range assets {

		// Check for occurrence that equals 0
		if assetOccurrence.Occurrence == 0 {

			if withRemove {
				if err = os.Remove(assetPath); err != nil {
					return fmt.Errorf("could not delete file %s", assetPath)
				}
				fmt.Printf("Removed asset at: %s \n", assetPath)
			} else {
				fmt.Printf("Found unused asset at: %s \n", assetPath)
			}
		}
	}

	return nil
}

// initAssetsOccurrence creates a map of assets from the source directory
func initAssetsOccurrence(assetsPath string, assetFileExt string) (map[string]*AssetOccurrence, error) {
	assets := make(map[string]*AssetOccurrence)

	err := filepath.WalkDir(assetsPath, func(path string, dirEntry os.DirEntry, err error) error {
		if !dirEntry.IsDir() && matchExtension(dirEntry.Name(), assetFileExt) {
			assets[path] = &AssetOccurrence{
				Name:       strings.SplitAfter(path, assetsPath)[1],
				Occurrence: 0,
			}
		}
		return nil
	})

	return assets, err
}

// searchForUsed walks through target directory and counts asset occurrences
func searchForUsed(assets map[string]*AssetOccurrence, targetPath string) error {
	return filepath.WalkDir(targetPath, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}

		bytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("could not read content of a file: %s", path)
		}

		for _, assetOccurrence := range assets {
			if strings.Contains(string(bytes), assetOccurrence.Name) {
				assetOccurrence.Occurrence += 1
				continue
			}
		}

		return nil
	})
}

// matchExtension checks if filename contains the specified extension
func matchExtension(fileName string, extension string) bool {
	return strings.Contains(fileName, extension)
}
