package posts

import (
	"context"
	"fmt"
	"os"
	"strings"
	"website/internal/views"
)

func getFileNames() []string {
	var files []string

	for _, file := range getMarkdownFiles() {
		withoutExtension := strings.Split(file, ".md")[0]
		withoutPostsPrefix := strings.Split(withoutExtension, "posts/")[1]

		files = append(files, withoutPostsPrefix)
	}

	return files
}

func Make() {
	generateFiles()

	// TODO change this to display based on year
	files := getFileNames()

	htmlFile, err := os.Create(fmt.Sprintf("%s/posts/index.html", outputPath))

	if err != nil {
		panic(err)
	}

	views.PostsIndex(files).Render(context.Background(), htmlFile)

}
