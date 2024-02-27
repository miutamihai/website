package posts

import (
	"io/fs"
	"path/filepath"
)

const outputPath = "./out"

func getMarkdownFiles() []string {
	var files []string

	err := filepath.Walk(postsPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".md" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}
