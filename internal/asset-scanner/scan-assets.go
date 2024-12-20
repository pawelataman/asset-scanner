package asset_scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ScanAssets(assetsPath string, targetPath string, assetFileExt string, withRemove bool) error {
	assets, err := initAssetsOccurrence(assetsPath, assetFileExt)

	if err != nil {
		return err
	}

	err = searchForUsed(assets, targetPath)

	if err != nil {
		return err
	}

	for assetPath, assetOccurrence := range assets {
		if assetOccurrence.Occurrence == 0 {
			fmt.Println(assetPath, assetOccurrence.Occurrence)

			if withRemove {
				if err = os.Remove(assetPath); err != nil {
					return fmt.Errorf("could not delete file %s", assetPath)
				}
			}
		}
	}

	return nil
}

type AssetOccurrence struct {
	Name       string
	Occurrence int
}

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

func matchExtension(fileName string, extension string) bool {
	return strings.Contains(fileName, extension)
}
