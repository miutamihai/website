package posts

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"website/internal/views"

	"github.com/yuin/goldmark"
)

const postsPath = "./posts"

func makeNewPath(path string) string {
	trimmedPath := strings.Split(path, ".md")[0]

	return fmt.Sprintf("%s/%s.html", outputPath, trimmedPath)
}

func generateFiles() {
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

		htmlFile, fileCreationErr := os.Create(newPath)

		if fileCreationErr != nil {
			panic(fileCreationErr)
		}

		views.Page(buf.String()).Render(context.Background(), htmlFile)
	}
}
