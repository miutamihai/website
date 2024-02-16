package main

import (
	"bytes"
	"context"
	"os"
	"website/views"

	"github.com/yuin/goldmark"
)

func main() {
	markdownBytes, err := os.ReadFile("posts/test.md")

	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(markdownBytes, &buf); err != nil {
		panic(err)
	}

	os.WriteFile("out/test.html", buf.Bytes(), 0644)

	indexFile, err := os.Create("out/index.html")

	if err != nil {
		panic(err)
	}

	home := views.Home()

	home.Render(context.Background(), indexFile)
}
