package posts

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

const postsPath = "./posts"
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

func makeNewPath(path string) string {
	trimmedPath := strings.Split(path, ".md")[0]

	return fmt.Sprintf("%s/%s.html", outputPath, trimmedPath)
}

func Make() {
	files := getMarkdownFiles()

	for _, file := range files {
		markdownBytes, err := os.ReadFile(file)

		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		if err := goldmark.Convert(markdownBytes, &buf); err != nil {
			panic(err)
		}

		newPath := makeNewPath(file)

		directoryCreationErr := os.MkdirAll(filepath.Dir(newPath), 0770)

		if directoryCreationErr != nil {
			panic(err)
		}
		_, fileCreationErr := os.Create(newPath)

		if fileCreationErr != nil {
			panic(fileCreationErr)
		}

		os.WriteFile(newPath, buf.Bytes(), 0644)
	}
}
