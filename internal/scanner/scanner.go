package scanner

import (
	"io/fs"
	"path/filepath"
)

func ScanRepo(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Only Go files for now
		if !d.IsDir() && filepath.Ext(path) == ".go" {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
